package dto

type GetTopCreatorsOutput struct {
	Data []TeamTopCreators `json:"data"`
}

type TeamTopCreators struct {
	TeamID   int64          `json:"team_id"`
	TeamName string         `json:"team_name"`
	TopUsers []TopUserStats `json:"top_users"`
}

type TopUserStats struct {
	UserID       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	TasksCreated int    `json:"tasks_created"`
	Rank         int    `json:"rank"`
}
