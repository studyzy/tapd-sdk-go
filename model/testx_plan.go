package model

import "encoding/json"

// ---------------------------------------------------------------------------
// 测试计划实体
// ---------------------------------------------------------------------------

// TestxPlanFolder testx 计划目录
type TestxPlanFolder struct {
	Uid         string            `json:"Uid,omitempty"`
	Namespace   string            `json:"Namespace,omitempty"`
	Audit       *TestxAudit       `json:"Audit,omitempty"`
	ParentUid   string            `json:"ParentUid,omitempty"`
	Name        string            `json:"Name,omitempty"`
	Description string            `json:"Description,omitempty"`
	PlanCount   uint32            `json:"PlanCount,omitempty"`
	ArchiveAuto bool              `json:"ArchiveAuto,omitempty"`
	Folders     []TestxPlanFolder `json:"Folders,omitempty"`
	Plans       []TestxPlanMeta   `json:"Plans,omitempty"`
	Nid         string            `json:"Nid,omitempty"`
	Path        string            `json:"Path,omitempty"`
}

// TestxPlanMeta testx 计划元信息
type TestxPlanMeta struct {
	Uid         string               `json:"Uid,omitempty"`
	Namespace   string               `json:"Namespace,omitempty"`
	Audit       *TestxAudit          `json:"Audit,omitempty"`
	FolderUid   string               `json:"FolderUid,omitempty"`
	Name        string               `json:"Name,omitempty"`
	Description string               `json:"Description,omitempty"`
	DataSource  *TestxPlanDataSource `json:"DataSource,omitempty"`
	State       string               `json:"State,omitempty"`
	Testers     []string             `json:"Testers,omitempty"`
	FolderPath  string               `json:"FolderPath,omitempty"`
	Version     string               `json:"Version,omitempty"`
	CasePath    *TestxCasePath       `json:"CasePath,omitempty"`
	Nid         string               `json:"Nid,omitempty"`
	Path        string               `json:"Path,omitempty"`
}

// TestxPlanDataSource testx 计划数据源
type TestxPlanDataSource struct {
	Mode         string   `json:"Mode,omitempty"`
	CaseUids     []string `json:"CaseUids,omitempty"`
	CasesetUids  []string `json:"CasesetUids,omitempty"`
	CaseSuiteUid string   `json:"CaseSuiteUid,omitempty"`
}

// TestxCasePath testx 用例存储路径
type TestxCasePath struct {
	RepoUid        string `json:"RepoUid,omitempty"`
	RepoVersionUid string `json:"RepoVersionUid,omitempty"`
	FolderUid      string `json:"FolderUid,omitempty"`
}

// TestxPlan testx 测试计划
type TestxPlan struct {
	Meta *TestxPlanMeta `json:"Meta,omitempty"`
	Spec *TestxPlanSpec `json:"Spec,omitempty"`
}

// TestxPlanSpec testx 计划规格
type TestxPlanSpec struct {
	Field        *TestxPlanFieldSpec `json:"Field,omitempty"`
	CustomFields []TestxProperty     `json:"CustomFields,omitempty"`
	Stories      []TestxIssue        `json:"Stories,omitempty"`
	Bugs         []TestxIssue        `json:"Bugs,omitempty"`
	Iterations   []TestxIssue        `json:"Iterations,omitempty"`
	Statistic    *TestxPlanStatistic `json:"Statistic,omitempty"`
	Properties   []TestxProperty     `json:"Properties,omitempty"`
	SystemFields []TestxProperty     `json:"SystemFields,omitempty"`
	Scope        *TestxPlanScope     `json:"Scope,omitempty"`
	Target       *TestxPlanTarget    `json:"Target,omitempty"`
	Versions     []TestxIssue        `json:"Versions,omitempty"`
}

// TestxPlanFieldSpec testx 计划字段
type TestxPlanFieldSpec struct {
	Developers        []string `json:"Developers,omitempty"`
	ProjectManagers   []string `json:"ProjectManagers,omitempty"`
	ProductManagers   []string `json:"ProductManagers,omitempty"`
	Summary           string   `json:"Summary,omitempty"`
	Version           string   `json:"Version,omitempty"`
	Type              string   `json:"Type,omitempty"`
	ExtranetPublish   string   `json:"ExtranetPublish,omitempty"`
	PublishReview     string   `json:"PublishReview,omitempty"`
	EstimateTestAt    string   `json:"EstimateTestAt,omitempty"`
	ActualTestAt      string   `json:"ActualTestAt,omitempty"`
	EstimatePublishAt string   `json:"EstimatePublishAt,omitempty"`
	ActualPublishAt   string   `json:"ActualPublishAt,omitempty"`
	EstimateStartedAt string   `json:"EstimateStartedAt,omitempty"`
	EstimateEndedAt   string   `json:"EstimateEndedAt,omitempty"`
	Env               string   `json:"Env,omitempty"`
}

// TestxPlanStatistic testx 计划统计信息
type TestxPlanStatistic struct {
	SucceedCaseCount  uint32 `json:"SucceedCaseCount,omitempty"`
	FailedCaseCount   uint32 `json:"FailedCaseCount,omitempty"`
	BlockCaseCount    uint32 `json:"BlockCaseCount,omitempty"`
	ErrorCaseCount    uint32 `json:"ErrorCaseCount,omitempty"`
	RetryCaseCount    uint32 `json:"RetryCaseCount,omitempty"`
	IgnoreCaseCount   uint32 `json:"IgnoreCaseCount,omitempty"`
	TodoCaseCount     uint32 `json:"TodoCaseCount,omitempty"`
	ManualSucceedRate string `json:"ManualSucceedRate,omitempty"`
	TestedCaseCount   uint32 `json:"TestedCaseCount,omitempty"`
	TotalCaseCount    uint32 `json:"TotalCaseCount,omitempty"`
	CaseSucceedRate   string `json:"CaseSucceedRate,omitempty"`
	StoryPassedRate   string `json:"StoryPassedRate,omitempty"`
	FailedTaskCount   uint32 `json:"FailedTaskCount,omitempty"`
	CaseCoverageRate  string `json:"CaseCoverageRate,omitempty"`
	StoryCount        uint32 `json:"StoryCount,omitempty"`
}

// TestxPlanScope testx 计划范围
type TestxPlanScope struct {
	Stories           []TestxIssue            `json:"Stories,omitempty"`
	StoryCaseStrategy *TestxStoryCaseStrategy `json:"StoryCaseStrategy,omitempty"`
}

// TestxStoryCaseStrategy testx 用例关联策略
type TestxStoryCaseStrategy struct {
	Mode       string   `json:"Mode,omitempty"`
	Priorities []string `json:"Priorities,omitempty"`
}

// TestxPlanTarget testx 计划目标
type TestxPlanTarget struct {
	Description       string   `json:"Description,omitempty"`
	Feature           string   `json:"Feature,omitempty"`
	Deliverables      string   `json:"Deliverables,omitempty"`
	SuccessConditions []string `json:"SuccessConditions,omitempty"`
}

// TestxPlanHistory testx 计划变更历史
type TestxPlanHistory struct {
	Uid      string                  `json:"Uid,omitempty"`
	Creator  string                  `json:"Creator,omitempty"`
	CreateAt string                  `json:"CreateAt,omitempty"`
	PlanUid  string                  `json:"PlanUid,omitempty"`
	Fields   []TestxPlanHistoryField `json:"Fields,omitempty"`
}

// TestxPlanHistoryField testx 计划变更字段
type TestxPlanHistoryField struct {
	Uid       string `json:"Uid,omitempty"`
	Name      string `json:"Name,omitempty"`
	PreValue  string `json:"PreValue,omitempty"`
	PostValue string `json:"PostValue,omitempty"`
}

// TestxPlanCaseDetail testx 计划用例详情
type TestxPlanCaseDetail struct {
	State         string               `json:"State,omitempty"`
	Tester        string               `json:"Tester,omitempty"`
	FinalTester   string               `json:"FinalTester,omitempty"`
	EndedAt       string               `json:"EndedAt,omitempty"`
	Source        string               `json:"Source,omitempty"`
	PlanUid       string               `json:"PlanUid,omitempty"`
	CaseUid       string               `json:"CaseUid,omitempty"`
	CaseNid       string               `json:"CaseNid,omitempty"`
	Bugs          []TestxIssue         `json:"Bugs,omitempty"`
	Events        []TestxPlanCaseEvent `json:"Events,omitempty"`
	Statistic     *TestxPlanStatistic  `json:"Statistic,omitempty"`
	Uid           string               `json:"Uid,omitempty"`
	CaseName      string               `json:"CaseName,omitempty"`
	PlanName      string               `json:"PlanName,omitempty"`
	PlanFolderUid string               `json:"PlanFolderUid,omitempty"`
	RunTimes      int                  `json:"RunTimes,omitempty"`
	ReviewState   string               `json:"ReviewState,omitempty"`
}

// TestxPlanCaseEvent testx 计划用例事件
type TestxPlanCaseEvent struct {
	Audit       *TestxAudit `json:"Audit,omitempty"`
	Type        string      `json:"Type,omitempty"`
	Detail      string      `json:"Detail,omitempty"`
	Source      string      `json:"Source,omitempty"`
	Attachments []string    `json:"Attachments,omitempty"`
}

// TestxPlanCasesInfo testx 计划用例信息
type TestxPlanCasesInfo struct {
	StoryGroup       []TestxPlanStoryGroup       `json:"StoryGroup,omitempty"`
	RepoVersionGroup []TestxPlanRepoVersionGroup `json:"RepoVersionGroup,omitempty"`
}

// TestxPlanStoryGroup testx 计划需求分组
type TestxPlanStoryGroup struct {
	Story         *TestxIssue       `json:"Story,omitempty"`
	Cases         []TestxCase       `json:"Cases,omitempty"`
	Folders       []TestxCaseFolder `json:"Folders,omitempty"`
	ExecCaseCount int               `json:"ExecCaseCount,omitempty"`
}

// TestxPlanRepoVersionGroup testx 计划用例库版本分组
type TestxPlanRepoVersionGroup struct {
	RepoVersion *TestxCaseRepoVersion `json:"RepoVersion,omitempty"`
	Folders     []TestxCaseFolder     `json:"Folders,omitempty"`
}

// TestxPlanCasesResult testx 计划用例搜索结果
type TestxPlanCasesResult struct {
	CaseUidToDetail map[string]TestxPlanCaseDetail `json:"CaseUidToDetail,omitempty"`
	PlanCasesInfo   *TestxPlanCasesInfo            `json:"PlanCasesInfo,omitempty"`
}

// TestxPlanStatisticsItem testx 计划统计项
type TestxPlanStatisticsItem struct {
	Uid       string              `json:"Uid,omitempty"`
	Statistic *TestxPlanStatistic `json:"Statistic,omitempty"`
	CaseRepos []TestxCaseRepo     `json:"CaseRepos,omitempty"`
}

// TestxPlanBugStatistics testx 计划缺陷统计
type TestxPlanBugStatistics struct {
	PlanUid  string `json:"PlanUid,omitempty"`
	BugCount uint32 `json:"BugCount,omitempty"`
}

// TestxPlanBug testx 计划关联的缺陷
type TestxPlanBug struct {
	WorkspaceId                 string               `json:"WorkspaceId,omitempty"`
	IterationId                 string               `json:"IterationId,omitempty"`
	Id                          string               `json:"Id,omitempty"`
	Summary                     string               `json:"Summary,omitempty"`
	Creator                     string               `json:"Creator,omitempty"`
	Priority                    *TestxBugLV          `json:"Priority,omitempty"`
	Created                     string               `json:"Created,omitempty"`
	Status                      *TestxBugLV          `json:"Status,omitempty"`
	LastUpdater                 string               `json:"LastUpdater,omitempty"`
	UpdateTime                  string               `json:"UpdateTime,omitempty"`
	Severity                    *TestxBugLV          `json:"Severity,omitempty"`
	Owners                      string               `json:"Owners,omitempty"`
	Begin                       string               `json:"Begin,omitempty"`
	Due                         string               `json:"Due,omitempty"`
	NamespaceId                 string               `json:"NamespaceId,omitempty"`
	Description                 string               `json:"Description,omitempty"`
	Url                         string               `json:"Url,omitempty"`
	Developers                  string               `json:"Developers,omitempty"`
	Testers                     string               `json:"Testers,omitempty"`
	Module                      string               `json:"Module,omitempty"`
	Attachments                 []TestxBugAttachment `json:"Attachments,omitempty"`
	Tags                        []string             `json:"Tags,omitempty"`
	RelatedType                 string               `json:"RelatedType,omitempty"`
	IconType                    string               `json:"IconType,omitempty"`
	Version                     []string             `json:"Version,omitempty"`
	ReleaseId                   string               `json:"ReleaseId,omitempty"`
	VersionReport               string               `json:"VersionReport,omitempty"`
	VersionTest                 string               `json:"VersionTest,omitempty"`
	VersionFix                  string               `json:"VersionFix,omitempty"`
	VersionClose                string               `json:"VersionClose,omitempty"`
	BaselineFind                string               `json:"BaselineFind,omitempty"`
	BaselineJoin                string               `json:"BaselineJoin,omitempty"`
	BaselineClose               string               `json:"BaselineClose,omitempty"`
	BaselineTest                string               `json:"BaselineTest,omitempty"`
	CC                          string               `json:"CC,omitempty"`
	Participator                string               `json:"Participator,omitempty"`
	Auditer                     string               `json:"Auditer,omitempty"`
	Confirmer                   string               `json:"Confirmer,omitempty"`
	Fixer                       string               `json:"Fixer,omitempty"`
	Closer                      string               `json:"Closer,omitempty"`
	ReopenTime                  string               `json:"ReopenTime,omitempty"`
	InProgressTime              string               `json:"InProgressTime,omitempty"`
	Resolved                    string               `json:"Resolved,omitempty"`
	VerifyTime                  string               `json:"VerifyTime,omitempty"`
	Closed                      string               `json:"Closed,omitempty"`
	RejectTime                  string               `json:"RejectTime,omitempty"`
	Modified                    string               `json:"Modified,omitempty"`
	Deadline                    string               `json:"Deadline,omitempty"`
	Os                          string               `json:"Os,omitempty"`
	Platform                    string               `json:"Platform,omitempty"`
	Testmode                    string               `json:"Testmode,omitempty"`
	Testphase                   string               `json:"Testphase,omitempty"`
	Testtype                    string               `json:"Testtype,omitempty"`
	Source                      string               `json:"Source,omitempty"`
	Frequency                   string               `json:"Frequency,omitempty"`
	Originphase                 string               `json:"Originphase,omitempty"`
	Sourcephase                 string               `json:"Sourcephase,omitempty"`
	Resolution                  *TestxBugLV          `json:"Resolution,omitempty"`
	Estimate                    string               `json:"Estimate,omitempty"`
	Lastmodify                  string               `json:"Lastmodify,omitempty"`
	CustomFields                map[string]string    `json:"CustomFields,omitempty"`
	TemplateId                  string               `json:"TemplateId,omitempty"`
	IsApplyTemplateDefaultValue bool                 `json:"IsApplyTemplateDefaultValue,omitempty"`
	CustomPlanFields            map[string]string    `json:"CustomPlanFields,omitempty"`
	Feature                     string               `json:"Feature,omitempty"`
	Effort                      string               `json:"Effort,omitempty"`
	Bugtype                     string               `json:"Bugtype,omitempty"`
	Size                        string               `json:"Size,omitempty"`
}

// TestxBugLV testx 缺陷标签/值对
type TestxBugLV struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}

// TestxPlanTemplate testx 计划模板
type TestxPlanTemplate struct {
	Uid         string                   `json:"Uid,omitempty"`
	Name        string                   `json:"Name,omitempty"`
	DescName    string                   `json:"DescName,omitempty"`
	Description string                   `json:"Description,omitempty"`
	Audit       *TestxAudit              `json:"Audit,omitempty"`
	Fields      []TestxPlanTemplateField `json:"Fields,omitempty"`
}

// TestxPlanTemplateField testx 计划模板字段
type TestxPlanTemplateField struct {
	Uid          string                       `json:"Uid,omitempty"`
	Field        *TestxPlanPropertyDefinition `json:"Field,omitempty"`
	Default      interface{}                  `json:"Default,omitempty"`
	VisibleScope string                       `json:"VisibleScope,omitempty"`
	Required     bool                         `json:"Required,omitempty"`
	Width        uint32                       `json:"Width,omitempty"`
	Order        uint32                       `json:"Order,omitempty"`
	FieldUid     string                       `json:"FieldUid,omitempty"`
}

// TestxPlanPropertyDefinition testx 计划属性定义
type TestxPlanPropertyDefinition struct {
	Widget      string                        `json:"Widget,omitempty"`
	Type        string                        `json:"Type,omitempty"`
	Uneditable  bool                          `json:"Uneditable,omitempty"`
	Name        string                        `json:"Name,omitempty"`
	Label       string                        `json:"Label,omitempty"`
	Description string                        `json:"Description,omitempty"`
	Help        string                        `json:"Help,omitempty"`
	Visible     bool                          `json:"Visible,omitempty"`
	Required    bool                          `json:"Required,omitempty"`
	Default     interface{}                   `json:"Default,omitempty"`
	Choices     []TestxPlanValueChoice        `json:"Choices,omitempty"`
	Flag        string                        `json:"Flag,omitempty"`
	Order       uint32                        `json:"Order,omitempty"`
	ValueSource json.RawMessage               `json:"ValueSource,omitempty"`
	Properties  []TestxPlanPropertyDefinition `json:"Properties,omitempty"`
	IsRichText  bool                          `json:"IsRichText,omitempty"`
}

// TestxPlanValueChoice testx 计划属性可选值
type TestxPlanValueChoice struct {
	ID    uint32      `json:"ID,omitempty"`
	Label string      `json:"Label,omitempty"`
	Value interface{} `json:"Value,omitempty"`
	Order uint32      `json:"Order,omitempty"`
}

// TestxFolderChildrenResult testx 目录子信息结果
type TestxFolderChildrenResult struct {
	Folders []TestxPlanFolder `json:"Folders,omitempty"`
	Plans   []TestxPlanMeta   `json:"Plans,omitempty"`
}

// ---------------------------------------------------------------------------
// 请求结构体
// ---------------------------------------------------------------------------

// TestxCreatePlanFolderRequest 创建计划目录请求
type TestxCreatePlanFolderRequest struct {
	Namespace   string `json:"-"`
	ParentUid   string `json:"ParentUid"`
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
}

// TestxUpdatePlanFolderRequest 更新计划目录请求
type TestxUpdatePlanFolderRequest struct {
	Namespace   string `json:"-"`
	FolderUid   string `json:"-"`
	ParentUid   string `json:"ParentUid"`
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
	ArchiveAuto *bool  `json:"ArchiveAuto,omitempty"`
}

// TestxListFolderChildrenRequest 获取计划目录子信息请求
type TestxListFolderChildrenRequest struct {
	Namespace      string   `json:"-"`
	Uid            string   `json:"-"`
	WithDescendant bool     `json:"-"`
	WithAncestor   bool     `json:"-"`
	Name           string   `json:"-"`
	PlanArchive    string   `json:"-"`
	PlanStates     []string `json:"-"`
	ItemType       string   `json:"-"`
}

// TestxGetPlanRequest 获取计划详情请求
type TestxGetPlanRequest struct {
	Namespace     string `json:"-"`
	Uid           string `json:"-"`
	WithStatistic bool   `json:"-"`
	WithDetail    string `json:"-"`
}

// TestxCreatePlanRequest 创建计划请求
type TestxCreatePlanRequest struct {
	Namespace string     `json:"-"`
	Plan      *TestxPlan `json:"Plan"`
}

// TestxUpdatePlanRequest 更新计划请求
type TestxUpdatePlanRequest struct {
	Namespace string     `json:"-"`
	Uid       string     `json:"-"`
	Plan      *TestxPlan `json:"Plan"`
}

// TestxUpdatePlanTargetScopeRequest 更新计划范围和目标请求
type TestxUpdatePlanTargetScopeRequest struct {
	Namespace string     `json:"-"`
	Uid       string     `json:"-"`
	Plan      *TestxPlan `json:"Plan"`
}

// TestxBatchUpdatePlanCaseInfo 批量更新用例信息项
type TestxBatchUpdatePlanCaseInfo struct {
	State   string       `json:"State"`
	Tester  string       `json:"Tester"`
	CaseUid string       `json:"CaseUid"`
	Bugs    []TestxIssue `json:"Bugs,omitempty"`
}

// TestxBatchUpdatePlanCaseEvent 批量更新用例事件
type TestxBatchUpdatePlanCaseEvent struct {
	Type   string `json:"Type"`
	Detail string `json:"Detail"`
	Source string `json:"Source"`
}

// TestxBatchUpdatePlanCaseRequest 批量更新计划用例请求
type TestxBatchUpdatePlanCaseRequest struct {
	Namespace string                          `json:"-"`
	PlanUid   string                          `json:"-"`
	CaseInfos []TestxBatchUpdatePlanCaseInfo  `json:"CaseInfos"`
	Events    []TestxBatchUpdatePlanCaseEvent `json:"Events,omitempty"`
}

// TestxBatchArchivePlanRequest 批量归档计划请求
type TestxBatchArchivePlanRequest struct {
	Namespace   string   `json:"-"`
	Uids        []string `json:"Uids"`
	ArchiveMode string   `json:"ArchiveMode"`
}

// TestxListPlansRequest 获取目录下计划列表请求
type TestxListPlansRequest struct {
	Namespace        string         `json:"-"`
	FolderUid        string         `json:"-"`
	WithDetail       bool           `json:"WithDetail,omitempty"`
	PageInfo         *TestxPageInfo `json:"PageInfo,omitempty"`
	Ordering         string         `json:"Ordering,omitempty"`
	States           []string       `json:"States,omitempty"`
	Testers          []string       `json:"Testers,omitempty"`
	Uids             []string       `json:"Uids,omitempty"`
	Archive          string         `json:"Archive,omitempty"`
	Name             string         `json:"Name,omitempty"`
	Description      string         `json:"Description,omitempty"`
	Assigned         string         `json:"Assigned,omitempty"`
	IterationUids    []string       `json:"IterationUids,omitempty"`
	IgnoreFolderUid  bool           `json:"IgnoreFolderUid,omitempty"`
	StoryUids        []string       `json:"StoryUids,omitempty"`
	StartAt          string         `json:"StartAt,omitempty"`
	EndAt            string         `json:"EndAt,omitempty"`
	TimeIntervalType string         `json:"TimeIntervalType,omitempty"`
	Fields           []string       `json:"Fields,omitempty"`
}

// TestxPlanCaseFilter testx 计划用例过滤条件
type TestxPlanCaseFilter struct {
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

// TestxListPlanCasesRequest 获取计划下用例列表请求
type TestxListPlanCasesRequest struct {
	Namespace           string               `json:"-"`
	Uid                 string               `json:"-"`
	ShowMode            string               `json:"ShowMode,omitempty"`
	PageInfo            *TestxPageInfo       `json:"PageInfo,omitempty"`
	Testers             []string             `json:"Testers,omitempty"`
	States              []string             `json:"States,omitempty"`
	Sources             []string             `json:"Sources,omitempty"`
	CaseFilter          *TestxPlanCaseFilter `json:"CaseFilter,omitempty"`
	WithAncestor        bool                 `json:"WithAncestor,omitempty"`
	RepoVersionUid      string               `json:"RepoVersionUid,omitempty"`
	CaseSelectFields    []string             `json:"CaseSelectFields,omitempty"`
	SelectFields        []string             `json:"SelectFields,omitempty"`
	Version             string               `json:"Version,omitempty"`
	WithFullPath        bool                 `json:"WithFullPath,omitempty"`
	CaseExtendFields    []string             `json:"CaseExtendFields,omitempty"`
	StoryShowMode       string               `json:"StoryShowMode,omitempty"`
	WithCaseReviewState bool                 `json:"WithCaseReviewState,omitempty"`
}

// TestxListPlanHistoriesRequest 获取计划变更历史请求
type TestxListPlanHistoriesRequest struct {
	Namespace string         `json:"-"`
	PlanUid   string         `json:"-"`
	PageInfo  *TestxPageInfo `json:"-"`
}

// TestxPlanStatisticsRequest 获取计划统计信息请求
type TestxPlanStatisticsRequest struct {
	Namespace string         `json:"-"`
	Uids      []string       `json:"Uids"`
	PageInfo  *TestxPageInfo `json:"PageInfo,omitempty"`
}

// TestxBatchBindPlanBugRequest 计划用例批量添加缺陷请求
type TestxBatchBindPlanBugRequest struct {
	Namespace string       `json:"-"`
	PlanUid   string       `json:"-"`
	CaseUids  []string     `json:"CaseUids"`
	BindBugs  []TestxIssue `json:"BindBugs"`
}

// TestxUnbindPlanBugRequest 移除计划用例关联缺陷请求
type TestxUnbindPlanBugRequest struct {
	Namespace string `json:"-"`
	PlanUid   string `json:"-"`
	CaseUid   string `json:"-"`
	IssueUid  string `json:"-"`
}

// TestxListPlanCaseIssuesRequest 获取计划用例关联缺陷列表请求
type TestxListPlanCaseIssuesRequest struct {
	Namespace string         `json:"-"`
	PlanUid   string         `json:"-"`
	CaseUid   string         `json:"-"`
	Type      string         `json:"-"`
	PageInfo  *TestxPageInfo `json:"-"`
}

// TestxListPlanCaseEventsRequest 获取计划用例事件列表请求
type TestxListPlanCaseEventsRequest struct {
	Namespace string         `json:"-"`
	PlanUid   string         `json:"-"`
	CaseUid   string         `json:"-"`
	PageInfo  *TestxPageInfo `json:"-"`
}

// TestxListPlanBugsRequest 获取计划关联缺陷列表请求
type TestxListPlanBugsRequest struct {
	Namespace    string         `json:"-"`
	PlanUid      string         `json:"-"`
	PageInfo     *TestxPageInfo `json:"-"`
	RelatedTypes []string       `json:"-"`
	Status       string         `json:"-"`
	Summary      string         `json:"-"`
	BugId        string         `json:"-"`
}

// TestxListPlanBugStatisticsRequest 批量查询计划关联缺陷统计请求
type TestxListPlanBugStatisticsRequest struct {
	Namespace string   `json:"-"`
	PlanUids  []string `json:"PlanUids"`
}

// TestxListPlanStoriesRequest 获取计划关联需求列表请求
type TestxListPlanStoriesRequest struct {
	Namespace string         `json:"-"`
	PlanUid   string         `json:"-"`
	PageInfo  *TestxPageInfo `json:"-"`
}

// TestxListPlanTemplatesRequest 获取计划模板请求
type TestxListPlanTemplatesRequest struct {
	Namespace string         `json:"-"`
	PageInfo  *TestxPageInfo `json:"-"`
}
