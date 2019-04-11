package myRouter

const FINDUSERS = "/FindUsers"
const FINDUSERS_EI = "/FindUsers/{data}"
var RouteGetUserEntity = Route{
  "Searching Users In DB",
  "GET",
  "/FindUsers/{data}",
  userController.FindUsers,
}

const FINDYOURCONTACTS = "/FindYourContacts"
const FINDUSERS_EI = "/FindYourContacts/{data}"
var RouteGetUserEntity = Route{
  "Searching Your Contacts In DB",
  "GET",
  "/FindYourContacts/{data}",
  userController.FindYourContacts,
}

const FINDYOURGAMES = "/FindYourGames"
const FINDUSERS_EI = "/FindYourGames/{data}"
var RouteGetUserEntity = Route{
  "Searching Your Games In DB",
  "GET",
  "/FindYourGames/{data}",
  userController.FindYourGames,
}

const ADDCONTACT = "/AddContact"
var RouteGetUserEntity = Route{
  "Add New Contact In DB",
  "PUT",
  "/AddContact",
  userController.AddContact,
}
