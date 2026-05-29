# 获取附件

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/attachment/get_attachments.html

-   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
-   [返回结果](#返回结果)

# 说明

返回符合查询条件的所有的附件

# url

`https://api.tapd.cn/attachments`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">工作项ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">filename</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">附件名称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">上传人</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/attachments?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/attachments?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Attachment": {
                "id": "1069993260504708083",
                "type": "story_custom_field",
                "entry_id": "1069993260856110917",
                "filename": "测试导出看板_需求_20230626094923.xlsx",
                "description": "",
                "content_type": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
                "created": "2023-07-19 15:07:20",
                "workspace_id": "69993260",
                "owner": "orangecyang"
            }
        },
        {
            "Attachment": {
                "id": "1069993260504708081",
                "type": "story_custom_field",
                "entry_id": "1069993260856110917",
                "filename": "70077675 (1).zip",
                "description": "",
                "content_type": "application/zip",
                "created": "2023-07-19 15:07:15",
                "workspace_id": "69993260",
                "owner": "orangecyang"
            }
        }
    ],
    "info": "success"
}
```


# 附件字段说明

## 附件重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">附件ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">工作项ID</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">filename</td><td style="text-align:center;">附件名称</td></tr><tr><td style="text-align:center;">content_type</td><td style="text-align:center;">内容类型</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">上传人</td></tr></tbody></table>
