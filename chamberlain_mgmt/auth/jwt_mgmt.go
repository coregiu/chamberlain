package auth

import (
	"chamberlain_mgmt/log"
	"errors"
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

var tokenMap = map[string]Token{}
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
	authMap["/rest/users"] = "admin"
	authMap["/rest/users/count"] = "all"
	authMap["/rest/users/login"] = "none"
	authMap["/rest/users/logout"] = "all"

	authMap["/rest/inputs"] = "admin"
	authMap["/rest/inputs/count"] = "admin"
	authMap["/rest/inputs/statistic"] = "admin"
	authMap["/rest/inputs/statistic/month"] = "admin"
	authMap["/rest/inputs/statistic/type"] = "admin"

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

func checkOperationAuth(role string, checkedToken Token) (bool, error) {
	// none role was checked before
	if strings.Compare(role, "all") == 0 || strings.Compare(role, checkedToken.User.Role) == 0 {
		return true, nil
	} else {
		log.Error("current user don't have operation right")
		return false, errors.New("current user don't have operation right")
	}
}

func (token *Token) CreateNewToken(user *User) error {
	rwLock.RLock()
	checkedToken, isExists := userTokenMap[user.Username]
	rwLock.RUnlock()

	if isExists {
		token.TokenId = checkedToken
		log.Info("using exist token")
		return nil
	} else {
		rwLock.Lock()
		token.TokenId = uuid.NewString()
		token.User = user
		token.IssueTime = time.Now()
		token.ExpireTime = time.Now().Add(time.Hour * HourOfDay)
		token.Issuer = user.Username
		tokenMap[token.TokenId] = *token

		userTokenMap[user.Username] = token.TokenId
		rwLock.Unlock()
		log.Info("create new token")
		return nil
	}
}

func (token *Token) DeleteToken() {
	rwLock.Lock()
	checkedToken, isExists := tokenMap[token.TokenId]
	if isExists {
		delete(userTokenMap, checkedToken.User.Username)
	}
	delete(tokenMap, token.TokenId)
	rwLock.Unlock()
}

func inspectToken() {
	log.Info("begin to inspect token map")
	for tokenId := range tokenMap {
		inspectToken := tokenMap[tokenId]
		if time.Now().Before(inspectToken.ExpireTime) {
			continue
		}

		rwLock.Lock()
		delete(tokenMap, tokenId)
		delete(userTokenMap, inspectToken.User.Username)
		rwLock.Unlock()
	}
}
