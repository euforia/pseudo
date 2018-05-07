package pseudo

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/euforia/pseudo/scope"
	"github.com/hashicorp/hil/ast"
	"github.com/stretchr/testify/assert"
)

var (
	testJobSpec       = "./test-fixtures/jobspec.nomad.hcl"
	testPlatformSpec  = "./test-fixtures/platform.hcl"
	testRandVars      = "./test-fixtures/vars.tfvars"
	testScopeVarsSpec = "./etc/context.hcl"
)

var testVariables = scope.Variables{
	"region.id": ast.Variable{
		Type:  ast.TypeString,
		Value: "us-west-2",
	},
	"region.datacenters": ast.Variable{
		Type: ast.TypeList,
		Value: []ast.Variable{
			ast.Variable{
				Type:  ast.TypeString,
				Value: "dc1",
			},
			ast.Variable{
				Type:  ast.TypeString,
				Value: "dc2",
			},
		},
	},
	"platform.enclave.id": ast.Variable{
		Type:  ast.TypeString,
		Value: "pci",
	},
	"platform.env.id": ast.Variable{
		Type:  ast.TypeString,
		Value: "dev",
	},
	"platform.env.internal_domain": ast.Variable{
		Type:  ast.TypeString,
		Value: "dev",
	},
	"app.version": ast.Variable{
		Type:  ast.TypeString,
		Value: "v1.0.1-3-a1b2c3d4",
	},
	"app.name": ast.Variable{
		Type:  ast.TypeString,
		Value: "testapp",
	},
	"app.tags": ast.Variable{
		Type: ast.TypeList,
		Value: []ast.Variable{
			ast.Variable{
				Type:  ast.TypeString,
				Value: "dev",
			},
		},
	},
}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	os.Exit(m.Run())
}

func Test_VM_basic(t *testing.T) {
	vm := NewVM()
	vm.SetVars(testVariables)

	_, err := vm.Eval()
	assert.NotNil(t, err)

	b, _ := ioutil.ReadFile(testJobSpec)
	err = vm.Parse(string(b))
	assert.Nil(t, err)

	_, err = vm.Eval()
	assert.Nil(t, err)
}

func Test_VM_script_dir(t *testing.T) {
	files, err := ReadDirFiles("./etc")
	assert.Nil(t, err)

	stream := make([]byte, 0)
	for k := range files {
		stream = append(stream, files[k]...)
	}

	vm := NewVM()
	err = vm.Parse(string(stream))
	assert.Nil(t, err)
}

func Test_VM(t *testing.T) {
	uri, _ := url.Parse(testScopeVarsSpec)
	vars, err := LoadVariables(uri, IndexOptions{ContentType: "pseudo"})
	//vars, err := LoadHCLScopeVarsFromFile(testScopeVarsSpec)
	assert.Nil(t, err)

	// vars := idx.Variables()
	vm := NewVM()
	vm.SetVars(vars)

	script, _ := NewScript(testJobSpec)
	err = vm.Parse(script.Contents())
	assert.Nil(t, err)

	_, err = vm.Eval()
	assert.Nil(t, err)

	assert.Equal(t, len(vars), len(vm.VarNames()))
	assert.Equal(t, len(vm.funcs), len(vm.FuncNames()))
}

func Test_VM_RegisterFunc(t *testing.T) {
	vm := NewVM()
	err := vm.RegisterFunc("tld", tldFunction())
	assert.NotNil(t, err)

	err = vm.RegisterFunc("test", ast.Function{})
	assert.Nil(t, err)
}
