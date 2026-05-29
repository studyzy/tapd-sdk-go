# 获取公司项目列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/projects.html

# 说明

返回符合查询条件的所有项目（无分页）

# url

`https://api.tapd.cn/workspaces/projects`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值,只能传公司参数，一次只能查一个公司

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">company_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">公司ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">category</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">项目类型, project:项目协作, mini_project:轻协作</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">with_extends</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">值=1可以返回自定义字段</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 获取公司项目列表

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/projects?company_id=20003261&category=project'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Workspace": {
                "id": "20026861",
                "name": "产品运营2015",
                "pretty_name": "20026861",
                "description": null,
                "status": "normal",
                "parent_id": "10104550",
                "category": "project",
                "secrecy": "0",
                "created": "2016-03-10 17:01:45",
                "creator_id": "0",
                "creator": "张贺",
                "begin_date": "2019-08-14",
                "end_date": "2019-07-15",
                "member_count": 271
            }
        },
        {
            "Workspace": {
                "id": "20033151",
                "name": "用户支持体系",
                "pretty_name": "20033151",
                "description": "",
                "status": "normal",
                "parent_id": "20026861",
                "category": "project",
                "secrecy": "0",
                "created": "2016-04-18 10:56:45",
                "creator_id": "0",
                "creator": "果子不吃果子",
                "begin_date": null,
                "end_date": null,
                "member_count": 22
            }
        }
    ],
    "info": "success"
}
```


# 项目字段说明

## 项目重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">项目 id</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">项目名称</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">项目状态</td></tr><tr><td style="text-align:center;">category</td><td style="text-align:center;">项目类型</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">父项目ID</td></tr><tr><td style="text-align:center;">begin_date</td><td style="text-align:center;">项目开始时间</td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">项目结束时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">项目创建时间</td></tr><tr><td style="text-align:center;">creator_id</td><td style="text-align:center;">项目创建者 id</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">项目创建者的名字和邮箱</td></tr><tr><td style="text-align:center;">member_count</td><td style="text-align:center;">项目人数</td></tr><tr><td style="text-align:center;">WorkspaceExtends</td><td style="text-align:center;">下列数据为自定义字段</td></tr></tbody></table>
