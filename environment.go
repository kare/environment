package environment

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dmotylev/goproperties"
	"os"
	"strings"
)

type Environment string

const (
	DEV  Environment = "dev"
	TEST Environment = "test"
	PROD Environment = "prod"
)

var currentEnvironment *string
var filename *string

func init() {
	currentEnvironment = flag.String("env", "test", "[dev|test|prod]")
	filename = flag.String("conf", "environment", "environment.properties file usage message")
}

func LoadProperties() *properties.Properties {
	flag.Parse()
	name := fmt.Sprintf("%s-%s.properties", *filename, Current().String())
	file, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
	defer file.Close()
	p := make(properties.Properties)
	p.Load(bufio.NewReader(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
	return &p
}

func Current() Environment {
	switch strings.ToLower(*currentEnvironment) {
	case "dev":
		return DEV
	case "test":
		return TEST
	case "prod":
		return PROD
	}
	panic("Environment (dev/test/prod) not defined.")
}

func (env Environment) String() string {
	return string(env)
}
