package common

const (
	Bsd2 = "bsd-2"
	Bsd3 = "bsd-3"
	Mozilla = "mozilla"
	Mit = "mit"
	Isc = "isc"
	Apache = "apache"
	Unknown = ""
)

type LicenseInfo struct {
	Name string
	Type string
	Version string
	Path string
	Value int
}

type LicenseWeight []LicenseInfo

func (l LicenseWeight) Len() int {
	return len(l)
}

func (l LicenseWeight) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l LicenseWeight) Less(i, j int) bool {
	return l[i].Value < l[j].Value
}