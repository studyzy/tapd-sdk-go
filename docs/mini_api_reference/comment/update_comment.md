# 更新评论

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/comment/update_comment.html

-   [curl 调用示例](#curl-调用示例)
-   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
-   [返回结果](#返回结果)

# 说明

更新评论，返回更新评论的数据

# url

`https://api.tapd.cn/comments`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">description</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">内容</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">评论id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">change_creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更人</td><td style="text-align:center;"></td></tr></tbody></table>

### curl 调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=69993260&description=<p>test1</p>&id=1069993260059089555' 'https://api.tapd.cn/comments'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' -d 'workspace_id=69993260&description=<p>test1</p>&id=1069993260059089555' 'https://api.tapd.cn/comments'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Comment": {
            "id": "1069993260059089555",
            "title": "在状态 [已完成] 添加",
            "description": "<p>test1</p>",
            "author": "orangecyang",
            "entry_type": "mini_items",
            "entry_id": "1069993260856110919",
            "reply_id": "0",
            "root_id": "0",
            "created": "2023-07-18 22:16:20",
            "modified": "2023-07-18 22:20:27",
            "workspace_id": "69993260"
        }
    },
    "info": "success"
}
```


# 评论字段说明

## 评论重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评论ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">内容</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">评论人</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">评论类型</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">评论所属工作项id</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后更改时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">空间ID</td></tr></tbody></table>
