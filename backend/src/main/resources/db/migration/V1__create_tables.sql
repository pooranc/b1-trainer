CREATE TABLE cards (
    id          BIGSERIAL PRIMARY KEY,
    type        VARCHAR(20) NOT NULL,
    question    TEXT NOT NULL,
    answer      TEXT NOT NULL,
    hint        TEXT,
    created_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_card_progress (
    id               BIGSERIAL PRIMARY KEY,
    card_id          BIGINT NOT NULL REFERENCES cards(id),
    repetition       INT DEFAULT 0,
    easiness         FLOAT DEFAULT 2.5,
    interval_days    INT DEFAULT 1,
    next_review_date DATE DEFAULT CURRENT_DATE,
    created_at       TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_progress_next_review ON user_card_progress(next_review_date);