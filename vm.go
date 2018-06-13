package pseudo

import (
	"errors"

	"github.com/euforia/pseudo/scope"
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

// VM is used to parse and eval a given script
type VM struct {
	// Root node available after parsing
	a ast.Node
	// Scoped variables
	vars scope.Variables
	// Scoped functions
	funcs map[string]ast.Function
}

// NewVM inits a new VM with core functions
func NewVM() *VM {
	vm := &VM{
		funcs: make(map[string]ast.Function),
	}

	vm.init()

	return vm
}

func (vm *VM) init() {
	vm.funcs = map[string]ast.Function{
		"length":  lengthFunction(),
		"replace": replaceFunction(),
		"tld":     tldFunction(),
	}
}

// FuncNames returns a list of function names in scope for this vm instance
func (vm *VM) FuncNames() []string {
	out := make([]string, 0, len(vm.funcs))
	for k := range vm.funcs {
		out = append(out, k)
	}
	return out
}

// VarNames returns a list of vars scoped for this vm instance
func (vm *VM) VarNames() []string {
	out := make([]string, 0, len(vm.vars))
	for k := range vm.vars {
		out = append(out, k)
	}
	return out
}

// SetVars sets the variables that are in scope for this vm instance
func (vm *VM) SetVars(vars map[string]ast.Variable) {
	vm.vars = vars
}

// RegisterFunc registers a function to the vm
func (vm *VM) RegisterFunc(name string, f ast.Function) error {
	if _, ok := vm.funcs[name]; ok {
		return errors.New("function registered: " + name)
	}
	vm.funcs[name] = f
	return nil
}

// Parse parses an input script and creates the ast root node
func (vm *VM) Parse(script string) (err error) {
	vm.a, err = hil.Parse(script)
	return err
}

// ParseEval is a helper function to set variables, parse the script and
// evaluate it
func (vm *VM) ParseEval(script string, vars scope.Variables) (result *hil.EvaluationResult, err error) {
	vm.SetVars(vars)

	err = vm.Parse(script)
	if err == nil {
		result, err = vm.Eval()
	}

	return result, err
}

// Eval evaluates a parsed script.  Parse must be called before a call to Eval
// can be made
func (vm *VM) Eval() (*hil.EvaluationResult, error) {
	if vm.a == nil {
		return nil, errors.New("must parse script")
	}

	// Build functions scope
	funcs := make(map[string]ast.Function)
	for k, v := range vm.funcs {
		funcs[k] = v
	}

	conf := &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			VarMap:  vm.vars,
			FuncMap: funcs,
		},
	}

	var (
		result *hil.EvaluationResult
		r, err = hil.Eval(vm.a, conf)
	)
	if err == nil {
		result = &r
	}

	return result, err
}
