# 更新级联自定义字段侯选值

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/update_cascade_field_options.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [更新自定义配置ID为 1010104801215158581 的侯选值](#更新自定义配置id为-1010104801215158581-的侯选值)
    -   [侯选值结构。注意要 json encode 成字符串](#侯选值结构。注意要-json-encode-成字符串)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

更新级联自定义字段侯选值

# url

`https://api.tapd.cn/custom_field_configs/update_cascade_field_options`

# HTTP请求方式

POST

# 请求数限制

-   每次只允许更新一个字段，并且是全量更新这个字段的所有侯选值
-   单层级不能超过100个
-   支持需求和缺陷的自定义字段

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;"></th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">19位长度的自定义字段配置ID，可以通过接口 <a href="/document/api-doc/API文档/api_reference/bug/get_bug_custom_fields_settings.html" class="">获取缺陷自定义字段配置</a> 或者 <a href="/document/api-doc/API文档/api_reference/story/get_story_custom_fields_settings.html" class="">获取需求自定义字段配置</a> 来获取</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">候选值, json 字符串结构。children 表示子项。示例结构：<code>[{"name":"a1","children":[{"name":"a11"},{"name":"a12","children":[{"name":"a123"}]}]},{"name":"a2"},{"name":"a3 "}]</code></td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 更新自定义配置ID为 1010104801215158581 的侯选值

### 侯选值结构。注意要 json encode 成字符串

```
[
    {
        "name": "a1",
        "children": [
            {
                "name": "a11"
            },
            {
                "name": "a12",
                "children": [
                    {
                        "name": "a123"
                    }
                ]
            }
        ]
    },
    {
        "name": "a2"
    },
    {
        "name": "a3 "
    }
]
```


### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&id=1010104801215158581&option=%5B%7B%22name%22%3A%22a1%22%2C%22children%22%3A%5B%7B%22name%22%3A%22a11%22%7D%2C%7B%22name%22%3A%22a12%22%2C%22children%22%3A%5B%7B%22name%22%3A%22a123%22%7D%5D%7D%5D%7D%2C%7B%22name%22%3A%22a2%22%7D%2C%7B%22name%22%3A%22a3%20%22%7D%5D' 'https://api.tapd.cn/custom_field_configs/update_cascade_field_options'`

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
