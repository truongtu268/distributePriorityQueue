CREATE TABLE ad (
    id text primary key,
    title text,
    description text,
    status text,
    genre text,
    target_audiences text[],
    visual_elements text[],
    analysis jsonb,
    call_to_action text,
    duration integer,
    priority integer,
    created_at timestamptz,
    retried_at timestamptz,
    completed_at timestamptz,
    retry_time integer
)
