package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	var out bytes.Buffer

	assert.Zero(t, run(&out, []string{"arg0", "arg1", "arg2"}))
	assert.Equal(t, out.String(), `Hello, world!
Got args: arg0,arg1,arg2
`)
}
