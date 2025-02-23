DROP INDEX IF EXISTS idx_messages_conversation_id;
ALTER TABLE messages DROP COLUMN IF EXISTS conversation_id;
