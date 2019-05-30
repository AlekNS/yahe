package webrest

// AuthModuleDependencyType .
type AuthModuleDependencyType int

const (
	// JwtControllerSymbol .
	JwtControllerSymbol AuthModuleDependencyType = iota

	// UserControllerSymbol .
	UserControllerSymbol
)
