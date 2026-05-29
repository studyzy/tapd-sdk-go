// Package tapd 封装了 TAPD API 的 HTTP 客户端，提供对 TAPD 平台所有资源的类型安全访问
package tapd

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/studyzy/tapd-sdk-go/model"
)

const (
	// DefaultAPIURL 是 TAPD API 的默认基础地址
	DefaultAPIURL = "https://api.tapd.cn"
	// DefaultWebURL 是 TAPD 前端页面的默认基础地址
	DefaultWebURL = "https://www.tapd.cn"
)

// TAPDError 表示 TAPD API 返回的错误
type TAPDError struct {
	HTTPStatus int
	ExitCode   int
	Message    string
}

// Error 返回错误描述信息
func (e *TAPDError) Error() string {
	return e.Message
}

// Client 是 TAPD API 的 HTTP 客户端封装，可安全地在多个 goroutine 间共享。
type Client struct {
	baseURL    string
	webURL     string
	httpClient *http.Client
	authHeader string // "Bearer <token>" or "Basic <base64>"

	mu   sync.RWMutex
	nick string // 当前用户昵称，Bearer Token 认证时通过 FetchNick() 自动获取
}

// NewClient 创建一个使用默认 URL 的 TAPD API 客户端。
// 如果 accessToken 非空则使用 Bearer 认证，否则使用 Basic 认证。
func NewClient(accessToken, apiUser, apiPassword string) *Client {
	return NewClientWithBaseURL(DefaultAPIURL, DefaultWebURL, accessToken, apiUser, apiPassword)
}

// NewClientWithBaseURL 创建一个指定 API 和 Web 基础地址的 TAPD API 客户端。
// apiURL 为 API 请求基础地址，webURL 为前端页面基础地址。
// 如果 accessToken 非空则使用 Bearer 认证，否则使用 Basic 认证。
func NewClientWithBaseURL(apiURL, webURL, accessToken, apiUser, apiPassword string) *Client {
	if apiURL == "" {
		apiURL = DefaultAPIURL
	}
	if webURL == "" {
		webURL = DefaultWebURL
	}
	// 去除末尾斜杠，确保拼接路径时格式正确
	apiURL = strings.TrimRight(apiURL, "/")
	webURL = strings.TrimRight(webURL, "/")

	return &Client{
		baseURL: apiURL,
		webURL:  webURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		authHeader: buildAuthHeader(accessToken, apiUser, apiPassword),
	}
}

// buildAuthHeader 根据凭据类型构建 Authorization 请求头值
func buildAuthHeader(accessToken, apiUser, apiPassword string) string {
	if accessToken != "" {
		return "Bearer " + accessToken
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(apiUser + ":" + apiPassword))
	return "Basic " + encoded
}

// WebURL 返回前端页面的基础地址
func (c *Client) WebURL() string {
	return c.webURL
}

// GetNick 返回当前用户昵称（并发安全）
func (c *Client) GetNick() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.nick
}

// SetNick 设置当前用户昵称（并发安全）
func (c *Client) SetNick(nick string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nick = nick
}

// FetchNick 通过 /users/info 接口获取当前用户昵称（Bearer Token 认证时使用）
func (c *Client) FetchNick(ctx context.Context) error {
	data, err := c.doGet(ctx, "/users/info", nil)
	if err != nil {
		return err
	}

	var result struct {
		Nick string `json:"nick"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("failed to parse user info: %w", err)
	}

	c.SetNick(result.Nick)
	return nil
}

// TestAuth 通过调用 /quickstart/testauth 接口验证凭据是否有效。
// 返回 nil 表示认证成功，否则返回错误信息。
func (c *Client) TestAuth(ctx context.Context) error {
	_, err := c.doGet(ctx, "/quickstart/testauth", nil)
	return err
}

// applyAuth 为请求设置 Authorization 头
func (c *Client) applyAuth(req *http.Request) {
	req.Header.Set("Authorization", c.authHeader)
}

// doGet 发送 GET 请求到指定端点，解析 TAPD 统一响应格式并返回 data 字段
func (c *Client) doGet(ctx context.Context, endpoint string, params map[string]string) (json.RawMessage, error) {
	reqURL, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if len(params) > 0 {
		q := reqURL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		reqURL.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(req)

	return c.doRequest(req)
}

// doPost 发送 POST 请求到指定端点，使用 form-urlencoded 编码请求体
func (c *Client) doPost(ctx context.Context, endpoint string, body map[string]string) (json.RawMessage, error) {
	form := url.Values{}
	for k, v := range body {
		form.Set(k, v)
	}

	reqURL, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.doRequest(req)
}

// doDelete 发送 DELETE 请求到指定端点
func (c *Client) doDelete(ctx context.Context, endpoint string, params map[string]string) (json.RawMessage, error) {
	reqURL, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if len(params) > 0 {
		q := reqURL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		reqURL.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(req)

	return c.doRequest(req)
}

// doRequest 执行 HTTP 请求并解析 TAPD 统一响应格式
func (c *Client) doRequest(req *http.Request) (json.RawMessage, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		snippet := truncateUTF8(string(bodyBytes), 200)
		return nil, &TAPDError{
			HTTPStatus: resp.StatusCode,
			ExitCode:   mapHTTPError(resp.StatusCode),
			Message:    fmt.Sprintf("HTTP %d: %s", resp.StatusCode, snippet),
		}
	}

	var tapdResp model.TAPDResponse
	if err := json.Unmarshal(bodyBytes, &tapdResp); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}

	if tapdResp.Status != 1 {
		return nil, &TAPDError{
			HTTPStatus: resp.StatusCode,
			ExitCode:   4,
			Message:    fmt.Sprintf("TAPD API error: %s", tapdResp.Info),
		}
	}

	return tapdResp.Data, nil
}

// doPostJSON 发送 JSON 格式的 POST 请求到指定 URL（用于企业微信等外部 API）
func (c *Client) doPostJSON(ctx context.Context, rawURL string, body []byte) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rawURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("HTTP %d: (failed to read response body)", resp.StatusCode)
		}
		snippet := truncateUTF8(string(bodyBytes), 200)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, snippet)
	}
	return nil
}

// truncateUTF8 将字符串截断到最多 maxRunes 个 Unicode 字符，确保不会在多字节字符中间截断
func truncateUTF8(s string, maxRunes int) string {
	if utf8.RuneCountInString(s) <= maxRunes {
		return s
	}
	var size int
	for i := 0; i < maxRunes; i++ {
		_, runeSize := utf8.DecodeRuneInString(s[size:])
		size += runeSize
	}
	return s[:size]
}

// mapHTTPError 将 HTTP 状态码映射到退出码
func mapHTTPError(statusCode int) int {
	switch statusCode {
	case 401:
		return 1
	case 404:
		return 2
	case 422:
		return 3
	default:
		return 4
	}
}
