package myRouter

const SIGNUP = "/SignUp"

var RouteSignUp = Route{
	"Inserting User Entity In DB",
	"POST",
	"/SignUp",
	userEntityController.InsertUserEntity,
}

const LOGIN = "/Login"

var RouteLogin = Route{
	"Inserting User Entity In DB",
	"POST",
	"/Login",
	userEntityController.LoginUserEntity,
}

const UPDATEUSERENTITY = "/UpdateUserEntity"

var RouteUpdateUserEntity = Route{
	"Updating User Entity In DB",
	"PUT",
	"/UpdateUserEntity",
	userEntityController.UpdateUserEntity,
}

const DISABLEUSERENTITY = "/DisableUserEntity"

var RouteDisableUserEntity = Route{
	"Disabling User From DB",
	"PUT",
	"/DisableUserEntity",
	userEntityController.DisableUserEntity,
}

const DELETEUSERENTITY = "/DeleteUserEntity"
const DELETEUSERENTITY_EI = "/DeleteUserEntity/{username}"

var RouteDeleteUserEntity = Route{
	"Deleting User From DB",
	"DELETE",
	"/DeleteUserEntity/{username}",
	userEntityController.DeleteUserEntity,
}
