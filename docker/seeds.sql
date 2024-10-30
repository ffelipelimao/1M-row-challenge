CREATE TABLE surveys (
    id VARCHAR PRIMARY KEY,
    merchant_id VARCHAR NOT NULL,
    user_id VARCHAR NOT NULL,
    rating FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE merchant_avg_ratings (
    merchant_id VARCHAR PRIMARY KEY,
    average_rating FLOAT
);