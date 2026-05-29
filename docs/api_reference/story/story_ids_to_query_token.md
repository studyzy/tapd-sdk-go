# 转换需求ID成列表queryToken

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/story_ids_to_query_token.html

# 说明

把一批需求ID转换成页面能用的 QueryToken

# url

`https://api.tapd.cn/stories/ids_to_query_token`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

为了保证页面显示效果，建议ID数量不超过500

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">ids</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">需求ID，使用英文逗号,做分隔</td></tr></tbody></table>

# 调用示例及返回结果

## 把需求ID转换成queryToken及返回列表链接

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&ids=1010104801102653321,1010104801085527301' 'https://api.tapd.cn/stories/ids_to_query_token'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "queryToken": "71ab88eeb45d084d8fbc85686a0d2399",
        "href": "https://www.tapd.cn/tapd_fe/10104801/story/list?page=1&queryToken=71ab88eeb45d084d8fbc85686a0d2399"
    },
    "info": "success"
}
```
