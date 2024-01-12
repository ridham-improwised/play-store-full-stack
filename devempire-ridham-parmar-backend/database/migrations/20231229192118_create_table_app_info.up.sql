-- +migrate Up
CREATE TABLE IF NOT EXISTS apps (
    id SERIAL PRIMARY KEY,
    app_name VARCHAR (500) NOT NULL,
    category VARCHAR (50) NOT NULL,
    rating VARCHAR (50) NOT NULL,
    reviews VARCHAR (50) NOT NULL,
    size VARCHAR (20) NOT NULL,
    installs VARCHAR (50) NOT NULL,
    type VARCHAR (20) NOT NULL,
    price VARCHAR (20) NOT NULL,
    content_rating VARCHAR (20) NOT NULL,
    generes VARCHAR (50) NOT NULL,
    last_updated VARCHAR (70) NOT NULL,
    current_version VARCHAR (100) NOT NULL,
    android_version VARCHAR (100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_apps_id ON apps(id);
