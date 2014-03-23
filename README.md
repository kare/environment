# Environment for Go [![Build Status](https://travis-ci.org/kare/environment.svg?branch=master)](https://travis-ci.org/kare/environment)

 This library parses two command line arguments on #init() function. Library depends on [https://github.com/dmotylev/goproperties](https://github.com/dmotylev/goproperties).

* -conf
	* Path to a properties file. Defaults to dev.properties.
* -env
	* Valid values are 'dev', 'test' and 'prod'. Defaults to 'dev'.

#Usage
Example:
```go
package main

import (
	"fmt"
	"github.com/kare/environment"
)

func main() {
	environment.LoadProperties()
	username := environment.String("username", "foobar")
	fmt.Println("Username:", username)
}
```
```sh
$ cat example-dev.properties
username=root
$ ./main -conf example-dev.properties -env dev
Username: root
```
