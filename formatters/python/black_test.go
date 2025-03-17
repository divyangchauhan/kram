package python

import (
	"os/exec"
	"testing"
)

func TestFormat(t *testing.T) {
	// Skip tests if Python is not installed
	if _, err := exec.LookPath("python3"); err != nil {
		t.Skip("Python3 not found, skipping Python formatting tests")
	}

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name: "format simple function",
			input: `def hello():
    print ('Hello World')
    return None`,
			want: `def hello():
    print("Hello World")
    return None
`,
		},
		{
			name: "format class definition",
			input: `class Example:
 def __init__(self,name):
  self.name=name
 def say_hello(self):
  return f'Hello {self.name}'`,
			want: `class Example:
    def __init__(self, name):
        self.name = name

    def say_hello(self):
        return f"Hello {self.name}"
`,
		},
		{
			name: "format imports and long lines",
			input: `import sys, os
from typing import List,   Optional,Dict
def process_data(data: List[Dict[str, Optional[str]]], default_value: str = 'unknown') -> List[str]:
    return [item.get('name', default_value) for item in data if item is not None]`,
			want: `import sys, os
from typing import List, Optional, Dict


def process_data(
    data: List[Dict[str, Optional[str]]], default_value: str = "unknown"
) -> List[str]:
    return [item.get("name", default_value) for item in data if item is not None]
`,
		},
		{
			name: "format list comprehension and comments",
			input: `# Generate squares of even numbers
squares = [x*x for x in range(10)   if x % 2 == 0]  # Using list comprehension

# Print results
for square in squares:
 print(f'Square: {square}')`,
			want: `# Generate squares of even numbers
squares = [x * x for x in range(10) if x % 2 == 0]  # Using list comprehension

# Print results
for square in squares:
    print(f"Square: {square}")
`,
		},
		{
			name: "format decorators and type hints",
			input: `from typing import Callable,TypeVar,Generic
T = TypeVar('T')

@property
@classmethod
def cached_value(cls)->int:
 return 42

@staticmethod
def process(func:Callable[[T],T],value:T)->T:
  return func(value)`,
			want: `from typing import Callable, TypeVar, Generic

T = TypeVar("T")


@property
@classmethod
def cached_value(cls) -> int:
    return 42


@staticmethod
def process(func: Callable[[T], T], value: T) -> T:
    return func(value)
`,
		},
		{
			name: "format match statement and walrus operator",
			input: `def process_value(value):
    match value:
        case str() as s if (length:=len(s))>10:
            return f"Long string: {length} chars"
        case list() | tuple() as seq if (total:=sum(seq))>100:
            return f"Large sequence sum: {total}"
        case _:
            return "No match"`,
			want: `def process_value(value):
    match value:
        case str() as s if (length := len(s)) > 10:
            return f"Long string: {length} chars"
        case list() | tuple() as seq if (total := sum(seq)) > 100:
            return f"Large sequence sum: {total}"
        case _:
            return "No match"
`,
		},
		{
			name: "format async functions and context managers",
			input: `async def fetch_data(url:str)->dict:
 async with aiohttp.ClientSession()as session:
  async with session.get(url)as response:
   if (data:=await response.json()).get('status')=='ok':
    return data
   raise ValueError("Bad response")`,
			want: `async def fetch_data(url: str) -> dict:
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            if (data := await response.json()).get("status") == "ok":
                return data
            raise ValueError("Bad response")
`,
		},
		{
			name: "format multiline strings and docstrings",
			input: `def complex_function(x:int,y:int)->str:
    """This is a complex function that does something.
    
    Args:
        x: The first number
        y: The second number
    
    Returns:
        A formatted string with the result
    """
    text = f"""Multiple
    line
    string with {x} and {y}"""
    return text`,
			want: `def complex_function(x: int, y: int) -> str:
    """This is a complex function that does something.

    Args:
        x: The first number
        y: The second number

    Returns:
        A formatted string with the result
    """
    text = f"""Multiple
    line
    string with {x} and {y}"""
    return text
`,
		},
		{
			name: "format empty input",
			input: "",
			want:  "",
		},
		{
			name:    "invalid Python syntax",
			input:   "def broken(:",
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
			if !tt.wantErr && got != tt.want {
				t.Errorf("Format() output mismatch:\nGot:\n%s\nWant:\n%s", got, tt.want)
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.LineLength != 88 {
		t.Errorf("Expected LineLength to be 88, got %d", config.LineLength)
	}
	if config.SkipString {
		t.Error("Expected SkipString to be false by default")
	}
	if config.FastMode {
		t.Error("Expected FastMode to be false by default")
	}
}
