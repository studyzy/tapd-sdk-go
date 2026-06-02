package model

// CodeCommitRelated 表示 commit 关联的业务对象
type CodeCommitRelated struct {
	Type        string `json:"type,omitempty"`
	ObjectID    string `json:"object_id,omitempty"`
	CommitID    string `json:"commit_id,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Code        string `json:"code,omitempty"`
}

// CodeCommitInfo 表示 GIT 提交信息
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/get_code_commit_infos.html
type CodeCommitInfo struct {
	ID              string              `json:"id,omitempty"`
	UserName        string              `json:"user_name,omitempty"`
	UserID          string              `json:"user_id,omitempty"`
	HookUserName    string              `json:"hook_user_name,omitempty"`
	CommitID        string              `json:"commit_id,omitempty"`
	WorkspaceID     string              `json:"workspace_id,omitempty"`
	Message         string              `json:"message,omitempty"`
	Author          string              `json:"author,omitempty"`
	Path            string              `json:"path,omitempty"`
	WebURL          string              `json:"web_url,omitempty"`
	HookProjectName string              `json:"hook_project_name,omitempty"`
	CommitTime      string              `json:"commit_time,omitempty"`
	Created         string              `json:"created,omitempty"`
	Ref             string              `json:"ref,omitempty"`
	RefStatus       string              `json:"ref_status,omitempty"`
	GitEnv          string              `json:"git_env,omitempty"`
	FileCommit      string              `json:"file_commit,omitempty"`
	RepoID          string              `json:"repo_id,omitempty"`
	BranchID        string              `json:"branch_id,omitempty"`
	FileSort        map[string]int      `json:"file_sort,omitempty"`
	Related         []CodeCommitRelated `json:"related,omitempty"`
}

// GetCodeCommitObjectsRequest 获取指定 commit 关联的 TAPD 业务对象的请求参数
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/source/get_commit_objects.html
type GetCodeCommitObjectsRequest struct {
	WorkspaceID string // 必填：项目 ID
	CommitID    string // 必填：提交 ID，支持多个，以英文逗号分隔
	EntityType  string // 必填：业务对象类型，story/bug/task
	ScmType     string // 可选：来源类型，p4/tgit/gitlab/github 等
	Limit       int    // 可选：返回数量限制，默认 30，最大 200
	Page        int    // 可选：页码，默认 1
	Order       string // 可选：排序规则，例如 "created desc"
	Fields      string // 可选：返回字段，多个以英文逗号分隔
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *GetCodeCommitObjectsRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id": r.WorkspaceID,
		"commit_id":    r.CommitID,
		"entity_type":  r.EntityType,
	}
	setOptional(params, "scm_type", r.ScmType)
	setOptionalInt(params, "limit", r.Limit)
	setOptionalInt(params, "page", r.Page)
	setOptional(params, "order", r.Order)
	setOptional(params, "fields", r.Fields)
	return params
}
