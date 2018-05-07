package pseudo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Script(t *testing.T) {
	scr, _ := NewScript(testJobSpec)
	vars := scr.Vars()
	t.Log(vars)
}

func Test_ReadDirFiles(t *testing.T) {
	_, err := ReadDirFiles("./does/not/exist")
	assert.NotNil(t, err)
}

func Test_Script_Vars(t *testing.T) {
	script, _ := NewScript(testRandVars)
	assert.Equal(t, 0, len(script.Vars()))

	script.contents = append(script.contents, []byte("\n${myvar}")...)
	assert.Equal(t, 1, len(script.Vars()))
}

func Test_Script_shabang(t *testing.T) {
	script := NewScriptBytes([]byte(`#!./my/bin`))
	assert.Equal(t, 0, len(script.contents))

	script = NewScriptBytes([]byte(`#!./m


`))
	assert.Equal(t, 2, len(script.contents))

	script = NewScriptBytes([]byte(`#!./bin/f
${var}`))
	assert.Equal(t, []byte("${var}"), script.contents)
}
