import i18next from "i18next"
import { initReactI18next } from "react-i18next"

i18next
    .use(initReactI18next)
    .init({
    lng: 'ru',
    resources: {
        ru: {
            translation: {
                credits: "Кредиты",
                sum: "Сумма",
                sumTenge: "Сумма в тенге",
                tenge: "тенге",
                pereiti: "Перейти",
                stavka: "Ставка",
                srok: "Срок",
                filters: "Фильтры",
                inputValue: "Введите значение",
                helloWorld: "Привет мир",
                months: "месяца",
                apply: "Применить",
                years14: "Можно взять с 14 лет"
            }
        },
        kz: {
            translation: {
                credits: "Несиелер",
                sum: "Сомасы",
                sumTenge: "Теңге сомасы",
                tenge: "теңге",
                pereiti: "Өту",
                stavka: "Ставка",
                srok: "Мерзімі",
                filters: "Сүзгілер",
                inputValue: "Мәнді енгізіңіз",
                helloWorld: "Сәлем Әлем",
                months: "ай",
                apply: "Қолдану",
                years14: "14 жастан бастап алуға болады"
            }
        }
    }
});