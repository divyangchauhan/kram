# Kram Formatter VSCode Extension

This extension integrates Kram, a universal code formatter, with Visual Studio Code. It provides formatting support for multiple programming languages including Python, JavaScript, and TypeScript.

## Features

- Format Python files using Black
- Format JavaScript files using Prettier
- Format on save support
- Configurable formatter settings
- Command palette integration

## Requirements

1. Install Kram:
```bash
go install github.com/divyangchauhan/kram@latest
```

2. Ensure Kram is in your PATH or configure the path in VSCode settings

## Extension Settings

This extension contributes the following settings:

* `kram.path`: Path to the Kram executable (default: "kram")
* `kram.formatOnSave`: Enable/disable format on save (default: true)

## Installation

### From VSCode Marketplace
1. Open VSCode
2. Go to Extensions (Ctrl+Shift+X)
3. Search for "Kram Formatter"
4. Click Install

### From VSIX File
1. Download the .vsix file
2. Open VSCode
3. Go to Extensions (Ctrl+Shift+X)
4. Click "..." at the top of the Extensions panel
5. Choose "Install from VSIX..."
6. Select the downloaded .vsix file

### Manual Installation
1. Clone the repository
2. Run `npm install`
3. Run `npm run compile`
4. Copy the extension files to your VSCode extensions folder:
   - Windows: %USERPROFILE%\.vscode\extensions
   - Linux/macOS: ~/.vscode/extensions

## Usage

### Format Current File
1. Open a supported file (Python, JavaScript, TypeScript)
2. Press `Shift+Alt+F` or right-click and select "Format Document"

### Format on Save
By default, files will be formatted automatically when saved. You can disable this in the settings.

## Troubleshooting

If you encounter any issues:

1. Check if Kram is installed and accessible from the command line
2. Verify the path in VSCode settings if Kram is installed in a custom location
3. Check the VSCode output panel for error messages
4. Make sure required language dependencies are installed:
   - Python: Black formatter
   - JavaScript: Prettier

## Contributing

Feel free to open issues or submit pull requests on the [GitHub repository](https://github.com/divyangchauhan/kram).

## License

MIT
