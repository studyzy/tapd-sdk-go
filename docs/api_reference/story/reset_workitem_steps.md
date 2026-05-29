# 流程重置

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/reset_workitem_steps.html

# 说明

并行工作流流程重置

# url

`https://api.tapd.cn/stories/reset_workitem_steps`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">reset_type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">重置类型（status: 按状态重置，step: 启用某个节点）</td></tr><tr><td style="text-align:center;">reset_dst</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">需启用的状态或者节点</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'story_id=1010104801871430407&reset_type=status&reset_dst=develop&workspace_id=10104801' 'https://api.tapd.cn/stories/reset_workitem_steps'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "Story": {
            "id": "1010104801871430407",
            "workitem_type_id": "1010104801000022091",
            "name": "aatt",
            "description": null,
            "workspace_id": "10104801",
            "creator": "v_xuanfang",
            "created": "2022-01-07 15:21:15",
            "modified": "2022-01-11 16:57:11",
            "status": "develop",
            "owner": "",
            "cc": "",
            "begin": null,
            "due": null,
            "size": null,
            "priority": "",
            "developer": "",
            "iteration_id": "0",
            "test_focus": "",
            "type": "",
            "source": "",
            "module": "",
            "version": "",
            "completed": null,
            "category_id": "-1",
            "path": "1010104801870636009:1010104801871430407:",
            "parent_id": "1010104801870636009",
            "children_id": "|",
            "ancestor_id": "1010104801870636009",
            "business_value": null,
            "effort": "0",
            "effort_completed": "4",
            "exceed": "10",
            "remain": "6",
            "release_id": "0",
            "confidential": "N",
            "templated_id": null,
            "custom_field_one": "",
            "custom_field_two": "",
            "custom_field_three": "",
            "custom_field_four": "",
            "custom_field_five": ""
        }
    },
    "info": "success"
}
```
