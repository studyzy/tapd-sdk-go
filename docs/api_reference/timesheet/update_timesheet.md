# 更新工时花费

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/update_timesheet.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [更新工时花费记录ID为 1010158231001169003 的花费成为 5](#更新工时花费记录id为-1010158231001169003-的花费成为-5)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [花费工时重要字段说明](#花费工时重要字段说明)

# 说明

更新花费工时，返回花费工时更新后的数据

# url

`https://api.tapd.cn/timesheets`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">timespent</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费工时</td></tr><tr><td style="text-align:center;">timeremain</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">剩余工时</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费描述</td></tr></tbody></table>

# 调用示例及返回结果

## 更新工时花费记录ID为 1010158231001169003 的花费成为 5

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'id=1010158231001169003&timespent=3&workspace_id=10158231' 'https://api.tapd.cn/timesheets'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Timesheet": {
            "id": "1010158231001169003",
            "entity_type": "story",
            "entity_id": "1010158231500709717",
            "timespent": "3",
            "spentdate": "2020-05-05",
            "owner": "anyechen",
            "created": "2020-05-06 22:08:35",
            "workspace_id": "10158231",
            "memo": "hey"
        }
    },
    "info": "success"
}
```


# 花费工时字段说明

## 花费工时重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">对象类型，如story、task等</td></tr><tr><td style="text-align:center;">entity_id</td><td style="text-align:center;">对象ID</td></tr><tr><td style="text-align:center;">timespent</td><td style="text-align:center;">花费工时</td></tr><tr><td style="text-align:center;">timeremain</td><td style="text-align:center;">剩余工时</td></tr><tr><td style="text-align:center;">spentdate</td><td style="text-align:center;">花费日期</td></tr></tbody></table>
