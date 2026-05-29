# 获取关联缺陷

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/get_related_bugs.html

-   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
-   [返回结果](#返回结果)

# 说明

返回符合查询条件的所有工作项关联的缺陷ID

# url

`https://api.tapd.cn/mini_items/get_related_bugs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有关系

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">mini_item_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">工作项ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_items/get_related_bugs?workspace_id=69993260&mini_item_id=1069993260856110919'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_items/get_related_bugs?workspace_id=69993260&mini_item_id=1069993260856110919'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "type": "direct_relate",
            "id": "1069990800856102923",
            "workspace_id": "69993260",
            "actas": "target",
            "created": "2023-07-18 19:07:27",
            "modified": "2023-07-18 19:07:27",
            "linked_workspace_id": 69990800
        }
    ],
    "info": "success"
}
```


# 字段说明

## 重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">type</td><td style="text-align:center;">关系类型</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">关联的缺陷ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">关联操作发生的空间ID</td></tr><tr><td style="text-align:center;">linked_workspace_id</td><td style="text-align:center;">关联缺陷所在的空间ID</td></tr><tr><td style="text-align:center;">actas</td><td style="text-align:center;">角色。target 为操作发起方</td></tr></tbody></table>

## 关联类型说明

<table><thead><tr><th style="text-align:center;">关联类型</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">direct_relate</td><td style="text-align:center;">直接关联</td></tr><tr><td style="text-align:center;">copy</td><td style="text-align:center;">转缺陷或缺陷转工作项</td></tr><tr><td style="text-align:center;">sync_copy</td><td style="text-align:center;">转缺陷或缺陷转工作项(同步字段)</td></tr></tbody></table>
