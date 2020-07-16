package makerlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"moul.io/godev"
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

func (c *Client) RawTasksCreate(ctx context.Context, req *makerlogtypes.TasksCreateRequest) (*makerlogtypes.TasksCreateReply, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Post("https://api.getmakerlog.com/tasks/", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//var reply makerlogtypes.TasksCreateReply
	var reply interface{}
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}
	fmt.Println(godev.PrettyJSON(reply))

	return nil, nil

	//return &reply, nil
}
