import * as vscode from 'vscode';
import { exec } from 'child_process';
import { promisify } from 'util';

const execAsync = promisify(exec);

export function activate(context: vscode.ExtensionContext) {
    // Register the formatter
    let disposable = vscode.languages.registerDocumentFormattingEditProvider(
        [
            { scheme: 'file', language: 'python' },
            { scheme: 'file', language: 'javascript' },
            { scheme: 'file', language: 'typescript' }
        ],
        {
            async provideDocumentFormattingEdits(document: vscode.TextDocument): Promise<vscode.TextEdit[]> {
                const config = vscode.workspace.getConfiguration('kram');
                const kramPath = config.get<string>('path', 'kram');

                try {
                    // Save the document content to a temporary file
                    const text = document.getText();
                    const { stdout, stderr } = await execAsync(`${kramPath} "${document.fileName}"`);

                    if (stderr) {
                        vscode.window.showErrorMessage(`Kram formatting error: ${stderr}`);
                        return [];
                    }

                    // Create a full document replacement edit
                    const firstLine = document.lineAt(0);
                    const lastLine = document.lineAt(document.lineCount - 1);
                    const range = new vscode.Range(
                        firstLine.range.start,
                        lastLine.range.end
                    );

                    return [vscode.TextEdit.replace(range, stdout)];
                } catch (error) {
                    vscode.window.showErrorMessage(`Failed to format: ${error}`);
                    return [];
                }
            }
        }
    );

    // Register format on save if enabled
    vscode.workspace.onWillSaveTextDocument((event) => {
        const config = vscode.workspace.getConfiguration('kram');
        if (config.get<boolean>('formatOnSave', true)) {
            event.waitUntil(
                vscode.commands.executeCommand('editor.action.formatDocument')
            );
        }
    });

    context.subscriptions.push(disposable);

    // Register the format command
    let formatCommand = vscode.commands.registerCommand('kram.format', () => {
        const editor = vscode.window.activeTextEditor;
        if (editor) {
            vscode.commands.executeCommand('editor.action.formatDocument');
        }
    });

    context.subscriptions.push(formatCommand);
}

export function deactivate() {}
