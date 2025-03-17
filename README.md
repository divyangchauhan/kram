# Kram

Kram is a universal code formatter that aims to format code in multiple programming languages. It provides a simple command-line interface to format your code files.

## Requirements

- Go 1.21 or higher

## Installation

```bash
go install github.com/divyangchauhan/kram@latest
```

## Usage

```bash
kram [flags] [file...]

# Format a single file
kram myfile.py

# Format multiple files
kram file1.js file2.py file3.go

# Format all files in a directory recursively
kram -r ./project
```

## Supported Languages

- Python
- JavaScript
- Go
- Java
- And more to come...


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License
