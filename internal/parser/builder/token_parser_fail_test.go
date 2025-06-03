package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type failTest struct {
	name  string
	input string
}

func runFailTokenTest(t *testing.T, parse parseXToken, tests []failTest) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := newTestTokenParser(test.input)
			parser.ExtractKeyword()
			res, err := parse(parser)

			assert.Nil(t, res)
			assert.Error(t, err)
		})
	}
}

func Test_SyntaxToken_FAIL(t *testing.T) {
	tests := []failTest{
		{
			name:  "Missing equals",
			input: "syntax \"proto3\";",
		}, {
			name:  "Missing semicolon",
			input: "syntax = \"proto3\"",
		}, {
			name:  "Missing matching quote",
			input: "syntax = \"proto3;",
		}, {
			name:  "Missing first quote",
			input: "syntax = proto3\";",
		}, {
			name:  "Quote on next line",
			input: "syntax = \"proto3\n\";",
		}, {
			name:  "Missing all",
			input: "syntax",
		},
	}
	runFailTokenTest(t, (*TokenParser).ParseSyntaxToken, tests)
}

func Test_PackageToken_FAIL(t *testing.T) {
	tests := []failTest{
		{
			name:  "Missing semicolon",
			input: "package example",
		}, {
			name:  "Wrong semicolon",
			input: "package; example",
		}, {
			name:  "No spaces",
			input: "packageexample;",
		}, {
			name:  "Wrong symbol between",
			input: "package & example;",
		}, {
			name:  "Missing all",
			input: "package",
		},
	}
	runFailTokenTest(t, (*TokenParser).ParseSyntaxToken, tests)
}

func Test_ImportToken_FAIL(t *testing.T) {
	tests := []failTest{
		{
			name:  "Missing semicolon",
			input: "import \"google/protobuf/timestamp.proto\"",
		}, {
			name:  "Missing matching quote",
			input: "import \"google/protobuf/timestamp.proto;",
		}, {
			name:  "Missing first quote",
			input: "import google/protobuf/timestamp.proto\";",
		}, {
			name:  "Quote on next line",
			input: "import \"google/protobuf/timestamp.proto\n\";",
		},
	}
	runFailTokenTest(t, (*TokenParser).ParseImportToken, tests)
}

func Test_OptionToken_FAIL(t *testing.T) {
	tests := []failTest{
		{
			name:  "Missing equals",
			input: "option go_package \"gitlab.ozon.ru/example/api/example;example\"",
		}, {
			name:  "Missing semicolon",
			input: "option go_package = \"gitlab.ozon.ru/example/api/example;example\"",
		}, {
			name:  "Missing matching quote",
			input: "option go_package = \"gitlab.ozon.ru/example/api/example;example;",
		}, {
			name:  "Missing first quote",
			input: "option go_package = gitlab.ozon.ru/example/api/example;example\";",
		}, {
			name:  "Quote on next line",
			input: "option go_package = \"gitlab.ozon.ru/example/api/example;example;\n\";",
		},
		// {
		// 	name: "No space",
		// 	input: "optiongo_package = \"gitlab.ozon.ru/example/api/example;example;\";", // Fix this bug
		// },
	}
	runFailTokenTest(t, (*TokenParser).ParseOptionToken, tests)
}

// TODO: custom method
// func Test_ServiceToken_Fail(t *testing.T) {
// 	tests := []failTest{
// 		{
// 			name:  "Missing brace pair",
// 			input: "service Example {",
// 		}, {
// 			name:  "Missing brace pair",
// 			input: "service Example }",
// 		}, {
// 			name: "Wrong keywords in braces",
// 			input: `service Example {
// 						aboba ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
// 					}`,
// 		}, {
// 			name:  "Quote on next line",
// 			input: "import \"google/protobuf/timestamp.proto\n\";",
// 		},
// 	}
// 	runFailTokenTest(t, (*TokenParser).ParseImportToken, tests)
// }

func Test_RpcToken_FAIL(t *testing.T) {
	tests := []failTest{
		{
			name:  "Missing semicolon",
			input: "rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {}",
		}, {
			name:  "Missing returns",
			input: "rpc ExampleRPC(ExampleRPCRequest) (ExampleRPCResponse) {};",
		}, {
			name:  "Wrong returns",
			input: "rpc ExampleRPC(ExampleRPCRequest) aboba (ExampleRPCResponse) {};",
		}, {
			name:  "Missing brace",
			input: "rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse {};",
		}, {
			name:  "Missing brace",
			input: "rpc ExampleRPC(ExampleRPCRequest) returns ExampleRPCResponse) {};",
		}, {
			name:  "Missing brace",
			input: "rpc ExampleRPC(ExampleRPCRequest returns (ExampleRPCResponse) {};",
		}, {
			name:  "Missing brace",
			input: "rpc ExampleRPCExampleRPCRequest) returns (ExampleRPCResponse) {};",
		}, {
			name:  "Missing name",
			input: "rpc (ExampleRPCRequest) returns (ExampleRPCResponse) {};",
		}, {
			name:  "Missing input",
			input: "rpc ExampleRPC returns (ExampleRPCResponse) {};",
		}, {
			name:  "Missing output",
			input: "rpc ExampleRPC(ExampleRPCRequest) returns {};",
		},
	}
	runFailTokenTest(t, (*TokenParser).ParseRpcToken, tests)
}

