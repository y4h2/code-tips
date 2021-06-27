package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	input := "test"
	cmd.SetArgs([]string{"--in", input})
	err := cmd.Execute()
	assert.NoError(err)

	output, err := ioutil.ReadAll(b)
	assert.NoError(err)
	assert.Equal(string(output), input)
}
