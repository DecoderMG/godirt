# GoDIRT

![Code Quality](https://www.code-inspector.com/project/21923/score/svg)
![Code Grade](https://www.code-inspector.com/project/21923/status/svg)
## Table of Contents
+ [About](#about)
+ [Getting Started](#getting_started)
+ [Usage](#usage)
+ [Contributing](../CONTRIBUTING.md)

## About <a name = "about"></a>
GoDIRT, otherwise known as Go Directory Inspection and Retroprective Tool is a GoLang based CLI tool used to compute the size of multiple directories.

## Getting Started <a name = "getting_started"></a>
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

Ensure you have [GoLang](https://golang.org/) installed on your system before attempting to run this module. If you wish to run the [pre-commit](https://pre-commit.com/) checks, you will also need toinstall pre-commit and all golang modules required for your pre-commit settings.

### Installing

You should run ```go install godirt``` to install the CLI tool as needed on your system.

## Usage <a name = "usage"></a>

To compute directory sizes after installing, in any terminal you need to run the following:

```godirt directorySize <list of directories to compute separated by spaces>```