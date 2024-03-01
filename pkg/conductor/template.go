package cond

import (
	"io/ioutil"
	"os"
	"strings"
)

func openHTMLTemplate(templatePath string) (string, error) {

	var (
		file           *os.File
		err            error
		data           []byte
		templateString string
	)

	if file, err = os.Open(templatePath); err != nil {
		return "", err
	}
	defer file.Close()

	if data, err = ioutil.ReadAll(file); err != nil {
		return "", err
	}

	templateString = string(data)

	if !strings.Contains(templateString, "{{.VerificationCode}}") {
		return "", ErrInvalidTemplate
	}

	return templateString, nil
}
