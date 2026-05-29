# 获取项目自定义字段

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/workspace_custom_field_settings.html

# 说明

返回符合查询条件的所有项目自定义字段（无分页）

# url

`https://api.tapd.cn/workspaces/workspace_custom_field_settings`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次返回所有符合条件的值,只能传项目参数，一次只能查一个项目

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">公司ID</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目自定义字段

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workspaces/workspace_custom_field_settings?workspace_id=20001871'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Workspace": {
            "id": "10104801",
            "name": "TAPD 乌云",
            "pretty_name": "tapd_security",
            "category": "product",
            "status": "normal",
            "description": "",
            "begin_date": null,
            "end_date": null,
            "external_on": "0",
            "creator": "",
            "created": "2015-03-27 16:02:02"
        }
    },
    "info": "success"
}
```


# 项目相关字段说明

## 项目相关重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">自定义字段ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">公司ID</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">归属类型</td></tr><tr><td style="text-align:center;">custom_field</td><td style="text-align:center;">自定义字段标识</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">自定义字段类型</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">自定义字段名称</td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;">候选值</td></tr><tr><td style="text-align:center;">extra_config</td><td style="text-align:center;">额外配置</td></tr><tr><td style="text-align:center;">enabled</td><td style="text-align:center;">是否启用</td></tr></tbody></table>
