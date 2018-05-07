package scope

import "github.com/mitchellh/reflectwalk"

// ReflectBuilder is a new scoped variable builder
type ReflectBuilder struct {
	walker *walker
}

// NewReflectBuilder inits a new builder with the starting prefix and delimiter to use
// for key paths
func NewReflectBuilder(prefix, delim string) *ReflectBuilder {
	return &ReflectBuilder{
		walker: &walker{
			prefix: prefix,
			delim:  delim,
			vars:   make(Variables),
		},
	}
}

// Build builds scope variables from the object
func (rb *ReflectBuilder) Build(obj interface{}) error {
	return reflectwalk.Walk(obj, rb.walker)
}

// Variables returns the built scope variables
func (rb *ReflectBuilder) Variables() Variables {
	return rb.walker.vars
}
