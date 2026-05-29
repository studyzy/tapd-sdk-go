# 过滤条件转换成列表queryToken

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/filter_to_query_token.html

# 说明

将过滤条件转换成页面能用的 QueryToken

# url

`https://api.tapd.cn/stories/filter_to_query_token`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

过滤条件不能超过10个

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">filter[字段名]</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">过滤条件，格式为 filter[字段名] = 过滤值</td></tr><tr><td style="text-align:center;">block_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">分组字段</td></tr><tr><td style="text-align:center;">show_fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">显示字段，以,号分割</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=20358527&filter[custom_field_one]=是&show_fields=title,current_owner,status,custom_field_one&block_type=current_owner' 'https://api.tapd.cn/stories/filter_to_query_token'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "queryToken": "71ab88eeb45d084d8fbc85686a0d2399",
        "href": "http://www.tapd.cn/tapd_fe/10104801/story/list?page=1&queryToken=71ab88eeb45d084d8fbc85686a0d2399&displayEmptyGroup=display&groupType=owner"
    },
    "info": "success"
}
```
