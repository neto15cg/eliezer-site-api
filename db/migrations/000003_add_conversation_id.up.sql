ALTER TABLE messages
ADD COLUMN conversation_id UUID NULL;

CREATE INDEX idx_messages_conversation_id ON messages(conversation_id);
