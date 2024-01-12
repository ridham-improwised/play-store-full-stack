-- +migrate Up
CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    app_name VARCHAR (500) NOT NULL,
    translated_review VARCHAR (5000) DEFAULT NULL,
    sentiment VARCHAR (50) NOT NULL,
    sentiment_polarity VARCHAR (50) DEFAULT NULL,
    sentiment_subjectivity VARCHAR (50) DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_reviews_app_name ON reviews(app_name);
