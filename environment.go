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

var props *properties.Properties
var conf *string
var current *Environment

func init() {
	conf = flag.String("conf", "dev.properties", "Full path to configuration file")
	paramEnv := flag.String("env", "dev", "Environment [dev|test|prod]")
	current = StringToEnvironment(*paramEnv)
}

func Current() *Environment {
	return current
}

func StringToEnvironment(s string) *Environment {
	env := new(Environment)
	switch strings.ToLower(s) {
	case "dev":
		*env = DEV
		return env
	case "test":
		*env = TEST
		return env
	case "prod":
		*env = PROD
		return env
	}
	panic("Environment (dev/test/prod) not defined.")
}

func LoadProperties() *properties.Properties {
	flag.Parse()
	file, err := os.Open(*conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(500)
	}
	defer file.Close()
	p := make(properties.Properties)
	p.Load(bufio.NewReader(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(501)
	}
	props = &p
	return &p
}

func (env Environment) String() string {
	return string(env)
}

func StringValue(key, defaultValue string) string {
	return props.String(key, defaultValue)
}
