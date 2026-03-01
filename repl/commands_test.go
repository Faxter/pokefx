package repl

import "testing"

var callbackCalled = false
var commandParameter = ""

func mockCallback(cfg *Config, param string) error {
	callbackCalled = true
	commandParameter = param
	return nil
}

func TestRegisterCommand(t *testing.T) {
	r := CreateRepl()
	r.registerCommand("test", "some description", mockCallback)
	_, ok := r.commands["test"]
	if !ok {
		t.Errorf("did not register test command!")
	}
}

func TestExecuteCommand(t *testing.T) {
	r := CreateRepl()
	r.registerCommand("test", "some description", mockCallback)
	r.ExecuteCommand("test", "testparam")
	if !callbackCalled {
		t.Errorf("could not execute test command!")
	}
	if commandParameter != "testparam" {
		t.Errorf("parameter was not read correctly!")
	}
}
