# 获取迭代仪表盘自定义卡片内容

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_custom_dash_board_content.html

# 说明

获取迭代自定义卡片内容。

# url

`https://api.tapd.cn/iterations/get_custom_dash_board_content`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只支持查询一个迭代的

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代ID</td></tr></tbody></table>

# 调用示例及返回结果

## 获取迭代自定义卡片内容

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/iterations/get_custom_dash_board_content?workspace_id=10104801&iteration_id=1010104801000723579'`

### 返回结果

```json
{
    "status": 1,
    "data": [
        {
            "id": "1010104801000003949",
            "template": "Custom",
            "title": "自定义aaa",
            "component_data": "[]",
            "width": "6",
            "height": "3",
            "card_type": "RichContent",
            "data": {
                "content": "<p>自定义卡片内容。支持 <strong>HTML</strong>。</p>",
                "description_type": "1",
                "value": "<p>自定义卡片内容。支持 <strong>HTML</strong>。</p>"
            }
        }
    ],
    "info": "success"
}
```
