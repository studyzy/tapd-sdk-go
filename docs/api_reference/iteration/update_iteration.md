# 更新迭代

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/update_iteration.html

# 说明

更新迭代，返回更新迭代的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
updateIteration
```


# url

`https://api.tapd.cn/iterations`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">变更人</td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/iteration/get_iteration_custom_fields_settings.html" class="">获取迭代自定义字段配置</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

## 在项目下创建迭代

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&current_user=v_xuanfang&description=test111' 'https://api.tapd.cn/iterations'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Iteration": {
            "id": "1010104801001171471",
            "name": "test",
            "workspace_id": "10104801",
            "startdate": "2020-06-05",
            "enddate": "2020-06-15",
            "status": "open",
            "release_id": null,
            "description": "test111",
            "creator": "v_xuanfang",
            "created": "2020-06-05 16:11:53",
            "modified": "2020-06-05 16:51:48",
            "completed": null,
            "cus_自定义字段的名称": "custom_field_value",
            "custom_field_1": null,
            "custom_field_2": null,
            "custom_field_3": null,
            "custom_field_4": null,
            "custom_field_5": null,
            "custom_field_6": null,
            "custom_field_7": null,
            "custom_field_8": null,
            "custom_field_9": null,
            "custom_field_10": null,
            "custom_field_11": null,
            "custom_field_12": null,
            "custom_field_13": null,
            "custom_field_14": null,
            "custom_field_15": null,
            "custom_field_16": null,
            "custom_field_17": null,
            "custom_field_18": null,
            "custom_field_19": null,
            "custom_field_20": null,
            "custom_field_21": null,
            "custom_field_22": null,
            "custom_field_23": null,
            "custom_field_24": null,
            "custom_field_25": null,
            "custom_field_26": null,
            "custom_field_27": null,
            "custom_field_28": null,
            "custom_field_29": null,
            "custom_field_30": null,
            "custom_field_31": null,
            "custom_field_32": null,
            "custom_field_33": null,
            "custom_field_34": null,
            "custom_field_35": null,
            "custom_field_36": null,
            "custom_field_37": null,
            "custom_field_38": null,
            "custom_field_39": null,
            "custom_field_40": null,
            "custom_field_41": null,
            "custom_field_42": null,
            "custom_field_43": null,
            "custom_field_44": null,
            "custom_field_45": null,
            "custom_field_46": null,
            "custom_field_47": null,
            "custom_field_48": null,
            "custom_field_49": null,
            "custom_field_50": null
        }
    },
    "info": "success"
}
```


# 迭代字段说明

## 迭代重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">完成时间</td></tr></tbody></table>

## 迭代状态(status)字段说明

<table><thead><tr><th style="text-align:center;">字段值</th><th style="text-align:center;">状态名</th></tr></thead><tbody><tr><td style="text-align:center;">open</td><td style="text-align:center;">开启</td></tr><tr><td style="text-align:center;">done</td><td style="text-align:center;">已关闭</td></tr></tbody></table>
