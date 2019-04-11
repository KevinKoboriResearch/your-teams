package myRouter

import (
	"lfkk/be/yourUserEntity"
)

var userEntityController = &yourUserEntity.UserEntityController{UserEntityRepository: yourUserEntity.UserEntityRepository{}}

//var userController = &yourUserEntity.UserEntityController{UserEntityRepository: yourUserEntity.UserEntityRepository{}}
