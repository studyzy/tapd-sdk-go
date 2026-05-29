# 获取需求前后置关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_time_relative_stories.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [需求重要字段说明](#需求重要字段说明)

# 说明

获取需求前后置关系

# url

`https://api.tapd.cn/stories/get_time_relative_stories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP 请求方式

GET

# 请求数限制

默认返回所有数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源项目ID</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源需求ID</td></tr></tbody></table>

# 调用示例及返回结果

## curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password''https://api.tapd.cn/stories/get_time_relative_stories?workspace_id=10104801&story_id=1010104801854917775'`

## curl 使用 OAuth Access Token 鉴权调用示例

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemTimeRelation": {
                "id": "1210104801000007813",
                "workspace_id": "10104801",
                "workitem_type": "story",
                "workitem_id": "1010104801854915911",
                "src_field": "begin",
                "dst_workspace_id": "10104801",
                "dst_workitem_type": "story",
                "dst_workitem_id": "1010104801854917775",
                "dst_field": "due",
                "relation_type": "after"
            }
        },
        {
            "WorkitemTimeRelation": {
                "id": "1210104801000007815",
                "workspace_id": "10104801",
                "workitem_type": "story",
                "workitem_id": "1010104801854917775",
                "src_field": "due",
                "dst_workspace_id": "10104801",
                "dst_workitem_type": "story",
                "dst_workitem_id": "1010104801854917809",
                "dst_field": "begin",
                "relation_type": "after"
            }
        }
    ],
    "info": "success"
}
```


# 需求字段说明

## 需求重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">源项目ID</td></tr><tr><td style="text-align:center;">workitem_type</td><td style="text-align:center;">业务对象类型，固定为 story</td></tr><tr><td style="text-align:center;">workitem_id</td><td style="text-align:center;">源需求ID</td></tr><tr><td style="text-align:center;">src_field</td><td style="text-align:center;">源需求被依赖的字段</td></tr><tr><td style="text-align:center;">dst_workspace_id</td><td style="text-align:center;">被依赖的项目ID</td></tr><tr><td style="text-align:center;">dst_workitem_type</td><td style="text-align:center;">被依赖的业务对象类型，固定为 story</td></tr><tr><td style="text-align:center;">dst_workitem_id</td><td style="text-align:center;">被依赖的需求ID</td></tr><tr><td style="text-align:center;">dst_field</td><td style="text-align:center;">被依赖的字段</td></tr><tr><td style="text-align:center;">relation_type</td><td style="text-align:center;">依赖类型。before 为前置依赖，after 为后置依赖</td></tr></tbody></table>
