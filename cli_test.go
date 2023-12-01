package main

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err = root.ExecuteC()

	return buf.String(), err
}

func TestRootCmd(t *testing.T) {
	_, err := executeCommand(rootCmd)
	assert.NoError(t, err)
}

func TestTaskCmd(t *testing.T) {
	_, err := executeCommand(taskCmd)
	assert.NoError(t, err)
}

func TestUserCmd(t *testing.T) {
	_, err := executeCommand(userCmd)
	assert.NoError(t, err)
}

func TestChatCmd(t *testing.T) {
	_, err := executeCommand(chatCmd)
	assert.NoError(t, err)
}

// Additional tests for individual commands would go here
// For example:
// func TestAddTaskCmd(t *testing.T) {
//     _, err := executeCommand(rootCmd, "task", "add", "TaskName", "TaskDescription", "Pending", "username")
//     assert.NoError(t, err)
// }
