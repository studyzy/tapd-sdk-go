package tapd

import (
	"encoding/json"
	"fmt"
)

// ParseList 解析 TAPD 列表响应格式 [{"Key": {...}}, ...]
// key 为 TAPD 响应中的包裹键名，如 "Story"、"Bug" 等。
// 不含指定 key 的条目会被跳过；含指定 key 但反序列化失败的条目会返回错误。
func ParseList[T any](data json.RawMessage, key string) ([]T, error) {
	return parseList[T](data, key)
}

// parseList 解析 TAPD 列表响应格式 [{"Key": {...}}, ...]
// key 为 TAPD 响应中的包裹键名，如 "Story"、"Bug" 等。
// 不含指定 key 的条目会被跳过；含指定 key 但反序列化失败的条目会返回错误。
func parseList[T any](data json.RawMessage, key string) ([]T, error) {
	var rawList []map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, fmt.Errorf("failed to parse list response: %w", err)
	}
	results := make([]T, 0, len(rawList))
	for i, item := range rawList {
		raw, ok := item[key]
		if !ok {
			continue
		}
		var v T
		if err := json.Unmarshal(raw, &v); err != nil {
			return nil, fmt.Errorf("failed to parse %s at index %d: %w", key, i, err)
		}
		results = append(results, v)
	}
	return results, nil
}

// parseOne 解析 TAPD 单实体响应格式 {"Key": {...}}
// key 为 TAPD 响应中的包裹键名，如 "Story"、"Bug" 等
func parseOne[T any](data json.RawMessage, key string) (*T, error) {
	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	raw, ok := wrapper[key]
	if !ok {
		return nil, fmt.Errorf("unexpected response format: missing %q key", key)
	}
	var v T
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", key, err)
	}
	return &v, nil
}

// parseCount 解析 TAPD 计数响应格式 {"count": N}
func parseCount(data json.RawMessage) (int, error) {
	var result map[string]int
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("failed to parse count response: %w", err)
	}
	if count, ok := result["count"]; ok {
		return count, nil
	}
	return 0, fmt.Errorf("unexpected count response format: missing 'count' key")
}
