package scope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testScopeVarsSpec = "../test-fixtures/platform.hcl"
)

func Test_LoadHCLScopeVarsFromFile(t *testing.T) {
	m, err := BuildHCLScopeVarsFromFile(testScopeVarsSpec)
	assert.Nil(t, err)
	t.Log(m.Names())
}
