# 获取工作项动态数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/get_mini_item_changes_count.html

-   [获取空间下工作项动态数量](#获取空间下工作项动态数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取工作项ID为 1069993260856110917 的工作项动态数量](#获取工作项id为-1069993260856110917-的工作项动态数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

计算符合查询条件的工作项动态数量并返回

# url

`https://api.tapd.cn/mini_item_changes/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回工作项动态数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">动态id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">mini_item_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">工作项ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人（操作人）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间（变更时间）</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">工作项变更描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">changes</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更详细记录</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取空间下工作项动态数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_item_changes/count?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_item_changes/count?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 21
    },
    "info": "success"
}
```


## 获取工作项ID为 1069993260856110917 的工作项动态数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_item_changes/count?workspace_id=69993260&mini_item_id=1069993260856110917'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_item_changes/count?workspace_id=69993260&mini_item_id=1069993260856110917'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 1
    },
    "info": "success"
}
```

