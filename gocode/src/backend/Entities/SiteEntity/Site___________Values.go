package SiteEntity

import (
	"backend/HyperText"
)
/*
const (
	DOCNAME     = "site_entity"
	ID_MASTERUSERNAME = "masterusername"
)*/

var (
	controller = &SiteEntityController{SiteEntityRepository: SiteEntityRepository{}}
	routes     = HyperText.Routes{{}}
	/*ue   UserEntity
	uev  UserEntityVerify
	ueus UserEntityUpdateSingle
	ueu  UserEntityUpdate
	uep  UserEntityProtected
	ues  UserEntities*/
)
