# 获取指定项目成员

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_workspace_users.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目成员列表](#获取项目成员列表)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取项目成员邮箱的列表](#获取项目成员邮箱的列表)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [项目相关重要字段说明](#项目相关重要字段说明)

# 说明

获取指定项目成员（无分页）

# url

`https://api.tapd.cn/workspaces/users`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值 一次只能查一个项目

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目 id</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需要查的字段值</td><td style="text-align:center;">user,role_id,email,tof_id 可选，以,分隔</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目成员列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/users?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "UserWorkspace": {
                "user": "anyechen",
                "role_id": [
                    "1000000000000000002",
                    "1000000000000000009"
                ],
                "name": "陈安业",
                "join_project_time": "2024-09-10",
                "real_join_time": "2024-09-10",
                "status": "1",
                "allocation": "100"
            }
        }
    ],
    "info": "success"
}
```


## 获取项目成员邮箱的列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/users?workspace_id=10158231&fields=user,email'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "UserWorkspace": {
                "email": "",
                "user": "anyechen"
            }
        }
    ],
    "info": "success"
}
```


# 项目相关字段说明

## 项目相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">user</td><td style="text-align:center;">成员昵称</td></tr><tr><td style="text-align:center;">role_id</td><td style="text-align:center;">成员角色 id</td></tr><tr><td style="text-align:center;">email</td><td style="text-align:center;">成员邮箱</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">中文名称</td></tr></tbody></table>
