package pseudo

import (
	"io/ioutil"
	"sort"
)

// Script is an evaluatable script
type Script struct {
	contents []byte
}

// NewScript loads a script from the given file returning an error on failure
func NewScript(filename string) (*Script, error) {
	b, err := ioutil.ReadFile(filename)
	if err == nil {
		return &Script{b}, nil
	}
	return nil, err
}

// Contents returns the script contents as a string
func (scr *Script) Contents() string {
	return string(scr.contents)
}

// Vars returns a list of variables needed to be passed in to the 'script'
func (scr *Script) Vars() []string {
	data := scr.contents
	vars := map[string]struct{}{}

	s := -1

	for i := range data {
		switch data[i] {

		case '$':
			if data[i-1] != '$' && data[i+1] == '{' {
				s = i
			}

		case '[', '(', '+', '>', '<':
			if s >= 0 {
				s = -1
			}

		case '}':
			if s >= 0 {
				v := string(data[s+2 : i])
				vars[v] = struct{}{}
				s = -1
			}

		}

	}

	out := make([]string, 0, len(vars))
	for k := range vars {
		out = append(out, k)
	}
	sort.Strings(out)

	return out
}
