# 更新项目信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/update_workspace_info.html

# 说明

更新项目信息

# url

`https://api.tapd.cn/workspaces/update_workspace_info`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次只能更新一个项目的一个字段

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">字段名。可选：description（项目描述）、begin_date（开始时间）、end_date（结束时间）、begin_end（开始结束时间）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">value</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">字段值</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 更新项目 69999237 结束时间成 2025-03-03

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=69999237&field=end_date&value=2025-03-03' 'https://api.tapd.cn/workspaces/update_workspace_info'`

### 返回结果

```json
{
    "status": 1,
    "data": "update workspace success",
    "info": "success"
}
```
