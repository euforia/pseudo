package pseudo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tldFunc(t *testing.T) {
	f := tldFunction()

	r, err := f.Callback([]interface{}{"dummy.com"})
	assert.Nil(t, err)
	assert.Equal(t, r.(string), "com")

	r, err = f.Callback([]interface{}{"dummy.com:1234"})
	assert.Nil(t, err)
	assert.Equal(t, r.(string), "com")

	r, err = f.Callback([]interface{}{"172.18.9.32:1234"})
	assert.Nil(t, err)
	assert.Equal(t, r.(string), "32")

	_, err = f.Callback([]interface{}{""})
	assert.Equal(t, errInvalidFQDN, err)

	_, err = f.Callback([]interface{}{"foo:bar:baz"})
	assert.Equal(t, errInvalidFQDN, err)

	_, err = f.Callback([]interface{}{""})
	assert.Equal(t, errInvalidFQDN, err)
}

func Test_lengthFunc(t *testing.T) {
	f := lengthFunction()
	c, err := f.Callback([]interface{}{"foo", "bar"})
	assert.Nil(t, err)
	assert.Equal(t, 2, c.(int))
}
