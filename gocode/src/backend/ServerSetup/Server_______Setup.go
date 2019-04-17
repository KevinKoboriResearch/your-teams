package ServerSetup

import (
	"backend/Entities/UserEntity"
	"backend/Entities/SiteEntity"
	"backend/HyperText"
	"backend/Interface"
	"backend/Units/UserFriend"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"log"
)

const (
	SERVER_IP   = "localhost"
	SERVER_PORT = ":8080"
	SERVER_HOST = SERVER_IP + SERVER_PORT
)

func StartServer() {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nServerSetup.StartServer()\n")

	Interface.StartConectionDatabase()
	HyperText.StartValidator()
	StartValidatorUserEntity()
	StartValidatorInterfaceEntity()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: true,
	})
	router := CreateAllRoutes()
	handler := c.Handler(router)
	log.Print("\n		 					Starting Server...\n\n")
	if err := http.ListenAndServe(SERVER_HOST, handler); err != nil {
		panic(err)
	}
}

func CreateAllRoutes() (routes *mux.Router) {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nServerSetup.CreateAllRoutes()\n")
	
	siteEntityRoutes := SiteEntity.SiteEntityRoutes()
	userEntityRoutes := UserEntity.UserEntityRoutes()
	userEntityFriendRoutes := UserFriend.UserFriendRoutes()
	appRoutes := append(siteEntityRoutes, userEntityRoutes...)
	appRoutes = append(appRoutes, userEntityFriendRoutes...)
	routes = HyperText.NewRouter(appRoutes)
	return routes
}

func StartValidatorUserEntity() {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nServerSetup.StartValidatorUserEntity()\n")

	HyperText.Validate.RegisterValidation("username-used", UserEntity.ValidateUsernameUsed)
	HyperText.Validate.RegisterValidation("username-exist", UserEntity.ValidateUsernameExist)
	HyperText.Validate.RegisterValidation("username-length", UserEntity.ValidateUsernameLength)
	HyperText.Validate.RegisterValidation("password-length", UserEntity.ValidatePasswordLength)
	HyperText.Validate.RegisterValidation("email-used", UserEntity.ValidateEmailUsed)
}

func StartValidatorInterfaceEntity() {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nServerSetup.StartValidatorInterfaceEntity()\n")
	
	HyperText.Validate.RegisterValidation("name-used", Interface.ValidateNameUsed)
	HyperText.Validate.RegisterValidation("name-exist", Interface.ValidateNameExist)
}
