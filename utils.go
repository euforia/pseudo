package pseudo

import (
	"io/ioutil"
	"path/filepath"
)

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
