-- +goose Up
-- Ensure the pgcrypto extension is enabled (required for gen_random_uuid)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Drop the old primary key constraint (if needed)
ALTER TABLE workout_entries DROP CONSTRAINT IF EXISTS workout_entries_pkey;

-- Change the type of the ID column to UUID (if it's not already)
ALTER TABLE workout_entries
    ALTER COLUMN id SET DATA TYPE UUID
    USING (id::uuid);

-- Set default value for ID to gen_random_uuid()
ALTER TABLE workout_entries
    ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- Re-add the primary key constraint on ID
ALTER TABLE workout_entries
    ADD PRIMARY KEY (id);

-- +goose Down
-- Revert: Remove default and convert back to TEXT or SERIAL, depending on your old setup

ALTER TABLE workout_entries
    ALTER COLUMN id DROP DEFAULT,
ALTER COLUMN id SET DATA TYPE TEXT
    USING (id::text);

ALTER TABLE workout_entries
    ADD PRIMARY KEY (id);
