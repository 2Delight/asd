import React, { useState } from 'react';
import classes from "./Modal.module.css";

const Modal = ({ props }) => {
    const [sum, setSum] = useState(300000);
    const [term, setTerm] = useState(3);

    const handleSumChange = (event) => {
        setSum(event.target.value);
    };

    const handleTermChange = (event) => {
        setTerm(event.target.value);
    };
    return (
        <div className={classes.modal}>
            <div className={classes.overlay} onClick={props.onClose}></div>
            <div className={classes.modalContent}>
                <div className={classes.modalHeader}>
                    <h2>{props.filtersTitle}</h2>
                    <button className={classes.closeButton} onClick={props.onClose}>X</button>
                </div>

                <div className={classes.inputGroup}>
                    <input type="text" placeholder={props.inputValue} className={classes.inputField} />
                    <button className={classes.inputButton}>{props.sumTengeTitle}</button>
                </div>

                <div className={classes.sliders}>
                    <div className={classes.sliderGroup}>
                        <label>{props.sumTitle}</label>
                        <input
                            type="range"
                            min="1000"
                            max="1000000"
                            value={sum}
                            onChange={handleSumChange}
                            className={classes.slider}
                        />
                        <span>{sum.toLocaleString()} {props.tengeTitle}</span> {/* Display formatted value */}
                    </div>

                    <div className={classes.sliderGroup}>
                        <label>{props.srokTitle}</label>
                        <input
                            type="range"
                            min="1"
                            max="12"
                            value={term}
                            onChange={handleTermChange}
                            className={classes.slider}
                        />
                        <span>{term} {props.monthTitle}</span> {/* Display singular/plural */}
                    </div>
                </div>

                <div className={classes.checkboxGrid}>
                    <label>
                        <input type="checkbox"/> {props.yearsTitle}
                    </label>
                    <label>
                        <input type="checkbox"/> {props.yearsTitle}
                    </label>
                    <label>
                        <input type="checkbox"/> {props.yearsTitle}
                    </label>
                    <label>
                        <input type="checkbox"/> {props.yearsTitle}
                    </label>
                    <label>
                        <input type="checkbox"/> {props.yearsTitle}
                    </label>
                    <label>
                        <input type="checkbox"/> {props.yearsTitle}
                    </label>
                </div>

                <button
                    className={classes.submitButton}
                    onClick={props.onClose}
                >{props.applyTitle}</button>
            </div>
        </div>
    );
};

export default Modal;