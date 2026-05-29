# 修改迭代仪表盘自定义卡片内容

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/update_custom_dash_board_content.html

# 说明

修改迭代/发布计划/计划应用 仪表盘自定义卡片内容

# url

`https://api.tapd.cn/iterations/update_custom_dash_board_content`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次只能修改一张卡片，只能修改已存在卡片的内容，不支持添加卡片

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源项目ID</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">card_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">卡片ID，在迭代仪表盘卡片上获取</td></tr><tr><td style="text-align:center;">content</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">text</td><td style="text-align:center;">卡片内容，支持富文本</td></tr><tr><td style="text-align:center;">plan_app_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">计划应用ID, 可从页面url参数获取, 默认为0, 代表迭代应用</td></tr></tbody></table>

# 调用示例及返回结果

## 修改迭代仪表盘自定义卡片内容

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&iteration_id=xxxxxxx&card_id=xxxxxx&content=xxxxxxxx&plan_app_id=xxxxx' 'https://api.tapd.cn/iterations/update_custom_dash_board_content'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "id": "xxxxxxxxxxxxxx"
    },
    "info": "success"
}
```
