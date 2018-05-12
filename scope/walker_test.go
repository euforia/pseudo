package scope

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/hil/ast"
	"github.com/mitchellh/reflectwalk"
	"github.com/stretchr/testify/assert"
)

type testStruct2 struct {
	Labels []string
}

type testStruct1 struct {
	Name      string
	Precision float64
}

type testStruct struct {
	Disabled    bool
	Map         map[int64]string
	NestedMap   map[string]interface{}
	Desc        *string
	NullPointer *testStruct1
	Pointer     *testStruct1
	Struct      testStruct2
	Iface       interface{}
	Slices      []*testStruct1
}

func testData() *testStruct {
	ii := 10
	data := &testStruct{
		Map: map[int64]string{
			1: "v1",
			2: "v2",
		},
		NestedMap: map[string]interface{}{
			"foo": map[int]interface{}{
				3: testStruct1{Precision: 3.1415},
				4: &testStruct2{},
			},
			"bar": map[string]interface{}{
				"five": &testStruct2{},
				"six":  &ii,
			},
		},
		Pointer: &testStruct1{},
		Iface:   &testStruct2{Labels: []string{"label1"}},
		Slices: []*testStruct1{
			&testStruct1{Name: "slice-0"},
			&testStruct1{Name: "slice-1"},
		},
	}
	return data
}

var expectedKeys = []string{
	"Disabled",
	"Map.1",
	"Map.2",
	"NestedMap.foo.3.Name",
	"NestedMap.foo.3.Precision",
	"NestedMap.foo.4.Labels",
	"NestedMap.bar.five.Labels",
	"NestedMap.bar.six",
	"Pointer.Name",
	"Pointer.Precision",
	"Struct.Labels",
	"Iface.Labels",
	"Slices.0.Name",
	"Slices.0.Precision",
	"Slices.1.Name",
	"Slices.1.Precision",
}

func Test_reflectWalker(t *testing.T) {

	data := testData()
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("%s\n", b)

	walker := &walker{
		delim: ".",
		vars:  make(map[string]ast.Variable),
	}
	err := reflectwalk.Walk(data, walker)
	assert.Nil(t, err)

	for k, v := range walker.vars {
		t.Logf("%s (%v)", k, v.Type)
	}

	vars := walker.vars
	for _, v := range expectedKeys {
		_, ok := vars[v]
		if !ok {
			t.Errorf("key missing: %s", v)
		}
	}

}

func Test_hilVarTypeFromValue_error(t *testing.T) {
	data := testData()
	v := reflect.ValueOf(data)
	_, err := hilVarTypeFromValue(v)
	assert.NotNil(t, err)
}
