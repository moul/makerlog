package makerlog

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"moul.io/roundtripper"
)

type Client struct {
	http     http.Client
	hasToken bool
}

func New(token string) *Client {
	client := Client{}

	transport := roundtripper.Transport{
		ExtraHeader: http.Header{"Content-Type": []string{"application/json"}},
	}
	if token != "" {
		transport.ExtraHeader["Authorization"] = []string{"Token " + token}
		client.hasToken = true
	}
	client.http = http.Client{Transport: &transport}

	return &client
}

func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("missing username or password")
	}

	formData := url.Values{
		"username": {username},
		"password": {password},
	}

	// FIXME: use context
	resp, err := http.PostForm("https://api.getmakerlog.com/api-token-auth/", formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var reply struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return "", err
	}

	return reply.Token, nil
}

func (c *Client) RawNotificationsList(ctx context.Context) (*RawNotificationsListReply, error) {
	resp, err := c.http.Get("https://api.getmakerlog.com/notifications/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reply RawNotificationsListReply
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}

	return &reply, nil
}

type RawNotificationsListReply []struct {
	ID        int    `json:"id"`
	Key       string `json:"key"`
	Read      bool   `json:"read"`
	Verb      string `json:"verb"`
	Recipient struct {
		ID                 int       `json:"id"`
		Username           string    `json:"username"`
		FirstName          string    `json:"first_name"`
		LastName           string    `json:"last_name"`
		Status             string    `json:"status"`
		Description        string    `json:"description"`
		Verified           bool      `json:"verified"`
		Private            bool      `json:"private"`
		Avatar             string    `json:"avatar"`
		Streak             string    `json:"streak"`
		Timezone           string    `json:"timezone"`
		WeekTda            string    `json:"week_tda"`
		TwitterHandle      string    `json:"twitter_handle"`
		InstagramHandle    string    `json:"instagram_handle"`
		ProductHuntHandle  string    `json:"product_hunt_handle"`
		GithubHandle       string    `json:"github_handle"`
		TelegramHandle     string    `json:"telegram_handle"`
		NomadlistHandle    string    `json:"nomadlist_handle"`
		BmcHandle          string    `json:"bmc_handle"`
		Header             string    `json:"header"`
		IsStaff            bool      `json:"is_staff"`
		Donor              bool      `json:"donor"`
		ShipstreamsHandle  string    `json:"shipstreams_handle"`
		Website            string    `json:"website"`
		Tester             bool      `json:"tester"`
		IsLive             bool      `json:"is_live"`
		Digest             bool      `json:"digest"`
		Gold               bool      `json:"gold"`
		Accent             string    `json:"accent"`
		MakerScore         string    `json:"maker_score"`
		DarkMode           bool      `json:"dark_mode"`
		WeekendsOff        bool      `json:"weekends_off"`
		HardcoreMode       bool      `json:"hardcore_mode"`
		EmailNotifications bool      `json:"email_notifications"`
		OgImage            string    `json:"og_image"`
		DateJoined         time.Time `json:"date_joined"`
	} `json:"recipient"`
	Actor         string    `json:"actor"`
	Unread        bool      `json:"unread"`
	Target        string    `json:"target"`
	BroadcastLink string    `json:"broadcast_link"`
	Created       time.Time `json:"created"`
	TargetType    string    `json:"target_type"`
}
