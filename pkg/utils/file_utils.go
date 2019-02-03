package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFromFile(fileName string) (jsonMsgs []string, err error) {
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		return
	}

	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return jsonMsgs, fmt.Errorf("error reading file %s. Error %+v", fileName, err)
	}

	jsonMsgs = strings.Split(string(fileBytes), "\n")

	return
}
