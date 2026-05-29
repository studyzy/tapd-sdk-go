# 获取工作项自定义字段配置

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/get_mini_item_custom_fields_settings.html

-   [获取工作项自定义字段配置](#获取工作项自定义字段配置)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取工作项自定义字段配置

# url

`https://api.tapd.cn/mini_items/custom_fields_settings`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个空间的配置

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取工作项自定义字段配置

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_items/custom_fields_settings?workspace_id=69993260'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_items/custom_fields_settings?workspace_id=69993260'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "CustomFieldConfig": {
                "id": "1069993260215045623",
                "workspace_id": "69993260",
                "entry_type": "mini_item",
                "custom_field": "custom_field_one",
                "type": "file",
                "name": "图片与文件",
                "options": "{}",
                "extra_config": null,
                "enabled": "1",
                "freeze": "0",
                "sort": "0",
                "memo": ""
            }
        },
        {
            "CustomFieldConfig": {
                "id": "1069993260215046499",
                "workspace_id": "69993260",
                "entry_type": "mini_item",
                "custom_field": "custom_field_four",
                "type": "user_chooser",
                "name": "开发人员",
                "options": "{}",
                "extra_config": null,
                "enabled": "1",
                "freeze": "0",
                "sort": "0",
                "memo": ""
            }
        },
        {
            "CustomFieldConfig": {
                "id": "1069993260215046895",
                "workspace_id": "69993260",
                "entry_type": "mini_item",
                "custom_field": "custom_field_six",
                "type": "user_chooser",
                "name": "测试人员",
                "options": "{}",
                "extra_config": null,
                "enabled": "1",
                "freeze": "0",
                "sort": "0",
                "memo": ""
            }
        },
        {
            "CustomFieldConfig": {
                "id": "1069993260215048539",
                "workspace_id": "69993260",
                "entry_type": "mini_item",
                "custom_field": "custom_field_9",
                "type": "datetime",
                "name": "测试时间",
                "options": "{}",
                "extra_config": null,
                "enabled": "1",
                "freeze": "0",
                "sort": "0",
                "memo": ""
            }
        },
        {
            "CustomFieldConfig": {
                "id": "1069993260215049365",
                "workspace_id": "69993260",
                "entry_type": "mini_item",
                "custom_field": "custom_field_10",
                "type": "user_chooser",
                "name": "测试人",
                "options": "{}",
                "extra_config": null,
                "enabled": "1",
                "freeze": "0",
                "sort": "0",
                "memo": ""
            }
        }
    ],
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">自定义字段配置的ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属空间ID</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">工作项ID</td></tr><tr><td style="text-align:center;">custom_field</td><td style="text-align:center;">自定义字段标识（英文名）</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">输入类型</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">自定义字段显示名称</td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;">自定义字段可选值</td></tr><tr><td style="text-align:center;">enabled</td><td style="text-align:center;">是否启用</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">显示时排序系数</td></tr><tr><td style="text-align:center;">freeze</td><td style="text-align:center;">是否冻结</td></tr><tr><td style="text-align:center;">extra_config</td><td style="text-align:center;">额外配置</td></tr></tbody></table>
