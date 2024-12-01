-- +goose Up
-- Создание таблицы `specifications`
CREATE TABLE specifications (
    id SERIAL PRIMARY KEY,               -- Уникальный идентификатор
    name VARCHAR(255) NOT NULL,          -- Название спецификации
    git_path VARCHAR(512) NOT NULL,      -- Путь в репозитории Git
    status VARCHAR(50) NOT NULL,         -- Статус спецификации (например, "draft", "committed")
    created_at TIMESTAMP DEFAULT NOW(),  -- Время создания записи
    updated_at TIMESTAMP DEFAULT NOW()   -- Время последнего обновления записи
);