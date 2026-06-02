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
// 用例仓库操作
// ---------------------------------------------------------------------------

// TestxCreateCaseRepo 创建用例仓库
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/create_repo.html
func (c *Client) TestxCreateCaseRepo(ctx context.Context, req *model.TestxCreateCaseRepoRequest) (*model.TestxCaseRepo, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos", req.Namespace)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxCaseRepo
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx repo: %w", err)
	}
	return &result, nil
}

// TestxUpdateCaseRepo 更新用例仓库
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/update_repo.html
func (c *Client) TestxUpdateCaseRepo(ctx context.Context, req *model.TestxUpdateCaseRepoRequest) (*model.TestxCaseRepo, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s", req.Namespace, req.RepoUid)
	resp, err := c.doTestxPut(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxCaseRepo
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx repo: %w", err)
	}
	return &result, nil
}

// TestxGetCaseRepo 获取用例仓库
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/get_repo.html
func (c *Client) TestxGetCaseRepo(ctx context.Context, req *model.TestxGetCaseRepoRequest) (*model.TestxCaseRepo, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s", req.Namespace, req.RepoUid)
	resp, err := c.doTestxGet(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	var result model.TestxCaseRepo
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx repo: %w", err)
	}
	return &result, nil
}

// TestxListCaseRepos 获取用例仓库列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/list_repo.html
func (c *Client) TestxListCaseRepos(ctx context.Context, req *model.TestxListCaseReposRequest) ([]model.TestxCaseRepo, int, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos", req.Namespace)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	if req.Search != "" {
		q.Set("Search", req.Search)
	}
	if req.Reverse {
		q.Set("Reverse", "true")
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxCaseRepo
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx repos: %w", err)
	}
	return result, resp.TotalCount, nil
}

// ---------------------------------------------------------------------------
// 用例目录操作
// ---------------------------------------------------------------------------

// TestxCreateCaseFolder 创建用例目录
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/create_folder.html
func (c *Client) TestxCreateCaseFolder(ctx context.Context, req *model.TestxCreateCaseFolderRequest) (*model.TestxCaseFolder, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/folders",
		req.Namespace, req.RepoUid, req.RepoVersionUid)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxCaseFolder
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx folder: %w", err)
	}
	return &result, nil
}

// TestxUpdateCaseFolder 更新用例目录
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/update_folder.html
func (c *Client) TestxUpdateCaseFolder(ctx context.Context, req *model.TestxUpdateCaseFolderRequest) (*model.TestxCaseFolder, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/folders/%s",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.FolderUid)
	resp, err := c.doTestxPut(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxCaseFolder
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx folder: %w", err)
	}
	return &result, nil
}

// ---------------------------------------------------------------------------
// 用例 CRUD
// ---------------------------------------------------------------------------

// TestxCreateCase 创建用例
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/create_case.html
func (c *Client) TestxCreateCase(ctx context.Context, req *model.TestxCreateCaseRequest) (*model.TestxCase, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases",
		req.Namespace, req.RepoUid, req.RepoVersionUid)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxCase
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx case: %w", err)
	}
	return &result, nil
}

// TestxBatchCreateCases 批量创建用例
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/batch_create_case.html
func (c *Client) TestxBatchCreateCases(ctx context.Context, req *model.TestxBatchCreateCasesRequest) (*model.TestxCase, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/batch-create",
		req.Namespace, req.RepoUid, req.RepoVersionUid)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxCase
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx case: %w", err)
	}
	return &result, nil
}

// TestxUpdateCase 更新用例
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/update_case.html
func (c *Client) TestxUpdateCase(ctx context.Context, req *model.TestxUpdateCaseRequest) (*model.TestxCase, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/%s",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.CaseUid)
	resp, err := c.doTestxPut(ctx, path, req)
	if err != nil {
		return nil, err
	}
	var result model.TestxCase
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx case: %w", err)
	}
	return &result, nil
}

// TestxBatchUpdateCases 批量更新用例
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/batch_update_case.html
func (c *Client) TestxBatchUpdateCases(ctx context.Context, req *model.TestxBatchUpdateCasesRequest) error {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/batch-update",
		req.Namespace, req.RepoUid, req.RepoVersionUid)
	_, err := c.doTestxPostJSON(ctx, path, req)
	return err
}

// TestxSearchCases 搜索用例列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/search_case.html
func (c *Client) TestxSearchCases(ctx context.Context, req *model.TestxSearchCasesRequest) (*model.TestxSearchCasesResponse, int, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/search",
		req.Namespace, req.RepoUid, req.RepoVersionUid)
	resp, err := c.doTestxPostJSON(ctx, path, req)
	if err != nil {
		return nil, 0, err
	}
	var result model.TestxSearchCasesResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx search cases: %w", err)
	}
	return &result, resp.TotalCount, nil
}

// ---------------------------------------------------------------------------
// 用例历史/执行/评审
// ---------------------------------------------------------------------------

// TestxListCaseHistorys 获取用例变更历史
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/list_case_historys.html
func (c *Client) TestxListCaseHistorys(ctx context.Context, req *model.TestxListCaseHistorysRequest) ([]model.TestxCaseHistory, int, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/%s/history",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.CaseUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxCaseHistory
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx case historys: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxListCaseExecutions 获取用例执行记录
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/list_case_executions.html
func (c *Client) TestxListCaseExecutions(ctx context.Context, req *model.TestxListCaseExecutionsRequest) ([]model.TestxCaseExecution, int, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/%s/executions",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.CaseUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	if req.Ordering != "" {
		q.Set("Ordering", req.Ordering)
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxCaseExecution
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx case executions: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxListCaseReviews 获取用例评审记录
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/list_case_reviews.html
func (c *Client) TestxListCaseReviews(ctx context.Context, req *model.TestxListCaseReviewsRequest) ([]model.TestxCaseReview, int, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/%s/reviews",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.CaseUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	if req.Source != "" {
		q.Set("Source", req.Source)
	}
	if req.MainUid != "" {
		q.Set("MainUid", req.MainUid)
	}
	if req.SourceKind != "" {
		q.Set("SourceKind", req.SourceKind)
	}
	if req.SourceUid != "" {
		q.Set("SourceUid", req.SourceUid)
	}
	if req.IsLastReview {
		q.Set("IsLastReview", "true")
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxCaseReview
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx case reviews: %w", err)
	}
	return result, resp.TotalCount, nil
}

// ---------------------------------------------------------------------------
// 用例 Bug 关联
// ---------------------------------------------------------------------------

// TestxListCaseBugs 获取用例关联的 Bug 列表
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/list_case_bugs.html
func (c *Client) TestxListCaseBugs(ctx context.Context, req *model.TestxListCaseBugsRequest) ([]model.TestxCaseBugItem, int, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/%s/bugs",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.CaseUid)
	q := url.Values{}
	if req.PageInfo != nil {
		q.Set("PageInfo.Offset", strconv.FormatUint(uint64(req.PageInfo.Offset), 10))
		q.Set("PageInfo.Limit", strconv.FormatUint(uint64(req.PageInfo.Limit), 10))
	}
	if req.Status != "" {
		q.Set("Status", req.Status)
	}
	if req.Priority != "" {
		q.Set("Priority", req.Priority)
	}
	if req.Handler != "" {
		q.Set("Handler", req.Handler)
	}
	if req.Name != "" {
		q.Set("Name", req.Name)
	}
	resp, err := c.doTestxGet(ctx, path, q)
	if err != nil {
		return nil, 0, err
	}
	var result []model.TestxCaseBugItem
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, 0, fmt.Errorf("failed to parse testx case bugs: %w", err)
	}
	return result, resp.TotalCount, nil
}

// TestxBatchBindCaseBug 批量关联 Bug
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/batch_bind_bug.html
func (c *Client) TestxBatchBindCaseBug(ctx context.Context, req *model.TestxBatchBindCaseBugRequest) error {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/%s/bugs/batch-bind",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.CaseUid)
	_, err := c.doTestxPostJSON(ctx, path, req)
	return err
}

// TestxBatchUnbindCaseBug 批量解绑 Bug
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/batch_unbind_bug.html
func (c *Client) TestxBatchUnbindCaseBug(ctx context.Context, req *model.TestxBatchUnbindCaseBugRequest) error {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/repos/%s/versions/%s/cases/%s/bugs/batch-unbind",
		req.Namespace, req.RepoUid, req.RepoVersionUid, req.CaseUid)
	_, err := c.doTestxPostJSON(ctx, path, req)
	return err
}

// ---------------------------------------------------------------------------
// 用例模板
// ---------------------------------------------------------------------------

// TestxListCaseTemplates 获取用例模板
// https://open.tapd.cn/document/api-doc/API文档/api_reference/testx/case/list_case_template.html
func (c *Client) TestxListCaseTemplates(ctx context.Context, req *model.TestxListCaseTemplatesRequest) ([]model.TestxCaseTemplate, error) {
	path := fmt.Sprintf("/api/testx/case/v1/namespaces/%s/case-templates", req.Namespace)
	resp, err := c.doTestxGet(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	var result []model.TestxCaseTemplate
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse testx case templates: %w", err)
	}
	return result, nil
}
