package tapd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetCommitMsg 获取源码提交关键字，返回提交关键字字符串
func (c *Client) GetCommitMsg(ctx context.Context, req *model.GetCommitMsgRequest) (string, error) {
	data, err := c.doGet(ctx, "/svn_commits/get_scm_copy_keywords", req.ToParams())
	if err != nil {
		return "", err
	}

	var result string
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("failed to parse commit keyword: %w", err)
	}
	return result, nil
}

// ListReleases 查询发布计划列表
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_new_releases.html
func (c *Client) ListReleases(ctx context.Context, req *model.ListReleasesRequest) ([]model.Release, error) {
	data, err := c.doGet(ctx, "/releases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Release](data, "Release")
}

// GetTodoStories 获取用户待办需求，返回强类型 Story 切片
func (c *Client) GetTodoStories(ctx context.Context, req *model.GetTodoRequest) ([]model.Story, error) {
	data, err := c.doGet(ctx, "/user_oauth/get_user_todo_story", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Story](data, "Story")
}

// GetTodoTasks 获取用户待办任务，返回强类型 Task 切片
func (c *Client) GetTodoTasks(ctx context.Context, req *model.GetTodoRequest) ([]model.Task, error) {
	data, err := c.doGet(ctx, "/user_oauth/get_user_todo_task", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Task](data, "Task")
}

// GetTodoBugs 获取用户待办缺陷，返回强类型 Bug 切片
func (c *Client) GetTodoBugs(ctx context.Context, req *model.GetTodoRequest) ([]model.Bug, error) {
	data, err := c.doGet(ctx, "/user_oauth/get_user_todo_bug", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.Bug](data, "Bug")
}

// SendQiweiMessage 发送消息到企业微信群
// 注意：此功能需要配置企业微信机器人 webhook URL
func (c *Client) SendQiweiMessage(ctx context.Context, webhookURL, msg string) error {
	if webhookURL == "" {
		return fmt.Errorf("qiwei webhook URL is not configured")
	}
	parsed, err := url.Parse(webhookURL)
	if err != nil {
		return fmt.Errorf("invalid webhook URL: %w", err)
	}
	if parsed.Scheme != "https" {
		return fmt.Errorf("webhook URL must use HTTPS scheme, got %q", parsed.Scheme)
	}

	// 构造请求体
	msgType := "markdown"
	payload := map[string]interface{}{
		"msgtype": msgType,
		msgType: map[string]string{
			"content": msg,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return c.doPostJSON(ctx, webhookURL, payloadBytes)
}
