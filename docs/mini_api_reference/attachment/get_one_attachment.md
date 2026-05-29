# 获取单个附件下载链接

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/attachment/get_one_attachment.html

-   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
-   [返回结果](#返回结果)
-   [返回字段说明](#返回字段说明)

# 说明

获取单个附件下载链接

# url

`https://api.tapd.cn/attachments/down`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

每次只能获取一个附件下载链接

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">附件ID</td></tr></tbody></table>

# 调用示例及返回结果

## curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/attachments/down?workspace_id=69993260&id=1069993260504708081'`

## curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/attachments/down?workspace_id=69993260&id=1069993260504708081'`

## 返回结果

```
{
    "status": 1,
    "data": {
        "Attachment": {
            "id": "1069993260504708081",
            "type": "story_custom_field",
            "entry_id": "1069993260856110917",
            "filename": "70077675 (1).zip",
            "description": "",
            "content_type": "application/zip",
            "created": "2023-07-19 15:07:15",
            "workspace_id": "69993260",
            "owner": "orangecyang",
            "download_url": "http://file.tapd.oa.com/attachments/tmp_download/ca30fd7a50e4b9bae194424307b10949?salt=acadb52e4c397fe3f3f1e1c12372d938&time=163642746"
        }
    },
    "info": "success"
}
```


# 字段说明

## 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">附件ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型(固定为story_custom_field)</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">工作项的id</td></tr><tr><td style="text-align:center;">filename</td><td style="text-align:center;">文件名</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">附件描述</td></tr><tr><td style="text-align:center;">content_type</td><td style="text-align:center;">文件类型</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">空间id</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">附件上传人</td></tr><tr><td style="text-align:center;">download_url</td><td style="text-align:center;">下载链接</td></tr></tbody></table>
