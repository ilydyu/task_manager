-- +goose Up
create table if not exists tasks (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(500) NOT NULL,
    description TEXT NULL,
    status ENUM('backlog', 'todo', 'in_progress', 'review', 'done', 'cancelled') NOT NULL DEFAULT 'todo',
    priority ENUM('low', 'medium', 'high', 'urgent') NOT NULL DEFAULT 'medium',
    team_id BIGINT UNSIGNED NOT NULL,
    assignee_id BIGINT UNSIGNED NULL,
    created_by BIGINT UNSIGNED NOT NULL,
    deadline DATE NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_tasks_assignee
        FOREIGN KEY (assignee_id) REFERENCES users(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,
    CONSTRAINT fk_tasks_team
        FOREIGN KEY (team_id) REFERENCES teams(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_tasks_created_by
        FOREIGN KEY (created_by) REFERENCES users(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

-- +goose Down
drop table if exists tasks;
