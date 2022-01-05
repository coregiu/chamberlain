package auth

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestToken_Should_Successfully_When_GetToken(t *testing.T) {
	token := new(Token)
	user := new(User)
	user.Username = "coregiu"
	user.Password = "123456"
	user.Role = "admin"
	err := token.CreateNewToken(user)
	if err != nil {
		t.Error(err.Error())
	}

	newToken, getErr := token.GetToken()
	if getErr != nil {
		t.Error(getErr.Error())
	}
	jsonByte, _ := json.Marshal(newToken)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)
	if strings.Compare(jsonStr, "admin") == -1 {
		t.Error("failed go get user by token")
	}
}

func TestToken_Should_Successfully_When_CreateToken(t *testing.T) {
	token := new(Token)
	user := new(User)
	user.Username = "coregiu"
	user.Password = "123456"
	user.Role = "admin"
	err := token.CreateNewToken(user)
	if err != nil {
		t.Error(err.Error())
	}

	newToken, getErr := token.GetToken()
	if getErr != nil {
		t.Error(getErr.Error())
	}
	jsonByte, _ := json.Marshal(newToken)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)
	if strings.Compare(jsonStr, "admin") == -1 {
		t.Error("failed go get user by token")
	}
}