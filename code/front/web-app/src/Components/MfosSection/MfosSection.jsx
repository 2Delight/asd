import React from 'react';
import classes from "./MfosSection.module.css";
import MfoItem from "./../MfoItem/MfoItem";

const MfosSection = ({props}) => {
    return (
        <div className={classes.mfosSection}>
            {props.mfoItems.map((mfoItem, index) => (
                <MfoItem key={index} props={mfoItem} />
            ))}
        </div>
    );
};

export default MfosSection;