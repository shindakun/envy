package envy

import (
	"os"
	"testing"
)

// TestGet tests envy.Get()
func TestGet(t *testing.T) {
	os.Setenv("ENVYTEST", "envy-test")

	testenv, _ := Get("ENVYTEST")
	if testenv != "envy-test" {
		t.Errorf("Get(\"ENVYTEST\") returned incorrect variable, got: %s, want: %s.", testenv, "envy-test")
	}

	testerror, err := Get("ENVYERROR")
	if err == nil {
		t.Errorf("Get(\"ENVYERROR\") returned incorrect variable, got: %s, want: %s.", testerror, "")
	}
}
