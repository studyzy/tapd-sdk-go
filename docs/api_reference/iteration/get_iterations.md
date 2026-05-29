# 获取迭代

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations.html

# 说明

返回符合查询条件的所有迭代（分页显示，默认一页30条）

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getIterations
```


# url

`https://api.tapd.cn/iterations`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">开始时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">结束时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代类别</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">plan_app_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">计划应用ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态（系统状态open/done，自定义状态可传中文）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">完成时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">locker</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">锁定人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下迭代

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/iterations?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Iteration": {
                "id": "1010158231000388075",
                "name": "迭代2",
                "workspace_id": "10158231",
                "startdate": "2017-06-26",
                "enddate": "2017-07-07",
                "status": "open",
                "release_id": null,
                "description": "熟悉敏捷迭代开发",
                "creator": "anyechen",
                "created": "2017-06-20 16:49:05",
                "modified": "2017-06-20 16:49:05",
                "completed": null,
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
        {
            "Iteration": {
                "id": "1010158231000388073",
                "name": "迭代1",
                "workspace_id": "10158231",
                "startdate": "2017-06-12",
                "enddate": "2017-06-23",
                "status": "open",
                "release_id": null,
                "description": "熟悉敏捷迭代开发",
                "creator": "anyechen",
                "created": "2017-06-20 16:49:03",
                "modified": "2017-06-20 16:49:03",
                "completed": null,
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
        }
    ],
    "info": "success"
}
```


# 迭代字段说明

## 迭代重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">完成时间</td></tr><tr><td style="text-align:center;">lock_info</td><td style="text-align:center;">锁定内容</td></tr><tr><td style="text-align:center;">locker</td><td style="text-align:center;">锁定人</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">迭代类别</td></tr><tr><td style="text-align:center;">plan_app_id</td><td style="text-align:center;">计划应用ID</td></tr></tbody></table>

## 迭代状态(status)字段说明

### 注：自定义状态可以直接传中文字符

<table><thead><tr><th style="text-align:center;">字段值</th><th style="text-align:center;">状态名</th></tr></thead><tbody><tr><td style="text-align:center;">open</td><td style="text-align:center;">开启</td></tr><tr><td style="text-align:center;">done</td><td style="text-align:center;">已关闭</td></tr></tbody></table>
