package app

import (
	"github.com/alekns/yahe/internal/auth/app"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/alekns/yahe/internal/logger"
	"github.com/sirupsen/logrus"
)

type userAppImpl struct {
	settings *config.Settings
	events   app.DomainEvents
	logger   *logrus.Entry

	userRepository  app.UserRepositoryService
	passwordService app.PasswordService
}

func (ua *userAppImpl) setUserPassword(password string, user *app.User) error {
	if len(password) < 3 {
		return app.ErrorPasswordIsVeryBasic
	}

	passwordHashed, err := ua.passwordService.Create(
		ua.settings.Users.PasswordsSalt,
		password)
	if err != nil {
		return err
	}

	user.Password = string(passwordHashed)
	return nil
}

func (ua *userAppImpl) checkUserPassword(password string, user *app.User) (bool, error) {
	var hashedPasswordBytes = []byte(user.Password)

	passwordComparsion, err := ua.passwordService.Verify(
		ua.settings.Users.PasswordsSalt,
		password,
		hashedPasswordBytes)

	return passwordComparsion, err
}

// Create .
func (ua *userAppImpl) Create(tenantID app.TenantID, user *app.User) (*app.User, error) {
	logger := ua.logger.WithField("method", "Create").
		WithField("tenantId", tenantID).
		WithField("userLogin", user.Login)

	logger.Debug("Setup password")
	var err = ua.setUserPassword(user.Password, user)
	if err != nil {
		logger.Errorf("Failed to set user password: %v", err)
		return nil, err
	}

	logger.Debug("Save to repository")
	user, err = ua.userRepository.Save(user)
	if err != nil {
		logger.Errorf("Failed to save: %v", err)
		return nil, err
	}

	ua.events.UserCreated().Emit(tenantID, user)

	return user, nil
}

// GetUserBy .
func (ua *userAppImpl) GetUserBy(tenantID app.TenantID, login string, password string) (*app.User, error) {
	logger := ua.logger.WithField("method", "GetUserBy").
		WithField("tenantId", tenantID).
		WithField("userLogin", login)

	logger.Debug("Get from repository")

	user, err := ua.userRepository.GetByLogin(tenantID, login)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, app.ErrorNotFound
	}

	logger.Debug("Check password")

	result, err := ua.checkUserPassword(password, user)
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, app.ErrorPasswordMismatch
	}

	return user, nil
}

// ChangePassword .
func (ua *userAppImpl) ChangePassword(tenantID app.TenantID, userID app.UserID, oldPassword, newPassword string) (*app.User, error) {
	logger := ua.logger.WithField("method", "ChangePassword").
		WithField("tenantId", tenantID).
		WithField("userId", userID)

	logger.Debug("Get from repository")

	user, err := ua.userRepository.GetByIDs(tenantID, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, app.ErrorNotFound
	}

	if !user.IsActive {
		return nil, app.ErrorUserIsNotActive
	}

	logger.Debug("Check password")

	result, err := ua.checkUserPassword(oldPassword, user)
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, app.ErrorPasswordMismatch
	}

	err = ua.setUserPassword(newPassword, user)
	if err != nil {
		return nil, err
	}

	logger.Debug("Save user to repository")

	user, err = ua.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	ua.events.UserPasswordChanged().Emit(tenantID, user)

	return user, nil
}

// ResetPassword .
func (ua *userAppImpl) ResetPassword(tenantID app.TenantID, userID app.UserID, newPassword string) (*app.User, error) {
	logger := ua.logger.WithField("method", "ResetPassword").
		WithField("tenantId", tenantID).
		WithField("userId", userID)

	logger.Debug("Get from repository")

	user, err := ua.userRepository.GetByIDs(tenantID, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, app.ErrorNotFound
	}

	if user.IsActive {
		return nil, app.ErrorUserIsActive
	}

	err = ua.setUserPassword(newPassword, user)
	if err != nil {
		return nil, err
	}

	logger.Debug("Save user to repository")

	user, err = ua.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	ua.events.UserResetPassword().Emit(tenantID, user)

	return user, nil
}

// NewUserApp .
func NewUserApp(
	settings *config.Settings,
	events app.DomainEvents,
	userRepository app.UserRepositoryService,
	passwordService app.PasswordService) app.UserApp {

	return &userAppImpl{
		settings:        settings,
		logger:          logger.Get("UserApp"),
		events:          events,
		userRepository:  userRepository,
		passwordService: passwordService,
	}
}
