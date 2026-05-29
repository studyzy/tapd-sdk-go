# 在缺陷下的工作流中增加状态

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/add_new_step_for_bug.html

# 说明

在缺陷下的工作流中增加状态

# url

`https://api.tapd.cn/workflows/add_new_step_for_bug`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workflow_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">工作流ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">step_name</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">状态名称</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 在缺陷下的工作流中增加状态

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=48464494&workflow_id=1148464494001000011&step_name=新增状态' 'https://api.tapd.cn/workflows/add_new_step_for_bug'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "sys_name": "status_10",
        "step_name": "新增状态"
    },
    "info": "success"
}
```
