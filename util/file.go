package util

import (
	"io/ioutil"
	"path/filepath"
)

// ReadFile reads all file content
func ReadFile(fpath string) ([]byte, error) {
	fname, err := filepath.Abs(fpath)
	if err != nil {
		return nil, err
	}
	fdata, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return fdata, nil
}
