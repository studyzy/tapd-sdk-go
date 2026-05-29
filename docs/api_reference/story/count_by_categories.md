# 获取指定分类下需求数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/count_by_categories.html

# 说明

获取指定分类下需求数量

# url

`https://api.tapd.cn/stories/count_by_categories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

返回需求数量，会返回子分类的数量，及做子分类下面需求数量的累加。

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求分类，支持多ID查询，如id1,id2,id3</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/count_by_categories?workspace_id=10158231&category_id=1010104801000079035,1010104801000086901'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "1010104801000079035": 4,
        "1010104801000086901": 2,
        "1010104801003796129": 1,
        "1010104801000086915": 0,
        "1010104801000079037": 2
    },
    "info": "success"
}
```
