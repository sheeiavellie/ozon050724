CREATE DATABASE forum_db;

-- Tables
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id INTEGER NOT NULL,
    content text NOT NULL,
    created_at TIMESTAMP NOT NULL,
    can_be_commented BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    parent_id INTEGER,
    author_id INTEGER NOT NULL,
    content VARCHAR(2000) NOT NULL,
    created_at TIMESTAMP NOT NULL,

    CONSTRAINT post_comment_fk FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE NO ACTION DEFERRABLE INITIALLY DEFERRED,
    CONSTRAINT parent_comment_fk FOREIGN KEY (parent_id) REFERENCES comments (id) ON DELETE NO ACTION DEFERRABLE INITIALLY DEFERRED
);

-- 🐘🐘🐘
CREATE OR REPLACE FUNCTION check_can_be_commented()
RETURNS TRIGGER AS $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM posts
        WHERE id = NEW.post_id AND can_be_commented = true
    ) THEN
        RAISE EXCEPTION 'Cannot insert into comments: post is not commentable';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_comment
BEFORE INSERT ON comments
FOR EACH ROW
EXECUTE FUNCTION check_can_be_commented();

CREATE OR REPLACE FUNCTION check_post_id()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.parent_id IS NOT NULL THEN
        PERFORM 1
        FROM comments
        WHERE id = NEW.parent_id
        AND post_id = NEW.post_id;

        IF NOT FOUND THEN
            RAISE EXCEPTION 'Invalid parent_id: post_id of the new comment does not match post_id of the referenced comment';
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_post_id_before_insert
BEFORE INSERT ON comments
FOR EACH ROW
EXECUTE FUNCTION check_post_id();

-- Nice data
INSERT INTO posts (title, author_id, content, created_at, can_be_commented) VALUES
    ('Amogus 1', 1, 'Amognus', NOW(), TRUE),
    ('Amogus 2', 1, 'Mongus', NOW(), TRUE),
    ('Amogus 3', 1, 'Sus', NOW(), FALSE);

INSERT INTO comments (post_id, parent_id, author_id, content, created_at) VALUES
    (1, NULL, 1, 'nice comment 1 cool content', NOW()),
    (2, NULL, 1, 'nice comment 2 cool content', NOW()),
    (2, 2, 1, 'nice comment 3 cool nested comment content', NOW());

