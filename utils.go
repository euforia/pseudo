package pseudo

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
)

func loadURI(uri *url.URL) (string, []byte, error) {

	var (
		contentType string
		data        []byte
		err         error
	)

	switch uri.Scheme {
	case "http", "https":
		contentType, data, err = loadHTTP(uri)

	default:
		// Load from file by default
		contentType, data, err = loadFile(uri.Path)

	}

	return contentType, data, err
}

func loadFile(filename string) (string, []byte, error) {
	data, err := ioutil.ReadFile(filename)

	return "", data, err
}

func loadHTTP(uri *url.URL) (contentType string, b []byte, err error) {
	var resp *http.Response

	resp, err = http.Get(uri.String())
	if err == nil {
		defer resp.Body.Close()

		contentType = resp.Header.Get("Content-Type")
		b, err = ioutil.ReadAll(resp.Body)
	}

	return
}

// ReadDirFiles reads all files in a given directory returning a map of filename
// to its contents
func ReadDirFiles(dirpath string) (map[string][]byte, error) {
	dirfiles, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	files := make(map[string][]byte)

	for _, f := range dirfiles {
		var b []byte
		b, err = ioutil.ReadFile(filepath.Join(dirpath, f.Name()))
		if err == nil {
			files[f.Name()] = b
			continue
		}
		break
	}

	return files, err
}
