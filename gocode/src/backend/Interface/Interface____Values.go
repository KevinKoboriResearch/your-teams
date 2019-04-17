package Interface

import(
	"log"
)

type Name struct {
	Name string `json:"name"`
}

type Unit struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

const (
	ID_Name     = "name"
	ID_Username = "username"
)

var docname string

func PutDOCNAME(doc string) {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.PutDOCNAME()\n")

	docname = doc
}
