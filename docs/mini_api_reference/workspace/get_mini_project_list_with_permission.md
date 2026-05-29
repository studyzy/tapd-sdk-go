# 获取用户所有参与的空间

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/workspace/get_mini_project_list_with_permission.html

-   [获取orange参与的所有有编辑权限的空间列表](#获取orange参与的所有有编辑权限的空间列表)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [空间相关重要字段说明](#空间相关重要字段说明)

# 说明

获取用户所有参与的空间

# url

`https://api.tapd.cn/workspaces/get_mini_project_list_with_permission`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">user</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">用户名</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">company_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">公司ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">can_edit</td><td style="text-align:center;">否</td><td style="text-align:center;">boolean</td><td style="text-align:center;">是否过滤出有编辑权限的空间，默认值false</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取orange参与的所有有编辑权限的空间列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/get_mini_project_list_with_permission?user=orange&can_edit=1'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/workspaces/get_mini_project_list_with_permission?user=orange&can_edit=1'`

### 返回结果

```
{
    "status": 1,
    "data": {
        {
            "created": "2021-09-27 20:00:24",
            "creator": "orange",
            "id": "20375687",
            "name": "mini_test_09",
            "status": "normal"
        },        
        {
            "created": "2022-11-16 15:57:17",
            "creator": "orange",
            "id": "69992082",
            "name": "webhook测试",
            "status": "normal"
        },
    },
    "info": "success"
}
```


# 空间相关字段说明

## 空间相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">空间 id</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">空间名称</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">空间状态: normal 正常，closed 关闭</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">空间创建者的名字</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">空间创建时间</td></tr></tbody></table>
