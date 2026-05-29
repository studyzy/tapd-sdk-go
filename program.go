package tapd

import (
	"context"

	"github.com/studyzy/tapd-sdk-go/model"
)

// BindProgramEntities 项目集批量关联/取消关联业务对象（story/bug）
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/program/program_bind_entities.html
func (c *Client) BindProgramEntities(ctx context.Context, req *model.BindEntitiesRequest) error {
	_, err := c.doPost(ctx, "/program/program_bind_entities", req.ToParams())
	return err
}

// RelateProgramWorkspace 项目集批量关联/取消关联、修改授权范围项目
// API 文档：https://open.tapd.cn/document/api-doc/API文档/api_reference/program/program_relate_workspace.html
func (c *Client) RelateProgramWorkspace(ctx context.Context, req *model.RelateWorkspaceRequest) error {
	_, err := c.doPost(ctx, "/program/program_relate_workspace", req.ToParams())
	return err
}
