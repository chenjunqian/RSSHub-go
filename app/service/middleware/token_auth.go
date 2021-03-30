package middleware

import (
	"encoding/json"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	response "rsshub/middleware"
	"strings"
)

type TokenModel struct {
	UserId         string `json:"userId"`
	UserName       string `json:"username"`
	NickName       string `json:"nickname"`
	Mobile         string `json:"mobile"`
	CreateDate     string `json:"createDate"`
	UpdateDateTime string `json:"updateDateTime"`
	Role           string `json:"role"`
	Token          string `json:"token"`
}

var privateKey = "rsshub-tm01-12-1"

func AuthToken(request *ghttp.Request) {
	authorization := request.Header.Get("Authorization")
	authorizationArray := strings.Split(authorization, "@@")
	if len(authorizationArray) < 2 {
		g.Log().Println("Token or uid is null")
		request.Response.WriteStatus(http.StatusUnauthorized)
		response.JsonExit(request, 0, "StatusUnauthorized", nil)
	}
	token := authorizationArray[0]
	uid := authorizationArray[1]
	if len(token) < 0 || len(uid) < 0 {
		g.Log().Println("AuthToken or uid is null")
		request.Response.WriteStatus(http.StatusUnauthorized)
		response.JsonExit(request, 0, "StatusUnauthorized", nil)
	}

	if tokenModel, err := ParseToken(token); err != nil {
		g.Log().Println("AuthToken invalid")
		request.Response.WriteStatus(http.StatusUnauthorized)
		response.JsonExit(request, 0, "StatusUnauthorized", nil)
	} else {
		if tokenModel.UserId != uid {
			g.Log().Println("token invalid tokenModel : ", tokenModel, " ,uid : ", uid)
			request.Response.WriteStatus(http.StatusUnauthorized)
			response.JsonExit(request, 0, "StatusUnauthorized", nil)
		}
	}
	request.Middleware.Next()
}

func ParseToken(tokenString string) (*TokenModel, error) {
	decodeToken, _ := gbase64.Decode([]byte(tokenString))
	decResult, err := gaes.Decrypt(decodeToken, []byte(privateKey))
	if err != nil {
		g.Log().Println("token decrypt error : ", tokenString)
		return nil, err
	}
	tokenModel := new(TokenModel)
	if err := json.Unmarshal(decResult, tokenModel); err != nil {
		g.Log().Println("token string decode to json error , token: ", decResult, " ,error : ", err)
		return nil, err
	}
	return tokenModel, nil
}

func CreateToken(tokenData TokenModel) (string, error) {
	if jsonToken, err := json.Marshal(tokenData); err != nil {
		g.Log().Println("decode token to json error : ", err)
		return "", err
	} else {
		if token, err := gaes.Encrypt(jsonToken, []byte(privateKey)); err != nil {
			g.Log().Println("aes encrypt string error: ", err)
			return "", err
		} else {
			encodeToken := gbase64.EncodeToString(token)
			return encodeToken, nil
		}
	}
}
