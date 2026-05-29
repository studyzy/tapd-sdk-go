# 获取自定义工作日历详情

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_custom_work_calendar.html

# 说明

获取自定义工作日历详情

# url

`https://api.tapd.cn/workspaces/get_custom_work_calendar`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目id</td></tr><tr><td style="text-align:center;">year</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">表示哪一年</td></tr></tbody></table>

# 调用示例及返回结果

## 获取自定义工作日历详情

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/get_custom_work_calendar?workspace_id=48464494&year=2025'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "weekdays": [
            "1",
            "2",
            "3",
            "4",
            "5",
            "6",
            "7"
        ],
        "holidays": [
            "2025-01-01"
        ],
        "workdays": [
            "2025-01-02",
            "2025-01-03",
            "2025-01-04"
        ]
    },
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">weekdays</td><td style="text-align:center;">所设置的一周内的哪些天</td></tr><tr><td style="text-align:center;">holidays</td><td style="text-align:center;">假日</td></tr><tr><td style="text-align:center;">workdays</td><td style="text-align:center;">工作日</td></tr></tbody></table>
