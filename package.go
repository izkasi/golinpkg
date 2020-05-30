package dpkg

// DebianPackage respresents a single "Package" entry of dpkg status database.
type DebianPackage struct {
	Package       string
	Status        string
	Priority      string
	Section       string
	InstalledSize int
	Maintainer    string
	Architecture  string
	Version       string
	Depends       string
	Suggests      string
	Conffiles     []string
	Description   []string
}
