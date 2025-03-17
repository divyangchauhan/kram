package javascript

import (
	"os/exec"
	"testing"
)

func TestFormat(t *testing.T) {
	// Skip tests if Node.js or npm is not installed
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("Node.js not found, skipping JavaScript formatting tests")
	}
	if _, err := exec.LookPath("npm"); err != nil {
		t.Skip("npm not found, skipping JavaScript formatting tests")
	}

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "format simple function",
			input: "function hello(){console.log('hello');}",
			want:  "function hello() {\n  console.log(\"hello\");\n}\n",
		},
		{
			name:  "format multiple statements",
			input: "const x=1;const y=2;\nfunction add(){return x+y}",
			want:  "const x = 1;\nconst y = 2;\nfunction add() {\n  return x + y;\n}\n",
		},
		{
			name:  "format object and array",
			input: "const obj={a:1,b:[1,2,3],c:{d:4}};const arr=[1,2,3]",
			want:  "const obj = { a: 1, b: [1, 2, 3], c: { d: 4 } };\nconst arr = [1, 2, 3];\n",
		},
		{
			name:  "format class definition",
			input: "class Example{constructor(name){this.name=name}sayHello(){return`Hello ${this.name}`}}",
			want:  "class Example {\n  constructor(name) {\n    this.name = name;\n  }\n  sayHello() {\n    return `Hello ${this.name}`;\n  }\n}\n",
		},
		{
			name:  "format async/await",
			input: "async function getData(){try{const response=await fetch('api/data');return await response.json()}catch(error){console.error(error)}}",
			want:  "async function getData() {\n  try {\n    const response = await fetch(\"api/data\");\n    return await response.json();\n  } catch (error) {\n    console.error(error);\n  }\n}\n",
		},
		{
			name:  "format JSX-like syntax",
			input: "function App(){return(<div><h1>Hello</h1><p>World</p></div>)}",
			want:  "function App() {\n  return (\n    <div>\n      <h1>Hello</h1>\n      <p>World</p>\n    </div>\n  );\n}\n",
		},
		{
			name:  "format empty input",
			input: "",
			want:  "",
		},
		{
			name:    "invalid JavaScript syntax",
			input:   "function broken{{",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got != tt.want {
					t.Errorf("Format() output mismatch:\nGot:\n%s\nWant:\n%s", got, tt.want)
				}
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	// Test default values
	if !config.Semi {
		t.Error("Expected Semi to be true by default")
	}
	if config.SingleQuote {
		t.Error("Expected SingleQuote to be false by default")
	}
	if config.TabWidth != 2 {
		t.Errorf("Expected TabWidth to be 2, got %d", config.TabWidth)
	}
	if config.PrintWidth != 80 {
		t.Errorf("Expected PrintWidth to be 80, got %d", config.PrintWidth)
	}
	if config.TrailingComma != "es5" {
		t.Errorf("Expected TrailingComma to be 'es5', got %s", config.TrailingComma)
	}
	if !config.BracketSpacing {
		t.Error("Expected BracketSpacing to be true by default")
	}
}
