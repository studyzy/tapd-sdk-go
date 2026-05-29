# 获取发布评审自定义字段

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_launch_forms_custom_fields_settings.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取发布评审自定义字段配置](#获取发布评审自定义字段配置)
    -   [curl 调用示例](#curl-调用示例)
    -   [返回结果](#返回结果)

# 说明

获取发布评审自定义字段配置

# url

`https://api.tapd.cn/launch_forms/custom_fields_settings`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个项目的配置

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取发布评审自定义字段配置

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/launch_forms/custom_fields_settings?workspace_id=20003271'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "CustomFieldConfig": {
                "id": "1120003271001000004",
                "workspace_id": "20003271",
                "entry_type": "launchform",
                "custom_field": "custom_field_one",
                "type": "textarea",
                "name": "DB变更",
                "options": null,
                "enabled": "1",
                "sort": null,
                "memo": null
            }
        }
    ],
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">自定义字段配置的ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属项目ID</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">所属实体对象</td></tr><tr><td style="text-align:center;">custom_field</td><td style="text-align:center;">自定义字段标识（英文名）</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">输入类型</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">自定义字段显示名称</td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;">自定义字段可选值</td></tr><tr><td style="text-align:center;">enabled</td><td style="text-align:center;">是否启用</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">显示时排序系数</td></tr></tbody></table>
