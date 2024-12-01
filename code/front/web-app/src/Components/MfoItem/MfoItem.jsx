import React from 'react';
import classes from "./MfoItem.module.css";
import Button from "./../Button/Button";

const MfoItem = ({props}) => {
    return (
        <div className={classes.mfoCard}>
            <div className={classes.mfoHeader}>
                <img src={props.imageUrl} alt={props.name} className={classes.mfoAvatar}/>
                <div className={classes.mfoTexts}>
                    <div className={classes.mfoTitle}>{props.title}</div>
                    <div className={classes.mfoSubtitle}>{props.subtitle}</div>
                </div>
            </div>
            {props.features && props.features.length > 0 && (
                <ul>
                    {props.features.map((feature, index) => (
                        <li key={index}>
                            <div className={classes.mfoFeatureTitle}>{feature.title}</div>
                            <div className={classes.mfoFeatureSubtitle}>{feature.subtitle}</div>
                        </li>
                    ))}
                </ul>
            )}
            {props.button && (
                <Button href="http://localhost:3000/mfo/1" props={{title: props.button.title, link: props.button.link}} />
            )}
        </div>
    );
};

export default MfoItem;