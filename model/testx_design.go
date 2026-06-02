package model

// ---------------------------------------------------------------------------
// 测试设计实体（snake_case JSON tag）
// ---------------------------------------------------------------------------

// TestxDesignAudit testx 测试设计审计信息
type TestxDesignAudit struct {
	Creator   string `json:"creator,omitempty"`
	Updater   string `json:"updater,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Tenant    string `json:"tenant,omitempty"`
}

// TestxDesignMeta testx 测试设计元数据
type TestxDesignMeta struct {
	Audit       *TestxDesignAudit `json:"audit,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Uid         string            `json:"uid,omitempty"`
	ParentUid   string            `json:"parent_uid,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Thumbnail   string            `json:"thumbnail,omitempty"`
	Source      string            `json:"source,omitempty"`
	UseAI       bool              `json:"use_ai,omitempty"`
	TemplateUid string            `json:"template_uid,omitempty"`
}

// TestxDesignIssue testx 测试设计关联问题（snake_case 版）
type TestxDesignIssue struct {
	IssueUid     string      `json:"issue_uid,omitempty"`
	Namespace    string      `json:"namespace,omitempty"`
	WorkspaceUid string      `json:"workspace_uid,omitempty"`
	IssueUrl     string      `json:"issue_url,omitempty"`
	Type         string      `json:"type,omitempty"`
	Source       string      `json:"source,omitempty"`
	Detail       interface{} `json:"detail,omitempty"`
	IssueName    string      `json:"issue_name,omitempty"`
	IsDeleted    bool        `json:"is_deleted,omitempty"`
	Uid          string      `json:"uid,omitempty"`
}

// TestxDesignLabel testx 测试设计标签（snake_case 版）
type TestxDesignLabel struct {
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Color       string `json:"color,omitempty"`
	Uneditable  bool   `json:"uneditable,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Source      string `json:"source,omitempty"`
}

// TestxDesignSpec testx 测试设计规格信息
type TestxDesignSpec struct {
	Theme        string             `json:"theme,omitempty"`
	Labels       []TestxDesignLabel `json:"labels,omitempty"`
	Stories      []TestxDesignIssue `json:"stories,omitempty"`
	Bugs         []TestxDesignIssue `json:"bugs,omitempty"`
	CustomFields map[string]string  `json:"custom_fields,omitempty"`
}

// TestxDesignComment testx 测试设计评论
type TestxDesignComment struct {
	Uid       string `json:"uid,omitempty"`
	Message   string `json:"message,omitempty"`
	ReplyUid  string `json:"reply_uid,omitempty"`
	ReviewUid string `json:"review_uid,omitempty"`
}

// TestxDesignReview testx 测试设计总评信息
type TestxDesignReview struct {
	Audit     *TestxDesignAudit    `json:"audit,omitempty"`
	Namespace string               `json:"namespace,omitempty"`
	DesignUid string               `json:"design_uid,omitempty"`
	Uid       string               `json:"uid,omitempty"`
	State     string               `json:"state,omitempty"`
	Comments  []TestxDesignComment `json:"comments,omitempty"`
	Reviewers []string             `json:"reviewers,omitempty"`
	Stat      interface{}          `json:"stat,omitempty"`
}

// TestxDesignStatistics testx 测试设计节点数据统计
type TestxDesignStatistics struct {
	CustomCount    uint32 `json:"custom_count,omitempty"`
	StoryCount     uint32 `json:"story_count,omitempty"`
	SceneCount     uint32 `json:"scene_count,omitempty"`
	TestPointCount uint32 `json:"test_point_count,omitempty"`
	CaseCount      uint32 `json:"case_count,omitempty"`
	FeatureCount   uint32 `json:"feature_count,omitempty"`
	BugCount       uint32 `json:"bug_count,omitempty"`
	TotalCount     uint32 `json:"total_count,omitempty"`
	NodeUid        string `json:"node_uid,omitempty"`
	ExecCount      uint32 `json:"exec_count,omitempty"`
	ExecPassCount  uint32 `json:"exec_pass_count,omitempty"`
	AiCaseCount    uint32 `json:"ai_case_count,omitempty"`
}

// TestxDesign testx 测试设计
type TestxDesign struct {
	Meta   *TestxDesignMeta       `json:"meta,omitempty"`
	Spec   *TestxDesignSpec       `json:"spec,omitempty"`
	Review *TestxDesignReview     `json:"review,omitempty"`
	Nodes  []interface{}          `json:"nodes,omitempty"`
	Stat   *TestxDesignStatistics `json:"stat,omitempty"`
}

// ---------------------------------------------------------------------------
// 统计信息实体
// ---------------------------------------------------------------------------

// TestxDesignArchiveState testx 测试设计归档状态
type TestxDesignArchiveState struct {
	Status  int32 `json:"status,omitempty"`
	Count   int32 `json:"count,omitempty"`
	CaseNum int32 `json:"case_num,omitempty"`
}

// TestxDesignReviewState testx 测试设计评审状态
type TestxDesignReviewState struct {
	State               string `json:"state,omitempty"`
	ReviewedIssueCount  uint32 `json:"reviewed_issue_count,omitempty"`
	TotalIssueCount     uint32 `json:"total_issue_count,omitempty"`
	ReviewedCaseCount   uint32 `json:"reviewed_case_count,omitempty"`
	TotalCaseCount      uint32 `json:"total_case_count,omitempty"`
	AIReviewedCaseCount uint32 `json:"ai_reviewed_case_count,omitempty"`
	TotalBugCount       uint32 `json:"total_bug_count,omitempty"`
}

// TestxDesignExtendFeatures testx 测试设计扩展信息
type TestxDesignExtendFeatures struct {
	UsedAI  bool                     `json:"used_ai,omitempty"`
	Archive *TestxDesignArchiveState `json:"archive,omitempty"`
	Review  *TestxDesignReviewState  `json:"review,omitempty"`
}

// TestxDesignStat testx 测试设计统计
type TestxDesignStat struct {
	DesignUid      string                     `json:"design_uid,omitempty"`
	ExtendFeatures *TestxDesignExtendFeatures `json:"extend_features,omitempty"`
}

// ---------------------------------------------------------------------------
// 请求结构体
// ---------------------------------------------------------------------------

// TestxDesignTimeFilter testx 测试设计时间过滤条件
type TestxDesignTimeFilter struct {
	StartAt  string `json:"start_at,omitempty"`
	EndAt    string `json:"end_at,omitempty"`
	TimeType string `json:"time_type,omitempty"`
}

// TestxDesignOrder testx 测试设计排序条件
type TestxDesignOrder struct {
	OrderField string `json:"order_field,omitempty"`
	OrderType  string `json:"order_type,omitempty"`
}

// TestxDesignPeriod testx 测试设计时间范围
type TestxDesignPeriod struct {
	BeginTime string `json:"begin_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Type      string `json:"type,omitempty"`
}

// TestxSearchDesignsFilter testx 搜索测试设计过滤条件
type TestxSearchDesignsFilter struct {
	Uids         []string                `json:"uids,omitempty"`
	Search       string                  `json:"search,omitempty"`
	Creators     []string                `json:"creators,omitempty"`
	Updaters     []string                `json:"updaters,omitempty"`
	TimeFilters  []TestxDesignTimeFilter `json:"time_filters,omitempty"`
	Orders       []TestxDesignOrder      `json:"orders,omitempty"`
	Period       *TestxDesignPeriod      `json:"period,omitempty"`
	RefUid       string                  `json:"ref_uid,omitempty"`
	States       []string                `json:"states,omitempty"`
	Labels       []string                `json:"labels,omitempty"`
	StoryIDs     []string                `json:"story_ids,omitempty"`
	Unscoped     bool                    `json:"unscoped,omitempty"`
	CustomFields map[string]string       `json:"custom_fields,omitempty"`
}

// TestxDesignPageInfo testx 测试设计分页参数
type TestxDesignPageInfo struct {
	Offset uint32 `json:"Offset"`
	Limit  uint32 `json:"Limit"`
}

// TestxSearchDesignsRequest 查询测试设计列表请求
type TestxSearchDesignsRequest struct {
	Namespace string                    `json:"-"`
	Filter    *TestxSearchDesignsFilter `json:"filter,omitempty"`
	PageInfo  *TestxDesignPageInfo      `json:"page_info,omitempty"`
}

// TestxListDesignStatsRequest 查询测试设计列表统计信息请求
type TestxListDesignStatsRequest struct {
	Namespace  string   `json:"-"`
	DesignUids []string `json:"design_uids,omitempty"`
}

// TestxListDesignLabelsRequest 查询测试设计标签列表请求
type TestxListDesignLabelsRequest struct {
	Namespace string `json:"-"`
	DesignUid string `json:"-"`
	Kind      string `json:"-"`
	Name      string `json:"-"`
}
