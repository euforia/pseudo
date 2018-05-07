package pseudo

import (
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testURI = "https://api.github.com/repos/euforia/pseudo"
)

func Test_LoadURI_http_json(t *testing.T) {
	uri, _ := url.Parse(testURI)
	_, err := LoadVariables(uri)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_LoadURI_file_hcl(t *testing.T) {
	uri, err := url.Parse(testScopeVarsSpec)
	assert.Nil(t, err)

	opt := IndexOptions{ContentType: "pseudo"}
	vars, err := LoadVariables(uri, opt)
	assert.Nil(t, err)

	keys := []string{
		"platform.env.id",
		"platform.env.name",
		"app.name",
		"registry.container.ecr.address",
	}

	// vars := idx.Variables()
	for _, v := range keys {
		_, ok := vars[v]
		assert.True(t, ok)
	}
}

func Test_LoadURI_error(t *testing.T) {
	uri, _ := url.Parse("./does/not/exist")
	_, err := LoadVariables(uri)
	assert.NotNil(t, err)
}

func Test_unmarshal(t *testing.T) {
	err := unmarshal("", nil, nil)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "unsupported Content-Type"))
}

func Test_buildIndexOptions(t *testing.T) {
	opt := buildIndexOptions("")
	assert.Equal(t, "json", opt.ContentType)
}
