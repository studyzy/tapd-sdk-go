# 创建发布评审依据

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/add_launch_accessories.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [返回结果](#返回结果)
-   [发布评审依据字段重要字段说明](#发布评审依据字段重要字段说明)

# 说明

创建发布评审依据，返回创建发布评审依据的数据

# url

`https://api.tapd.cn/launch_accessories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">form_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">发布评审ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">类型（仅支持launch_url）</td></tr><tr><td style="text-align:center;">content</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">url地址</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&form_id=1010104801079533889&type=launch_url&content=https://www.tapd.cn/' 'https://api.tapd.cn/launch_accessories'`

## 返回结果

```
{
    "status": 1,
    "data": {
        "LaunchAccessory": {
            "id": "1010104801000254035",
            "form_id": "1010104801079533889",
            "workspace_id": "10104801",
            "type": "launch_url",
            "tag": "",
            "title": "URL",
            "content": "https://www.tapd.cn/",
            "description": null,
            "content_type": "",
            "created_by": "tapd",
            "created": "2022-09-08 16:45:30",
            "group_id": null,
            "source": ""
        }
    },
    "info": "success"
}
```


# 发布评审依据字段说明

## 发布评审依据字段重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">依据ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属项目ID</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">form_id</td><td style="text-align:center;">发布评审ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">content</td><td style="text-align:center;">url地址</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">created_by</td><td style="text-align:center;">创建人</td></tr></tbody></table>
