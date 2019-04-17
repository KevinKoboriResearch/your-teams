package SiteEntity

import (
	"backend/HyperText"
	//"backend/Interface"
	//"encoding/json"
	//"github.com/gorilla/mux"
	"net/http"
	//"strings"
	"log"
)

//____________________________ Index _________________________________________//
func (c *SiteEntityController) Index(w http.ResponseWriter, r *http.Request) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\nSignUp")

	HyperText.HttpTemplateResponse(w, http.StatusOK, HyperText.TemplateResponses["index"], nil)
	return
}

//____________________________ SignUp ________________________________________//
func (c *SiteEntityController) SignUp(w http.ResponseWriter, r *http.Request) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\nSignUp")

	HyperText.HttpTemplateResponse(w, http.StatusOK, HyperText.TemplateResponses["sign-up"], nil)
	return
}