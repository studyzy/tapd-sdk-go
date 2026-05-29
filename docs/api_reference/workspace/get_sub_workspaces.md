# 获取子项目信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_sub_workspaces.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [项目相关重要字段说明](#项目相关重要字段说明)

# 说明

获取子项目信息（无分页）

# url

`https://api.tapd.cn/workspaces/sub_workspaces`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">父项目ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目模板ID</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/sub_workspaces?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Workspace": {
            "id": "10104801",
            "name": "TAPD 乌云",
            "pretty_name": "tapd_security",
            "category": "product",
            "status": "normal",
            "description": "",
            "begin_date": null,
            "end_date": null,
            "external_on": "0",
            "creator": "",
            "created": "2015-03-27 16:02:02"
            "template_id": "1000000000000000020"
        }
    },
    "info": "success"
}
```


# 项目相关字段说明

## 项目相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">项目 id</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">项目名称</td></tr><tr><td style="text-align:center;">pretty_name</td><td style="text-align:center;">项目英文昵称</td></tr><tr><td style="text-align:center;">category</td><td style="text-align:center;">项目类别</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">项目状态: normal 正常，closed 关闭，suspend 挂起</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">项目描述</td></tr><tr><td style="text-align:center;">begin_date</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">external_on</td><td style="text-align:center;">是否开通外网</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">项目创建者的名字</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">项目模板ID</td></tr></tbody></table>
