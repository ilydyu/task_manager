-- +goose Up
create table if not exists task_comments (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    task_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    content TEXT NOT NULL,
    parent_comment_id BIGINT UNSIGNED NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_task_comments_task (task_id),
    CONSTRAINT fk_task_comments_task
        FOREIGN KEY (task_id) REFERENCES tasks(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_task_comments_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_task_comments_parent
        FOREIGN KEY (parent_comment_id) REFERENCES task_comments(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- +goose Down
drop table if exists task_comments;
