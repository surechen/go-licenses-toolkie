package licenses

import (
	"errors"
	"fmt"
	_ "github.com/surechen/go-licenses-toolkie/licenses/bsd"
	"github.com/surechen/go-licenses-toolkie/licenses/common"
	"github.com/surechen/go-licenses-toolkie/licenses/handler"
	"sort"

	//"github.com/surechen/go-licenses-toolkie/licenses/handler"
	"io/ioutil"
)

const (
	LicenseName = "LICENSE"
	NoticenameName = "NOTICE"
	LicenseMatchLow1 = "license"
	LicenseMatchLow2 = "licenses"
	NoticeMatchLow1 = "notice"
	NoticeMatchLow2 = "notices"
)

var Licenses map[string]common.LicenseInfo

func AddLicense(license common.LicenseInfo) error {
	if _, ok := Licenses[license.Name]; ok {
		return errors.New("duplicated license file:" + license.Name)
	}
	Licenses[license.Name] = license
	return nil
}

func DeleteLicense(name string) (common.LicenseInfo, error) {
	if license, ok := Licenses[name]; !ok {
		return license, errors.New("no license file:" + license.Name)
	} else {
		return license, nil
	}
}

func AnalyseLicenseStrType(info string) (string, error) {
	result := handler.Analyse(info)
	var licenseweight common.LicenseWeight
	for name, value := range result {
		licenseweight = append(licenseweight, common.LicenseInfo{Name:name, Value:value})
	}
	sort.Sort(licenseweight)
	return "", nil
}

func AnalyseLicenseFileType(name string) (string, error) {
	fileInfo, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return common.Unknown, err
	}
	return AnalyseLicenseStrType(string(fileInfo))
}

func GetLicenses() (map[string]common.LicenseInfo) {
	return Licenses
}



func init(){
	Licenses = make(map[string]common.LicenseInfo, 10)
}