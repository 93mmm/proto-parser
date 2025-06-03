package builder

import (
	"strings"
	"testing"

	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/parser/constants"
	"github.com/stretchr/testify/assert"
)

func withSpaces(parts ...string) string {
	spaces := "\n \t\n \t"
	return spaces + strings.Join(parts, spaces)
}

type result struct {
	name string
	kind string
}

type okTest struct {
	name  string
	input string
	want  result
}

type parseXToken func(*TokenParser) (*symbols.Symbol, error)

func assertResult(t *testing.T, expected result, actual *symbols.Symbol, err error) {
	assert.NotNil(t, actual)
	assert.NoError(t, err)
	if actual == nil {
		return
	}
	assert.Equal(t, expected.name, actual.Name)
	assert.Equal(t, expected.kind, actual.Type)
}

func runOkTokenTest(t *testing.T, parseFunc parseXToken, tests []okTest) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := newTestTokenParser(test.input)
			parser.ExtractKeyword()
			res, err := parseFunc(parser)

			t.Logf("input: %v, expected output: %v", test.input, test.want)
			assertResult(t, test.want, res, err)
			assert.True(t, parser.EOF())
		})
	}
}

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

func Test_SyntaxToken_OK(t *testing.T) {
	res := result{
		"proto3", constants.Syntax,
	}

	tests := []okTest{
		{
			name:  "Normal spaces",
			input: `syntax = "proto3";`,
			want:  res,
		}, {
			name:  "Without spaces",
			input: `syntax="proto3";`,
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("syntax", "=", `"proto3"`, ";"),
			want:  res,
		},
	}
	runOkTokenTest(t, (*TokenParser).ParseSyntaxToken, tests)
}

func Test_PackageToken_OK(t *testing.T) {
	res := result{
		"example", constants.Package,
	}

	tests := []okTest{
		{
			name:  "Normal spaces",
			input: "package example;",
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("package", "example", ";"),
			want:  res,
		},
	}
	runOkTokenTest(t, (*TokenParser).ParsePackageToken, tests)
}

func Test_ImportToken_OK(t *testing.T) {
	res := result{
		"google/protobuf/timestamp.proto", constants.Import,
	}

	tests := []okTest{
		{
			name:  "Normal spaces",
			input: "import \"google/protobuf/timestamp.proto\";",
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("import", `"google/protobuf/timestamp.proto"`, ";"),
			want:  res,
		},
	}
	runOkTokenTest(t, (*TokenParser).ParseImportToken, tests)
}

func Test_OptionToken_OK(t *testing.T) {
	res := result{
		"go_package", constants.Option,
	}

	tests := []okTest{
		{
			name:  "Normal spaces",
			input: "option go_package = \"gitlab.ozon.ru/example/api/example;example\";",
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("option", "go_package", "=", `"gitlab.ozon.ru/example/api/example;example"`, ";"),
			want:  res,
		},
	}
	runOkTokenTest(t, (*TokenParser).ParseOptionToken, tests)
}

func Test_ServiceToken_OK(t *testing.T) {
	type serviceTokenTest struct {
		name  string
		input string
		want  []result
	}

	serviceRes := result{
		"Example", constants.Service,
	}
	rpcRes := result{
		"ExampleRPC", constants.Rpc,
	}

	tests := []serviceTokenTest{
		{
			name:  "Normal spaces, w/o rpcs",
			input: "service Example {}",
			want: []result{
				serviceRes,
			},
		}, {
			name:  "No spaces, w/o rpcs",
			input: "service Example {}",
			want: []result{
				serviceRes,
			},
		}, {
			name:  "Maximal spaces",
			input: withSpaces("service", "Example", "{", "}"),
			want: []result{
				serviceRes,
			},
		}, {
			name: "Normal spaces, with one rpc",
			input: `service Example {
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
					}`,
			want: []result{
				serviceRes,
				rpcRes,
			},
		}, {
			name: "Normal spaces, with 2 rpcs",
			input: `service Example {
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
					}`,
			want: []result{
				serviceRes,
				rpcRes,
				rpcRes,
			},
		}, {
			name: "Normal spaces, with some rpcs",
			input: `service Example {
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
						rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
					}`,
			want: []result{
				serviceRes,
				rpcRes,
				rpcRes,
				rpcRes,
				rpcRes,
				rpcRes,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := newTestTokenParser(test.input)
			parser.ExtractKeyword()
			res, err := parser.ParseServiceToken()

			t.Logf("input: %v, expected output: %v", test.input, test.want)
			assert.Equal(t, len(test.want), len(res))
			assert.NoError(t, err)

			for i, out := range test.want {
				assertResult(t, out, res[i], nil)
			}
		})
	}
}

func Test_RpcToken_OK(t *testing.T) {
	res := result{
		"ExampleRPC", constants.Rpc,
	}

	tests := []okTest{
		{
			name:  "Normal spaces",
			input: "rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};",
			want:  res,
		}, {
			name:  "Missing braces",
			input: "rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse);",
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("rpc", "ExampleRPC", "(", "ExampleRPCRequest", ")", "returns", "(", "ExampleRPCResponse", ")", "{", "}", ";"),
			want:  res,
		},
	}
	runOkTokenTest(t, (*TokenParser).ParseRpcToken, tests)
}

func Test_EnumToken_OK(t *testing.T) {
	res := result{
		"ExampleEnum", constants.Enum,
	}

	tests := []okTest{
		{
			name: "Normal spaces",
			input: `enum ExampleEnum {
				ONE = 0;
				TWO = 1;
				THREE = 2;
			}`,
			want: res,
		}, {
			name:  "No spaces",
			input: "enum ExampleEnum{}",
			want:  res,
		}, {
			name:  "No fields",
			input: "enum ExampleEnum {}",
			want:  res,
		}, {
			name:  "Correct brace sequence",
			input: "enum ExampleEnum {{{{{{{{{{}}}}}}}}}}",
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("enum", "ExampleEnum", "{", "}"),
			want:  res,
		},
	}
	runOkTokenTest(t, (*TokenParser).ParseEnumToken, tests)
}

func Test_MessageToken_OK(t *testing.T) {
	res := result{
		"ExampleRPCResponse", constants.Message,
	}

	tests := []okTest{
		{
			name: "Normal spaces",
			input: `message ExampleRPCResponse {
				message Emb { string field11 = 1; }
				ExampleEnum field1 = 1;
				Emb filed2 = 2;
				google.protobuf.Timestamp filed3 = 3;
			}`,
			want: res,
		}, {
			name:  "No spaces",
			input: "message ExampleRPCResponse{}",
			want:  res,
		}, {
			name:  "No fields",
			input: "message ExampleRPCResponse {}",
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("message", "ExampleRPCResponse", "{", "}"),
			want:  res,
		},
	}
	runOkTokenTest(t, (*TokenParser).ParseMessageToken, tests)
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
		// 	input: "optiongo_package = \"gitlab.ozon.ru/example/api/example;example;\";", // FIXME: optiongo_package -> optiongo + _package
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

