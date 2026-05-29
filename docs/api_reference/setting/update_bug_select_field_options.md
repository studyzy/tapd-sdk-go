# 更新缺陷下拉类型自定义字段候选值

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/update_bug_select_field_options.html

# 说明

更新缺陷下拉类型自定义字段候选值

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
updateBugSelectFieldOptions
```


# url

`https://api.tapd.cn/custom_field_configs/update_bug_select_field_options`

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一个字段，并且是全量更新这个字段的所有侯选值。

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;"></th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">19位长度的自定义字段配置ID</td><td style="text-align:center;">可以通过接口<a href="/document/api-doc/API文档/api_reference/bug/get_bug_custom_fields_settings.html" class="">获取缺陷自定义字段配置</a> 来获取</td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">候选值。以英文竖线隔开</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 更新需求自定义配置ID为 1010104801214991279 的侯选值为 开发，测试，产品，运营

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'options=开发|测试|产品|运营&id=1010104801214991279&workspace_id=10104801' 'https://api.tapd.cn/custom_field_configs/update_bug_select_field_options'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "status": 1
    },
    "info": "success"
}
```


# 返回说明

<table><thead><tr><th style="text-align:center;">返回</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">status</td><td style="text-align:center;">取值1为成功，0为失败</td></tr></tbody></table>
