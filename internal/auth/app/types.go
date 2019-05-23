package app

type (
	//
	// Authentication
	//

	// TenantID .
	TenantID = string

	// UserID .
	UserID = string

	// JwtID .
	JwtID = string

	// JwtAttrs .
	JwtAttrs = map[string]interface{}

	// User .
	User struct {
		ID       UserID `json:"id"`
		Name     string `json:"name,omitempty"`
		Login    string `json:"login"`
		Password string `json:"password"`
		IsActive bool   `json:"is_active"`
	}

	// JwtEncoded .
	JwtEncoded = string

	// Jwt .
	Jwt struct {
		TenantID TenantID `json:"-"`
		JwtDecoded

		SessionID string `json:"session,omitempty"`
		ExpiredAt int    `json:"expired,omitempty"`
	}

	// JwtDecoded .
	JwtDecoded struct {
		ID     JwtID             `json:"id"`
		UserID UserID            `json:"user"`
		Scope  string            `json:"scope,omitempty"`
		Iat    int               `json:"iat"`
		Attrs  map[string]string `json:"attrs,omitempty"`
	}

	// JwtAppCreateRequest .
	JwtAppCreateRequest struct {
		Scope string
		TTL   int
		Attrs map[string]string
	}

	//
	// Autherization
	//
/*
	// PolicySubject .
	PolicySubject struct {
	    UserID string
	}

	PolicyTuple = []interface{}
	PolicyObject = map[string]string{}

	PolicyModel struct {
	    ID      string         `json:"id"`
	    Tuples  []PolicyTuple   ``
	    Objects []PolicyObject
	}

	PolicyResults struct {
	    Decisions []bool
	}

	RequestPolicyVerifySubject struct {
	    UserID string `json:"userId"`
	}

	RequestPolicyVerifyPolicy struct {
	    ModelID string `json:"modelId"`
	    Tuples  
	    Objects 
	}

	RequestPolicyVerify struct {
	    Subject  RequestPolicyVerifySubject  `json:"subject"`
	    Policies []RequestPolicyVerifyPolicy `json:"policies"`
	}
*/
)
