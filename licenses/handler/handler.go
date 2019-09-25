package handler

var HandlerChain = make(map[string]LicenseHandler)

type LicenseHandler interface {
	Parse(licenseInfo string) error
	Analyse(licenseInfo string) int
}

func InstallHandler(name string, h LicenseHandler) {
	HandlerChain[name] = h
}

func GetHandlerChain() (make(map[string]LicenseHandler)) {
	return HandlerChain
}

func Analyse(licneseInfo string) map[string]int {
	result := make(map[string]int)
	for name, licneseHander := range HandlerChain {
		count := licneseHander.Analyse(licneseInfo)
		result[name] = count
	}
	return result
}

func init() {

}