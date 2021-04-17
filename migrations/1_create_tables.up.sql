CREATE TABLE threads (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);
-- Changed votes to not null, since go will default to 0
-- if the value is not provided :)
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    -- Deletes all threads posts if a thread is deleted
    thread_id UUID NOT NULL REFERENCES threads (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    votes INT NOT NULL
);
CREATE TABLE comments(
    id UUID PRIMARY KEY,
    post_id UUID NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    votes INT NOT NULL
);