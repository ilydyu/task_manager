package domain

type TeamStatistics struct {
	ID                 uint64 `json:"id"`
	Name               string `json:"name"`
	MemberCount        int    `json:"member_count"`
	DoneTasksLast7Days int    `json:"done_tasks_last_7_days"`
}
