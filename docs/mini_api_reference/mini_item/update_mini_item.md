# 更新工作项

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/update_mini_item.html

-   [更新工作项 1069993260856097281 的处理人为 orangecyang，开始时间为 2023-07-18](#更新工作项-1069993260856097281-的处理人为-orangecyang-开始时间为-2023-07-18)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [curl 使用 OAuth Access Token 鉴权调用示例](#curl-使用-oauth-access-token-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [工作项重要字段说明](#工作项重要字段说明)

# 说明

更新工作项，返回工作项更新后的数据

# url

`https://api.tapd.cn/mini_items`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">progress_manual</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">进度</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">分组</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签，标签不存在时将自动创建，多个以英文坚线分格</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/mini_api_reference/mini_item/get_mini_item_custom_fields_settings.html" class="">获取工作项自定义字段配置</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

## 更新工作项 1069993260856097281 的处理人为 orangecyang，开始时间为 2023-07-18

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'id=1069993260856097281&owner=orangecyang;&begin=2023-07-18&workspace_id=69993260' 'https://api.tapd.cn/mini_items'`

### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' -d 'id=1069993260856097281&owner=orangecyang;&begin=2023-07-18&workspace_id=69993260' 'https://api.tapd.cn/mini_items'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "MiniItem": {
            "id": "1069993260856097281",
            "name": "test",
            "workspace_id": "69993260",
            "category_id": "-1",
            "status": "open",
            "owner": "orangecyang;",
            "begin": "2023-07-18",
            "due": null,
            "priority": "",
            "label": "",
            "description_type": "1",
            "description": null,
            "markdown_description": null,
            "ancestor_id": "1069993260856097281",
            "parent_id": "0",
            "children_id": "|",
            "level": "0",
            "creator": "orangecyang",
            "created": "2023-07-06 16:16:21",
            "modifier": "",
            "modified": "2023-07-18 11:04:30",
            "completed": null,
            "has_attachment": "0",
            "sort": "85546902000000",
            "is_archived": "0",
            "progress_manual": "0",
            "participator": "",
            "custom_field_one": "",
            "custom_field_two": "",
            "custom_field_three": "",
            "custom_field_four": "",
            "custom_field_five": "",
            "custom_field_six": "",
            "custom_field_seven": "",
            "custom_field_eight": "",
            "custom_field_9": "",
            "custom_field_10": "",
            "custom_field_11": "",
            "custom_field_12": "",
            "custom_field_13": "",
            "custom_field_14": "",
            "custom_field_15": "",
            "custom_field_16": "",
            "custom_field_17": "",
            "custom_field_18": "",
            "custom_field_19": "",
            "custom_field_20": "",
            "custom_field_21": "",
            "custom_field_22": "",
            "custom_field_23": "",
            "custom_field_24": "",
            "custom_field_25": "",
            "custom_field_26": "",
            "custom_field_27": "",
            "custom_field_28": "",
            "custom_field_29": "",
            "custom_field_30": "",
            "custom_field_31": "",
            "custom_field_32": "",
            "custom_field_33": "",
            "custom_field_34": "",
            "custom_field_35": "",
            "custom_field_36": "",
            "custom_field_37": "",
            "custom_field_38": "",
            "custom_field_39": "",
            "custom_field_40": "",
            "custom_field_41": "",
            "custom_field_42": "",
            "custom_field_43": "",
            "custom_field_44": "",
            "custom_field_45": "",
            "custom_field_46": "",
            "custom_field_47": "",
            "custom_field_48": "",
            "custom_field_49": "",
            "custom_field_50": "",
            "custom_field_51": "",
            "custom_field_52": "",
            "custom_field_53": "",
            "custom_field_54": "",
            "custom_field_55": "",
            "custom_field_56": "",
            "custom_field_57": "",
            "custom_field_58": "",
            "custom_field_59": "",
            "custom_field_60": "",
            "custom_field_61": "",
            "custom_field_62": "",
            "custom_field_63": "",
            "custom_field_64": "",
            "custom_field_65": "",
            "custom_field_66": "",
            "custom_field_67": "",
            "custom_field_68": "",
            "custom_field_69": "",
            "custom_field_70": "",
            "custom_field_71": "",
            "custom_field_72": "",
            "custom_field_73": "",
            "custom_field_74": "",
            "custom_field_75": "",
            "custom_field_76": "",
            "custom_field_77": "",
            "custom_field_78": "",
            "custom_field_79": "",
            "custom_field_80": "",
            "custom_field_81": "",
            "custom_field_82": "",
            "custom_field_83": "",
            "custom_field_84": "",
            "custom_field_85": "",
            "custom_field_86": "",
            "custom_field_87": "",
            "custom_field_88": "",
            "custom_field_89": "",
            "custom_field_90": "",
            "custom_field_91": "",
            "custom_field_92": "",
            "custom_field_93": "",
            "custom_field_94": "",
            "custom_field_95": "",
            "custom_field_96": "",
            "custom_field_97": "",
            "custom_field_98": "",
            "custom_field_99": "",
            "custom_field_100": ""
        }
    },
    "info": "success"
}
```


# 工作项字段说明

## 工作项重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">优先级</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">标签</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">完成时间</td></tr><tr><td style="text-align:center;">description_type</td><td style="text-align:center;">详细描述类型</td></tr><tr><td style="text-align:center;">level</td><td style="text-align:center;">距离根节点的深度</td></tr><tr><td style="text-align:center;">ancestor_id</td><td style="text-align:center;">层级</td></tr><tr><td style="text-align:center;">progress_manual</td><td style="text-align:center;">进度</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">分组</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">父工作项</td></tr><tr><td style="text-align:center;">children_id</td><td style="text-align:center;">子工作项</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">is_archived</td><td style="text-align:center;">是否归档</td></tr></tbody></table>
