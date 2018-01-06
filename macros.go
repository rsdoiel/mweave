//
// macros.go is a thin wrapper adding macro definition and expantion
// via the shorthand package.  It's just part of the experiment and
// may get removed if it doesn't prove useful. RSD, 2018-01-06
//
package mweave

import (
	"fmt"
	"github.com/rsdoiel/shorthand"
	"strings"
)

// ApplyMacros creates a shorthand VM and evaluates the
// byte array passed in. This adds a Macro ability to mweave
// hopefully making it more literate in the process :-)
func ApplyMacros(text []byte) ([]byte, error) {
	output := []string{}
	vm := shorthand.New()
	vm.SetPrompt("")

	// This is ugly because shorthand works with lines of strings...
	for i, line := range strings.Split(fmt.Sprintf("%s", text), "\n") {
		if s, err := vm.Eval(line, i); err == nil {
			output = append(output, s)
		} else {
			return nil, err
		}
	}
	return []byte(strings.Join(output, "\n")), nil
}
