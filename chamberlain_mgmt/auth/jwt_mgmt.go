package auth

import (
	"chamberlain_mgmt/log"
	"errors"
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

var tokenMap = map[string]*Token{}
var userTokenMap = map[string]string{}
var authMap = map[string]string{}
var rwLock = sync.RWMutex{}

const HourOfDay = 24

type Token struct {
	TokenId    string
	IssueTime  time.Time
	ExpireTime time.Time
	Issuer     string
	User       *User
}

func init() {
	authMap["/users"] = "admin"
	authMap["/users/count"] = "all"
	authMap["/users/token"] = "all"
	authMap["/users/password"] = "all"
	authMap["/users/login"] = "none"
	authMap["/users/logout"] = "all"
	authMap["/notebooks"] = "all"
	authMap["/notesummaryies"] = "all"
	authMap["/notesummaryies/content"] = "all"

	authMap["/inputs"] = "admin"
	authMap["/syslogs"] = "admin"
	authMap["/inputs/count"] = "admin"
	authMap["/inputs/statistic"] = "admin"
	authMap["/inputs/statistic/month"] = "admin"
	authMap["/inputs/statistic/type"] = "admin"
	authMap["/blogs"] = "admin"

	timer := time.NewTimer(time.Hour)
	go func() {
		for {
			timer.Reset(time.Hour)
			select {
			case <-timer.C:
				inspectToken()
			}
		}
	}()
}

type TokenFunc interface {
	CheckAuth(operation string) (bool, error)
	CreateNewToken(user *User) error
	DeleteToken()
	GetToken() (*Token, error)
}

func (token *Token) CheckAuth(operation string) (bool, error) {
	if operation == "" {
		log.Error("operation is nil")
		return false, errors.New("username is nil")
	}
	role := authMap[operation]
	if role == "" {
		log.Error("operation is has no operation right")
		return false, errors.New("operation is has no operation right")
	}
	log.Info("required role = %s", role)
	if strings.Compare(role, "none") == 0 {
		return true, nil
	}

	if token.TokenId == "" {
		log.Error("token id is empty")
		return false, errors.New("no authorization for no auth token")
	}

	rwLock.RLock()
	checkedToken, isExists := tokenMap[token.TokenId]
	rwLock.RUnlock()

	if isExists {
		if time.Now().After(checkedToken.ExpireTime) {
			token.DeleteToken()
			log.Error("token has been expired")
			return false, errors.New("token has been expired")
		}

		return checkOperationAuth(role, checkedToken)
	} else {
		return false, errors.New("token is invalid or expired")
	}
}

func checkOperationAuth(role string, checkedToken *Token) (bool, error) {
	// none role was checked before
	if strings.Compare(role, "all") == 0 || strings.Compare(role, checkedToken.User.Role) == 0 {
		return true, nil
	} else {
		log.Error("current user don't have operation right")
		return false, errors.New("current user don't have operation right")
	}
}

func (token *Token) GetToken() (*Token, error) {
	if token.TokenId == "" {
		log.Error("token id is nil")
		return token, errors.New("token id is nil")
	}
	mapToken := tokenMap[token.TokenId]
	if mapToken == nil {
		log.Error("no token exists")
		return token, errors.New("no token exists")
	}
	log.Info("Get token %v", mapToken.User.Username)
	return mapToken, nil
}

func (token *Token) CreateNewToken(user *User) error {
	rwLock.RLock()
	checkedToken, isExists := userTokenMap[user.Username]
	rwLock.RUnlock()

	if isExists {
		token.TokenId = checkedToken
		user.Password = ""
		token.User = user
		log.Info("using exist token")
		return nil
	} else {
		return storeToken(user, token)
	}
}

func storeToken(user *User, token *Token) error {
	defer rwLock.Unlock()

	rwLock.Lock()
	user.Password = ""
	token.TokenId = uuid.NewString()
	token.User = user
	token.IssueTime = time.Now()
	token.ExpireTime = time.Now().Add(time.Hour * HourOfDay)
	token.Issuer = user.Username
	tokenMap[token.TokenId] = token

	userTokenMap[user.Username] = token.TokenId
	log.Info("create new token")
	return nil
}

func (token *Token) DeleteToken() {
	defer rwLock.Unlock()

	rwLock.Lock()
	checkedToken, isExists := tokenMap[token.TokenId]
	if isExists {
		delete(userTokenMap, checkedToken.User.Username)
	}
	delete(tokenMap, token.TokenId)
}

func inspectToken() {
	defer rwLock.Unlock()
	log.Info("begin to inspect token map")
	rwLock.Lock()
	for tokenId := range tokenMap {
		inspectToken := tokenMap[tokenId]
		if time.Now().Before(inspectToken.ExpireTime) {
			continue
		}

		delete(tokenMap, tokenId)
		delete(userTokenMap, inspectToken.User.Username)
	}
}