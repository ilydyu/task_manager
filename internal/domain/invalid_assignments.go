package domain

type InvalidAssignment struct {
	TaskID       int64
	Title        string
	TeamID       int64
	TeamName     string
	AssigneeID   int64
	AssigneeName string
}
