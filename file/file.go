package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateDirectory(direcName string) error {
	err := os.Mkdir(direcName, os.ModePerm)
	return err
}

func CopyFile(sourceName, destName string) error {
	fileInfo, err := ioutil.ReadFile(sourceName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(destName, fileInfo,  0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetCurrentPath() (string, error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return strings.Replace(path, "\\", "/", -1), err
}