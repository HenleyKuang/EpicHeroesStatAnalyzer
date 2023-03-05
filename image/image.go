package image

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func FromBase64(imgAsBase64 string) (*os.File, error) {
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return nil, err
	}
	tempfile, err := ioutil.TempFile("", "ocrserver"+"-")
	if err != nil {
		return nil, err
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	if len(body.Base64) == 0 {
		return nil, fmt.Errorf("base64 string required")
	}
	body.Base64 = regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(imgAsBase64, "")
	b, err := base64.StdEncoding.DecodeString(imgAsBase64)
	if err != nil {
		return nil, err
	}
	tempfile.Write(b)
	return tempfile, nil
}
