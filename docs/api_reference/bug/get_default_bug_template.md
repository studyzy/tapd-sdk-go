# 获取缺陷模板字段

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_default_bug_template.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [返回结果](#返回结果)
-   [模板重要字段说明](#模板重要字段说明)

# 说明

返回符合查询条件的所有缺陷模板字段

# url

`https://api.tapd.cn/bugs/get_default_bug_template`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有缺陷模板字段

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">模板ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">use_priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否替换优先级字段为 priority_label。取值0和1，默认值是 0</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/get_default_bug_template?template_id=1010104801000068639&workspace_id=10104801'`

## 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778831",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "title",
                "value": "",
                "required": "1",
                "sort": "0"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778833",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "description",
                "value": "",
                "required": "1",
                "sort": "0"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778835",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "current_owner",
                "value": "",
                "required": "1",
                "sort": "1"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778837",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "module",
                "value": "test2",
                "required": "0",
                "sort": "2"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778839",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "priority",
                "value": "",
                "required": "1",
                "sort": "3"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778841",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "severity",
                "value": "",
                "required": "0",
                "sort": "4"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778847",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "version_report",
                "value": "",
                "required": "0",
                "sort": "5"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778843",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "begin",
                "value": "",
                "required": "0",
                "sort": "6"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1010104801000778845",
                "workspace_id": "10104801",
                "type": "bug",
                "template_id": "1010104801000068639",
                "field": "due",
                "value": "",
                "required": "0",
                "sort": "7"
            }
        }
    ],
    "info": "success"
}
```


# 模板字段说明

## 模板重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">模板字段ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">字段英文名</td></tr><tr><td style="text-align:center;">value</td><td style="text-align:center;">默认值</td></tr><tr><td style="text-align:center;">required</td><td style="text-align:center;">是否必填</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">计数</td></tr></tbody></table>
