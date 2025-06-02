package builder

import (
	"strings"
	"testing"

	"github.com/93mmm/proto-parser/internal/symbols"
	"github.com/93mmm/proto-parser/internal/token"
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

type tokenTest struct {
	name  string
	input string
	want  result
}

type serviceTokenTest struct {
	name  string
	input string
	want  []result
}

type parseXToken func(*TokenParser) (*symbols.Symbol, error)

func assertResult(t *testing.T, expected result, actual *symbols.Symbol, err error) {
	assert.NotNil(t, actual)
	assert.NoError(t, err)
	if actual == nil {
		return
	}
	assert.Equal(t, expected.name, actual.Name())
	assert.Equal(t, expected.kind, actual.Type())
}

func runTokenTest(t *testing.T, parseFunc parseXToken, tests []tokenTest) {
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

func Test_SyntaxToken_OK(t *testing.T) {
	res := result{
		"proto3", token.Syntax,
	}

	tests := []tokenTest{
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
	runTokenTest(t, (*TokenParser).ParseSyntaxToken, tests)
}

func Test_PackageToken_OK(t *testing.T) {
	res := result{
		"example", token.Package,
	}

	tests := []tokenTest{
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
	runTokenTest(t, (*TokenParser).ParsePackageToken, tests)
}

func Test_ImportToken_OK(t *testing.T) {
	res := result{
		"google/protobuf/timestamp.proto", token.Import,
	}

	tests := []tokenTest{
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
	runTokenTest(t, (*TokenParser).ParseImportToken, tests)
}

func Test_OptionToken_OK(t *testing.T) {
	res := result{
		"go_package", token.Option,
	}

	tests := []tokenTest{
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
	runTokenTest(t, (*TokenParser).ParseOptionToken, tests)
}

func Test_ServiceToken_OK(t *testing.T) {
	serviceRes := result{
		"Example", token.Service,
	}
	rpcRes := result{
		"ExampleRPC", token.Rpc,
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

	t.Run("Normal, with rpcs inside", func(t *testing.T) {
		input := `service Example {
			rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};
			rpc ExampleRPC1(ExampleRPCRequest) returns (ExampleRPCResponse) {};
		}`
		parser := newTestTokenParser(input)
		parser.ExtractKeyword()
		result, err := parser.ParseServiceToken()

		assert.Equal(t, 3, len(result))

		assert.Equal(t, "service", result[0].Type())
		assert.Equal(t, "Example", result[0].Name())

		assert.Equal(t, "rpc", result[1].Type())
		assert.Equal(t, "ExampleRPC", result[1].Name())

		assert.Equal(t, "rpc", result[2].Type())
		assert.Equal(t, "ExampleRPC1", result[2].Name())
		assert.NoError(t, err)
	})
}

func Test_RpcToken_OK(t *testing.T) {
	res := result{
		"ExampleRPC", token.Rpc,
	}

	tests := []tokenTest{
		{
			name:  "Normal spaces",
			input: "rpc ExampleRPC(ExampleRPCRequest) returns (ExampleRPCResponse) {};",
			want:  res,
		}, {
			name:  "Maximal spaces",
			input: withSpaces("rpc", "ExampleRPC", "(", "ExampleRPCRequest", ")", "returns", "(", "ExampleRPCResponse", ")", "{", "}", ";"),
			want:  res,
		},
	}
	runTokenTest(t, (*TokenParser).ParseRpcToken, tests)
}

func Test_EnumToken_OK(t *testing.T) {
	res := result{
		"ExampleEnum", token.Enum,
	}

	tests := []tokenTest{
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
			name:  "Maximal spaces",
			input: withSpaces("enum", "ExampleEnum", "{", "}"),
			want:  res,
		},
	}
	runTokenTest(t, (*TokenParser).ParseEnumToken, tests)
}

func Test_MessageToken_OK(t *testing.T) {
	res := result{
		"ExampleRPCResponse", token.Message,
	}

	tests := []tokenTest{
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
	runTokenTest(t, (*TokenParser).ParseMessageToken, tests)
}
