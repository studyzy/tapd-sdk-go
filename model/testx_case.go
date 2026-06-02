package model

// TestxCaseRepo testx 用例仓库
type TestxCaseRepo struct {
	Audit       *TestxAudit            `json:"Audit,omitempty"`
	Namespace   string                 `json:"Namespace,omitempty"`
	Uid         string                 `json:"Uid,omitempty"`
	Name        string                 `json:"Name,omitempty"`
	Description string                 `json:"Description,omitempty"`
	Versions    []TestxCaseRepoVersion `json:"Versions,omitempty"`
	Nid         string                 `json:"Nid,omitempty"`
	Type        string                 `json:"Type,omitempty"`
}

// TestxCaseRepoVersion testx 用例仓库版本
type TestxCaseRepoVersion struct {
	Audit       *TestxAudit    `json:"Audit,omitempty"`
	Uid         string         `json:"Uid,omitempty"`
	Repo        *TestxCaseRepo `json:"Repo,omitempty"`
	Name        string         `json:"Name,omitempty"`
	Description string         `json:"Description,omitempty"`
}

// TestxCaseFolder testx 用例目录
type TestxCaseFolder struct {
	Audit          *TestxAudit       `json:"Audit,omitempty"`
	Uid            string            `json:"Uid,omitempty"`
	UUID           string            `json:"UUID,omitempty"`
	RepoUid        string            `json:"RepoUid,omitempty"`
	RepoVersionUid string            `json:"RepoVersionUid,omitempty"`
	FolderUid      string            `json:"FolderUid,omitempty"`
	FullPath       string            `json:"FullPath,omitempty"`
	Name           string            `json:"Name,omitempty"`
	Owners         []string          `json:"Owners,omitempty"`
	Description    string            `json:"Description,omitempty"`
	CaseCount      int               `json:"CaseCount,omitempty"`
	Path           string            `json:"Path,omitempty"`
	Folders        []TestxCaseFolder `json:"Folders,omitempty"`
	Cases          []TestxCase       `json:"Cases,omitempty"`
	ArchiveUid     string            `json:"ArchiveUid,omitempty"`
}

// TestxCaseStep testx 用例步骤
type TestxCaseStep struct {
	Id             string `json:"Id,omitempty"`
	Content        string `json:"Content,omitempty"`
	ExpectedResult string `json:"ExpectedResult,omitempty"`
	NID            string `json:"NID,omitempty"`
}

// TestxCaseAttachment testx 用例附件
type TestxCaseAttachment struct {
	Id       string `json:"Id,omitempty"`
	Key      string `json:"Key,omitempty"`
	FileName string `json:"FileName,omitempty"`
	Size     string `json:"Size,omitempty"`
}

// TestxCase testx 用例
type TestxCase struct {
	Audit            *TestxAudit           `json:"Audit,omitempty"`
	Uid              string                `json:"Uid,omitempty"`
	UUID             string                `json:"UUID,omitempty"`
	RepoUid          string                `json:"RepoUid,omitempty"`
	RepoVersionUid   string                `json:"RepoVersionUid,omitempty"`
	FolderUid        string                `json:"FolderUid,omitempty"`
	FullPath         string                `json:"FullPath,omitempty"`
	Name             string                `json:"Name,omitempty"`
	Description      string                `json:"Description,omitempty"`
	Priority         string                `json:"Priority,omitempty"`
	PreConditions    string                `json:"PreConditions,omitempty"`
	Type             string                `json:"Type,omitempty"`
	StepType         string                `json:"StepType,omitempty"`
	Steps            []TestxCaseStep       `json:"Steps,omitempty"`
	StepText         *TestxCaseStep        `json:"StepText,omitempty"`
	Attachments      []TestxCaseAttachment `json:"Attachments,omitempty"`
	CustomFields     []TestxProperty       `json:"CustomFields,omitempty"`
	Labels           []TestxLabel          `json:"Labels,omitempty"`
	Source           string                `json:"Source,omitempty"`
	IsManualRelation bool                  `json:"IsManualRelation,omitempty"`
	Issues           []TestxIssue          `json:"Issues,omitempty"`
	Owners           []string              `json:"Owners,omitempty"`
	ManHourEstimated string                `json:"ManHourEstimated,omitempty"`
	Path             string                `json:"Path,omitempty"`
	RunTimes         string                `json:"RunTimes,omitempty"`
	IsFolder         bool                  `json:"IsFolder,omitempty"`
	ReviewAt         string                `json:"ReviewAt,omitempty"`
	ReviewState      string                `json:"ReviewState,omitempty"`
	Nid              string                `json:"Nid,omitempty"`
	BugCount         string                `json:"BugCount,omitempty"`
	Bugs             []interface{}         `json:"Bugs,omitempty"`
	Executions       []interface{}         `json:"Executions,omitempty"`
	Reviews          []interface{}         `json:"Reviews,omitempty"`
	CaseBug          interface{}           `json:"CaseBug,omitempty"`
}

// TestxCaseHistory testx 用例变更历史
type TestxCaseHistory struct {
	Audit          *TestxAudit              `json:"Audit,omitempty"`
	Uid            string                   `json:"Uid,omitempty"`
	ChangeType     string                   `json:"ChangeType,omitempty"`
	ChangeContents []TestxCaseChangeContent `json:"ChangeContents,omitempty"`
}

// TestxCaseChangeContent testx 用例变更内容
type TestxCaseChangeContent struct {
	FieldName string `json:"FieldName,omitempty"`
	PreValue  string `json:"PreValue,omitempty"`
	PostValue string `json:"PostValue,omitempty"`
}

// TestxCaseExecution testx 用例执行记录
type TestxCaseExecution struct {
	SourceName       string      `json:"SourceName,omitempty"`
	SourceUid        string      `json:"SourceUid,omitempty"`
	Executor         string      `json:"Executor,omitempty"`
	StartExecuteTime string      `json:"StartExecuteTime,omitempty"`
	EndExecuteTime   string      `json:"EndExecuteTime,omitempty"`
	ExecuteState     string      `json:"ExecuteState,omitempty"`
	Message          string      `json:"Message,omitempty"`
	BugCount         string      `json:"BugCount,omitempty"`
	LinkData         interface{} `json:"LinkData,omitempty"`
	Source           string      `json:"Source,omitempty"`
}

// TestxCaseReview testx 用例评审记录
type TestxCaseReview struct {
	SourceName  string      `json:"SourceName,omitempty"`
	SourceUid   string      `json:"SourceUid,omitempty"`
	Reviewer    string      `json:"Reviewer,omitempty"`
	ReviewTime  string      `json:"ReviewTime,omitempty"`
	ReviewState string      `json:"ReviewState,omitempty"`
	Message     string      `json:"Message,omitempty"`
	LinkData    interface{} `json:"LinkData,omitempty"`
	Source      string      `json:"Source,omitempty"`
	Uid         string      `json:"Uid,omitempty"`
	MainUid     string      `json:"MainUid,omitempty"`
	SourceKind  string      `json:"SourceKind,omitempty"`
	Total       uint32      `json:"Total,omitempty"`
	CaseUid     string      `json:"CaseUid,omitempty"`
}

// TestxBugStatusField testx 缺陷状态字段
type TestxBugStatusField struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}

// TestxBugPriorityField testx 缺陷优先级字段
type TestxBugPriorityField struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}

// TestxBugSeverityField testx 缺陷严重程度字段
type TestxBugSeverityField struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}

// TestxBugResolutionField testx 缺陷解决方法字段
type TestxBugResolutionField struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}

// TestxBugAttachment testx 缺陷附件
type TestxBugAttachment struct {
	Id          string `json:"Id,omitempty"`
	DownloadUrl string `json:"DownloadUrl,omitempty"`
	Type        string `json:"Type,omitempty"`
	EntryId     string `json:"EntryId,omitempty"`
	Filename    string `json:"Filename,omitempty"`
	Description string `json:"Description,omitempty"`
	ContentType string `json:"ContentType,omitempty"`
	Created     string `json:"Created,omitempty"`
	WorkspaceId string `json:"WorkspaceId,omitempty"`
	Owner       string `json:"Owner,omitempty"`
}

// TestxBug testx 缺陷信息
type TestxBug struct {
	WorkspaceId                 string                   `json:"WorkspaceId,omitempty"`
	IterationId                 string                   `json:"IterationId,omitempty"`
	Id                          string                   `json:"Id,omitempty"`
	Summary                     string                   `json:"Summary,omitempty"`
	Creator                     string                   `json:"Creator,omitempty"`
	Priority                    *TestxBugPriorityField   `json:"Priority,omitempty"`
	Created                     string                   `json:"Created,omitempty"`
	Status                      *TestxBugStatusField     `json:"Status,omitempty"`
	LastUpdater                 string                   `json:"LastUpdater,omitempty"`
	UpdateTime                  string                   `json:"UpdateTime,omitempty"`
	Severity                    *TestxBugSeverityField   `json:"Severity,omitempty"`
	Owners                      string                   `json:"Owners,omitempty"`
	Begin                       string                   `json:"Begin,omitempty"`
	Due                         string                   `json:"Due,omitempty"`
	NamespaceId                 string                   `json:"NamespaceId,omitempty"`
	Description                 string                   `json:"Description,omitempty"`
	Url                         string                   `json:"Url,omitempty"`
	Developers                  string                   `json:"Developers,omitempty"`
	Testers                     string                   `json:"Testers,omitempty"`
	Module                      string                   `json:"Module,omitempty"`
	Attachments                 []TestxBugAttachment     `json:"Attachments,omitempty"`
	Tags                        []string                 `json:"Tags,omitempty"`
	RelatedType                 string                   `json:"RelatedType,omitempty"`
	IconType                    string                   `json:"IconType,omitempty"`
	Version                     []string                 `json:"Version,omitempty"`
	ReleaseId                   string                   `json:"ReleaseId,omitempty"`
	VersionReport               string                   `json:"VersionReport,omitempty"`
	VersionTest                 string                   `json:"VersionTest,omitempty"`
	VersionFix                  string                   `json:"VersionFix,omitempty"`
	VersionClose                string                   `json:"VersionClose,omitempty"`
	BaselineFind                string                   `json:"BaselineFind,omitempty"`
	BaselineJoin                string                   `json:"BaselineJoin,omitempty"`
	BaselineClose               string                   `json:"BaselineClose,omitempty"`
	BaselineTest                string                   `json:"BaselineTest,omitempty"`
	CC                          string                   `json:"CC,omitempty"`
	Participator                string                   `json:"Participator,omitempty"`
	Auditer                     string                   `json:"Auditer,omitempty"`
	Confirmer                   string                   `json:"Confirmer,omitempty"`
	Fixer                       string                   `json:"Fixer,omitempty"`
	Closer                      string                   `json:"Closer,omitempty"`
	ReopenTime                  string                   `json:"ReopenTime,omitempty"`
	InProgressTime              string                   `json:"InProgressTime,omitempty"`
	Resolved                    string                   `json:"Resolved,omitempty"`
	VerifyTime                  string                   `json:"VerifyTime,omitempty"`
	Closed                      string                   `json:"Closed,omitempty"`
	RejectTime                  string                   `json:"RejectTime,omitempty"`
	Modified                    string                   `json:"Modified,omitempty"`
	Deadline                    string                   `json:"Deadline,omitempty"`
	Os                          string                   `json:"Os,omitempty"`
	Platform                    string                   `json:"Platform,omitempty"`
	Testmode                    string                   `json:"Testmode,omitempty"`
	Testphase                   string                   `json:"Testphase,omitempty"`
	Testtype                    string                   `json:"Testtype,omitempty"`
	BugSource                   string                   `json:"Source,omitempty"`
	Frequency                   string                   `json:"Frequency,omitempty"`
	Originphase                 string                   `json:"Originphase,omitempty"`
	Sourcephase                 string                   `json:"Sourcephase,omitempty"`
	Resolution                  *TestxBugResolutionField `json:"Resolution,omitempty"`
	Estimate                    string                   `json:"Estimate,omitempty"`
	Lastmodify                  string                   `json:"Lastmodify,omitempty"`
	CustomFields                map[string]string        `json:"CustomFields,omitempty"`
	TemplateId                  string                   `json:"TemplateId,omitempty"`
	IsApplyTemplateDefaultValue bool                     `json:"IsApplyTemplateDefaultValue,omitempty"`
	CustomPlanFields            map[string]string        `json:"CustomPlanFields,omitempty"`
	Feature                     string                   `json:"Feature,omitempty"`
	Effort                      string                   `json:"Effort,omitempty"`
	Bugtype                     string                   `json:"Bugtype,omitempty"`
	Size                        string                   `json:"Size,omitempty"`
}

// TestxCaseBugSource testx 用例缺陷来源
type TestxCaseBugSource struct {
	Uid   string `json:"Uid,omitempty"`
	Type  string `json:"Type,omitempty"`
	Name  string `json:"Name,omitempty"`
	Label string `json:"Label,omitempty"`
	Url   string `json:"Url,omitempty"`
}

// TestxCaseBugItem testx 用例关联缺陷条目
type TestxCaseBugItem struct {
	Bug     *TestxBug            `json:"Bug,omitempty"`
	Sources []TestxCaseBugSource `json:"Sources,omitempty"`
}

// TestxValueChoice testx 模板字段可选值
type TestxValueChoice struct {
	ID    uint32      `json:"ID,omitempty"`
	Label string      `json:"Label,omitempty"`
	Value interface{} `json:"Value,omitempty"`
	Order uint32      `json:"Order,omitempty"`
}

// TestxPropertyDefinition testx 模板属性定义
type TestxPropertyDefinition struct {
	Widget        string                    `json:"Widget,omitempty"`
	Type          string                    `json:"Type,omitempty"`
	Uneditable    bool                      `json:"Uneditable,omitempty"`
	Name          string                    `json:"Name,omitempty"`
	Label         string                    `json:"Label,omitempty"`
	Description   string                    `json:"Description,omitempty"`
	Help          string                    `json:"Help,omitempty"`
	Visible       bool                      `json:"Visible,omitempty"`
	Required      bool                      `json:"Required,omitempty"`
	Default       interface{}               `json:"Default,omitempty"`
	Choices       []TestxValueChoice        `json:"Choices,omitempty"`
	Flag          string                    `json:"Flag,omitempty"`
	Order         uint32                    `json:"Order,omitempty"`
	EncryptOption interface{}               `json:"EncryptOption,omitempty"`
	ValueSource   interface{}               `json:"ValueSource,omitempty"`
	Properties    []TestxPropertyDefinition `json:"Properties,omitempty"`
	IsRichText    bool                      `json:"IsRichText,omitempty"`
}

// TestxCaseTemplateField testx 用例模板字段
type TestxCaseTemplateField struct {
	Uid           string                   `json:"Uid,omitempty"`
	Field         *TestxPropertyDefinition `json:"Field,omitempty"`
	Default       interface{}              `json:"Default,omitempty"`
	VisibleScopes []string                 `json:"VisibleScopes,omitempty"`
	Required      bool                     `json:"Required,omitempty"`
	Order         int32                    `json:"Order,omitempty"`
	Width         int32                    `json:"Width,omitempty"`
}

// TestxCaseTemplate testx 用例模板
type TestxCaseTemplate struct {
	Uid    string                   `json:"Uid,omitempty"`
	Name   string                   `json:"Name,omitempty"`
	Fields []TestxCaseTemplateField `json:"Fields,omitempty"`
}

// ----- 请求结构体 -----

// TestxCreateCaseRepoRequest 创建用例仓库请求
type TestxCreateCaseRepoRequest struct {
	Namespace   string `json:"-"`
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
}

// TestxUpdateCaseRepoRequest 更新用例仓库请求
type TestxUpdateCaseRepoRequest struct {
	Namespace   string `json:"-"`
	RepoUid     string `json:"-"`
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
}

// TestxGetCaseRepoRequest 获取用例仓库请求
type TestxGetCaseRepoRequest struct {
	Namespace string `json:"-"`
	RepoUid   string `json:"-"`
}

// TestxListCaseReposRequest 获取用例仓库列表请求
type TestxListCaseReposRequest struct {
	Namespace string         `json:"-"`
	PageInfo  *TestxPageInfo `json:"PageInfo,omitempty"`
	Search    string         `json:"Search,omitempty"`
	Uids      []string       `json:"Uids,omitempty"`
	Reverse   bool           `json:"Reverse,omitempty"`
}

// TestxCreateCaseFolderRequestFolder 创建用例目录中的目录信息
type TestxCreateCaseFolderRequestFolder struct {
	Uid            string   `json:"Uid,omitempty"`
	RepoUid        string   `json:"RepoUid,omitempty"`
	RepoVersionUid string   `json:"RepoVersionUid,omitempty"`
	FolderUid      string   `json:"FolderUid,omitempty"`
	Name           string   `json:"Name"`
	Owners         []string `json:"Owners,omitempty"`
	Description    string   `json:"Description,omitempty"`
}

// TestxCreateCaseFolderRequest 创建用例目录请求
type TestxCreateCaseFolderRequest struct {
	Namespace      string                              `json:"-"`
	RepoUid        string                              `json:"-"`
	RepoVersionUid string                              `json:"-"`
	Folder         *TestxCreateCaseFolderRequestFolder `json:"Folder"`
}

// TestxUpdateCaseFolderRequest 更新用例目录请求
type TestxUpdateCaseFolderRequest struct {
	Namespace      string                              `json:"-"`
	RepoUid        string                              `json:"-"`
	RepoVersionUid string                              `json:"-"`
	FolderUid      string                              `json:"-"`
	Folder         *TestxCreateCaseFolderRequestFolder `json:"Folder"`
}

// TestxCreateCaseRequestCase 创建用例中的用例信息
type TestxCreateCaseRequestCase struct {
	Uid              string                `json:"Uid,omitempty"`
	RepoUid          string                `json:"RepoUid,omitempty"`
	RepoVersionUid   string                `json:"RepoVersionUid,omitempty"`
	FolderUid        string                `json:"FolderUid,omitempty"`
	Name             string                `json:"Name"`
	Description      string                `json:"Description,omitempty"`
	Priority         string                `json:"Priority"`
	PreConditions    string                `json:"PreConditions,omitempty"`
	Type             string                `json:"Type,omitempty"`
	StepType         string                `json:"StepType,omitempty"`
	Steps            []TestxCaseStep       `json:"Steps,omitempty"`
	StepText         *TestxCaseStep        `json:"StepText,omitempty"`
	Attachments      []TestxCaseAttachment `json:"Attachments,omitempty"`
	CustomFields     []TestxProperty       `json:"CustomFields,omitempty"`
	Labels           []TestxLabel          `json:"Labels,omitempty"`
	Issues           []TestxIssue          `json:"Issues,omitempty"`
	Owners           []string              `json:"Owners,omitempty"`
	ManHourEstimated string                `json:"ManHourEstimated,omitempty"`
	RunTimes         string                `json:"RunTimes,omitempty"`
}

// TestxCreateCaseRequest 创建用例请求
type TestxCreateCaseRequest struct {
	Namespace      string                      `json:"-"`
	RepoUid        string                      `json:"-"`
	RepoVersionUid string                      `json:"-"`
	Case           *TestxCreateCaseRequestCase `json:"Case"`
}

// TestxUpdateCaseRequest 更新用例请求
type TestxUpdateCaseRequest struct {
	Namespace      string                      `json:"-"`
	RepoUid        string                      `json:"-"`
	RepoVersionUid string                      `json:"-"`
	CaseUid        string                      `json:"-"`
	Case           *TestxCreateCaseRequestCase `json:"Case"`
}

// TestxBatchCreateCaseFolderItem 批量创建用例中的目录项（支持嵌套）
type TestxBatchCreateCaseFolderItem struct {
	Uid            string                           `json:"Uid,omitempty"`
	RepoUid        string                           `json:"RepoUid,omitempty"`
	RepoVersionUid string                           `json:"RepoVersionUid,omitempty"`
	FolderUid      string                           `json:"FolderUid,omitempty"`
	Name           string                           `json:"Name"`
	Owners         []string                         `json:"Owners,omitempty"`
	Description    string                           `json:"Description,omitempty"`
	Folders        []TestxBatchCreateCaseFolderItem `json:"Folders,omitempty"`
	Cases          []TestxCreateCaseRequestCase     `json:"Cases,omitempty"`
}

// TestxBatchCreateCasesRequest 批量创建用例请求
type TestxBatchCreateCasesRequest struct {
	Namespace       string                           `json:"-"`
	RepoUid         string                           `json:"-"`
	RepoVersionUid  string                           `json:"-"`
	TargetFolderUid string                           `json:"TargetFolderUid"`
	Folders         []TestxBatchCreateCaseFolderItem `json:"Folders,omitempty"`
	Cases           []TestxCreateCaseRequestCase     `json:"Cases,omitempty"`
}

// TestxBatchUpdateCaseInfo 批量更新用例的更新信息
type TestxBatchUpdateCaseInfo struct {
	FieldName  string      `json:"FieldName"`
	FieldValue interface{} `json:"FieldValue"`
	Action     string      `json:"Action"`
}

// TestxBatchUpdateCasesRequest 批量更新用例请求
type TestxBatchUpdateCasesRequest struct {
	Namespace      string                     `json:"-"`
	RepoUid        string                     `json:"-"`
	RepoVersionUid string                     `json:"-"`
	CaseUids       []string                   `json:"CaseUids"`
	UpdateInfos    []TestxBatchUpdateCaseInfo `json:"UpdateInfos"`
}

// TestxSearchCasesFilter testx 搜索用例过滤条件
type TestxSearchCasesFilter struct {
	Name         string          `json:"Name,omitempty"`
	Uuid         string          `json:"Uuid,omitempty"`
	Priorities   []string        `json:"Priorities,omitempty"`
	Description  string          `json:"Description,omitempty"`
	Labels       []string        `json:"Labels,omitempty"`
	ReviewStates []string        `json:"ReviewStates,omitempty"`
	Owners       []string        `json:"Owners,omitempty"`
	CustomFields []TestxProperty `json:"CustomFields,omitempty"`
	Issues       []string        `json:"Issues,omitempty"`
	ItemType     string          `json:"ItemType,omitempty"`
	Creators     []string        `json:"Creators,omitempty"`
}

// TestxSearchCasesRequest 搜索用例请求
type TestxSearchCasesRequest struct {
	Namespace          string                  `json:"-"`
	RepoUid            string                  `json:"-"`
	RepoVersionUid     string                  `json:"-"`
	PageInfo           *TestxPageInfo          `json:"PageInfo,omitempty"`
	FolderUid          string                  `json:"FolderUid,omitempty"`
	CaseUids           []string                `json:"CaseUids,omitempty"`
	ExcludeCaseUids    []string                `json:"ExcludeCaseUids,omitempty"`
	Filter             *TestxSearchCasesFilter `json:"Filter"`
	ShowMode           string                  `json:"ShowMode,omitempty"`
	ExtendFields       []string                `json:"ExtendFields,omitempty"`
	IncludeDescendants bool                    `json:"IncludeDescendants,omitempty"`
	IncludeAncestors   bool                    `json:"IncludeAncestors,omitempty"`
	SelectFields       []string                `json:"SelectFields,omitempty"`
}

// TestxSearchCasesResponse 搜索用例响应
type TestxSearchCasesResponse struct {
	Folders []TestxCaseFolder `json:"Folders,omitempty"`
	Cases   []TestxCase       `json:"Cases,omitempty"`
	Repos   []TestxCaseRepo   `json:"Repos,omitempty"`
}

// TestxListCaseHistorysRequest 获取用例变更历史请求
type TestxListCaseHistorysRequest struct {
	Namespace      string         `json:"-"`
	RepoUid        string         `json:"-"`
	RepoVersionUid string         `json:"-"`
	CaseUid        string         `json:"-"`
	PageInfo       *TestxPageInfo `json:"PageInfo,omitempty"`
}

// TestxListCaseExecutionsRequest 获取用例执行记录请求
type TestxListCaseExecutionsRequest struct {
	Namespace      string         `json:"-"`
	RepoUid        string         `json:"-"`
	RepoVersionUid string         `json:"-"`
	CaseUid        string         `json:"-"`
	PageInfo       *TestxPageInfo `json:"PageInfo,omitempty"`
	Ordering       string         `json:"Ordering,omitempty"`
}

// TestxListCaseReviewsRequest 获取用例评审记录请求
type TestxListCaseReviewsRequest struct {
	Namespace      string         `json:"-"`
	RepoUid        string         `json:"-"`
	RepoVersionUid string         `json:"-"`
	CaseUid        string         `json:"-"`
	PageInfo       *TestxPageInfo `json:"PageInfo,omitempty"`
	Source         string         `json:"Source,omitempty"`
	MainUid        string         `json:"MainUid,omitempty"`
	SourceKind     string         `json:"SourceKind,omitempty"`
	SourceUid      string         `json:"SourceUid,omitempty"`
	MainUids       []string       `json:"MainUids,omitempty"`
	IsLastReview   bool           `json:"IsLastReview,omitempty"`
	CaseUids       []string       `json:"CaseUids,omitempty"`
	SourceUids     []string       `json:"SourceUids,omitempty"`
}

// TestxListCaseBugsRequest 获取用例关联缺陷请求
type TestxListCaseBugsRequest struct {
	Namespace      string         `json:"-"`
	RepoUid        string         `json:"-"`
	RepoVersionUid string         `json:"-"`
	CaseUid        string         `json:"-"`
	Status         string         `json:"Status,omitempty"`
	Priority       string         `json:"Priority,omitempty"`
	Handler        string         `json:"Handler,omitempty"`
	PageInfo       *TestxPageInfo `json:"PageInfo,omitempty"`
	Name           string         `json:"Name,omitempty"`
	SourceTypes    []string       `json:"SourceTypes,omitempty"`
}

// TestxBindBug testx 绑定 Bug 信息
type TestxBindBug struct {
	IssueUid     string `json:"IssueUid"`
	WorkspaceUid string `json:"WorkspaceUid"`
}

// TestxBatchBindCaseBugRequest 批量关联 Bug 请求
type TestxBatchBindCaseBugRequest struct {
	Namespace      string         `json:"-"`
	RepoUid        string         `json:"-"`
	RepoVersionUid string         `json:"-"`
	CaseUid        string         `json:"-"`
	BindBugs       []TestxBindBug `json:"BindBugs"`
}

// TestxBatchUnbindCaseBugRequest 批量解绑 Bug 请求
type TestxBatchUnbindCaseBugRequest struct {
	Namespace      string   `json:"-"`
	RepoUid        string   `json:"-"`
	RepoVersionUid string   `json:"-"`
	CaseUid        string   `json:"-"`
	BugUids        []string `json:"BugUids"`
}

// TestxListCaseTemplatesRequest 获取用例模板请求
type TestxListCaseTemplatesRequest struct {
	Namespace string `json:"-"`
}
