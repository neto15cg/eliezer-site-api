-- First, ensure any existing NULL values are set to current timestamp
UPDATE messages 
SET created_at = NOW() 
WHERE created_at IS NULL;

UPDATE messages 
SET updated_at = NOW() 
WHERE updated_at IS NULL;

-- Add NOT NULL constraints
ALTER TABLE messages 
    ALTER COLUMN created_at SET NOT NULL,
    ALTER COLUMN updated_at SET NOT NULL;

-- Add constraint to ensure timestamps are never NULL
ALTER TABLE messages
    ADD CONSTRAINT check_timestamps_not_null 
    CHECK (created_at IS NOT NULL AND updated_at IS NOT NULL);
