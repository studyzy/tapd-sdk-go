# 获取并行工作节点和状态的对应关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/get_workflow_step_map.html

# 说明

获取并行工作流节点信息及其与状态的归属关系。仅需求支持并行工作流

# url

`https://api.tapd.cn/workflows/step_map`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个项目的工作流的节点信息

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">system</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">系统名，当前仅支持需求 story</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求类别id</td><td style="text-align:center;">查询需求状态时需必传</td></tr></tbody></table>

# 调用示例及返回结果

## 获取并行工作节点和状态的对应关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows/step_map?system=story&workspace_id=10158231'`

### 返回结果

```json
{
    "status": 1,
    "data": [
        {
            "name": "new",
            "label": "开始阶段",
            "steps": [
                {
                    "name": "step_begin",
                    "label": "创建"
                },
                {
                    "name": "step_2970811_1",
                    "label": "待处理"
                }
            ]
        },
        {
            "name": "in_progress",
            "label": "执行阶段",
            "steps": [
                {
                    "name": "step_2970811_2",
                    "label": "执行节点1"
                },
                {
                    "name": "step_2970811_3",
                    "label": "执行节点2"
                }
            ]
        },
        {
            "name": "closed",
            "label": "结束阶段",
            "steps": [
                {
                    "name": "step_end",
                    "label": "结束"
                }
            ]
        },
        {
            "name": "workflow_suspended",
            "label": "流程挂起",
            "steps": []
        },
        {
            "name": "workflow_end",
            "label": "流程终止",
            "steps": []
        }
    ],
    "info": "success"
}
```
