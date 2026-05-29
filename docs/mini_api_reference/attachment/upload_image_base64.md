# 上传base64图片

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/attachment/upload_image_base64.html

-   [把本地的 uu.jpg 的 base64 上传到空间 69995768](#把本地的-uu-jpg-的-base64-上传到空间-69995768)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [返回字段说明](#返回字段说明)

# 说明

以 base64 格式上传图片，大小限制为15MB以内。

# url

`https://api.tapd.cn/files/upload_image_base64`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

-   每次只允许上传一个图片
-   文件大小限小于15MB

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">base64_data</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">图片 base64 格式数据</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">类型(固定为story_custom_field)</td></tr><tr><td style="text-align:center;">custom_field</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">字段英文名</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">工作项ID</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">附件创建人</td></tr></tbody></table>

# 调用示例及返回结果

## 把本地的 uu.jpg 的 base64 上传到空间 69995768

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -F 'workspace_id=69995768' -F 'type=story_custom_field' -F 'custom_field=custom_field_one' -F 'entry_id=1069995768115415038' -F 'base64_data=xxxxx' 'https://api.tapd.cn/files/upload_image_base64'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' -F 'workspace_id=69995768' -F 'type=story_custom_field' -F 'custom_field=custom_field_one' -F 'entry_id=1069995768115415038' -F 'base64_data=xxxx' 'https://api.tapd.cn/files/upload_image_base64'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Attachment": {
            "id": "1069995768523406033",
            "type": "story_custom_field",
            "entry_id": 1069995768115415038,
            "filename": "tapd_base64_3031935_1701845640.jpg",
            "description": "",
            "content_type": "image/png",
            "created": "2023-12-05 14:54:01"
        }
    },
    "info": "success"
}
```


# 字段说明

## 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">附件id</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型(固定为story_custom_field)</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">工作项id</td></tr><tr><td style="text-align:center;">filename</td><td style="text-align:center;">文件名</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">附件描述</td></tr><tr><td style="text-align:center;">content_type</td><td style="text-align:center;">文件类型</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr></tbody></table>
