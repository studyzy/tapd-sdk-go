# 解除需求缺陷关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/remove_story_bug_raletions.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

解除需求缺陷关联关系

# url

`https://api.tapd.cn/stories/remove_story_bug_raletions`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许解除一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求id</td></tr><tr><td style="text-align:center;">bug_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">缺陷id</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">操作人</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&story_id=1010104801855713221&bug_id=1010104801500637481' 'https://api.tapd.cn/stories/remove_story_bug_raletions'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "success": true
    },
    "info": "success"
}
```

