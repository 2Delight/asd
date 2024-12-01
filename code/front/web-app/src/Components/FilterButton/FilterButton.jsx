import React from 'react';
import classes from "./FilterButton.module.css";
import {Image} from "react-bootstrap";

const FilterButton = ({props}) => {
    return (
        <button
            onClick={props.onClick}
            className={classes.buttonBackground}
        >
            <a className={classes.buttonIcon}>
                {/* <Image width={26}
                       height={26}
                       src="/icons/filterIcon.svg"
                >
                </Image> */}
            </a>
            <div className={classes.buttonTitle}>
                {props.title}
            </div>
        </button>
    );
};

export default FilterButton;