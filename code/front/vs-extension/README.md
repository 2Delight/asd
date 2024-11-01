# Mline: Basic YAML Linter for VS Code

Mline is a VS Code extension that provides basic linting for YAML files. It checks for syntax errors and ensures the presence of required keys.

## Features

• **Syntax Error Detection:** Highlights syntax errors during YAML parsing.
• **Required Key Validation:** Verifies that the keys `present_model`, `pipeline`, and `stages` are present in your YAML file.
• **Real-time Linting:**  Provides immediate feedback as you type.


## Installation

1. Open VS Code.
2. Go to the Extensions view (Ctrl+Shift+X).
3. Search for "Mline" and click "Install".

## Usage

Mline automatically validates any open YAML files.  Errors will be displayed as underlines in the editor.  The extension currently checks for:

• **YAML Syntax Errors:**  Incorrect YAML formatting will be highlighted.
• **Missing Required Keys:**  The keys `present_model`, `pipeline`, and `stages` are required at the top level of the YAML.  If any are missing, an error will be reported.


## Limitations

• This extension currently performs only basic validation.  It does not perform comprehensive schema validation beyond checking for required keys.  More sophisticated checks (data type validation, etc.) are not yet implemented.
• Error reporting might not be perfectly accurate in all cases, especially with complex YAML structures.


## Release Notes

### 1.0.0
• Initial release
