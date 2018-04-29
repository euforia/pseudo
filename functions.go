package pseudo

import (
	"errors"
	"strings"

	"github.com/hashicorp/hil/ast"
)

var (
	errInvalidFQDN = errors.New("invalid FQDN")
)

func replaceFunction() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString, ast.TypeString, ast.TypeString},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			s := inputs[0].(string)
			o := inputs[1].(string)
			n := inputs[2].(string)
			return strings.Replace(s, o, n, -1), nil
		},
	}
}

// returns the tld given a fqdn
func tldFunction() ast.Function {
	return ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			fqdn := inputs[0].(string)
			if len(fqdn) == 0 {
				return nil, errInvalidFQDN
			}

			hp := strings.Split(fqdn, ":")
			switch len(hp) {

			case 1, 2:
				parts := strings.Split(hp[0], ".")
				return parts[len(parts)-1], nil

			default:
				return nil, errInvalidFQDN

			}
		},
	}
}
