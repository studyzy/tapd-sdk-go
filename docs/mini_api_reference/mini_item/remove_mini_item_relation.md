# 解除工作项与其他业务对象的关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/remove_mini_item_relation.html

-   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
-   [返回结果](#返回结果)

# 说明

解除工作项与其他业务对象的关联关系

# url

`https://api.tapd.cn/mini_items/remove_mini_item_relation`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许解除一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">mini_item_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">工作项id</td></tr><tr><td style="text-align:center;">target_type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">业务对象类型</td></tr><tr><td style="text-align:center;">target_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">业务对象id</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=69993260&mini_item_id=1069993260856110919&target_type=bug&target_id=1010022001500746693' 'https://api.tapd.cn/mini_items/remove_mini_item_relation'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' -d 'workspace_id=69993260&mini_item_id=1069993260856110919&bug_id=1010022001500746693' 'https://api.tapd.cn/mini_items/remove_mini_item_relation'`

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

