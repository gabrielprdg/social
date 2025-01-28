CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    title TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
)
