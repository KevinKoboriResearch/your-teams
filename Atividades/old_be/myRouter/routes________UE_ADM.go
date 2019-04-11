package myRouter

const GETUSERENTITIES = "/GetUserEntities"

var RouteGetUserEntities = Route{
	"Geting All User Entitites From DB",
	"GET",
	"/admin/GetUserEntities",
	userEntityController.GetUserEntities,
}

const GETUSERENTITY = "/GetUserEntity"
const GETUSERENTITY_EI = "/GetUserEntity/{username}"

var RouteGetUserEntity = Route{
	"Searching User Entity In DB",
	"GET",
	"/GetUserEntity/{username}",
	userEntityController.GetUserEntity,
}
