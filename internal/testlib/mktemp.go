// Package testlib contains test helpers for goreleaser tests.
package testlib

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
)

// Mktmp creates a new tempdir, cd into it and provides a back function that
// cd into the previous directory.
func Mktmp(t *testing.T) (folder string, back func()) {
	folder, err := ioutil.TempDir("", "goreleasertest")
	assert.NoError(t, err)
	current, err := os.Getwd()
	assert.NoError(t, err)
	assert.NoError(t, os.Chdir(folder))
	return folder, func() {
		assert.NoError(t, os.Chdir(current))
	}
}

func CreateTempFile(content string, prefix string) (*os.File, error) {
	file, err := ioutil.TempFile(os.TempDir(), prefix)
	if err != nil {
		return nil, errors.Wrap(err, "error creating temp file")
	}
	_, err = file.WriteString(content)
	if err != nil {
		return nil, errors.Wrap(err, "error writing to file")
	}
	return file, nil
}
