package services

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/alekns/yahe/internal/auth/app"
	"github.com/go-redis/redis"
)

const keyFieldDelimiter = ":"

type jwtRepositoryServiceRedis struct {
	keyPrefix string
	client    *redis.Client
}

func (jr *jwtRepositoryServiceRedis) formatKey(tenantID app.TenantID,
	userID app.UserID, tokenID app.JwtID) string {

	return jr.keyPrefix + tenantID + keyFieldDelimiter +
		userID + keyFieldDelimiter + tokenID
}

func (jr *jwtRepositoryServiceRedis) formatKeyByToken(token *app.Jwt) string {
	return jr.formatKey(token.TenantID, token.UserID, token.ID)
}

// GetAll .
func (jr *jwtRepositoryServiceRedis) GetAll(tenantID app.TenantID,
	userID app.UserID) ([]app.Jwt, error) {

	keys, err := jr.client.Keys(jr.formatKey(tenantID, userID, "*")).Result()
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, nil
	}

	tokensRaw, err := jr.client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	if len(tokensRaw) == 0 {
		return nil, nil
	}

	var result = make([]app.Jwt, 0, len(tokensRaw))
	for _, tokenRaw := range tokensRaw {
		if tokenRaw == nil {
			continue
		}
		token := app.Jwt{}
		err := json.Unmarshal([]byte(tokenRaw.(string)), &token)
		if err != nil {
			return nil, err
		}
		result = append(result, token)
	}

	return result, nil
}

// GetByIDs .
func (jr *jwtRepositoryServiceRedis) GetByIDs(tenantID app.TenantID,
	userID app.UserID, tokenID app.JwtID) (*app.Jwt, error) {

	tokenRaw, err := jr.client.Get(jr.formatKey(tenantID, userID, tokenID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, app.ErrorNotFound
		}
		return nil, err
	}

	if len(tokenRaw) == 0 {
		return nil, app.ErrorInternalStorageInconsistent
	}

	var token = &app.Jwt{}

	err = json.Unmarshal([]byte(tokenRaw), token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Save .
func (jr *jwtRepositoryServiceRedis) Save(token *app.Jwt) (*app.Jwt, error) {
	if len(token.ID) == 0 {
		token.ID = uuid.NewV4().String()
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	_, err = jr.client.SetNX(jr.formatKeyByToken(token),
		tokenJSON, time.Duration(token.ExpiredAt-token.Iat)*time.Second).Result()
	if err != nil {
		return nil, err
	}

	return token, nil
}

// RemoveAll .
func (jr *jwtRepositoryServiceRedis) RemoveAll(tenantID app.TenantID,
	userID app.UserID) error {

	keys, err := jr.client.Keys(jr.formatKey(tenantID, userID, "*")).Result()
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}

	_, err = jr.client.Del(keys...).Result()
	return err
}

// Remove .
func (jr *jwtRepositoryServiceRedis) Remove(token *app.Jwt) error {
	_, err := jr.client.Del(jr.formatKeyByToken(token)).Result()
	return err
}

// NewJwtRepositoryServiceRedis .
func NewJwtRepositoryServiceRedis(keyPrefix string, client *redis.Client) app.JwtRepositoryService {
	return &jwtRepositoryServiceRedis{
		keyPrefix: keyPrefix + "jwt.",
		client:    client,
	}
}
