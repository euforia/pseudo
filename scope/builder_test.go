package scope

import (
	"io/ioutil"
	"testing"

	"github.com/hashicorp/hil/ast"
	"github.com/stretchr/testify/assert"
)

var (
	testScopeVarsSpec = "../test-fixtures/platform.hcl"
)

func Test_LoadHCLScopeVarsFromFile(t *testing.T) {
	in, _ := ioutil.ReadFile(testScopeVarsSpec)
	b := NewPseudoBuilder("")
	err := b.Build(in)

	assert.Nil(t, err)
	t.Log(b.Variables().Names())
}

func Test_ReflectBuilder(t *testing.T) {
	b := NewReflectBuilder("", ".")
	m := make(map[string]interface{})
	m["test"] = []string{"1"}

	err := b.Build(m)
	assert.Nil(t, err)

	vars := b.Variables()
	vt := vars["test"]
	assert.Equal(t, ast.TypeList, vt.Type)
}
