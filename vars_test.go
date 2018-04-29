package pseudo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildVarScope(t *testing.T) {
	m, err := LoadScopeVarsFromFile(testScopeVarsSpec)
	assert.Nil(t, err)
	t.Log(m.Names())

}

func Test_Script(t *testing.T) {
	scr, _ := NewScript(testJobSpec)
	vars := scr.Vars()
	t.Log(vars)
}
