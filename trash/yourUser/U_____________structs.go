package yourUser

type UserController struct {
	 UserRepository UserRepository
}

type UserRepository struct{
}

type User struct {
	Username     string 		  `json:"username"    validate:"required"`
	Disable			 bool					`json:"disable"		  validate:"required"`
	Status       bool 				`json:"status"		  validate:"required"`
}

type Users []User
/*
type UserContact {
	Username  string `json:"username"    validate:"required"`
	Contact  string `json:"contact"    validate:"required"`
}

type UserContacts []Contact

type UserGame {
	Username  string `json:"username"    validate:"required"`
	GameName  string `json:"gamename"    validate:"required"`
  GameRank  string `json:"gamerank"    validate:"required"`
	GameLevel int 	 `json:"gamelevel"   validate:"required,level-between"`
	GameStars int 	 `json:"gamestars"		validate:"required,starts-between"`
}

type UserGames []UserGame
*/
/*
type User struct {
	Username     string 		  `json:"username"    validate:"required"`
	UserContacts UserContacts `json:"usercontacts,omitempty"`
	UserGames    UserGames    `json:"usergames,omitempty"`
  UserTeams    UserTeams    `json:"userteams,omitempty"`
//	UserGraphics  UserGraphics  `json:"usergrafics,omitempty"`
//	UserChats  UserChats  `json:"userchats,omitempty"`
	Status       bool 				`json:"status"		  validate:"required"`
}
*/
/*
type UserGraphic {
	GameName  string 			`json:"gamename"    validate:"required"`
  GameRank  string 			`json:"gamerank"    validate:"required"`
	GameLevel int 				`json:"gamelevel"   validate:"required,level-between"`
	GameStars int 				`json:"gamestars"		validate:"required,starts-between"`
}

type UserGraphics []UserGraphic

type UserChat {
	GameName  string 			`json:"gamename"    validate:"required"`
  GameRank  string 			`json:"gamerank"    validate:"required"`
	GameLevel int 				`json:"gamelevel"   validate:"required,level-between"`
	GameStars int 				`json:"gamestars"		validate:"required,starts-between"`
}

type UserChats []UserChats
*/
