# 获取评论

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/comment/get_comments.html

-   [获取空间下评论](#获取空间下评论)
    -   [curl 调用示例](#curl-调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取某条工作项的评论](#获取某条工作项的评论)
    -   [curl 调用示例](#curl-调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [评论重要字段说明](#评论重要字段说明)

# 说明

返回符合查询条件的所有评论（分页显示，默认一页30条）

# url

`https://api.tapd.cn/comments`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">内容</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论所属工作项id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后更改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">root_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">根评论ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">reply_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论回复的ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取空间下评论

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/comments?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Comment": {
                "id": "1069993260058859425",
                "title": "在状态 [未开始] 添加",
                "description": "<p>test1</p>",
                "author": "orangecyang",
                "entry_type": "mini_items",
                "entry_id": "1069993260855484441",
                "reply_id": "0",
                "root_id": "0",
                "created": "2023-03-16 14:19:34",
                "modified": "2023-03-16 14:19:34",
                "workspace_id": "69993260"
            }
        },
        {
            "Comment": {
                "id": "1069993260058848877",
                "title": "在状态 [未开始] 添加",
                "description": "<p>test2</p>",
                "author": "orangecyang",
                "entry_type": "mini_items",
                "entry_id": "1069993260855310177",
                "reply_id": "0",
                "root_id": "0",
                "created": "2023-01-13 15:29:24",
                "modified": "2023-01-13 15:29:24",
                "workspace_id": "69993260"
            }
        },
        {
            "Comment": {
                "id": "1069993260058848801",
                "title": "在状态 [未开始] 添加",
                "description": "<p>test3</p>",
                "author": "orangecyang",
                "entry_type": "mini_items",
                "entry_id": "1069993260855310179",
                "reply_id": "0",
                "root_id": "0",
                "created": "2023-01-13 14:44:04",
                "modified": "2023-01-13 14:44:04",
                "workspace_id": "69993260"
            }
        }
    ],
    "info": "success"
}
```


## 获取某条工作项的评论

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments?entry_id=1069993260856110919&limit=2&workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/comments?entry_id=1069993260856110919&limit=2&workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Comment": {
                "id": "1069993260059089547",
                "title": "在状态 [已完成] 添加",
                "description": "<p>测试一下</p>",
                "author": "orangecyang",
                "entry_type": "mini_items",
                "entry_id": "1069993260856110919",
                "reply_id": "0",
                "root_id": "0",
                "created": "2023-07-18 21:39:34",
                "modified": "2023-07-18 21:39:34",
                "workspace_id": "69993260"
            }
        }
    ],
    "info": "success"
}
```


# 评论字段说明

## 评论重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评论ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">内容</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">评论人</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">评论类型</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">评论所属工作项id</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后更改时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">root_id</td><td style="text-align:center;">根评论ID</td></tr><tr><td style="text-align:center;">reply_id</td><td style="text-align:center;">评论回复的ID</td></tr></tbody></table>
