# 设置启用工作日历

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/enable_work_calendar.html

# 说明

设置启用工作日历

# url

`https://api.tapd.cn/workspaces/enable_work_calendar`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次只能启用一个项目下的工作日历

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">启用类型。可选范围system或custom</td></tr></tbody></table>

# 调用示例及返回结果

## 设置启用工作日历

### curl 使用 Basic Auth 鉴权调用示例

`curl -u "api_user:api_password" -d 'workspace_id=48464494&type=system' 'https://api.tapd.cn/workspaces/enable_work_calendar'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "success": true
    },
    "info": "success"
}
```

