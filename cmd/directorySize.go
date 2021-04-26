package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/reactivex/rxgo/v2"
	"github.com/spf13/cobra"
)

// directorySizeCmd represents the directorySize command
var directorySizeCmd = &cobra.Command{
	Use:   "directorySize",
	Short: "Computes the total size of a directory",
	Long:  `Will `,
	Run: func(cmd *cobra.Command, args []string) {
		handleCommand(args)
	},
}

func init() {
	rootCmd.AddCommand(directorySizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// directorySizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// directorySizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Handle the command for computing the size of all provided directories
func handleCommand(paths []string) (size int) {
	sizeSum := 0

	// RXGo leverages GoRoutines to become concurrent, should we enconter an error we should stop processing as we may have invaild filepaths
	observable := rxgo.Just(paths)().Map(getFullPath, rxgo.WithCPUPool(), rxgo.WithErrorStrategy(rxgo.StopOnError))
	for path := range observable.Observe() {
		size := computeSize(path.V.(string))
		fmt.Printf("Directory size of %s is: %d \n", path.V, size)
		sizeSum += size
	}
	fmt.Printf("Total size of all directories is: %d\n", sizeSum)

	return
}

// Recursive function to compute the total size of all child directories and files.
func computeSize(path string) (size int) {
	files, err := ioutil.ReadDir(path)
	size = 0

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if err != nil {
			log.Fatal(err)
		}
		if file.IsDir() {
			newPath := path + "/" + file.Name() + "/"
			size += computeSize(newPath)
		}
		size += int(file.Size())
	}

	return
}

// Constructs the absolute path value of a provided relative path.
// Meant to be used within a rxos stream, otherwise pass in a nil context.
func getFullPath(ctx context.Context, path interface{}) (fullPath interface{}, err error) {

	fullPath, err = filepath.Abs(path.(string))

	if err != nil {
		log.Fatal(err)
	}

	return
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
