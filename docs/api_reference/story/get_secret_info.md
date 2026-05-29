# 获取需求保密信息

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_secret_info.html

# 说明

获取指定需求的保密信息

# url

`https://api.tapd.cn/stories/get_secret_info`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

每次只允许查询一条需求

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr></tbody></table>

# 调用示例及返回结果

## 获取需求1010104801871430407的保密信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_secret_info?workspace_id=10104801&story_id=1010104801871430407'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "creator": "xinweihe",
        "allow_list": "xinweihe;1000000000000000002",
        "secret_root_id": "1010104801871430407",
        "add_participant_fields": "true",
        "secret_scrope": "secret"
    },
    "info": "success"
}
```
