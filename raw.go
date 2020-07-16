package makerlog

import (
	"context"
	"encoding/json"

	"github.com/google/go-querystring/query"
	"moul.io/makerlog/makerlogtypes"
)

func (c *Client) RawNotificationsList(ctx context.Context) (*makerlogtypes.NotificationsListReply, error) {
	resp, err := c.http.Get("https://api.getmakerlog.com/notifications/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reply makerlogtypes.NotificationsListReply
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}

	return &reply, nil
}

func (c *Client) RawTasksList(ctx context.Context, req *makerlogtypes.TasksListRequest) (*makerlogtypes.TasksListReply, error) {
	qs := ""
	if req != nil {
		v, err := query.Values(req)
		if err != nil {
			return nil, err
		}
		qs = v.Encode()
	}

	resp, err := c.http.Get("https://api.getmakerlog.com/tasks/?" + qs)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reply makerlogtypes.TasksListReply
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
