package app

//go:generate mockgen -source=auth_interfaces.go -package=app -destination=auth_interfaces_mocks.go

import "github.com/alekns/yahe/pkg/subscribs"

type (
	//
	// Domain services
	//

	// UserRepositoryService .
	UserRepositoryService interface {
		GetByLogin(tenantID TenantID, login string) (*User, error)
		GetByIDs(tenantID TenantID, userID UserID) (*User, error)
		Save(user *User) (*User, error)
	}

	// JwtRepositoryService .
	JwtRepositoryService interface {
		GetAll(tenantID TenantID, userID UserID) ([]Jwt, error)
		GetByIDs(tenantID TenantID, userID UserID, tokenID JwtID) (*Jwt, error)

		Save(token *Jwt) (*Jwt, error)

		RemoveAll(tenantID TenantID, userID UserID) error
		Remove(token *Jwt) error
	}

	// PasswordService .
	PasswordService interface {
		Create(salt []byte, password string) ([]byte, error)
		Verify(salt []byte, srcPassword string, hashedPassword []byte) (bool, error)
	}

	// JwtService .
	JwtService interface {
		SignAndEncode(attrs JwtAttrs) (string, error)
		VerifyAndDecode(tokenString string) (JwtAttrs, error)
	}

	//
	// Events
	//

	// DomainEvents .
	DomainEvents interface {
		UserCreated() subscribs.EventHandler
		UserStatusChanged() subscribs.EventHandler
		UserResetPassword() subscribs.EventHandler
		UserPasswordChanged() subscribs.EventHandler
	}

	//
	// Applications
	//

	// UserApp .
	UserApp interface {
		Create(tenantID TenantID, user *User) (*User, error)
		GetUserBy(tenantID TenantID, login string, password string) (*User, error)
		ChangePassword(tenantID TenantID, userID UserID, oldPassword, newPassword string) (*User, error)
		ResetPassword(tenantID TenantID, userID UserID, newPassword string) (*User, error)
	}

	// JwtApp .
	JwtApp interface {
		GetAllByUserID(tenantID TenantID, userID UserID) ([]JwtEncoded, error)
		Create(tenantID TenantID, login, password string, request *JwtAppCreateRequest) (JwtEncoded, error)
		Validate(tenantID TenantID, token JwtEncoded) (*JwtDecoded, error)
		Invalidate(tenantID TenantID, token JwtEncoded) error
		InvalidateAllTokens(tenantID TenantID, userID UserID) error
	}

	// AuthorizeApp .
	//AuthorizeApp interface {
	//
	//}
)
