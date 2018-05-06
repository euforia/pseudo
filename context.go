package pseudo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/euforia/pseudo/ewok"
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

	default:
		err = fmt.Errorf("unsupported Content-Type='%s'", ct)

	}

	return err
}

// LoadIndex loads an index from the URL. It reads the data, parses and indexes
// the data structure
func LoadIndex(uri *url.URL, opts ...IndexOptions) (*ewok.Ewok, error) {
	contentType, b, err := loadURI(uri)
	if err != nil {
		return nil, err
	}

	var (
		opt  = buildIndexOptions(contentType, opts...)
		conf = ewok.Config{TrimRoot: true, ScalarsOnly: true}
		ew   = ewok.New(conf)
		data = make(map[string]interface{})
	)

	err = unmarshal(opt.ContentType, b, &data)
	if err == nil {
		ew.Index(data)
	}

	return ew, err
}
