package auth

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

/*local cache of token, key is tokenId, value is token struct*/
var tokenMap = map[string]*Token{}

/*local cache of token with LoginId = user+host mapping, key is user+host, value is tokenId*/
var userTokenMap = map[string]string{}
var authMap = map[string]string{}
var rwLock = sync.RWMutex{}

const HourOfDay = 24

type Token struct {
	TokenId    string
	RemoteAddr string
	LoginId    string
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
	GetTokenById() (*Token, error)
	GetTokenByUser(user *User) (*Token, error)
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

	checkedToken, err := token.GetTokenById()

	if err == nil {
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

func (token *Token) GetTokenById() (*Token, error) {
	if token.TokenId == "" {
		log.Error("token id is empty")
		return nil, errors.New("no authorization for no auth token")
	}

	rwLock.RLock()
	checkedToken, isExists := tokenMap[token.TokenId]
	rwLock.RUnlock()

	// no local cache, turn to get from redis
	if !isExists {
		checkedToken = getValueObjFromRedis(token.TokenId)
		if checkedToken != nil {
			tokenMap[token.TokenId] = checkedToken
			userTokenMap[token.LoginId] = checkedToken.TokenId
			return checkedToken, nil
		}
		return nil, errors.New("token not exists")
	}
	return checkedToken, nil
}

func (token *Token) GetTokenByUser(user *User) (string, error) {
	if user.Username == "" {
		log.Error("username is empty")
		return "", errors.New("username is empty")
	}

	token.LoginId = strings.Join([]string{user.Username, token.RemoteAddr}, "-")
	rwLock.RLock()
	checkedTokenId, isExists := userTokenMap[token.LoginId]
	rwLock.RUnlock()
	token.TokenId = checkedTokenId
	token.User = user

	// no local cache, turn to get from redis
	if !isExists {
		return getValueStrFromRedis(token.LoginId)
	}
	return checkedTokenId, nil
}

func (token *Token) CreateNewToken(user *User) error {
	checkedToken, err := token.GetTokenByUser(user)

	if err == nil {
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
	userTokenMap[token.LoginId] = token.TokenId
	tokenByte, _ := json.Marshal(token)
	_ = setValueToRedis(token.TokenId, string(tokenByte), time.Hour*HourOfDay)
	_ = setValueToRedis(token.LoginId, token.TokenId, time.Hour*HourOfDay)
	log.Info("create new token")
	return nil
}

func (token *Token) DeleteToken() {
	defer rwLock.Unlock()

	rwLock.Lock()
	checkedToken, isExists := tokenMap[token.TokenId]
	if isExists {
		delete(userTokenMap, checkedToken.LoginId)
		_ = deleteKeyToRedis(checkedToken.LoginId)
	}
	delete(tokenMap, token.TokenId)
	_ = deleteKeyToRedis(token.TokenId)
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
		delete(userTokenMap, inspectToken.LoginId)
	}
}

func getValueObjFromRedis(key string) *Token {
	redisDb := config.GetRedisConnection()
	if redisDb == nil {
		log.Warn("there is no redis connection")
		return nil
	}
	ctx := context.Background()
	tokenJson, err := redisDb.Get(ctx, key).Result()
	if err != nil {
		log.Error("failed to query redis")
		return nil
	}
	token := &Token{}
	err = json.Unmarshal([]byte(tokenJson), &token)
	if err != nil {
		log.Error("failed to get token from redis")
		return nil
	}
	return token
}

func getValueStrFromRedis(key string) (string, error) {
	redisDb := config.GetRedisConnection()
	if redisDb == nil {
		log.Warn("there is no redis connection")
		return "", errors.New("there is no redis connection")
	}
	ctx := context.Background()
	return redisDb.Get(ctx, key).Result()
}

func setValueToRedis(key string, value string, expireTime time.Duration) error {
	redisDb := config.GetRedisConnection()
	if redisDb == nil {
		log.Warn("there is no redis connection")
		return errors.New("there is no redis connection")
	}
	ctx := context.Background()
	return redisDb.SetEX(ctx, key, value, expireTime).Err()
}

func deleteKeyToRedis(key string) error {
	redisDb := config.GetRedisConnection()
	if redisDb == nil {
		log.Warn("there is no redis connection")
		return errors.New("there is no redis connection")
	}
	ctx := context.Background()
	return redisDb.Del(ctx, key).Err()
}
