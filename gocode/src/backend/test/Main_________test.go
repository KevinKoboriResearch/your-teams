package test

import (
	"backend/HyperText"
	"testing"
)

//---------------------------- USER ENTITY -----------------------------------//
//---------------------------- USER ENTITY -----------------------------------//
//---------------------------- USER ENTITY -----------------------------------//
//---------------------------- USER ENTITY -----------------------------------//
//---------------------------- USER ENTITY -----------------------------------//

//____________________________ INSERT ________________________________________//
func TestUserEntityInsertWRONGBODY(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertWRONGBODY)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-body"])
}

func TestUserEntityInsertSUCCESS(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertSUCCESS)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-insert"])
}

//____________________________ VERIFY ________________________________________//
func TestUserEntityVerifyWRONGBODY(t *testing.T) {
	resp := sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyWRONGBODY)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-body"])
}

func TestUserEntityVerifyWRONGPASSWORD(t *testing.T) {
	resp := sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyWRONGPASSWORD)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["error-database"])
}

func TestUserEntityVerifySUCCESS(t *testing.T) {
	auth = sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifySUCCESS)
	response := responseToString(auth)
	compareResults(t, response, HyperText.CustomResponses["success-login"])
}

//____________________________ UPDATE SINGLE _________________________________//
func TestUserEntityUpdateSingleWRONGBODY(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=UpdateSingle", UserEntityUpdateSingleWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-body"])
}

func TestUserEntityUpdateSingleWRONGEMAIL(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=UpdateSingle", UserEntityUpdateSingleWRONGEMAIL, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-validation"])
}

func TestUserEntityUpdateSingleSUCCESS(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=UpdateSingle", UserEntityUpdateSingleSUCCESS, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-update"])
}

//____________________________ UPDATE ________________________________________//
func TestUserEntityUpdateWRONGBODY(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=Update", UserEntityUpdateWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-body"])
}

func TestUserEntityUpdateWRONGIMAGE(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=Update", UserEntityUpdateWRONGIMAGE, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-validation"])
}

func TestUserEntityUpdateSUCCESS(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=Update", UserEntityUpdateSUCCESS, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-update"])
}

//____________________________ DISABLE _______________________________________//
func TestUserEntityDisableWRONGBODY(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=Disable", UserEntityDisableWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-body"])
}

func TestUserEntityDisableWRONGPASSWORD(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=Disable", UserEntityDisableWRONGPASSWORD, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["error-disable"])
}

func TestUserEntityDisableSUCCESS(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/TESTING/YourAccount&q=Disable", UserEntityDisableSUCCESS, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-disabled"])
}

//____________________________ GET UNIQUE ____________________________________//
func TestUserEntityGetUniqueNOTEXIST(t *testing.T) {
	resp := sendGet("http://localhost:8080/Search/User&q=NOTEXIST")
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["not-found-entity"])
}

func TestUserEntityGetUniqueEXIST(t *testing.T) {
	resp := sendGet("http://localhost:8080/Search/User&q=TESTING")
	response := responseToString(resp)
	compareResults(t, response, ResponseUserEntityGetUnique)
}

//____________________________ GET ALL WHILE _________________________________//
func TestUserEntityGetAllWhileSUCCESS(t *testing.T) {
	resp := sendGet("http://localhost:8080/Search/Users&p=email&v=testing@company.com")
	response := responseToString(resp)
	compareResults(t, response, ResponseUserEntityAllEnabled)
}

//____________________________ GET ALL ENABLED _______________________________//
func TestUserEntityGetAllEnabledSUCCESS(t *testing.T) {
	resp := sendGet("http://localhost:8080/Search/Users")
	response := responseToString(resp)
	compareResults(t, response, ResponseUserEntityAllEnabled)
}

//____________________________ GET ALL _______________________________________//
func TestUserEntityGetAllSUCCESS(t *testing.T) {
	resp := sendGet("http://localhost:8080/Search/UsersAll")
	response := responseToString(resp)
	compareResults(t, response, ResponseUserEntityGetAll)
}

//____________________________ DELETE ________________________________________//
func TestUserEntityDeleteSUCCESS(t *testing.T) {
	resp, _ := sendDelete("http://localhost:8080/TESTING/YourAccount&q=Delete", UserEntityDeleteSUCCESS, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-delete"])
}

//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//

//____________________________ INSERT ________________________________________//
func TestUserEntityInsertEXAMPLE1(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertEXAMPLE1)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-insert"])
}

func TestUserEntityInsertEXAMPLE2(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertEXAMPLE2)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-insert"])
}

//____________________________ VERIFY ________________________________________//
func TestUserEntityVerifyEXAMPLE1(t *testing.T) {
	auth = sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyEXAMPLE1)
	response := responseToString(auth)
	compareResults(t, response, HyperText.CustomResponses["success-login"])
}

//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//

//____________________________ INSERT ________________________________________//
func TestUserFriendRequestWRONGBODY(t *testing.T) {
	resp := sendPost("http://localhost:8080/EXAMPLE2&q=solicitation", APPJASON_UTF_8, UserFriendRequestWRONGBODY)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-body"])
}

func TestUserFriendRequestEXAMPLE1(t *testing.T) {
	resp := sendPost("http://localhost:8080/EXAMPLE2&q=solicitation", APPJASON_UTF_8, UserFriendRequestEXAMPLE1)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["friend-request"])
}

//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//

//____________________________ VERIFY ________________________________________//
func TestUserEntityVerifyEXAMPLE2(t *testing.T) {
	auth = sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyEXAMPLE2)
	response := responseToString(auth)
	compareResults(t, response, HyperText.CustomResponses["success-login"])
}

//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//
//---------------------------- USER FRIEND -----------------------------------//

//____________________________ INSERT ________________________________________//
func TestUserFriendAcceptEXAMPLE2(t *testing.T) {
	resp := sendPost("http://localhost:8080/EXAMPLE1&q=solicitation", APPJASON_UTF_8, UserFriendAcceptEXAMPLE2)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["friend-accpect"])
}

func TestUserFriendRequestALREADYEXIST(t *testing.T) {
	resp := sendPost("http://localhost:8080/EXAMPLE2&q=solicitation", APPJASON_UTF_8, UserFriendRequestALREADYEXIST)
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["error-database"])
}

//____________________________ DELETE ________________________________________//
func TestUserFriendDeleteWRONGBODY(t *testing.T) {
	resp, _ := sendDelete("http://localhost:8080/EXAMPLE2&q=undofriendship", UserFriendRequestWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["wrong-body"])
}

func TestUserFriendDeleteNOTEXIST(t *testing.T) {
	resp, _ := sendDelete("http://localhost:8080/ABCDEFGH&q=undofriendship", UserFriendRequestEXAMPLE1, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["not-found-entity"])
}

func TestUserFriendDeleteEXAMPLE1EXAMPLE2(t *testing.T) {
	resp, _ := sendDelete("http://localhost:8080/EXAMPLE2&q=undofriendship", UserFriendRequestEXAMPLE1, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-delete"])
}

//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//
//---------------------------- USER ENTITY - USER FRIEND ---------------------//

//____________________________ DELETE ________________________________________//
func TestUserEntityDeleteEXAMPLE1(t *testing.T) {
	resp, _ := sendDelete("http://localhost:8080/EXAMPLE1/YourAccount&q=Delete", UserEntityDeleteSUCCESS, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-delete"])
}

func TestUserEntityDeleteEXAMPLE2(t *testing.T) {
	resp, _ := sendDelete("http://localhost:8080/EXAMPLE2/YourAccount&q=Delete", UserEntityDeleteSUCCESS, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, HyperText.CustomResponses["success-delete"])
}
