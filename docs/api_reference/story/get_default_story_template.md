# 获取需求模板字段

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_default_story_template.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [返回结果](#返回结果)
-   [模板重要字段说明](#模板重要字段说明)

# 说明

返回符合查询条件的所有需求模板字段

# url

`https://api.tapd.cn/stories/get_default_story_template`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有需求模板字段

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">模板ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">use_priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否替换优先级字段为 priority_label。取值0和1，默认值是 0</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_default_story_template?template_id=1010104801000068641&workspace_id=10104801'`

## 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemTemplateField": {
                "id": "1010104801015287651",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "description",
                "value": "",
                "required": "1",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801015287653",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "name",
                "value": "",
                "required": "1",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243901",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "sub_stories_auto_succeed_name",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243903",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "auto_succeed_story_fields",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243905",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "category_id",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243907",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "priority",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243909",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "owner",
                "value": "",
                "required": "1",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243911",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "iteration_id",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243913",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "cc",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243915",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "developer",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243917",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "custom_field_14",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243919",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "begin",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": "",
                "default_value": ""
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801016243921",
                "workspace_id": "10104801",
                "type": "story",
                "template_id": "1010104801000850579",
                "field": "due",
                "value": "",
                "required": "0",
                "sort": "0",
                "linkage_rules": "",
                "default_value": ""
            }
        }
    ],
    "info": "success"
}
```


# 模板字段说明

## 模板重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">模板字段ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">字段名称</td></tr><tr><td style="text-align:center;">value</td><td style="text-align:center;">默认值</td></tr><tr><td style="text-align:center;">required</td><td style="text-align:center;">是否必填</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">计数</td></tr><tr><td style="text-align:center;">linkage_rules</td><td style="text-align:center;">配置显示规则</td></tr><tr><td style="text-align:center;">default_value</td><td style="text-align:center;">默认值</td></tr></tbody></table>
