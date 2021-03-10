package dpkg

import (
	"bufio"
	"os"
	"testing"
)

func TestParseStatus(t *testing.T) {

	fileHandle, _ := os.Open("tests/status")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	pkgs := ParseStatus(fileScanner)

	if len(pkgs) != 23 {
		t.Errorf("Number of parsed packages was incorrect, got: %d, want: %d.", len(pkgs), 23)
	}

	if pkgs[10].Package != "bind9-host" {
		t.Errorf("Parsed package was incorrect, got: %s, want: %s.", pkgs[10].Package, "bind9-host")
	}

	if len(pkgs[9].Conffiles) != 2 {
		t.Errorf("Parsed package Conffiles was incorrect, got: %d, want: %d.", len(pkgs[9].Conffiles), 2)
	}

}
