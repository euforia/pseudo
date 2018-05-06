package pseudo

import (
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hil/ast"
	"github.com/stretchr/testify/assert"
)

func Test_BuildVarScope(t *testing.T) {
	m, err := LoadHCLScopeVarsFromFile(testScopeVarsSpec)
	assert.Nil(t, err)
	t.Log(m.Names())

}

func Test_Script(t *testing.T) {
	scr, _ := NewScript(testJobSpec)
	vars := scr.Vars()
	t.Log(vars)
}

func Test_walk(t *testing.T) {
	in := []byte(`
        number = 4
        float = 1.1
        bool = true
`)
	tree, _ := hcl.ParseBytes(in)
	vars := walk("", tree.Node)

	v, _ := vars[".number"]
	assert.Equal(t, ast.TypeInt, v.Type)

	v, _ = vars[".float"]
	assert.Equal(t, ast.TypeFloat, v.Type)

	v, _ = vars[".bool"]
	assert.True(t, v.Value.(bool))
}

func Test_ReadDirFiles(t *testing.T) {
	_, err := ReadDirFiles("./does/not/exist")
	assert.NotNil(t, err)
}
