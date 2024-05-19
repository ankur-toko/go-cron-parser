# Cron Expression Parser


## Overview
The CRON Expression Parser is a command-line tool that parses CRON expressions and outputs in stdout in the below format 

```
minutes       [0 15 30 45]
hours         [0]
days          [1 15]
month         [1 2 3 4 5 6 7 8 9 10 11 12]
day of week   [1 2 3 4 5]
command       /usr/bin/find
```
 
## Building the Project
Install latest go version for this to work from [here](https://go.dev/doc/install)

To build the project, use the following command:
```sh
make build
```


## Testing
Examples are added in the makefile to test the program. To run example 2, use the below command
```
make ex2
```
which is equivalent to
```
go run . "*/15 0 1,15 * 1-5 /usr/bin/find"
```


Feel free to checkout examples in the makeFile, or run it directly using "go run command" or "go-cron-parser" binary


