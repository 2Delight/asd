import React, {useState, useEffect, useRef} from 'react';
import Header from "../../Components/Header/Header";
import {useParams} from "react-router-dom";
import classes from "./MfosScreen.module.css";
import {useTranslation, Trans} from 'react-i18next'
import FilterButton from "../../Components/FilterButton/FilterButton";
import Editor, { DiffEditor, useMonaco, loader } from '@monaco-editor/react';
import { SpecificationService } from './SpecificationService';
import { REACT_APP_API_URL, REACT_APP_PORT } from "../../Constants"

const UserScreen = () => {
    const { id } = useParams();
    const editorRef = useRef(null);
    const {t, i18n} = useTranslation();
    const [specification, setSpecification] = useState("");
    const url = REACT_APP_API_URL + ":" + REACT_APP_PORT
    const specificationService = new SpecificationService(url);
    const [validationResult, setValidationResult] = useState(null);

    useEffect(() => {
        fetchSpecification();
    }, [i18n.resolvedLanguage]);

    const fetchSpecification = async () => {
        try {
            const spec = await specificationService.getSpecificationById(1, false);  // true для мока
            setSpecification(spec.content);
            console.log(spec.content)
        } catch (error) {
            console.error("Ошибка при получении спецификации:", error);
        }
    };

    const header = {
        texts: [],
        buttons: []
    }

    const validateSpecification = async () => {
        try {
            const result = await specificationService.validateSpecification(id, specification, false); // true для мока
            setValidationResult(result);
            if (result.isValid) {
                console.log("Validation successful!");
            } else {
                console.error("Validation failed:", result.errors);
                console.log("Hints for improvement:", result.hints);

                if (window.monaco && editorRef.current) {
                    const model = editorRef.current.getModel();
                    window.monaco.editor.setModelMarkers(model, "owner", [
                        {
                            startLineNumber: 1,
                            startColumn: 1,
                            endLineNumber: 1,
                            endColumn: 50,
                            message: "Invalid syntax detected in the first line.",
                            severity: window.monaco.MarkerSeverity.Error,
                        },
                    ]);
                }
            }
        } catch (error) {
            console.error("Ошибка при валидации спецификации:", error);
        }
    };

    const validateBtn = {
        title: `Validate id[${id}]`,
        onClick: () => {
            console.log(`Validate button clicked for ID: ${id}`);
            validateSpecification();
        },
    }

    const submitBtn = {
        title: `Submit id[${id}]`,
        onClick: () => {
            specificationService.updateSpecification(id, specification)
            console.log(`Submit button clicked for ID: ${id}`);
        },
    }

    function handleEditorChange(value, event) {
        // specificationService.validateSpecification(id, specification)
        // console.log('here is the current model value:', value);
    }

    function handleEditorDidMount(editor, monaco) {
        editorRef.current = editor;
        window.monaco = monaco;
    }

    if (id == undefined) {
        return (
            <div>Not found with id</div>
        )
    }

    return (
        <div className={classes.userScreen}>
            <Header props={header}/>
            <div className={classes.filterButtonStack}>
                <FilterButton props={validateBtn} />
                <FilterButton props={submitBtn} />
            </div>
            <Editor
                height="90vh"
                defaultLanguage="yml"
                value={specification}
                onChange={handleEditorChange}
                onMount={handleEditorDidMount}
            />
            {validationResult && (
                <div className={classes.validationResult}>
                    <h3>Validation Result</h3>
                    <p>Is Valid: {validationResult.isValid ? "Yes" : "No"}</p>
                    {validationResult.errors.length > 0 && (
                        <div>
                            <h4>Errors:</h4>
                            <ul>
                                {validationResult.errors.map((error, index) => (
                                    <li key={index}>
                                        {error.code}: {error.message}
                                    </li>
                                ))}
                            </ul>
                        </div>
                    )}
                    {validationResult.hints.length > 0 && (
                        <div>
                            <h4>Hints:</h4>
                            <ul>
                                {validationResult.hints.map((hint, index) => (
                                    <li key={index}>{hint.message}</li>
                                ))}
                            </ul>
                        </div>
                    )}
                </div>
            )}
        </div>
    );
};

export default UserScreen;