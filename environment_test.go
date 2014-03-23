package environment

import "testing"

func TestString(t *testing.T) {
	LoadProperties()
	if String("username") != "james" {
		t.Error("Username should be james")
	}
	if String("password") != "007" {
		t.Error("Password should be '007'")
	}
	if String("url") != "http://www.fi" {
		t.Error("Url should be 'http://www.fi'")
	}
}

func TestNonExistingKey(t *testing.T) {
	LoadProperties()
	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered!")
		}
	}()
	value := String("non-existing-key")
	t.Errorf("'non-existing-key' returned '%s'!", value)
}

func TestConvertStringToEnvironment(t *testing.T) {
	var env *Environment = StringToEnvironment("Test")
	if *env != TEST {
		t.Error("String converstion to environment failed")
	}
}

func TestCurrentEnvironment(t *testing.T) {
	if *Current() != DEV {
		t.Error("Current environment should default to 'dev'")
	}
}
