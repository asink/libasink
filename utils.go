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
    "os"
)

// Returns the current working directory
// as a string
func getWorkingDirectory() string {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    return dir
}
