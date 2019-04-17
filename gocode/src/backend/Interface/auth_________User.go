package Interface

import (
	"backend/HyperText"
	"crypto/rsa"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type MyClaimsType struct {
	jwt.StandardClaims
	User interface{} `json:"user"`
}

var routes = HyperText.Routes{{}}

func CreateAuthRoutes() HyperText.Routes {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nCreateAuthRoutes\n")

	routes := HyperText.Routes{}

	return routes
}

func SetAuthenticatedMiddleware(r func(http.ResponseWriter, *http.Request)) (n *negroni.Negroni) {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.SetAuthenticatedMiddleware()\n")

	n = negroni.New(negroni.HandlerFunc(ValidateToken), negroni.Wrap(http.HandlerFunc(r)))
	return
}

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.init()\n")

	privateBytes, err := ioutil.ReadFile("./Interface/auth_________private.rsa")

	if err != nil {
		log.Fatal("[ERROR] Can't read private file.")
	}

	publicBytes, err := ioutil.ReadFile("./Interface/auth_________public.rsa.pub")

	if err != nil {
		log.Fatal("[ERROR] Can't read public file.")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)

	if err != nil {
		log.Fatal("[ERROR] Can't parse privatekey.")
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)

	if err != nil {
		log.Fatal("[ERROR] Can't parse publickey.")
	}
}

func GenerateJWT(user interface{}) string {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	claims := MyClaimsType{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(30)).Unix(),
			Issuer:    "Your Teams",
		},
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	result, err := token.SignedString(privateKey)

	if err != nil {
		log.Fatal("[ERROR] Can't generate token.")
	}

	return result
}

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nValidateToken\n")

	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &MyClaimsType{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Print("[AUTH] Expired token. ", err)
		HyperText.HttpErrorResponse(w, http.StatusUnauthorized, HyperText.CustomResponses["unauthorized"])
		return
	}

	if !token.Valid {
		log.Print("[AUTH] Invalid token.")
		HyperText.HttpErrorResponse(w, http.StatusUnauthorized, HyperText.CustomResponses["unauthorized"])
		return
	}

	log.Print("[AUTH] Valid token.")
	next(w, r)
}
