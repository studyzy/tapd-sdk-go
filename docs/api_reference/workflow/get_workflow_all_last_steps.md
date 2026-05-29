# 获取所有结束状态

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/get_workflow_all_last_steps.html

# 说明

获取工作流所有结束状态

# url

`https://api.tapd.cn/workflows/all_last_steps`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个项目的工作流所有结束状态

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">system</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">系统名。目前只支持 story（需求的）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">group_key</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">分组字段，可选字段 workflow_id(工作流ID) 或 workitem_type_id (需求类别ID)</td><td style="text-align:center;">默认按workitem_type_id分组</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">节点类型。默认为 status，返回起始状态。step 返回结束节点，仅并行工作流支持。若需要同时返回结束状态和结束节点，支持数组type[]=status&amp;type[]=step</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目缺陷工作流结束状态中英文名对应关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows/all_last_steps?system=story&workspace_id=10104801&group_key=workitem_type_id'`

`curl 'https://api.tapd.cn/workflows/all_last_steps?system=story&workspace_id=10104801&group_key=workitem_type_id&access_token=ACCESS_TOKEN'`

### 返回结果

```
{
  "status": 1,
  "data": {
    "1069990230000187079": {
      "status_70": "已完成",
      "status_69": "未开始"
    },
    "1069990230000131609": {
      "done": "已完成-任务"
    }
  },
  "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">键</td><td style="text-align:center;">工作流ID 或者 需求类别ID , 根据group_key确定</td></tr><tr><td style="text-align:center;">值</td><td style="text-align:center;">状态英文名 : 状态中文名</td></tr></tbody></table>
