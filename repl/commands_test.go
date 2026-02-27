package repl

import "testing"

var callbackCalled = false

func mockCallback(cfg *Config) error {
	callbackCalled = true
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
	r.ExecuteCommand("test")
	if !callbackCalled {
		t.Errorf("could not execute test command!")
	}
}
