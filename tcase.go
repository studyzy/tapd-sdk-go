package tapd

import (
	"context"
	"encoding/json"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ListTCases 查询测试用例列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcases.html
func (c *Client) ListTCases(ctx context.Context, req *model.ListTCasesRequest) ([]model.TCase, error) {
	data, err := c.doGet(ctx, "/tcases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.TCase](data, "Tcase")
}

// CountTCases 查询测试用例数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcases_count.html
func (c *Client) CountTCases(ctx context.Context, req *model.CountTCasesRequest) (int, error) {
	data, err := c.doGet(ctx, "/tcases/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// CreateTCase 创建测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_tcase.html
func (c *Client) CreateTCase(ctx context.Context, req *model.CreateTCaseRequest) (*model.TCase, error) {
	data, err := c.doPost(ctx, "/tcases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.TCase](data, "Tcase")
}

// BatchCreateTCases 批量创建测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/batch_add_tcase.html
func (c *Client) BatchCreateTCases(ctx context.Context, req *model.BatchCreateTCasesRequest) ([]model.TCase, error) {
	body, err := req.ToJSON()
	if err != nil {
		return nil, err
	}
	data, err := c.doPostJSONBody(ctx, "/tcases/batch_save", body)
	if err != nil {
		return nil, err
	}
	return parseList[model.TCase](data, "Tcase")
}

// UpdateTCase 更新测试用例
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/update_tcase.html
func (c *Client) UpdateTCase(ctx context.Context, req *model.UpdateTCaseRequest) (*model.TCase, error) {
	data, err := c.doPost(ctx, "/tcases", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.TCase](data, "Tcase")
}

// ListTCaseCategories 查询测试用例目录列表
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_categories.html
func (c *Client) ListTCaseCategories(ctx context.Context, req *model.ListTCaseCategoriesRequest) ([]model.TCaseCategory, error) {
	data, err := c.doGet(ctx, "/tcase_categories", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseList[model.TCaseCategory](data, "TcaseCategory")
}

// CountTCaseCategories 查询测试用例目录数量
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_categories_count.html
func (c *Client) CountTCaseCategories(ctx context.Context, req *model.ListTCaseCategoriesRequest) (int, error) {
	data, err := c.doGet(ctx, "/tcase_categories/count", req.ToParams())
	if err != nil {
		return 0, err
	}

	return parseCount(data)
}

// CreateTCaseCategory 创建测试用例目录
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_tcase_category.html
func (c *Client) CreateTCaseCategory(ctx context.Context, req *model.CreateTCaseCategoryRequest) (*model.TCaseCategory, error) {
	data, err := c.doPost(ctx, "/tcase_categories", req.ToParams())
	if err != nil {
		return nil, err
	}

	return parseOne[model.TCaseCategory](data, "TcaseCategory")
}

// GetStoryByTCaseID 获取测试用例关联的需求
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_story_by_tcase_id.html
func (c *Client) GetStoryByTCaseID(ctx context.Context, req *model.GetStoryByTCaseIDRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/tcases/get_story_by_tcase_id", req.ToParams())
}

// GetTCaseCustomFieldsSettings 获取测试用例自定义字段配置
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_custom_fields_settings.html
func (c *Client) GetTCaseCustomFieldsSettings(ctx context.Context, req *model.WorkspaceIDRequest) ([]model.CustomFieldConfig, error) {
	data, err := c.doGet(ctx, "/tcases/custom_fields_settings", req.ToParams())
	if err != nil {
		return nil, err
	}
	return parseList[model.CustomFieldConfig](data, "CustomFieldConfig")
}

// GetTCaseFieldsInfo 获取测试用例所有字段及候选值
// API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_fields_info.html
func (c *Client) GetTCaseFieldsInfo(ctx context.Context, req *model.WorkspaceIDRequest) (json.RawMessage, error) {
	return c.doGet(ctx, "/tcases/get_fields_info", req.ToParams())
}
