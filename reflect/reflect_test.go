package reflect

import (
	"strings"
	"testing"
)

/*
	this function does not perform any tests,
 	it is here as it is required for testing the IsFunction and GetFunctionName methods
*/
func iAmAFunc() {}

func TestIsFunction(t *testing.T) {
	res := IsFunction(1)

	if res != false {
		t.Errorf("utils/reflect: IsFunction: expected true, got false for IsFunction(1)")
	}

	res = IsFunction(iAmAFunc)
	if res != true {
		t.Errorf("utils/reflect: IsFunction: expected true, got false for IsFunction(iAmAFunc)")
	}
}

func TestPanicOfGetFunctionName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("utils/reflect: GetFunctionName: expected a panic, code did NOT panic for GetFunctionName(1)")
		}
	}()

	GetFunctionName(1)
}

func TestGetFunctionName(t *testing.T) {
	name := GetFunctionName(iAmAFunc)

	// ignoring any package names and extracting just the function name, thus enabling project independent testing
	parts := strings.Split(name, ".")
	name = parts[len(parts)-1]

	if name != "iAmAFunc" {
		t.Errorf("utils/reflect: GetFunctionName: expected \"iAmAFunc\", got %s for GetFunctionName(iAmAFunc)", name)
	}
}
