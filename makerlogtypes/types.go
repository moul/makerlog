package makerlogtypes

import "time"

// POST /tasks

type TasksCreateRequest struct {
	Content     string      `json:"content,omitempty"`
	Done        bool        `json:"done,omitempty"`
	InProgress  bool        `json:"in_progress,omitempty"`
	DueAt       *time.Time  `json:"due_at,omitempty"`
	User        *User       `json:"user,omitempty"`
	Description string      `json:"description,omitempty"`
	Attachment  *Attachment `json:"-"`
}
type TasksCreateReply Task

func (r *TasksCreateReply) CanonicalURL() string {
	t := Task(*r)
	return t.CanonicalURL()
}

// GET /tasks/

type TasksListRequest struct {
	Done      bool       `json:"done,omitempty" url:"done,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	User      string     `json:"user,omitempty" url:"user,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty" url:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty" url:"end_date,omitempty"`
	Limit     int        `json:"limit,omitempty" url:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty" url:"offset,omitempty"`
	DateRange string     `json:"date_range,omitempty" url:"date_range,omitempty"`
}
type TasksListReply struct {
	Count    int     `json:"count,omitempty"`
	Next     string  `json:"next,omitempty"`
	Previous string  `json:"previous,omitempty"`
	Results  []*Task `json:"results,omitempty"`
}

// GET /notifications/

type NotificationsListReply []struct {
	ID            int         `json:"id,omitempty"`
	Key           string      `json:"key,omitempty"`
	Read          bool        `json:"read,omitempty"`
	Verb          string      `json:"verb,omitempty"`
	Recipient     *User       `json:"recipient,omitempty"`
	Actor         *User       `json:"actor,omitempty"`
	Target        *Task       `json:"target,omitempty"`
	BroadcastLink interface{} `json:"broadcast_link,omitempty"`
	Created       time.Time   `json:"created,omitempty"`
	TargetType    string      `json:"target_type,omitempty"`
}
