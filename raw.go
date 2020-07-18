package makerlog

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"

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

func (c *Client) RawTasksCreate(ctx context.Context, req *makerlogtypes.TasksCreateRequest) (*makerlogtypes.TasksCreateReply, error) {
	fields := map[string]string{
		"content": req.Content,
	}
	if req.Description != "" {
		fields["description"] = req.Description
	}
	if req.Done {
		fields["done"] = "true"
	}
	if req.InProgress {
		fields["in_progress"] = "true"
	}
	if req.DueAt != nil {
		fields["due_at"] = req.DueAt.String()
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for k, v := range fields {
		fw, err := w.CreateFormField(k)
		if err != nil {
			return nil, err
		}
		_, err = fw.Write([]byte(v))
		if err != nil {
			return nil, err
		}
	}
	if req.Attachment != nil {
		fw, err := w.CreateFormFile("attachment", req.Attachment.Filename)
		if err != nil {
			return nil, err
		}
		_, err = fw.Write(req.Attachment.Bytes)
		if err != nil {
			return nil, err
		}
		w.Close()
	}

	request, err := http.NewRequest("POST", "https://api.getmakerlog.com/tasks/", &b)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reply makerlogtypes.TasksCreateReply
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}
	return &reply, nil
}
