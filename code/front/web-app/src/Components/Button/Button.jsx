import React from 'react';
import classes from "./Button.module.css";

const Button = ({props}) => {
    return (
        <button onClick={() => window.location.href = props.link}  className={classes.buttonBackground}>
            <div className={classes.buttonTitle}>
                { props.title }
            </div>
        </button>
    );
};

export default Button;