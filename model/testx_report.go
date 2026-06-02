package model

import "encoding/json"

// ---------------------------------------------------------------------------
// 报告实体
// ---------------------------------------------------------------------------

// TestxReportPlanMeta testx 报告关联计划元数据
type TestxReportPlanMeta struct {
	Uid         string      `json:"Uid,omitempty"`
	Namespace   string      `json:"Namespace,omitempty"`
	Audit       *TestxAudit `json:"Audit,omitempty"`
	FolderUid   string      `json:"FolderUid,omitempty"`
	Name        string      `json:"Name,omitempty"`
	Description string      `json:"Description,omitempty"`
	DataSource  interface{} `json:"DataSource,omitempty"`
	State       string      `json:"State,omitempty"`
	Testers     []string    `json:"Testers,omitempty"`
	FolderPath  string      `json:"FolderPath,omitempty"`
	Version     string      `json:"Version,omitempty"`
	CasePath    interface{} `json:"CasePath,omitempty"`
	Nid         string      `json:"Nid,omitempty"`
	Path        string      `json:"Path,omitempty"`
}

// TestxReportPlan testx 报告关联计划
type TestxReportPlan struct {
	Meta *TestxReportPlanMeta `json:"Meta,omitempty"`
	Spec interface{}          `json:"Spec,omitempty"`
}

// TestxReportCron testx 定期任务表达式
type TestxReportCron struct {
	Period      string   `json:"Period,omitempty"`
	MonthlyDays []uint32 `json:"MonthlyDays,omitempty"`
	WeeklyDays  []string `json:"WeeklyDays,omitempty"`
	TimeAt      string   `json:"TimeAt,omitempty"`
}

// TestxReportPeriodicTask testx 定期任务
type TestxReportPeriodicTask struct {
	Id        string           `json:"Id,omitempty"`
	StartedAt string           `json:"StartedAt,omitempty"`
	ExpiredAt string           `json:"ExpiredAt,omitempty"`
	Enabled   bool             `json:"Enabled,omitempty"`
	Cron      *TestxReportCron `json:"Cron,omitempty"`
	Namespace string           `json:"Namespace,omitempty"`
}

// TestxReport testx 报告
type TestxReport struct {
	Audit           *TestxAudit               `json:"Audit,omitempty"`
	Uid             string                    `json:"Uid,omitempty"`
	Title           string                    `json:"Title,omitempty"`
	Namespace       string                    `json:"Namespace,omitempty"`
	Summary         string                    `json:"Summary,omitempty"`
	Plans           []TestxReportPlan         `json:"Plans,omitempty"`
	Source          string                    `json:"Source,omitempty"`
	TemplateUid     string                    `json:"TemplateUid,omitempty"`
	NotificationUid string                    `json:"NotificationUid,omitempty"`
	Tasks           []TestxReportPeriodicTask `json:"Tasks,omitempty"`
	Nid             string                    `json:"Nid,omitempty"`
	Stat            interface{}               `json:"Stat,omitempty"`
}

// TestxReportTemplate testx 报告模板
type TestxReportTemplate struct {
	Audit            *TestxAudit `json:"Audit,omitempty"`
	Uid              string      `json:"Uid,omitempty"`
	Title            string      `json:"Title,omitempty"`
	Namespace        string      `json:"Namespace,omitempty"`
	IsSystem         bool        `json:"IsSystem,omitempty"`
	Desc             string      `json:"Desc,omitempty"`
	AsDefault        bool        `json:"AsDefault,omitempty"`
	Order            uint32      `json:"Order,omitempty"`
	NotificationUid  string      `json:"NotificationUid,omitempty"`
	NotificationName string      `json:"NotificationName,omitempty"`
	Modules          []string    `json:"Modules,omitempty"`
	Count            uint32      `json:"Count,omitempty"`
}

// TestxReportData testx 报告详情数据（结构复杂，使用 json.RawMessage）
type TestxReportData struct {
	Raw json.RawMessage
}

// ---------------------------------------------------------------------------
// 请求结构体
// ---------------------------------------------------------------------------

// TestxListReportsRequest 获取报告列表请求
type TestxListReportsRequest struct {
	Namespace      string         `json:"-"`
	PageInfo       *TestxPageInfo `json:"-"`
	Search         string         `json:"-"`
	StartAt        string         `json:"-"`
	EndAt          string         `json:"-"`
	Creators       []string       `json:"-"`
	PlanUids       []string       `json:"-"`
	WithAssociated bool           `json:"-"`
	TemplateUid    string         `json:"-"`
	Source         string         `json:"-"`
	Sources        []string       `json:"-"`
}

// TestxGetReportRequest 获取报告详情请求
type TestxGetReportRequest struct {
	Namespace string `json:"-"`
	Uid       string `json:"-"`
}

// TestxGetReportDataRequest 获取报告详情数据请求
type TestxGetReportDataRequest struct {
	Namespace   string `json:"-"`
	ReportUid   string `json:"-"`
	TemplateUid string `json:"-"`
}

// TestxListReportTemplatesRequest 获取报告模板列表请求
type TestxListReportTemplatesRequest struct {
	Namespace string         `json:"-"`
	PageInfo  *TestxPageInfo `json:"-"`
}
