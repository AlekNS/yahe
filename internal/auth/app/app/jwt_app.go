package app

import (
	"time"

	"github.com/alekns/yahe/internal/auth/app"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/alekns/yahe/internal/logger"
	"github.com/sirupsen/logrus"
)

type jwtAppImpl struct {
	settings *config.Settings
	logger   *logrus.Entry

	userApp       app.UserApp
	jwtRepository app.JwtRepositoryService
	jwtService    app.JwtService
}

func jwtEntityToJwtAttrs(token *app.Jwt) app.JwtAttrs {
	return app.JwtAttrs{
		"id":    token.ID,
		"user":  token.UserID,
		"scope": token.Scope,
		"iat":   token.Iat,
		"attrs": token.Attrs,
	}
}

// GetAllByUserID .
func (ja *jwtAppImpl) GetAllByUserID(tenantID app.TenantID, userID app.UserID) ([]app.JwtEncoded, error) {
	logger := ja.logger.WithField("method", "GetAllByUserID").
		WithField("tenantId", tenantID).
		WithField("userId", userID)

	var tokens = make([]app.JwtEncoded, 0, 0)

	logger.Debug("GetAll from repository")

	tokenEntities, err := ja.jwtRepository.GetAll(tenantID, userID)
	if err != nil {
		return nil, err
	}
	if len(tokenEntities) == 0 {
		return tokens, nil
	}

	tokens = make([]app.JwtEncoded, 0, len(tokenEntities))
	for _, tokenEntity := range tokenEntities {
		// Cache (store in redis)?
		tokenEncoded, err := ja.jwtService.SignAndEncode(jwtEntityToJwtAttrs(&tokenEntity))
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tokenEncoded)
	}

	return tokens, nil
}

// Create .
func (ja *jwtAppImpl) Create(tenantID app.TenantID,
	login, password string, request *app.JwtAppCreateRequest) (app.JwtEncoded, error) {

	logger := ja.logger.WithField("method", "Create").
		WithField("tenantId", tenantID).
		WithField("userLogin", login)

	ttl := request.TTL
	if ttl == 0 {
		ttl = ja.settings.Jwt.DefaultExpiration
	}

	logger.Debug("Get user by login and password")

	user, err := ja.userApp.GetUserBy(tenantID, login, password)
	if err != nil {
		return "", err
	}
	if !user.IsActive {
		return "", app.ErrorUserIsNotActive
	}

	expiredAt := int(time.Now().Unix()) + ttl

	token := &app.Jwt{
		TenantID:  tenantID,
		ExpiredAt: expiredAt,
		JwtDecoded: app.JwtDecoded{
			UserID: user.ID,
			Scope:  request.Scope,
			Iat:    int(time.Now().Unix()),
			Attrs:  request.Attrs,
		},
	}

	logger.Debug("Store to jwt repository")

	_, err = ja.jwtRepository.Save(token)
	if err != nil {
		return "", err
	}

	tokenEncoded, err := ja.jwtService.SignAndEncode(jwtEntityToJwtAttrs(token))
	if err != nil {
		return "", err
	}

	return tokenEncoded, nil
}

// Validate .
func (ja *jwtAppImpl) Validate(tenantID app.TenantID, token app.JwtEncoded) (*app.JwtDecoded, error) {
	logger := ja.logger.WithField("method", "Validate").
		WithField("tenantId", tenantID)

	logger.Debug("Verify and decode token")

	tokenAttrs, err := ja.jwtService.VerifyAndDecode(token)
	if err != nil {
		return nil, app.ErrorInvalidJwtToken
	}

	logger.Debugf("Get by id: %v", tokenAttrs)

	tokenEntity, err := ja.jwtRepository.GetByIDs(tenantID,
		tokenAttrs["user"].(string), tokenAttrs["id"].(string))
	if err != nil {
		return nil, err
	}

	return &tokenEntity.JwtDecoded, nil
}

// Invalidate .
func (ja *jwtAppImpl) Invalidate(tenantID app.TenantID, token app.JwtEncoded) error {
	logger := ja.logger.WithField("method", "Invalidate").
		WithField("tenantId", tenantID)

	logger.Debug("Verify and decode token")

	tokenAttrs, err := ja.jwtService.VerifyAndDecode(token)
	if err != nil {
		return app.ErrorInvalidJwtToken
	}

	logger.Debugf("Get by id: %v", tokenAttrs)

	tokenEntity, err := ja.jwtRepository.GetByIDs(tenantID,
		tokenAttrs["user"].(string), tokenAttrs["id"].(string))
	if err != nil {
		return err
	}
	if tokenEntity == nil {
		return nil
	}

	tokenEntity.TenantID = tenantID

	return ja.jwtRepository.Remove(tokenEntity)
}

// InvalidateAllTokens .
func (ja *jwtAppImpl) InvalidateAllTokens(tenantID app.TenantID, userID app.UserID) error {
	logger := ja.logger.WithField("method", "InvalidateAllTokens").
		WithField("tenantId", tenantID).
		WithField("userId", userID)

	logger.Debug("Remove all tokens")

	return ja.jwtRepository.RemoveAll(tenantID, userID)
}

// NewJwtApp .
func NewJwtApp(settings *config.Settings,
	userApp app.UserApp,
	jwtRepository app.JwtRepositoryService,
	jwtService app.JwtService) app.JwtApp {

	return &jwtAppImpl{
		settings:      settings,
		logger:        logger.Get("UserApp"),
		userApp:       userApp,
		jwtRepository: jwtRepository,
		jwtService:    jwtService,
	}
}
