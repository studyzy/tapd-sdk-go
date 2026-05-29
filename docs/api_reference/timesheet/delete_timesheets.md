# 删除工时花费

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/delete_timesheets.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [删除工时花费](#删除工时花费)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

删除工时花费

# url

`https://api.tapd.cn/timesheets/delete_timesheets`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

删除工时花费，一次最多支持删除一百条数据。

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">对象类型，如story、task、bug等</td></tr><tr><td style="text-align:center;">entity_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">对象ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">cost_ids</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">array</td><td style="text-align:center;">花费id集合</td></tr></tbody></table>

# 调用示例及返回结果

## 删除工时花费

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'entity_type=story&entity_id=1148464494001000040&cost_ids[]=1148464494001000097&cost_ids[]=1148464494001000099&workspace_id=48464494' 'https://api.tapd.cn/timesheets/delete_timesheets'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "msg": "delete completed",
        "data": {
            "success": {
                "cost_ids": [
                    "1148464494001000111"
                ],
                "msg": "delete success"
            },
            "failed": [
                {
                    "cost_ids": [
                        "1148464494001000106"
                    ],
                    "msg": "the record does not belong to the specified entity (entity_type:story, entity_id:1148464494001000137)"
                },
                {
                    "cost_ids": [
                        "1148464494001000110",
                        "1148464494001000076",
                        "1141659004001000035",
                        "1141659004001000037"
                    ],
                    "msg": " the record does not exist in workspace 48464494"
                }
            ]
        }
    },
    "info": "success"
}
```


<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">success</td><td style="text-align:center;">删除成功的记录</td></tr><tr><td style="text-align:center;">failed</td><td style="text-align:center;">删除失败的记录</td></tr><tr><td style="text-align:center;">cost_ids</td><td style="text-align:center;">花费id集合</td></tr><tr><td style="text-align:center;">msg</td><td style="text-align:center;">具体的操作信息</td></tr></tbody></table>
