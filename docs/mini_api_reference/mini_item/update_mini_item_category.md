# 更新分组

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/update_mini_item_category.html

-   [更新分组](#更新分组)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [分组重要字段说明](#分组重要字段说明)

# 说明

更新分组，返回更新分组的数据

# url

`https://api.tapd.cn/mini_item_categories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">分组名称</td></tr></tbody></table>

# 调用示例及返回结果

## 更新分组

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=69993260&name=test222&id=1069993260001017705' 'https://api.tapd.cn/mini_item_categories'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' -d 'workspace_id=69993260&name=test222&id=1069993260001017705' 'https://api.tapd.cn/mini_item_categories'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Category": {
            "id": "1069993260001017705",
            "workspace_id": "69993260",
            "name": "test222",
            "modified": "2023-07-18 21:11:57",
            "created": "2023-07-18 21:08:33",
            "creator": "dev",
            "modifier": "dev"
        }
    },
    "info": "success"
}
```


# 分组字段说明

## 分组重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">分组名称</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr></tbody></table>
