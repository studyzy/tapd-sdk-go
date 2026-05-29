# 创建特性接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/add_feature.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [返回结果](#返回结果)
-   [特性字段重要字段说明](#特性字段重要字段说明)

# 说明

新建特性，返回新建特性的数据

# url

`https://api.tapd.cn/features`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&name=创建1' 'https://api.tapd.cn/features'`

## 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Feature": {
                "id": "1010104801000049305",
                "name": "创建1",
                "description": "jx",
                "created": "2020-11-24 16:58:15",
                "owner": "",
                "due": "1970-01-01",
                "completed": "0",
                "default": "0",
                "module_id": "1010104801007551425",
                "release_id": "0",
                "release_name": "",
                "priority": "0",
                "estimate": "0",
                "estimate_completed": "0",
                "modified": "2020-11-30 17:26:39",
                "workspace_id": "10104801"
            }
        }
    ],
    "info": "success"
}
```


# 特性字段说明

## 特性字段重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">完成状态</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">预计结束时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">module_id</td><td style="text-align:center;">模块id</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目id</td></tr></tbody></table>
