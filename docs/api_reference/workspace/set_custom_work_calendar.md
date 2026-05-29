# 设置自定义工作日历

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/set_custom_work_calendar.html

# 说明

设置自定义工作日历

# url

`https://api.tapd.cn/workspaces/set_custom_work_calendar`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">year</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">年</td></tr><tr><td style="text-align:center;">weekdays</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">array</td><td style="text-align:center;">表示周几，赋值范围为1-7，传入的值代表设置周几为工作日（如果不传默认周一到周五为工作日）</td></tr><tr><td style="text-align:center;">holidays</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">array</td><td style="text-align:center;">假日（在weekdays的基础上，设置哪天为额外的节假日）</td></tr><tr><td style="text-align:center;">workdays</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">array</td><td style="text-align:center;">工作日 (在weekdays的基础上，设置哪天为额外的工作日）</td></tr></tbody></table>

# 调用示例及返回结果

## 设置自定义工作日历

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=48464494&year=2025&week=[1,2,3,4,5]&holidays=[“2025-01-01”]&workdays=[“2025-01-04”]' 'https://api.tapd.cn/workspaces/set_custom_work_calendar'`

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

