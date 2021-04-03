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
	authMap["users"] = "all"
	authMap["inputs"] = "admin"

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

	rwLock.RLock()
	checkedToken, isExists := tokenMap[token.TokenId]
	rwLock.RUnlock()

	if isExists {
		if time.Now().After(checkedToken.ExpireTime) {
			token.DeleteToken()
			log.Error("token has been expired")
			return false, errors.New("token has been expired")
		}

		return checkOperationAuth(operation, checkedToken)
	} else {
		return false, errors.New("token is invalid or expired")
	}
}

func checkOperationAuth(operation string, checkedToken Token) (bool, error) {
	operationRole, operationExists := authMap[operation]
	if operationExists {
		if strings.Compare(operationRole, "all") == 0 || strings.Compare(operationRole, checkedToken.User.Role) == 0 {
			return true, nil
		} else {
			log.Error("current user don't have operation right")
			return false, errors.New("current user don't have operation right")
		}
	} else {
		log.Error("current operation don't have role right")
		return false, errors.New("current operation don't have role right")
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
