package services

import (
	"github.com/alekns/yahe/internal/auth/app"
	"github.com/alekns/yahe/pkg/subscribs"
)

// domainEventsImpl implementation of Events
type domainEventsImpl struct {
	userCreated       subscribs.EventHandler
	userStatusChanged subscribs.EventHandler

	userResetPassword   subscribs.EventHandler
	userPasswordChanged subscribs.EventHandler
}

// UserCreated .
func (e *domainEventsImpl) UserCreated() subscribs.EventHandler {
	return e.userCreated
}

// UserStatusChanged .
func (e *domainEventsImpl) UserStatusChanged() subscribs.EventHandler {
	return e.userStatusChanged
}

// UserResetPassword .
func (e *domainEventsImpl) UserResetPassword() subscribs.EventHandler {
	return e.userResetPassword
}

// UserPasswordChanged .
func (e *domainEventsImpl) UserPasswordChanged() subscribs.EventHandler {
	return e.userPasswordChanged
}

// NewSyncEventsImpl creates sync events.
func NewSyncEventsImpl() app.DomainEvents {
	return &domainEventsImpl{
		userCreated:       subscribs.NewSyncEventHandler(),
		userStatusChanged: subscribs.NewSyncEventHandler(),

		userResetPassword:   subscribs.NewSyncEventHandler(),
		userPasswordChanged: subscribs.NewSyncEventHandler(),
	}
}
