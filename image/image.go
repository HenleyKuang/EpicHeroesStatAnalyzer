package image

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// FromBase64 converts a base64 string representation of an iage to a file object.
func FromBase64(imgAsBase64 string) (*os.File, error) {
	tempfile, err := ioutil.TempFile("", "ocrserver"+"-")
	if err != nil {
		return nil, err
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	if len(imgAsBase64) == 0 {
		return nil, fmt.Errorf("base64 string required")
	}
	imgAsBase64 = regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(imgAsBase64, "")
	b, err := base64.StdEncoding.DecodeString(imgAsBase64)
	if err != nil {
		return nil, err
	}
	tempfile.Write(b)
	return tempfile, nil
}
