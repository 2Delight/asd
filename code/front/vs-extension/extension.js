const vscode = require('vscode');
const yaml = require('yaml');
const Ajv = require('ajv');

let diagnosticCollection;

const schema = {
    type: 'object',
    properties: {
        present_model: { type: 'object' },
        pipeline: { type: 'object' },
        stages: { type: 'object' },
        monitoring_services: {
            type: 'array',
            items: {
                type: 'string',
                enum: ["tensorboard", "ngrok", "notification_bot", "model_controller"]
            },
            uniqueItems: true //Added to ensure no duplicates
        }
    },
    required: ['present_model', 'pipeline', 'stages']
};

function activate(context) {
    diagnosticCollection = vscode.languages.createDiagnosticCollection('yaml');
    context.subscriptions.push(diagnosticCollection);

    const diagnosticProviderDisposable = vscode.workspace.onDidChangeTextDocument(event => {
        if (event.document.languageId === 'yaml') {
            diagnosticCollection.set(event.document.uri, yamlDiagnostics(event.document, schema));
        }
    });
    context.subscriptions.push(diagnosticProviderDisposable);

    context.subscriptions.push(vscode.workspace.onDidCloseTextDocument(document => {
        diagnosticCollection.delete(document.uri);
    }));
}

function yamlDiagnostics(document, schema) {
    const diagnostics = [];
    const yamlText = document.getText();

    try {
        const data = yaml.parse(yamlText);
        const validationErrors = validateYaml(data, schema);

        validationErrors.forEach(error => {
            const range = getRangeFromError(document, error); //Use helper function
            if (range) { // Only add if range is valid
                diagnostics.push(new vscode.Diagnostic(range, error.message, vscode.DiagnosticSeverity.Error));
            }
        });
    } catch (error) {
        const range = getRangeFromError(document, error); //Use helper function
        if(range){
            diagnostics.push(new vscode.Diagnostic(range, `YAML Parsing Error: ${error.message}`, vscode.DiagnosticSeverity.Error));
        } else {
            const range = new vscode.Range(0, 0, 0, 0); // Fallback to beginning if range cannot be determined
            diagnostics.push(new vscode.Diagnostic(range, `YAML Parsing Error: ${error.message}`, vscode.DiagnosticSeverity.Error));
        }
    }

    return diagnostics;
}

function validateYaml(data, schema) {
    const errors = [];
    // Simple validation: Check if required keys exist
    const requiredKeys = schema.required;
    for (const key of requiredKeys) {
        if (!(key in data)) {
            errors.push({line:0, col:0, message:`Missing required key: '${key}'`});
        }
    }

    // Add more sophisticated validation checks here as needed.  This example is very basic

    return errors;
}

function getRangeFromError(document, error) {
    if (error.mark && error.mark.position !== undefined && error.mark.length !== undefined) {
        //Use error.mark for accurate position if available
        const startPos = document.positionAt(error.mark.position);
        const endPos = document.positionAt(error.mark.position + error.mark.length);
        return new vscode.Range(startPos, endPos);
    } else if (error.line !== undefined && error.col !== undefined){
        // Attempt to use line and column information if mark is not available
        const startPos = new vscode.Position(error.line, error.col);
        const endPos = new vscode.Position(error.line, error.col + 1); // Adjust as needed
        return new vscode.Range(startPos, endPos);

    } else {
        return null; // Return null if range information is missing
    }
}

function deactivate() {
    if (diagnosticCollection) {
        diagnosticCollection.dispose();
    }
}

module.exports = {
    activate,
    deactivate
}