// libasink v0.0.1
//
// (c) Ground Six
//
// @package libasink
// @version 0.0.1
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package asink

// Filter represents a slice of commands.
type Filter struct {
    Dummy    bool
    commands []string
}

// A list of software packages defined for commands
// or configuration to be ran before the install.
var packages map[string]func(f *Filter) = map[string]func(f *Filter){

}

// NewFilter ceates a new instance of Filter with a
// default value. The task package string.
func NewFilter() Filter {
    return Filter{false, []string{}}
}

// Apply applies the filter before the package is
// installed.
func (f Filter) Apply(installs []string) {
    for _, p := range installs {
        packages[p](&f)
    }
}

// Commands returns the slice of commands held in
// object as strings.
func (f Filter) Commands() []string {
    return f.commands
}
