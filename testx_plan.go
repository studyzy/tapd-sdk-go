package tapd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/studyzy/tapd-sdk-go/model"
)

// ---------------------------------------------------------------------------
// 测试计划目录操作
// ---------------------------------------------------------------------------

// TestxCreatePlanFolder 创建计划目录
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/create_folder.html
func (c *Client) TestxCreatePlanFolder(ctx context.Context, req *model.TestxCreatePlanFolderRequest) (*model.TestxPlanFolder, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/folders", req.Namespace)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxPlanFolder
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx plan folder: %w", err)
	}
	return &result, nil
}

// TestxUpdatePlanFolder 更新计划目录
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/update_folder.html
func (c *Client) TestxUpdatePlanFolder(ctx context.Context, req *model.TestxUpdatePlanFolderRequest) (*model.TestxPlanFolder, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/folders/%s", req.Namespace, req.FolderUid)
	resp, err := c.doTestxPut(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxPlanFolder
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx plan folder: %w", err)
	}
	return &result, nil
}

// TestxListFolderChildren 获取计划目录子信息
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_folder_children.html
func (c *Client) TestxListFolderChildren(ctx context.Context, req *model.TestxListFolderChildrenRequest) (*model.TestxFolderChildrenResult, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/folders/children", req.Namespace)
	q := url.Values{}
	q.Set("Uid", req.Uid)
	if req.WithDescendant {
		q.Set("WithDescendant", "true")
	}
	if req.WithAncestor {
		q.Set("WithAncestor", "true")
	}
	if req.Name != "" {
		q.Set("Name", req.Name)
	}
	if req.PlanArchive != "" {
		q.Set("PlanArchive", req.PlanArchive)
	}
	if len(req.PlanStates) > 0 {
		for _, s := range req.PlanStates {
			q.Add("PlanStates", s)
		}
	}
	if req.ItemType != "" {
		q.Set("ItemType", req.ItemType)
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, err
	}
	var result model.TestxFolderChildrenResult
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx folder children: %w", err)
	}
	return &result, nil
}

// ---------------------------------------------------------------------------
// 测试计划操作
// ---------------------------------------------------------------------------

// TestxGetPlan 获取计划详情
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/get_plan.html
func (c *Client) TestxGetPlan(ctx context.Context, req *model.TestxGetPlanRequest) (*model.TestxPlan, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s", req.Namespace, req.Uid)
	q := url.Values{}
	if req.WithStatistic {
		q.Set("WithStatistic", "true")
	}
	if req.WithDetail != "" {
		q.Set("WithDetail", req.WithDetail)
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, err
	}
	var result model.TestxPlan
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx plan: %w", err)
	}
	return &result, nil
}

// TestxCreatePlan 创建计划
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/create_plan.html
func (c *Client) TestxCreatePlan(ctx context.Context, req *model.TestxCreatePlanRequest) (*model.TestxPlan, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans", req.Namespace)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxPlan
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx plan: %w", err)
	}
	return &result, nil
}

// TestxUpdatePlan 更新计划
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/update_plan.html
func (c *Client) TestxUpdatePlan(ctx context.Context, req *model.TestxUpdatePlanRequest) (*model.TestxPlan, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans", req.Namespace)
	resp, err := c.doTestxPut(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxPlan
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx plan: %w", err)
	}
	return &result, nil
}

// TestxUpdatePlanTargetScope 更新计划范围和目标
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/update_plan_target_scope.html
func (c *Client) TestxUpdatePlanTargetScope(ctx context.Context, req *model.TestxUpdatePlanTargetScopeRequest) (*model.TestxPlan, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/target-scope", req.Namespace, req.Uid)
	resp, err := c.doTestxPut(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxPlan
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx plan: %w", err)
	}
	return &result, nil
}

// TestxBatchUpdatePlanCase 批量更新计划用例
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/batch_update_case.html
func (c *Client) TestxBatchUpdatePlanCase(ctx context.Context, req *model.TestxBatchUpdatePlanCaseRequest) error {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/cases/batch-update", req.Namespace, req.PlanUid)
	_, err := c.doTestxPostJSON(ctx, path, req)
	return err
}

// TestxBatchArchivePlan 批量归档计划
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/batch_archive.html
func (c *Client) TestxBatchArchivePlan(ctx context.Context, req *model.TestxBatchArchivePlanRequest) error {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/batch-archive", req.Namespace)
	_, err := c.doTestxPut(ctx, path, req)
	return err
}

// TestxListPlans 获取目录下计划列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_plans.html
func (c *Client) TestxListPlans(ctx context.Context, req *model.TestxListPlansRequest) ([]model.TestxPlan, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/folders/%s/plans-list", req.Namespace, req.FolderUid)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxPlan
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plans: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxListPlanCases 获取计划下用例列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_plan_cases.html
func (c *Client) TestxListPlanCases(ctx context.Context, req *model.TestxListPlanCasesRequest) (*model.TestxPlanCasesResult, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/cases-search", req.Namespace, req.Uid)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, 0, err
	}
	var result model.TestxPlanCasesResult
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan cases: %w", err)
	}
	return &result, resp.TotalCount, nil
}

// TestxListPlanHistories 获取计划变更历史
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_plan_histories.html
func (c *Client) TestxListPlanHistories(ctx context.Context, req *model.TestxListPlanHistoriesRequest) ([]model.TestxPlanHistory, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/histories", req.Namespace, req.PlanUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxPlanHistory
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan histories: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxPlanStatistics 获取计划统计信息
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/statistics.html
func (c *Client) TestxPlanStatistics(ctx context.Context, req *model.TestxPlanStatisticsRequest) ([]model.TestxPlanStatisticsItem, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/statistics", req.Namespace)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxPlanStatisticsItem
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan statistics: %w", err)
	}
	return result, resp.TotalCount, nil
}

// ---------------------------------------------------------------------------
// 计划用例缺陷关联
// ---------------------------------------------------------------------------

// TestxBatchBindPlanBug 计划用例批量添加缺陷
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/batch_bind_bug.html
func (c *Client) TestxBatchBindPlanBug(ctx context.Context, req *model.TestxBatchBindPlanBugRequest) error {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/cases/bugs", req.Namespace, req.PlanUid)
	_, err := c.doTestxPostJSON(ctx, path, req)
	return err
}

// TestxUnbindPlanBug 移除计划用例关联的缺陷
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/unbind_bug.html
func (c *Client) TestxUnbindPlanBug(ctx context.Context, req *model.TestxUnbindPlanBugRequest) error {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/cases/%s/issues/%s",
		req.Namespace, req.PlanUid, req.CaseUid, req.IssueUid)
	_, err := c.doTestxDelete(ctx, path)
	return err
}

// TestxListPlanCaseIssues 获取计划下用例关联的缺陷列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_case_issue.html
func (c *Client) TestxListPlanCaseIssues(ctx context.Context, req *model.TestxListPlanCaseIssuesRequest) ([]model.TestxIssue, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/cases/%s/issues",
		req.Namespace, req.PlanUid, req.CaseUid)
	q := url.Values{}
	q.Set("Type", req.Type)
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxIssue
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan case issues: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxListPlanCaseEvents 获取计划下用例的事件列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_case_events.html
func (c *Client) TestxListPlanCaseEvents(ctx context.Context, req *model.TestxListPlanCaseEventsRequest) ([]model.TestxPlanCaseEvent, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/cases/%s/events",
		req.Namespace, req.PlanUid, req.CaseUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxPlanCaseEvent
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan case events: %w", err)
	}
	return result, resp.TotalCount, nil
}

// ---------------------------------------------------------------------------
// 计划缺陷和需求查询
// ---------------------------------------------------------------------------

// TestxListPlanBugs 获取计划关联的缺陷列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_plan_bugs.html
func (c *Client) TestxListPlanBugs(ctx context.Context, req *model.TestxListPlanBugsRequest) ([]model.TestxPlanBug, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/bugs", req.Namespace, req.PlanUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	for _, rt := range req.RelatedTypes {
		q.Add("RelatedTypes", rt)
	}
	if req.Status != "" {
		q.Set("Status", req.Status)
	}
	if req.Summary != "" {
		q.Set("Summary", req.Summary)
	}
	if req.BugId != "" {
		q.Set("BugId", req.BugId)
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxPlanBug
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan bugs: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxListPlanBugStatistics 批量查询计划关联缺陷统计数
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_plan_bug_statistics.html
func (c *Client) TestxListPlanBugStatistics(ctx context.Context, req *model.TestxListPlanBugStatisticsRequest) ([]model.TestxPlanBugStatistics, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/bug-statistics", req.Namespace)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result []model.TestxPlanBugStatistics
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx plan bug statistics: %w", err)
	}
	return result, nil
}

// TestxListPlanStories 获取计划关联需求列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_plan_stories.html
func (c *Client) TestxListPlanStories(ctx context.Context, req *model.TestxListPlanStoriesRequest) ([]model.TestxIssue, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plans/%s/stories", req.Namespace, req.PlanUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxIssue
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan stories: %w", err)
	}
	return result, resp.TotalCount, nil
}

// ---------------------------------------------------------------------------
// 计划模板
// ---------------------------------------------------------------------------

// TestxListPlanTemplates 获取计划模板
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/plan/list_plan_template.html
func (c *Client) TestxListPlanTemplates(ctx context.Context, req *model.TestxListPlanTemplatesRequest) ([]model.TestxPlanTemplate, int, error) {
	path := fmt.Sprintf("/api/testx/plan/v1/namespaces/%s/plan-templates", req.Namespace)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxPlanTemplate
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx plan templates: %w", err)
	}
	return result, resp.TotalCount, nil
}
