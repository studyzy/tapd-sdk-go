# 获取需求关联的缺陷

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_related_bugs.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [需求关联缺陷重要字段说明](#需求关联缺陷重要字段说明)

# 说明

返回符合查询条件的所有需求关联的缺陷ID

# url

`https://api.tapd.cn/stories/get_related_bugs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有关系

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求ID</td><td style="text-align:center;">支持多ID查询</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_related_bugs?workspace_id=10104801&story_id=1010104801866181263'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "workspace_id": 10104801,
            "story_id": "1010104801866181263",
            "bug_id": "1010104801083691309"
        },
        {
            "workspace_id": 10104801,
            "story_id": "1010104801866181263",
            "bug_id": "1010104801085894269"
        },
        {
            "workspace_id": 10104801,
            "story_id": "1010104801866181263",
            "bug_id": "1010104801083691321"
        },
        {
            "workspace_id": 10104801,
            "story_id": "1010104801866181263",
            "bug_id": "1010104801083691305"
        }
    ],
    "info": "success"
}
```


# 需求关联缺陷字段说明

## 需求关联缺陷重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">需求ID</td></tr><tr><td style="text-align:center;">bug_id</td><td style="text-align:center;">缺陷ID</td></tr></tbody></table>
