# 解除需求关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/remove_story_link_relation.html

# 说明

解除需求关联关系

# url

`https://api.tapd.cn/stories/remove_story_link_relation`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次只能删除一条关联关系

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">src_story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源需求ID</td></tr><tr><td style="text-align:center;">target_story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">目标需求ID</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&src_story_id=1010104801115147834&target_story_id=1010104801115129060' 'https://api.tapd.cn/stories/remove_story_link_relation'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "success": 1
    },
    "info": "success"
}
```
