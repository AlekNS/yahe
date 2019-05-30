package app

// AuthModuleDependencyType .
type AuthModuleDependencyType int

const (
	// UserRepositoryServiceSymbol .
	UserRepositoryServiceSymbol AuthModuleDependencyType = iota
	// JwtRepositoryServiceSymbol .
	JwtRepositoryServiceSymbol
	// PasswordServiceSymbol .
	PasswordServiceSymbol
	// JwtServiceSymbol .
	JwtServiceSymbol

	// UserAppSymbol .
	UserAppSymbol
	// JwtAppSymbol .
	JwtAppSymbol

	// DomainEventsSymbol .
	DomainEventsSymbol
)
