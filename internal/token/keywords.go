package token

const (
	Syntax  = "syntax"
	Package = "package"
	Import  = "import"
	Option  = "option"
	Service = "service"
	Rpc     = "rpc"
	Enum    = "enum"
	Message = "message"
)

// TODO: maybe we won't use it, think later
var setOfTokens = map[string]struct{}{
	Syntax: {},
	Package: {},
	Import: {},
	Option: {},
	Service: {},
	Rpc: {},
	Enum: {},
	Message: {},
}

func IsKeyword(kw string) bool {
	_, ok := setOfTokens[kw]
	return ok
}
