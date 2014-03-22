package environment

import "testing"

func TestEverything(t *testing.T) {
	LoadProperties()
	if StringValue("username", "") != "james" {
		t.Error("Username should be james")
	}
	if StringValue("password", "") != "007" {
		t.Error("Password should be '007'")
	}
	if StringValue("url", "") != "http://www.fi" {
		t.Error("Url should be 'http://www.fi'")
	}
	var env *Environment = StringToEnvironment("Test")
	if *env != TEST {
		t.Error("Current environment should be test")
	}
}
