# 获取工作项所有字段的中英文

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/get_mini_item_fields_label.html

-   [获取空间下的工作项字段](#获取空间下的工作项字段)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

返回工作项所有字段的中英文

# url

`https://api.tapd.cn/mini_items/get_fields_label`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取空间下的工作项字段

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_items/get_fields_label?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_items/get_fields_label?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "ID": "id",
        "标题": "name",
        "标签": "label",
        "优先级": "priority",
        "状态": "status",
        "处理人": "owner",
        "最后修改人": "modifier",
        "创建人": "creator",
        "预计开始": "begin",
        "预计结束": "due",
        "创建时间": "created",
        "最后修改时间": "modified",
        "完成时间": "completed",
        "需求分类": "category_id",
        "父需求": "parent_id",
        "子需求": "children_id",
        "所属项目": "workspace_id",
        "图片与文件": "custom_field_one",
        "开发人员": "custom_field_four",
        "测试人员": "custom_field_six",
        "测试时间": "custom_field_9",
        "测试人": "custom_field_10",
        "详细描述": "description"
    },
    "info": "success"
}
```

