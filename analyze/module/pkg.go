package module

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// A PackageID unique identifies a package.
type PackageID struct {
	Type     string
	Location string
	Name     string
	Revision string
}

func (id *PackageID) String() string {
	return id.Type + "+" + id.Name + "$" + id.Revision + "@" + id.Location
}

// ID returns an opaque identifier unique to this package.
func (id *PackageID) ID() string {
	h := sha256.New()
	h.Sum([]byte(strings.Join(
		[]string{
			"FOSSA-PackageID-Type: " + id.Type,
			"FOSSA-PackageID-Location: " + id.Location,
			"FOSSA-PackageID-Name: " + id.Name,
			"FOSSA-PackageID-Revision: " + id.Revision,
		},
		"\n",
	)))
	return hex.EncodeToString(h.Sum(nil))
}
