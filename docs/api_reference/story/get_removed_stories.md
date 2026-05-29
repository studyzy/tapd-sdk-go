# 获取回收站的需求

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_removed_stories.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [回收站重要字段说明](#回收站重要字段说明)

# 说明

返回符合查询条件的所有回收站中的需求

# url

`https://api.tapd.cn/stories/get_removed_stories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">is_archived</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否为归档。默认取 0，为不返回归档的需求。传 is_archived=1 参数则仅返回归档的需求</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">deleted</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">删除时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_removed_stories?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "RemovedStory": {
                "id": "1010104801854921589",
                "name": "cat",
                "creator": "tapd",
                "created": "2021-08-25 15:37:16",
                "operation_user": "v_xuanfang",
                "deleted": "2021-09-13 16:48:16",
                "is_archived": "0"
            }
        },
        {
            "RemovedStory": {
                "id": "1010104801854923037",
                "name": "嗷嗷嗷",
                "creator": "v_xuanfang",
                "created": "2021-08-31 11:01:51",
                "operation_user": "anyechen",
                "deleted": "2021-09-14 15:46:11",
                "is_archived": "0"
            }
        }
    ],
    "info": "success"
}
```


# 回收站字段说明

## 回收站重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">需求ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">operation_user</td><td style="text-align:center;">删除人</td></tr><tr><td style="text-align:center;">deleted</td><td style="text-align:center;">删除时间</td></tr><tr><td style="text-align:center;">is_archived</td><td style="text-align:center;">是否为归档</td></tr></tbody></table>
