package SiteEntity

import (
	"backend/HyperText"
	"log"
)

//__ ROUTES __________________________________________________________________//
func SiteEntityRoutes() HyperText.Routes {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nEntities.Site.SiteRoutes()\n")

	routes := HyperText.Routes{
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"Index - Site",
			"GET",
			"/",
			controller.Index,
		},
		HyperText.Route{ //___________________ FUNCIONANDO _______________________//
			"SignUp - Site",
			"GET",
			"/SignUp",
			controller.SignUp,
		},
	}
	return routes
}
