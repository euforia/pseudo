package scope

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/hil/ast"
	"github.com/mitchellh/reflectwalk"
)

// Walker walks a data structure constructing the context variables
type walker struct {
	// initial starting prefix
	prefix string
	// key path delimiter
	delim string

	// last visited key name
	lastkey string

	// Temp for holding current primitive slices.
	currslice []ast.Variable
	// This is true when primitive slices are being processed. This allows
	// being able to skip complex types
	processElems bool

	// All valid walked vars
	vars Variables
}

// Enter is required for reflectwalk interface in order to use Exit. It levels
// up on each complex structure
func (t *walker) Enter(l reflectwalk.Location) error {

	switch l {

	case reflectwalk.MapKey, reflectwalk.MapValue,
		reflectwalk.StructField, reflectwalk.SliceElem:

		t.prefix += t.lastkey

	}
	//log.Printf("Ent prefix=%s (%v)", t.prefix, l)
	return nil
}

// Exit trims the last part of the prefix for data structures only.  It levels
// down on each complex structure
func (t *walker) Exit(l reflectwalk.Location) error {

	switch l {

	case reflectwalk.MapKey, reflectwalk.MapValue,
		reflectwalk.StructField, reflectwalk.SliceElem:

		i := strings.LastIndex(t.prefix, t.delim)
		if i >= 0 {
			t.prefix = t.prefix[:i]
		}

	case reflectwalk.Slice:
		if t.processElems {
			// Build ast list variable and add to vars
			t.vars[t.prefix] = ast.Variable{
				Value: t.currslice, Type: ast.TypeList,
			}

			t.currslice = nil
			t.processElems = false
		}

	}
	//log.Printf("Ext prefix='%s' (%v)", t.prefix, l)
	return nil
}

// Slice initializes a slice to hold elements if the first element contains
// a primitive type.  It assumes all type are the same.  Non-primitive types
// are skipped
func (t *walker) Slice(v reflect.Value) error {
	length := v.Len()
	if length == 0 {
		// Even though empty add the variable for reference
		t.currslice = make([]ast.Variable, 0)
		t.processElems = true
		return nil
	}

	if isPrimitive(v.Index(0).Kind()) {
		t.currslice = make([]ast.Variable, length)
		t.processElems = true
	}

	return nil
}

func (t *walker) SliceElem(i int, v reflect.Value) error {
	t.lastkey = fmt.Sprintf(".%d", i)

	var err error

	if t.processElems {

		var avar ast.Variable
		avar, err = hilVarTypeFromValue(v)
		if err == nil {
			avar.Value = v.Interface()
			t.currslice[i] = avar
		}

	}

	return err
}

func (t *walker) Primitive(v reflect.Value) error {
	if !v.IsValid() {
		return nil
	}

	// Return if processing of slice primitive slice elements is enabled
	if t.processElems {
		return nil
	}

	iface := v.Interface()
	// We skip map keys.  If this is a map key the value will be the same as
	// the last key
	cmp := fmt.Sprintf("%v", iface)
	if t.lastkey[1:] == cmp {
		return nil
	}

	fmt.Println(t.prefix, v.Type(), v.Kind(), v)

	avar, err := hilVarTypeFromValue(v)
	if err == nil {
		avar.Value = iface
		t.vars[t.prefix] = avar
	}

	return err
}

// Struct is required for reflectwalk interface
func (t *walker) Struct(v reflect.Value) error {
	return nil
}

// StructField sets the prefix if the value is a nestable data structure and updates the index
func (t *walker) StructField(sf reflect.StructField, v reflect.Value) error {
	t.lastkey = "." + sf.Name

	return nil
}

// Map is required for reflectwalk interface in order to use MapElem
func (t *walker) Map(m reflect.Value) error {
	return nil
}

// MapElem converts a key to a string.  If the value is a data structure the
// prefix is also updated
func (t *walker) MapElem(m, k, v reflect.Value) error {
	var (
		kind  = k.Kind()
		iface = k.Interface()
	)

	if isNumber(kind) {
		t.lastkey = fmt.Sprintf(".%d", iface)
	} else {
		t.lastkey = "." + iface.(string)
	}

	return nil
}

// func (t *walker) Interface(reflect.Value) error { return nil }

func isPrimitive(kind reflect.Kind) bool {
	return (kind == reflect.String) || (kind >= reflect.Bool && kind <= reflect.Complex128)
}

func isNumber(kind reflect.Kind) bool {
	return (kind >= reflect.Int) && (kind <= reflect.Complex128)
}

func isInt(kind reflect.Kind) bool {
	return kind >= reflect.Int && kind < reflect.Uintptr
}

func isFloat(kind reflect.Kind) bool {
	return kind >= reflect.Float32 && kind <= reflect.Complex128
}

func hilVarTypeFromValue(v reflect.Value) (ast.Variable, error) {
	var (
		avar ast.Variable
		kind = v.Kind()
		err  error
	)
SWITCH:
	switch {
	case kind == reflect.Bool:
		avar.Type = ast.TypeBool

	case isInt(kind):
		avar.Type = ast.TypeInt

	case isFloat(kind):
		avar.Type = ast.TypeFloat

	case kind == reflect.String:
		avar.Type = ast.TypeString

	case kind == reflect.Interface:
		if v.IsNil() {
			avar.Type = ast.TypeString
			break
		}

		vi := v.Interface()
		kind = reflect.TypeOf(vi).Kind()
		goto SWITCH

	default:
		err = fmt.Errorf("unsupported type: %v", kind)
	}

	return avar, err
}
