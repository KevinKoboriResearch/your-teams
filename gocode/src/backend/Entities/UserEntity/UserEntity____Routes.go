package UserEntity

import (
	"backend/HyperText"
)

//__ ROUTES __________________________________________________________________//
func UserEntityRoutes() HyperText.Routes {
	routes := HyperText.Routes{
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"SignUp - User Entity",
			"POST",
			"/SignUp",
			controller.SignUp,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Login - User Entity",
			"POST",
			"/Login",
			controller.Login,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Update Single - User Entity",
			"PUT",
			"/{username}/YourAccount&q=UpdateSingle",
			controller.UpdateSingle,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Update - User Entity",
			"PUT",
			"/{username}/YourAccount&q=Update",
			controller.Update,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Disable - User Entity",
			"PUT",
			"/{username}/YourAccount&q=Disable",
			controller.Disable,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get Unique - User Entity",
			"GET",
			"/Search/User&q={username}",
			controller.GetUnique,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get All While - User Entities",
			"GET",
			"/Search/Users&p={position}&v={value}",
			controller.GetAllEnabledWhile,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get All Enabled - User Entities",
			"GET",
			"/Search/Users",
			controller.GetAllEnabled,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Get All - User Entities",
			"GET",
			"/Search/UsersAll",
			controller.GetAll,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Delete - User Entity",
			"DELETE",
			"/{username}/YourAccount&q=Delete",
			controller.Delete,
		},
	}
	return routes
}
