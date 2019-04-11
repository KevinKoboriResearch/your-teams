package UserFriend

import (
	"backend/HyperText"
)

//__ ROUTES __________________________________________________________________//
func UserFriendRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Add Friend - User Friend",
			"POST",
			"/{username}&q=solicitation",
			controller.AddFriend,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Delete - User Friend",
			"DELETE",
			"/{username}&q=undofriendship",
			controller.Delete,
		},
	}
	return routes
}
