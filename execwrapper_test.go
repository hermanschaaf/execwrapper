package execwrapper

import (
	"fmt"
	"strings"
	"testing"
)

func TestCommandBasic(t *testing.T) {
	out, err := Command(`echo "hi"`).Output()
	expected := "hi"
	if err != nil {
		fmt.Println(err)
	}
	output := strings.TrimRight(string(out), " \t\n")
	if output != expected {
		t.Errorf("Error: %s != %s", output, expected)
	}
}

func TestCommandPiped(t *testing.T) {
	out, err := Command("echo \"hello\nthere\" | grep \"he\"").Output()
	expected := "hello\nthere"
	if err != nil {
		fmt.Println(err)
	}
	output := strings.TrimRight(string(out), " \t\n")
	if output != expected {
		t.Errorf("Error: %s != %s", output, expected)
	}
}
