# 获取用户参与的项目列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/user_participant_projects.html

# 说明

获取用户参与的项目列表（无分页）

# url

`https://api.tapd.cn/workspaces/user_participant_projects`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值,只能传用户 nick 参数，一次只能查一个用户

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">nick</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">成员昵称</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">company_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">int</td><td style="text-align:center;">公司ID</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 获取 anyechen 参与的项目列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/user_participant_projects?nick=anyechen'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Workspace": {
                "id": "755",
                "name": "TAPD平台",
                "pretty_name": "tapd",
                "category": "product",
                "status": "normal",
                "description": "研发管理平台",
                "begin_date": "2006-04-13",
                "end_date": "2017-09-27",
                "external_on": "1",
                "creator": "",
                "created": "2007-05-01 00:00:00"
            }
        }
    ],
    "info": "success"
}
```


# 项目相关字段说明

## 项目相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">项目 id</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">项目名称</td></tr><tr><td style="text-align:center;">pretty_name</td><td style="text-align:center;">项目英文昵称</td></tr><tr><td style="text-align:center;">category</td><td style="text-align:center;">项目类别</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">项目状态</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">项目描述</td></tr><tr><td style="text-align:center;">begin_date</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">external_on</td><td style="text-align:center;">是否开通外网</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">项目创建者的名字</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">项目创建时间</td></tr></tbody></table>
