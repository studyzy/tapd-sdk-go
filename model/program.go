// Package model 中的 program.go 定义了 TAPD 项目集相关 API 的请求参数结构体
package model

// BindEntitiesRequest 项目集批量关联/取消关联业务对象的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/program/program_bind_entities.html
type BindEntitiesRequest struct {
	WorkspaceID string // 必填：项目集 id
	Action      string // 必填：bind/unbind
	EntityType  string // 必填：story/bug
	EntityIDs   string // 必填：关联/取消关联的 id（多个用逗号分隔）
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *BindEntitiesRequest) ToParams() map[string]string {
	return map[string]string{
		"workspace_id": r.WorkspaceID,
		"action":       r.Action,
		"entity_type":  r.EntityType,
		"entity_ids":   r.EntityIDs,
	}
}

// RelateWorkspaceRequest 项目集批量关联/取消关联、修改授权范围项目的请求参数
// 参考：https://open.tapd.cn/document/api-doc/API文档/api_reference/program/program_relate_workspace.html
type RelateWorkspaceRequest struct {
	WorkspaceID        string // 必填：项目集 id
	Action             string // 必填：bind/unbind
	RelateWorkspaceIDs string // 必填：关联/取消关联的项目 id（多个用逗号分隔）
	Auth               string // 可选：项目授权（unbind 时不生效），默认全部，可选值如 create_story,create_bug,edit_story,edit_bug
}

// ToParams 将请求结构体转换为 TAPD API 参数 map
func (r *RelateWorkspaceRequest) ToParams() map[string]string {
	params := map[string]string{
		"workspace_id":         r.WorkspaceID,
		"action":               r.Action,
		"relate_workspace_ids": r.RelateWorkspaceIDs,
	}
	setOptional(params, "auth", r.Auth)
	return params
}
