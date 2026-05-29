# 添加工作项与其他业务对象的关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/create_mini_item_relation.html

-   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)

# 说明

创建需求与缺陷关联关系

# URL

`https://api.tapd.cn/mini_items/create_mini_item_relation`

# 支持格式

JSON/XML (默认JSON格式)

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">所属TAPD空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">target_type</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">关联关系目标对象类型（story、bug、mini_item）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">source_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">关联关系源对象id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">target_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">关联关系目标对象id</td><td style="text-align:center;"></td></tr></tbody></table>

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=69993260&target_type=bug&target_id=1010022001500746693&source_id=1069993260856110919' 'https://api.tapd.cn/mini_items/create_mini_item_relation'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -u 'Authorization: Bearer ACCESS_TOKEN' -d 'workspace_id=69993260&target_type=bug&target_id=1010022001500746693&source_id=1069993260856110919' 'https://api.tapd.cn/mini_items/create_mini_item_relation'`

# 返回结果

以返回结果为准

```
{
    "status": 1,
    "data": {
        "success": true
    },
    "info": "success"
}
```

