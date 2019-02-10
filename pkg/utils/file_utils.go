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

func WriteDataToFile(fileName string, data []string) (err error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("failed to open file %s.Error %+v", fileName, err)
	}

	if _, err = f.WriteString(strings.Join(data, "\n")); err != nil {
		return fmt.Errorf("failed to write data to file %s. Error %+v", fileName, err)
	}

	f.WriteString("\n")
	return
}
