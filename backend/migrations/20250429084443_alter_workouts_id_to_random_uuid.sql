-- +goose Up
-- Enable pgcrypto for UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Drop foreign key constraint first
ALTER TABLE workout_entries DROP CONSTRAINT IF EXISTS workout_entries_workout_id_fkey;

-- Drop the primary key constraint
ALTER TABLE workouts DROP CONSTRAINT IF EXISTS workouts_pkey;

-- Change the ID column type to UUID and set default
ALTER TABLE workouts
    ALTER COLUMN id SET DATA TYPE UUID
    USING (id::uuid),
ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- Recreate primary key
ALTER TABLE workouts ADD PRIMARY KEY (id);

-- Recreate the foreign key
ALTER TABLE workout_entries
    ADD CONSTRAINT workout_entries_workout_id_fkey
        FOREIGN KEY (workout_id)
            REFERENCES workouts(id)
            ON DELETE CASCADE;

-- +goose Down
-- Drop foreign key constraint first
ALTER TABLE workout_entries DROP CONSTRAINT IF EXISTS workout_entries_workout_id_fkey;

-- Drop primary key constraint
ALTER TABLE workouts DROP CONSTRAINT IF EXISTS workouts_pkey;

-- Revert ID column back to TEXT (or whatever it was before)
ALTER TABLE workouts
    ALTER COLUMN id DROP DEFAULT,
ALTER COLUMN id SET DATA TYPE TEXT
    USING (id::text);

-- Recreate primary key
ALTER TABLE workouts ADD PRIMARY KEY (id);

-- Recreate the foreign key
ALTER TABLE workout_entries
    ADD CONSTRAINT workout_entries_workout_id_fkey
        FOREIGN KEY (workout_id)
            REFERENCES workouts(id)
            ON DELETE CASCADE;
