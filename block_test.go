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
    "testing"
    "reflect"
    "fmt"
)

func TestNewBlock(t *testing.T) {
    b := NewBlock(func() {
        fmt.Println("Hello, World!")
    });
    
    b.AsyncCount = 1
    b.RelCount   = 1

    tp := reflect.TypeOf(b).String()
    if tp != "asink.Block" {
        t.Error("Expected asink.Block, got ", tp)
    }
}

func TestExecBlock(t *testing.T) {
    b := NewBlock(func() {
        fmt.Println("Hello, World!")
    });
    result := b.Exec()
    if result != true {
        t.Error("Expected true, got ", result)
    }
}
