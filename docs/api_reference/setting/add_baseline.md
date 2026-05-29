# 创建基线接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/add_baseline.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [返回结果](#返回结果)
-   [基线重要字段说明](#基线重要字段说明)

# 说明

新建基线，返回新建基线的数据

# url

`https://api.tapd.cn/baselines`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">version_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">版本</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&name=创建基线' 'https://api.tapd.cn/baselines'`

## 返回结果

```
{
    "status": 1,
    "data": {
        "Baseline": {
            "id": "1010104801101308345",
            "name": "创建基线",
            "description": "",
            "created": "2020-12-08 16:22:36",
            "owner": "",
            "due": null,
            "completed": "0",
            "default": "0",
            "version_id": "0",
            "module_id": "0",
            "svn_tag": "0",
            "svn_project_id": "0",
            "svn_path_id": "0",
            "svn_sync_type": "",
            "workspace_id": "10104801"
        }
    },
    "info": "success"
}
```


# 基线字段说明

## 基线重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">version_id</td><td style="text-align:center;">版本ID</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">0（未完成）1（完成状态）</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr></tbody></table>
