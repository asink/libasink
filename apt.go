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

import (
    "strings"
)

// Apt contains the information needed to run apt
// get or install commands.
type Apt struct {
    Action   string
    Packages []string
    Callback func(command string)
    Dummy    bool
}

// NewApt creates a new instance of Apt with some
// default values. The action string is the
// only initial value that is required.
func NewApt(action string) Apt {
    a := Apt{}
    a.Action   = action
    a.Packages = []string{}
    a.Callback = func(command string){}
    a.Dummy    = false
    return a
}

// Exec generates the command string to be ran based
// on apt-get. Currently supports 'update' or
// 'install'.
func (a Apt) Exec() bool {
    c := NewCommand("apt-get")
    c.Args = append([]string{a.Action, "-y"}, a.Packages...)
    a.Callback("apt-get " + strings.Join(c.Args, " "))
    if a.Dummy == false {
        c.Exec()
    }
    return true
}

// AddPackage adds a package to install.
func (a *Apt) AddPackage(p string) {
    a.Packages = append(a.Packages, p)
}

// AddPackages adds multiple packages to install.
func (a *Apt) AddPackages(p []string) {
    for _, pa := range p {
        a.Packages = append(a.Packages, pa)
    }
}
