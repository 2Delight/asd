import React, { useState, useEffect } from 'react';
import {useParams} from "react-router-dom";
import Header from "../../Components/Header/Header";
import {useTranslation} from "react-i18next";

const lngs = {
    ru: {nativeName: "Рус"},
    kz: {nativeName: "Қаз"}
}

const MfoCardScreen = () => {

    const {id} = useParams();
    const {t, i18n} = useTranslation();

    console.log(id)

    if (!id) return (<div>Not found mfo</div>);

    const buttons = Object.keys(lngs).map((lng) => ({
        title: lngs[lng].nativeName,
        isActive: i18n.resolvedLanguage === lng,
        onActive: () => i18n.changeLanguage(lng)
    }));

    const header = {
        texts: [
            {
                title: t("credits"),
                link: "/mfo/1"
            }
        ],
        buttons: buttons
    }

    return (
        <div className="user-screen">
            <Header props={header}/>
            TODO: Card with id {id}
        </div>
    );
};

export default MfoCardScreen;