-- +goose Up
create table if not exists team_members (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    team_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    role ENUM('owner', 'admin', 'member') NOT NULL DEFAULT 'member',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    UNIQUE KEY uk_team_user (team_id, user_id),
    CONSTRAINT fk_team_members_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_team_members_team
        FOREIGN KEY (team_id) REFERENCES teams(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);


CREATE INDEX idx_team_members_team_id ON team_members(team_id);
CREATE INDEX idx_team_members_user_id ON team_members(user_id);

-- +goose Down
drop table if exists team_member;
