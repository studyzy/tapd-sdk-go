# 获取视图对应的缺陷列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs_by_view_conf_id.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的缺陷数据](#获取项目下的缺陷数据)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

返回视图下最新的30条缺陷（分页显示，默认一页30条）

# url

`https://api.tapd.cn/bugs/get_bugs_by_view_conf_id`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">view_conf_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">视图ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">当前登录用户视图</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的缺陷数据

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/get_bugs_by_view_conf_id?view_conf_id=1010104801030259563&workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Bug": {
                "id": "1010104801084955735",
                "title": "test",
                "version_report": "",
                "severity": "",
                "priority": "",
                "status": "new",
                "current_owner": "",
                "created": "2021-01-21 11:02:09",
                "reporter": "v_xuanfang",
                "resolution": ""
            }
        },
        {
            "Bug": {
                "id": "1010104801083011055",
                "title": "aaa",
                "version_report": "",
                "severity": "",
                "priority": "",
                "status": "new",
                "current_owner": "",
                "created": "2020-10-28 18:28:27",
                "reporter": "v_xuanfang",
                "resolution": ""
            }
        },
        {
            "Bug": {
                "id": "1010104801082968057",
                "title": "aaaa",
                "version_report": "",
                "severity": "",
                "priority": "",
                "status": "in_progress",
                "current_owner": "",
                "created": "2020-10-27 10:14:20",
                "reporter": "v_xuanfang",
                "resolution": ""
            }
        }
    ],
    "info": {
        "total": 13,
        "current_page": 1,
        "page_size": 3,
        "total_page": 5
    }
}
```


# 缺陷字段说明

缺陷字段说明，请参考 [缺陷说明](/document/api-doc/API文档/api_reference/bug/bug.html)
