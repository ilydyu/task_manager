package domain

import "time"

type TaskHistory struct {
	ID        int64     `json:"id"`
	TaskID    int64     `json:"task_id"`
	ChangedBy int64     `json:"changed_by"`
	FieldName string    `json:"field_name"`
	Action    string    `json:"action"`
	OldValue  *string   `json:"old_value"`
	NewValue  *string   `json:"new_value"`
	ChangedAt time.Time `json:"changed_at"`
}
