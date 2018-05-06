//
// Package ewok implements a data structure indexer. The key is the concatenated
// field name with the delimiter based on supplied data strucuture
//
package ewok

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Config is the walker configuration
type Config struct {
	// Initial prefix
	Prefix string
	// Key delimiter. Default '.'
	Delimiter string
	// If true the root delimiter is excluded from the key
	TrimRoot bool
	// Field tag to use
	FieldTag string
	// Only index scalars
	ScalarsOnly bool
}

// Ewok is used to walk and index a data structure
type Ewok struct {
	// initial prefix
	prefix string
	// level prefix
	delim string
	// field tag identifier
	ftag string
	// scalars only
	sonly bool

	m map[string]reflect.Value
	// Index function
	idx func(v reflect.Value)
}

// New inits a new object index.  Prefix is the prefix for all keys
func New(conf Config) *Ewok {
	idx := &Ewok{
		prefix: conf.Prefix,
		delim:  conf.Delimiter,
		ftag:   conf.FieldTag,
		sonly:  conf.ScalarsOnly,
		m:      make(map[string]reflect.Value),
	}

	idx.init(conf.TrimRoot)

	return idx
}

func (w *Ewok) init(woRoot bool) {
	// Set the index function up front for optimization
	if woRoot {
		w.idx = w.indexWithoutRoot
	} else {
		w.idx = w.indexWithRoot
	}

	// Set the default delimiter if not provided
	if len(w.delim) == 0 {
		w.delim = "."
	}
}

// Index indexes the object such that it can be references using dot delimitation
// i.e Foo.Name
func (w *Ewok) Index(obj interface{}) {
	v := reflect.ValueOf(obj)
	w.recurse(v)
}

// Get returns the reflect value of the specified key
func (w *Ewok) Get(key string) (reflect.Value, bool) {
	val, ok := w.m[key]
	return val, ok
}

// Iter iterates over each key and value in lexagraphical order of the key
func (w *Ewok) Iter(f func(key string, value reflect.Value) bool) {
	keys := make([]string, 0, len(w.m))
	for k := range w.m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if !f(k, w.m[k]) {
			break
		}
	}
}

// func (w *Ewok) index(v reflect.Value) {
// 	vt := reflect.TypeOf(v.Interface())
// 	fmt.Println(vt)
// 	//if !v.IsNil() {
// 	// 	fmt.Println(reflect.TypeOf(v.Interface()).Kind())
// 	//}
// 	w.idx(v)
// }

func (w *Ewok) indexWithRoot(v reflect.Value) {
	w.m[w.prefix] = v
}

// strips the first character
func (w *Ewok) indexWithoutRoot(v reflect.Value) {
	w.m[w.prefix[1:]] = v
}

func (w *Ewok) upLevel(s string) {
	w.prefix += w.delim + s
}

func (w *Ewok) downLevel() {
	parts := strings.Split(w.prefix, w.delim)
	w.prefix = strings.Join(parts[:len(parts)-1], w.delim)
}

// The original code is from: https://gist.github.com/hvoecking/10772475
func (w *Ewok) recurse(v reflect.Value) {

	switch v.Kind() {
	// The first cases handle nested structures and translate them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		w.indexPtr(v)

	case reflect.Interface:
		w.indexInterface(v)

	case reflect.Struct:
		w.indexStruct(v)

	case reflect.Slice:
		w.indexSlice(v)

	case reflect.Map:
		w.indexMap(v)

	// And everything else will simply be taken from the original
	default:
		// Base literals
		//fmt.Println(original)
		w.idx(v)

	}

}

// If it is an interface (which is very similar to a pointer), do
// basically the same as for the pointer. Though a pointer is not the
// same as an interface so note that we have to call Elem() after
// creating a new object otherwise we would end up with an actual
// pointer
func (w *Ewok) indexInterface(v reflect.Value) {
	// Get rid of the wrapping interface
	value := v.Elem()
	w.recurse(value)
}

func (w *Ewok) indexPtr(v reflect.Value) {
	// To get the actual value of the original we have to call Elem()
	// At the same time this unwraps the pointer so we don't end up in
	// an infinite recursion.
	value := v.Elem()
	// Check if the pointer is nil
	if !value.IsValid() {
		return
	}

	w.recurse(value)
}

func (w *Ewok) indexSlice(v reflect.Value) {
	for i := 0; i < v.Len(); i++ {
		w.recurse(v.Index(i))
	}
}

func (w *Ewok) indexMap(v reflect.Value) {
	var (
		mapkeys = v.MapKeys()
		kind    = mapkeys[0].Kind()
		f       func(reflect.Value) string
	)

	// Determine the function to normalize the key to a string
	if kind >= reflect.Int && kind <= reflect.Complex128 {
		f = rvitos
	} else {
		f = rvstos
	}

	for _, key := range mapkeys {
		w.upLevel(f(key))

		value := v.MapIndex(key)

		w.ird(value)
	}
}

func (w *Ewok) indexStruct(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		// Struct field metadata
		sf := v.Type().Field(i)

		stag := sf.Tag.Get(w.ftag)
		if len(stag) == 0 {
			w.upLevel(sf.Name)
		} else {
			w.upLevel(stag)
		}

		field := v.Field(i)

		w.ird(field)
	}
}

// index, recurse, downLevel
func (w *Ewok) ird(v reflect.Value) {
	if !w.sonly {
		w.idx(v)
	}

	w.recurse(v)
	w.downLevel()
}

func rvitos(v reflect.Value) string {
	i := v.Interface()
	return fmt.Sprintf("%d", i)
}

func rvstos(v reflect.Value) string {
	return v.Interface().(string)
}
