package dpkg

import (
	"bufio"
	"strconv"
	"strings"
)

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

func parseLine(line string) (string, string) {

	if len(line) == 0 {
		return "", ""
	}

	if line[0] == ' ' {
		return "", line
	}

	kv := strings.Split(line, ":")

	if len(kv[1]) == 0 {
		return kv[0], ""
	}

	return kv[0], kv[1][1:len(kv[1])]
}

func mapLine(key string, value string, pkg *DebianPackage) {

	switch key {
	case "Package":
		pkg.Package = value

	case "Status":
		pkg.Status = value

	case "Priority":
		pkg.Priority = value

	case "Section":
		pkg.Section = value

	case "Installed-Size":
		var err error
		pkg.InstalledSize, err = strconv.Atoi(value)

		if err != nil {
			pkg.InstalledSize = -1
		}

	case "Maintainer":
		pkg.Maintainer = value

	case "Architecture":
		pkg.Architecture = value

	case "Version":
		pkg.Version = value

	case "Depends":
		pkg.Depends = value

	case "Conffiles":
		if value != "" {
			pkg.Conffiles = append(pkg.Conffiles, value)
		}

	case "Description":
		if value == "." {
			pkg.Description = append(pkg.Description, "")
		} else {
			pkg.Description = append(pkg.Description, value)
		}

	}

}

//ParseStatus takes a *bufio.Scanner and returns a slice of type DebianPackage
func ParseStatus(scanner *bufio.Scanner) []DebianPackage {

	pkg := new(DebianPackage)
	pkgs := []DebianPackage{}
	previousKey := ""

	for scanner.Scan() {

		k, v := parseLine(scanner.Text())

		if len(k) == 0 && len(v) == 0 {
			// end of package
			pkgs = append(pkgs, *pkg)
			previousKey = ""

			// start a new package
			pkg = new(DebianPackage)

		} else if len(k) == 0 {
			// must be multiline add from past line
			mapLine(previousKey, v, pkg)

		} else {
			mapLine(k, v, pkg)
			previousKey = k
		}

	}

	return pkgs
}
