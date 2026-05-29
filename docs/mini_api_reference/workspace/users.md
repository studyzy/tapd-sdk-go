# 获取项目成员列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/workspace/users.html

-   [获取空间成员列表](#获取空间成员列表)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [空间相关重要字段说明](#空间相关重要字段说明)
-   [用户组ID说明](#用户组id说明)

# 说明

返回符合查询条件的所有空间成员（无分页）

# url

`https://api.tapd.cn/workspaces/users`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值 一次只能查一个空间

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">空间id 或者 公司id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需要查的字段值</td><td style="text-align:center;">user,role_id,email 可选，以,分隔</td></tr></tbody></table>

# 调用示例及返回结果

## 获取空间成员列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/users?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "UserWorkspace": {
                "user": "david",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "zhijie",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "orange",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "ajun",
                "role_id": [
                    "1000000000000000012"
                ]
            }
        }
    ],
    "info": "success"
}
```


# 空间相关接口字段说明

## 空间相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">user</td><td style="text-align:center;">成员昵称</td></tr><tr><td style="text-align:center;">role_id</td><td style="text-align:center;">用户组ID</td></tr><tr><td style="text-align:center;">email</td><td style="text-align:center;">成员邮箱号</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">真实姓名</td></tr></tbody></table>

## 用户组ID说明

<table><thead><tr><th style="text-align:center;">ID</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">1000000000000000012</td><td style="text-align:center;">项目所有者</td></tr><tr><td style="text-align:center;">1000000000000000002</td><td style="text-align:center;">空间管理员</td></tr><tr><td style="text-align:center;">1000000000000000089</td><td style="text-align:center;">普通用户</td></tr><tr><td style="text-align:center;">1000000000000000006</td><td style="text-align:center;">查看角色（无编辑权限）</td></tr></tbody></table>
