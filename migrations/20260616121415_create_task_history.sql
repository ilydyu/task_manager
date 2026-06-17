-- +goose Up
create table if not exists task_history (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    task_id BIGINT UNSIGNED NOT NULL,
    changed_by BIGINT UNSIGNED NOT NULL,
    field_name VARCHAR(100) NOT NULL,
    action VARCHAR(255) NOT NULL,
    old_value TEXT NULL,
    new_value TEXT NULL,
    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_task_history_task (task_id),
    CONSTRAINT fk_task_history_task
        FOREIGN KEY (task_id) REFERENCES tasks(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_task_history_user
        FOREIGN KEY (changed_by) REFERENCES users(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

-- +goose Down
drop table if exists task_history;
