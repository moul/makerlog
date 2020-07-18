package makerlogtypes

import (
	"fmt"
	"time"
)

type User struct {
	ID                 int         `json:"id,omitempty"`
	Username           string      `json:"username,omitempty"`
	FirstName          string      `json:"first_name,omitempty"`
	LastName           string      `json:"last_name,omitempty"`
	Status             interface{} `json:"status,omitempty"`
	Description        string      `json:"description,omitempty"`
	Verified           bool        `json:"verified,omitempty"`
	Private            bool        `json:"private,omitempty"`
	Avatar             string      `json:"avatar,omitempty"`
	Streak             int         `json:"streak,omitempty"`
	Timezone           string      `json:"timezone,omitempty"`
	WeekTda            float64     `json:"week_tda,omitempty"`
	TwitterHandle      string      `json:"twitter_handle,omitempty"`
	InstagramHandle    string      `json:"instagram_handle,omitempty"`
	ProductHuntHandle  string      `json:"product_hunt_handle,omitempty"`
	GithubHandle       string      `json:"github_handle,omitempty"`
	TelegramHandle     string      `json:"telegram_handle,omitempty"`
	NomadlistHandle    string      `json:"nomadlist_handle,omitempty"`
	BmcHandle          string      `json:"bmc_handle,omitempty"`
	Header             interface{} `json:"header,omitempty"`
	IsStaff            bool        `json:"is_staff,omitempty"`
	Donor              bool        `json:"donor,omitempty"`
	ShipstreamsHandle  string      `json:"shipstreams_handle,omitempty"`
	Website            string      `json:"website,omitempty"`
	Tester             bool        `json:"tester,omitempty"`
	IsLive             bool        `json:"is_live,omitempty"`
	Digest             bool        `json:"digest,omitempty"`
	Gold               bool        `json:"gold,omitempty"`
	Accent             string      `json:"accent,omitempty"`
	MakerScore         int         `json:"maker_score,omitempty"`
	DarkMode           bool        `json:"dark_mode,omitempty"`
	WeekendsOff        bool        `json:"weekends_off,omitempty"`
	HardcoreMode       bool        `json:"hardcore_mode,omitempty"`
	EmailNotifications bool        `json:"email_notifications,omitempty"`
	OgImage            string      `json:"og_image,omitempty"`
	DateJoined         time.Time   `json:"date_joined,omitempty"`
}

type Project struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Private bool   `json:"private,omitempty"`
	User    int    `json:"user,omitempty"`
}

type Task struct {
	ID           int         `json:"id,omitempty"`
	Event        interface{} `json:"event,omitempty"`
	Done         bool        `json:"done,omitempty"`
	InProgress   bool        `json:"in_progress,omitempty"`
	Content      string      `json:"content,omitempty"`
	CreatedAt    time.Time   `json:"created_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at,omitempty"`
	DueAt        interface{} `json:"due_at,omitempty"`
	DoneAt       time.Time   `json:"done_at,omitempty"`
	User         *User       `json:"user,omitempty"`
	Description  interface{} `json:"description,omitempty"`
	ProjectSet   []*Project  `json:"project_set,omitempty"`
	Praise       int         `json:"praise,omitempty"`
	Attachment   string      `json:"attachment,omitempty"`
	CommentCount int         `json:"comment_count,omitempty"`
	OgImage      string      `json:"og_image,omitempty"`
}

func (t *Task) CanonicalURL() string {
	return fmt.Sprintf("https://getmakerlog.com/tasks/%d", t.ID)
}
