package tapd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/studyzy/tapd-sdk-go/model"
)

// doTestxRequest 执行 HTTP 请求并解析 testx 统一响应格式
func (c *Client) doTestxRequest(req *http.Request) (*model.TestxResponse, error) {
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

	var testxResp model.TestxResponse
	if err := json.Unmarshal(bodyBytes, &testxResp); err != nil {
		return nil, fmt.Errorf("failed to parse testx response JSON: %w", err)
	}

	if string(testxResp.Error) != "null" && len(testxResp.Error) > 0 {
		return nil, &TAPDError{
			HTTPStatus: resp.StatusCode,
			ExitCode:   4,
			Message:    fmt.Sprintf("testx API error: %s", string(testxResp.Error)),
		}
	}

	return &testxResp, nil
}

// doTestxPostJSON 发送 JSON 格式的 POST 请求到 testx API 端点
func (c *Client) doTestxPostJSON(ctx context.Context, path string, body interface{}) (*model.TestxResponse, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	reqURL, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(req)
	req.Header.Set("Content-Type", "application/json")

	return c.doTestxRequest(req)
}

// doTestxGet 发送 GET 请求到 testx API 端点
func (c *Client) doTestxGet(ctx context.Context, path string, query url.Values) (*model.TestxResponse, error) {
	reqURL, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if len(query) > 0 {
		reqURL.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(req)

	return c.doTestxRequest(req)
}

// doTestxPut 发送 JSON 格式的 PUT 请求到 testx API 端点
func (c *Client) doTestxPut(ctx context.Context, path string, body interface{}) (*model.TestxResponse, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	reqURL, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, reqURL.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(req)
	req.Header.Set("Content-Type", "application/json")

	return c.doTestxRequest(req)
}

// doTestxDelete 发送 DELETE 请求到 testx API 端点
func (c *Client) doTestxDelete(ctx context.Context, path string) (*model.TestxResponse, error) {
	reqURL, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	c.applyAuth(req)

	return c.doTestxRequest(req)
}
