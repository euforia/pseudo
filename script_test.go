package pseudo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Script_Vars(t *testing.T) {
	script, _ := NewScript(testRandVars)
	assert.Equal(t, 0, len(script.Vars()))

	script.contents = append(script.contents, []byte("\n${myvar}")...)
	assert.Equal(t, 1, len(script.Vars()))
}
