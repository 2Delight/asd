import React from 'react';
import classes from "./Header.module.css";

const Header = ({props}) => {
    return (
        <div className={classes.header}>
            <a className={classes.linkLogo} href="/">
                <div className={classes.logo}>mline</div>
            </a>
            {props.texts && (
                <div className={classes.texts}>
                    {props.texts.map((textItem, index) => (
                        <a
                            className={classes.text}
                            href={textItem.link}
                            key={index}
                        >{textItem.title}
                        </a>
                    ))}
                </div>
            )}

            <div className={classes.buttons}>
                {props.buttons !== undefined && props.buttons && (
                    (
                        props.buttons.map((button, index) => (
                            <button
                                className={(button.isActive === true ? classes.buttonActive : classes.button)}
                                onClick={() => {
                                    button.onActive()
                                }}
                                key={index}
                            >
                                {button.title}
                            </button>
                        ))
                    )
                )}
            </div>
        </div>
    );
};


export default Header;