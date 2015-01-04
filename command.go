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
    "os/exec"
    "sync"
    "strings"
)

type Command struct {
    Name       string
    AsyncCount int
    RelCount   int
    Dir        string
    Args       []string
    Env        []string
    Callback   func(command string)
    Dummy      bool
}

// Creates a new instance of Command with some
// default values. The command string is the
// only initial value that is required
func NewCommand(name string) Command {
    return Command{
        Name: name,                  
        AsyncCount: 1,                     
        RelCount: 1,
        Dir: getWorkingDirectory(), 
        Args: []string{},            
        Env:  []string{},
        Callback: func(command string){}, 
        Dummy: false,
    }
}

// Implemented to satisfy the task's Execer
// interface. Loops through the AsyncCount
// to concurrently execute the command
func (c Command) Exec() bool {
    var wg sync.WaitGroup

    command := make(chan Command)

    c.setenv()
    c.getenv()
    c.chdir()

    for i := 0; i != c.AsyncCount; i++ {
        wg.Add(1)
        go runCommand(command, &wg)
        command <- c
    }

    close(command)
    wg.Wait()
    return true
}

// Sets env vars before executing a command
func (c Command) setenv() {
    for _, e := range c.Env {
        kv := strings.Split(e, "=")
        os.Setenv(kv[0], kv[1])
    }
}

// Gets env vars from the Name, Args and Env
// parts of the command object
func (c *Command) getenv() {
    c.Name = os.ExpandEnv(c.Name)
    c.Dir = os.ExpandEnv(c.Dir)
    for ai, a := range c.Args {
        c.Args[ai] = os.ExpandEnv(a)
    }
    for ei, e := range c.Env {
        c.Env[ei] = os.ExpandEnv(e)
    }
    c.setenv()
}

// Changes to a specified dir before executing
// a command
func (c Command) chdir() {
    os.Chdir(getWorkingDirectory())
    os.Chdir(c.Dir)
}

// Is called within Exec, the actual command
// execution happens in here
func runCommand(command chan Command, wg *sync.WaitGroup) {
    defer wg.Done()

    c := <- command

    for j := 0; j != c.RelCount; j++ {
        c.Callback(c.Name + " " + strings.Join(c.Args, " "))
        if c.Dummy == false {
            cmd := exec.Command(c.Name, c.Args...)
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            cmd.Run()
        }
    }
}
