package app

type AppDependencyType int

const (
	UserRepositoryServiceSymbol AppDependencyType = iota
	JwtRepositoryServiceSymbol
	PasswordServiceSymbol
	JwtServiceSymbol

	UserAppSymbol
	JwtAppSymbol

	DomainEventsSymbol
)
