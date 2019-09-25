package main

import (
	"fmt"
	"github.com/surechen/go-licenses-toolkie/licenses/common"
	"os"
	"github.com/surechen/go-licenses-toolkie/syscmd"
	"github.com/surechen/go-licenses-toolkie/file"
	"github.com/surechen/go-licenses-toolkie/licenses"
	"strings"
)

func main() {
	currentPath, err := file.GetCurrentPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	goPath := os.Getenv("GOPATH")
	fmt.Println(goPath)
	modPath :=  goPath + "/pkg/mod/"
	fmt.Println(currentPath)
	path := "./"
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}
	err = os.Chdir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmdStr := `go list -m all`
	arrMod, err := syscmd.LinuxCmd(cmdStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(arrMod))
	arrModList := make([]string, len(arrMod))
	for _, modInfo := range arrMod {
		arrModItem := strings.Split(modInfo, " ")
		if len(arrModItem) != 2 {
			continue
			fmt.Println("modInfo wrong. " + modInfo)
		}
		arrName := strings.Split(arrModItem[0], "/")
		if len(arrName) >= 3 {
			arrName = arrName[len(arrName)-2:]
		}
		fileNameEndFix := ""
		for _, name := range arrName {
			fileNameEndFix += name + "-"
		}
		fileNameEndFix = fileNameEndFix[:len(fileNameEndFix)-1]
		arrModList = append(arrModList, arrModItem[0] + "@" + arrModItem[1])
		license := common.LicenseInfo{Name:fileNameEndFix, Path:modPath + "/" + arrModItem[0] + "@" + arrModItem[1], Version:arrModItem[1], Type:"default"}
		licenses.AddLicense(license)
	}

	err = os.Chdir(currentPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	mkdirPath := "default"
	if path != "./" {
		arrPath := strings.Split(path, "/")
		if len(arrPath) > 0 {
			mkdirPath = arrPath[len(arrPath) - 1]
		}
	}
	fmt.Println(mkdirPath)
	err = file.CreateDirectory(mkdirPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, modItem := range licenses.GetLicenses() {
		cmdStr := `ls ` + modItem.Path
		arrFile, err := syscmd.LinuxCmd(cmdStr)
		if err != nil {
			fmt.Println("ls cmd err: " + err.Error())
			continue
		}
		for _, fileName := range arrFile {
			if strings.ToLower(fileName) == licenses.LicenseMatchLow1 || strings.ToLower(fileName) == licenses.LicenseMatchLow2 {
				sourceFile := modItem.Path + "/" + fileName
				destName := mkdirPath + "/" + licenses.LicenseName + "-" + modItem.Name
				file.CopyFile(sourceFile, destName)
			}
			if strings.ToLower(fileName) == licenses.NoticeMatchLow1 || strings.ToLower(fileName) == licenses.NoticeMatchLow2 {
				sourceFile := modItem.Path + "/" + fileName
				destName := mkdirPath + "/" + licenses.NoticenameName + "-" + modItem.Name
				file.CopyFile(sourceFile, destName)
			}
		}
	}
}
