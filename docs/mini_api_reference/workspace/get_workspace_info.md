# 获取空间信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/workspace/get_workspace_info.html

-   [获取当前用户参与的空间列表](#获取当前用户参与的空间列表)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [空间相关重要字段说明](#空间相关重要字段说明)

# 说明

获取空间信息（无分页）

# url

`https://api.tapd.cn/workspaces/get_workspace_info`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 获取当前用户参与的空间列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/get_workspace_info?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/workspaces/get_workspace_info?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Workspace": {
            "id": "69993260",
            "name": "api测试",
            "category": "mini_project",
            "status": "normal",
            "description": "",
            "product_type": "",
            "begin_date": "2022-12-22",
            "end_date": null,
            "creator": "orangecyang",
            "created": "2022-12-22 14:29:17"
        }
    },
    "info": "success"
}
```


# 空间相关字段说明

## 空间相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">空间 id</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">空间名称</td></tr><tr><td style="text-align:center;">category</td><td style="text-align:center;">空间类别</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">空间状态: normal 正常，closed 关闭</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">空间描述</td></tr><tr><td style="text-align:center;">begin_date</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">空间创建者的名字</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">空间创建时间</td></tr><tr><td style="text-align:center;">product_type</td><td style="text-align:center;">产品类型</td></tr></tbody></table>
