package auth

import (
	"encoding/json"
	"fmt"
	"gintest/constant"
	"gintest/dto"
	"gintest/util"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

type Permission struct {
	ListPermission map[string]string
}

func CreateToken(claim jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(constant.TOKEN_STRING))
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constant.TOKEN_STRING), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenData(r *http.Request) (*dto.ResponseToken, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constant.TOKEN_STRING), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	user := new(dto.ResponseToken)
	if !ok || !token.Valid {
		return nil, util.UnhandledError{ErrorMessage: "You Are Not Authorized"}
	}

	b, err := json.MarshalIndent(claims, "", " ")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("jwt : %v", string(b))

	err = json.Unmarshal([]byte(b), &user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("user : %v", user)

	//permission := claims["permission"].(Permission)
	//log.Printf("claim permission : %v", permission)
	////for _, data := range permission[0].ListPermission{
	////	permissionDto := new(dto.Permission)
	////	permissionDto.PermissionID = data.PermissionID
	////	permissionDto.PermissionName = data.PermissionName
	////	user.Permission.ListPermission = append(user.Permission.ListPermission, *permissionDto)
	////}
	//
	//user.UserID = claims["user_id"].(string)
	//user.Username = claims["username"].(string)
	//user.AppID = claims["app_id"].(string)
	//user.AppName = claims["app_name"].(string)
	//user.SubscriptionType = claims["subscription_type"].(string)
	//user.StatusDesign = claims["status_design"].(string)
	return user, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}