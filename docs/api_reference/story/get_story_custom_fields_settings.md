# 获取需求自定义字段配置

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_custom_fields_settings.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [SDK 方法名](#sdk-方法名)
-   [获取需求自定义字段配置](#获取需求自定义字段配置)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取需求自定义字段配置

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getStoryCustomFieldsSettings
```


# url

`https://api.tapd.cn/stories/custom_fields_settings`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个项目的配置

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取需求自定义字段配置

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/custom_fields_settings?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "CustomFieldConfig": {
                "id": "1010104801216209053",
                "workspace_id": "10104801",
                "app_id": "1",
                "entry_type": "story",
                "custom_field": "custom_field_17",
                "type": "cascade_radio",
                "name": "联动字段测试",
                "options": "[{\"name\":\"a1\",\"children\":[{\"name\":\"a11\"},{\"name\":\"a12\",\"children\":[{\"name\":\"a123\"}]}]},{\"name\":\"a2\"},{\"name\":\"a3\"}]",
                "extra_config": null,
                "enabled": "1",
                "creator": "",
                "created": "0000-00-00 00:00:00",
                "modified": "0000-00-00 00:00:00",
                "freeze": "0",
                "sort": "0",
                "memo": null,
                "open_extension_id": "",
                "is_out": 0,
                "is_uninstall": 0,
                "app_name": ""
            }
        }
    ],
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">自定义字段配置的ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属项目ID</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">所属实体对象</td></tr><tr><td style="text-align:center;">custom_field</td><td style="text-align:center;">自定义字段标识（英文名）</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">输入类型</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">自定义字段显示名称</td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;">自定义字段可选值</td></tr><tr><td style="text-align:center;">enabled</td><td style="text-align:center;">是否启用</td></tr><tr><td style="text-align:center;">app_id</td><td style="text-align:center;">检查项</td></tr><tr><td style="text-align:center;">extra_config</td><td style="text-align:center;">额外配置(颜色信息等)</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">显示时排序系数</td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">备注</td></tr><tr><td style="text-align:center;">open_extension_id</td><td style="text-align:center;">插件扩展字段标识</td></tr><tr><td style="text-align:center;">is_out</td><td style="text-align:center;">已弃用</td></tr><tr><td style="text-align:center;">is_uninstall</td><td style="text-align:center;">应用是否安装到当前项目</td></tr><tr><td style="text-align:center;">app_name</td><td style="text-align:center;">应用名</td></tr></tbody></table>
