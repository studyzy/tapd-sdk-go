# 获取评论数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/comment/get_comments_count.html

-   [获取空间下评论数量](#获取空间下评论数量)
    -   [curl 调用示例](#curl-调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取某条工作项评论的数量](#获取某条工作项评论的数量)
    -   [curl 调用示例](#curl-调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

计算符合查询条件的评论数量并返回

# url

`https://api.tapd.cn/comments/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回评论数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">内容</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论所依附的业务对象实体id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后更改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取空间下评论数量

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments/count?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/comments/count?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 67
    },
    "info": "success"
}
```


## 获取某条工作项评论的数量

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments/count?entry_id=1010104801074085199&amp;workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/comments/count?entry_id=1010104801074085199&amp;workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 5
    },
    "info": "success"
}
```

