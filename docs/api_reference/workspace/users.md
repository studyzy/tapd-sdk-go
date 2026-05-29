# 获取项目成员列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/users.html

# 说明

返回符合查询条件的所有项目成员（无分页）

# url

`https://api.tapd.cn/workspaces/users`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值 一次只能查一个项目

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目id 或者 公司id</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">user</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用户昵称</td><td style="text-align:center;">支持多昵称过滤，用,分隔</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需要查的字段值</td><td style="text-align:center;">user,user_id,role_id,name,email,real_join_time 可选，以,分隔</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目成员列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/users?workspace_id=20003271'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "UserWorkspace": {
                "user": "wiki",
                "role_id": [
                    "1000000000000000002"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "zgd",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "laili",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "Leozhang",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "evonneliu",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "rafer",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "anye",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "瑞波",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "fogeen",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "lijun1",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "gingkolv",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "cleverliu",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "六月月月",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        },
        {
            "UserWorkspace": {
                "user": "二卫",
                "role_id": [
                    "1000000000000000089"
                ]
            }
        }
    ],
    "info": "success"
}
```


# 项目相关接口字段说明

## 项目相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">user</td><td style="text-align:center;">成员昵称</td></tr><tr><td style="text-align:center;">role_id</td><td style="text-align:center;">用户组ID</td></tr><tr><td style="text-align:center;">email</td><td style="text-align:center;">成员邮箱号</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">真实姓名</td></tr><tr><td style="text-align:center;">real_join_time</td><td style="text-align:center;">加入项目（公司）时间</td></tr></tbody></table>
