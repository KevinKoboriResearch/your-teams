package myHttp

//import ("net/http")

var responses = map[int]string{
	219: `{"code": 220, "message":"usuário logado com sucesso!"}`,
	220: `{"code": 220, "message":"usuário inserido com sucesso!"}`,
	240: `{"code": 240, "message":"usuário excluído com sucesso!"}`,
	407: `{"code": 407, "message":"This username already exists!"}`,
	470: `{"code": 476, "message":"Email already Exist!"}`,
	471: `{"code": 476, "message":"Wrong Username and Email!"}`,
	472: `{"code": 476, "message":"Wrong Username and Password!"}`,
	473: `{"code": 474, "message":"name needs to contain 5 caracters"}`,
	474: `{"code": 474, "message":"Password needs to contain 6 caracters"}`,
	475: `{"code": 475, "message":"Error, try it again later"}`,
	476: `{"code": 476, "message":"Wrong Email and Password!"}`,
	477: `{"code": 477, "message":"Invalid Email"}`,
	//	477: `{"code":` + string(http.StatusBadRequest) + `, "message":"Invalid Email"}`,
	478: `{"code": 478, "message":"Where is the body?"}`,
	479: `{"code": 479 , "message":"You Are Under 18"}`,
	480: `{"code": 480 , "message":"Data Was Entered Incorrectly"}`,
	507: `{"code": 507, "message":"Several error on DB server"}`,
}
