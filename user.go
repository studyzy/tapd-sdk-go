package tapd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/studyzy/tapd-sdk-go/model"
)

// GetPersonalSetting 获取用户个人配置
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/user/get_personal_setting.html
func (c *Client) GetPersonalSetting(ctx context.Context, req *model.GetPersonalSettingRequest) (*model.PersonalSetting, error) {
	data, err := c.doGet(ctx, "/users/get_personal_setting", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.PersonalSetting
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse personal setting: %w", err)
	}
	return &result, nil
}

// GetThirdUserMapping 获取用户关联的第三方系统的 user_id 和类型
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/user/get_third_user_mapping.html
func (c *Client) GetThirdUserMapping(ctx context.Context, req *model.GetThirdUserMappingRequest) (*model.ThirdUserMapping, error) {
	data, err := c.doGet(ctx, "/users/get_third_user_mapping", req.ToParams())
	if err != nil {
		return nil, err
	}

	var result model.ThirdUserMapping
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse third user mapping: %w", err)
	}
	return &result, nil
}
