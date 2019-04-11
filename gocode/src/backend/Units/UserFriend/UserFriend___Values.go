package UserFriend

import (
	"backend/Entities/UserEntity"
	"backend/HyperText"
	"gopkg.in/mgo.v2"
)

const (
	DOCNAME     = "user_friend"
	ID_USERNAME = "username"
)

var (
	collection *mgo.Collection
	controller = &UserFriendController{UserFriendRepository: UserFriendRepository{}}
	routes     = HyperText.Routes{{}}
	uf         UserFriend
	ufind      UserFriend
	ue         UserEntity.UsernameExist
)
