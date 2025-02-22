-- Remove the constraint and NULL checks
ALTER TABLE messages
    DROP CONSTRAINT IF EXISTS check_timestamps_not_null;

ALTER TABLE messages 
    ALTER COLUMN created_at DROP NOT NULL,
    ALTER COLUMN updated_at DROP NOT NULL;
