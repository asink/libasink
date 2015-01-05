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
    "sync"
)

// Execer represents the Exec method that needs to
// be implemented.
type Execer interface {
    Exec() bool
}

// TaskMap is a map of all created tasks.
var TasksMap map[string]Task

// Task holds the information required to run
// a kind of task that implements the Execer.
type Task struct {
    Name    string
    Process Execer
    Require string
    Group   string
}

// NewTask creates a new instance of Task with some
// default values. The task name string and
// the Execer process are the only initial
// values that are required.
func NewTask(name string, process Execer) Task {
    return Task{name, process, "", ""}
}

// Exec executes a single task, given that there are
// no required tasks attached to it.
func (t Task) Exec() bool {
    p := t.Process

    // Check for any required tasks to execute first
    executeRequiredTask(t)

    if executeGroupedTasks(t) != true {
        if (p != nil) {
            p.Exec()
            delete(TasksMap, t.Name)
        }
    }
    return true
}

// ExecMulti executes multiple tasks from a slice of
// tasks which are organised into a key value
// map first.
func ExecMulti(taskSlice []Task) bool {
    TasksMap = createTasksMap(taskSlice)
    for _, t := range TasksMap {
        t.Exec()
    }
    return true
}

// Converts the initial tasks slice into a key value
// map using the task name as the key and the instance
// as the value.
func createTasksMap(tasks []Task) map[string]Task {
    tasksMap := make(map[string]Task)
    for _, task := range tasks {
        tasksMap[task.Name] = task
    }
    return tasksMap
}

// If a required task has been specefied it will be
// found and ran at this point.
func executeRequiredTask(t Task) {
    if (t.Require != "") {
        task := TasksMap[t.Require]
        task.Exec()
    }
}

// If grouped tasks have been found they will be ran
// asynchronously at this point.
func executeGroupedTasks(task Task) bool {
    if (task.Group != "") {
        group := task.Group
        var wg sync.WaitGroup
        for _, block := range TasksMap {
            if block.Group == group {
                wg.Add(1)
                go executeGroupConcurrently(block, &wg)
            }
        }
        wg.Wait()
        return true
    }
    return false
}

// Allows grouped tasks to be ran without
// any blocking.
func executeGroupConcurrently(t Task, wg *sync.WaitGroup) {
    defer wg.Done()
    process := t.Process
    process.Exec()
    delete(TasksMap, t.Name)
}
