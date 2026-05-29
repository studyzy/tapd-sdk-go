# 获取用户个人配置

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_personal_setting.html

# 说明

获取用户个人配置

# url

`https://api.tapd.cn/users/get_personal_setting`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">公司ID</td></tr><tr><td style="text-align:center;">nick</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">用户唯一标识</td></tr></tbody></table>

# 调用示例及返回结果

## 获取用户个人配置

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/users/get_personal_setting?workspace_id=48464494&nick="ocenhu"'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "language": "zh_CN",
        "message_setting": [
            {
                "type": "LETTER_STORY",
                "disable": [
                    "LETTER_STORY_FOLLOW_UPDATE"
                ],
                "enable": [
                    "LETTER_STORY_FOLLOW_UPDATE_STATUS",
                    "LETTER_STORY_FOLLOW_DELETE",
                    "LETTER_STORY_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "WX_STORY",
                "disable": [
                    "WX_STORY_FOLLOW_UPDATE"
                ],
                "enable": [
                    "WX_STORY_FOLLOW_UPDATE_STATUS",
                    "WX_STORY_FOLLOW_DELETE",
                    "WX_STORY_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "MAIL_TASK",
                "disable": [],
                "enable": [
                    "MAIL_TASK_FOLLOW_UPDATE",
                    "MAIL_TASK_FOLLOW_UPDATE_STATUS",
                    "MAIL_TASK_FOLLOW_DELETE",
                    "MAIL_TASK_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "LETTER_TASK",
                "disable": [],
                "enable": [
                    "LETTER_TASK_FOLLOW_UPDATE",
                    "LETTER_TASK_FOLLOW_UPDATE_STATUS",
                    "LETTER_TASK_FOLLOW_DELETE",
                    "LETTER_TASK_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "WX_TASK",
                "disable": [],
                "enable": [
                    "WX_TASK_FOLLOW_UPDATE",
                    "WX_TASK_FOLLOW_UPDATE_STATUS",
                    "WX_TASK_FOLLOW_DELETE",
                    "WX_TASK_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "MAIL_BUG",
                "disable": [],
                "enable": [
                    "MAIL_BUG_FOLLOW_UPDATE",
                    "MAIL_BUG_FOLLOW_UPDATE_STATUS",
                    "MAIL_BUG_FOLLOW_DELETE",
                    "MAIL_BUG_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "LETTER_BUG",
                "disable": [],
                "enable": [
                    "LETTER_BUG_FOLLOW_UPDATE",
                    "LETTER_BUG_FOLLOW_UPDATE_STATUS",
                    "LETTER_BUG_FOLLOW_DELETE",
                    "LETTER_BUG_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "WX_BUG",
                "disable": [],
                "enable": [
                    "WX_BUG_FOLLOW_UPDATE",
                    "WX_BUG_FOLLOW_UPDATE_STATUS",
                    "WX_BUG_FOLLOW_DELETE",
                    "WX_BUG_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "MAIL_ITEM",
                "disable": [],
                "enable": [
                    "MAIL_ITEM_FOLLOW_UPDATE",
                    "MAIL_ITEM_FOLLOW_UPDATE_STATUS",
                    "MAIL_ITEM_FOLLOW_DELETE",
                    "MAIL_ITEM_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "LETTER_ITEM",
                "disable": [],
                "enable": [
                    "LETTER_ITEM_FOLLOW_UPDATE",
                    "LETTER_ITEM_FOLLOW_UPDATE_STATUS",
                    "LETTER_ITEM_FOLLOW_DELETE",
                    "LETTER_ITEM_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "WX_ITEM",
                "disable": [],
                "enable": [
                    "WX_ITEM_FOLLOW_UPDATE",
                    "WX_ITEM_FOLLOW_UPDATE_STATUS",
                    "WX_ITEM_FOLLOW_DELETE",
                    "WX_ITEM_FOLLOW_ADD_COMMENT"
                ]
            },
            {
                "type": "MAIL_STORY",
                "disable": [
                    "MAIL_STORY_FOLLOW_UPDATE",
                    "MAIL_STORY_FOLLOW_UPDATE_STATUS",
                    "MAIL_STORY_FOLLOW_DELETE",
                    "MAIL_STORY_FOLLOW_ADD_COMMENT"
                ],
                "enable": []
            }
        ]
    },
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">language</td><td style="text-align:center;">当前用户设置的系统语种</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类别</td></tr><tr><td style="text-align:center;">enable</td><td style="text-align:center;">该类别下启用的功能项</td></tr><tr><td style="text-align:center;">disable</td><td style="text-align:center;">该类别下禁用的功能项</td></tr></tbody></table>

# 返回值说明

<table><thead><tr><th style="text-align:center;">返回值</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">LETTER_STORY</td><td style="text-align:center;">关注人需要收到站内信，当需求</td></tr><tr><td style="text-align:center;">WX_STORY</td><td style="text-align:center;">关注人需要收到公众号消息，当需求</td></tr><tr><td style="text-align:center;">MAIL_STORY</td><td style="text-align:center;">关注人需要收到邮件，当需求</td></tr><tr><td style="text-align:center;">MAIL_TASK</td><td style="text-align:center;">关注人需要收到邮件，当任务</td></tr><tr><td style="text-align:center;">LETTER_TASK</td><td style="text-align:center;">关注人需要收到站内信，当任务</td></tr><tr><td style="text-align:center;">WX_TASK</td><td style="text-align:center;">关注人需要收到公众号消息，当任务</td></tr><tr><td style="text-align:center;">MAIL_BUG</td><td style="text-align:center;">关注人需要收到邮件，当缺陷</td></tr><tr><td style="text-align:center;">WX_BUG</td><td style="text-align:center;">关注人需要收到公众号消息，当缺陷</td></tr><tr><td style="text-align:center;">LETTER_BUG</td><td style="text-align:center;">关注人需要收到站内信，当缺陷</td></tr><tr><td style="text-align:center;">MAIL_ITEM</td><td style="text-align:center;">关注人需要收到邮件，当工作项</td></tr><tr><td style="text-align:center;">LETTER_ITEM</td><td style="text-align:center;">关注人需要收到站内信，当工作项</td></tr><tr><td style="text-align:center;">WX_ITEM</td><td style="text-align:center;">关注人需要收到公众号消息，当工作项</td></tr></tbody></table>
