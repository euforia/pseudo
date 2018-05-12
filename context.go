package pseudo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/euforia/pseudo/scope"
	"github.com/hashicorp/hcl"
)

// IndexOptions are index loader options
type IndexOptions struct {
	// Content type of context data
	ContentType string
}

func buildIndexOptions(ct string, opt ...IndexOptions) IndexOptions {
	var o IndexOptions
	if len(opt) > 0 {
		o = opt[0]
	}

	if len(o.ContentType) == 0 {
		if len(ct) > 0 {
			// Set determined content type
			o.ContentType = ct
		} else {
			// Set default to json
			o.ContentType = "json"
		}
	}

	return o
}

func unmarshal(ct string, b []byte, data interface{}) error {
	var err error

	switch {

	case strings.Contains(ct, "json"):
		err = json.Unmarshal(b, data)

	case strings.Contains(ct, "hcl"):
		err = hcl.Unmarshal(b, data)

	case strings.Contains(ct, "pseudo"):

	default:
		err = fmt.Errorf("unsupported Content-Type='%s'", ct)

	}

	return err
}

// LoadVariables loads an index from the URL. It reads the data, parses and
// indexes the data structure
func LoadVariables(uri *url.URL, opts ...IndexOptions) (scope.Variables, error) {
	contentType, b, err := loadURI(uri)
	if err != nil {
		return nil, err
	}

	opt := buildIndexOptions(contentType, opts...)
	return LoadVariableBytes(opt.ContentType, b)
}

// LoadVariableBytes loads an index from the given bytes.
func LoadVariableBytes(contentType string, b []byte) (scope.Variables, error) {
	var (
		vars scope.Variables
		err  error
	)

	if contentType == "pseudo" {
		// custom hcl parser
		vars, err = scope.BuildHCLScopeVars(b)
	} else {

		builder := scope.NewReflectBuilder("", ".")
		v := make(map[string]interface{})

		err = unmarshal(contentType, b, &v)
		if err == nil {
			if err = builder.Build(v); err == nil {
				vars = builder.Variables()
			}
		}

	}

	return vars, err
}
