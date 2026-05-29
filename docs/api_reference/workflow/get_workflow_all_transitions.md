# 获取工作流流转细则

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/get_workflow_all_transitions.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目缺陷工作流](#获取项目缺陷工作流)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [appendfield字段说明](#appendfield字段说明)
-   [default\_value字段说明](#default-value字段说明)
-   [并行工作流说明](#并行工作流说明)

# 说明

获取工作流流转细则

# url

`https://api.tapd.cn/workflows/all_transitions`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次只能获取一个项目的工作流流转细则

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">system</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">系统名。取 bug （缺陷的）或者 story（需求的）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求类别</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目缺陷工作流

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workflows/all_transitions?system=bug&workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
            "Name": "start-new",
            "StepPrevious": "start",
            "StepNext": "new",
            "Appendfield": [
                {
                    "DBModel": "Bug",
                    "FieldName": "BugStoryRelation_relative_id",
                    "Notnull": "yes",
                    "Sort": "2"
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "begin",
                    "Notnull": "yes",
                    "Sort": "3"
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "due",
                    "Notnull": "yes",
                    "Sort": "4"
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "current_owner",
                    "Notnull": "yes",
                    "Sort": "1"
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "module",
                    "Notnull": "yes",
                    "Sort": "8"
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "priority",
                    "Notnull": "yes",
                    "Sort": "10"
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "severity",
                    "Notnull": "yes",
                    "Sort": "11"
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "version_report",
                    "Notnull": "yes",
                    "Sort": "12"
                }
            ]
        },
        {
            "Name": "new-new",
            "StepPrevious": "new",
            "StepNext": "new",
            "Inform": [
                {
                    "InformType": "RTX",
                    "InformId": "1716"
                },
                {
                    "InformType": "Email",
                    "InformId": "1717"
                }
            ],
            "Appendfield": [
                {
                    "DBModel": "Bug",
                    "FieldName": "remarks",
                    "Notnull": "yes",
                    "Sort": "1",
                    "DefaultValue": [
                        {
                            "Type": "default_value",
                            "Value": ""
                        }
                    ]
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "current_owner",
                    "Notnull": "yes",
                    "Sort": "2",
                    "DefaultValue": [
                        {
                            "Type": "record_value",
                            "DBModel": "Bug",
                            "Field": "current_owner"
                        }
                    ]
                }
            ]
        },
        {
            "Name": "new-in_progress",
            "StepPrevious": "new",
            "StepNext": "in_progress",
            "Inform": [
                {
                    "InformType": "RTX",
                    "InformId": "1703"
                },
                {
                    "InformType": "Email",
                    "InformId": "1715"
                }
            ],
            "Appendfield": [
                {
                    "DBModel": "Bug",
                    "FieldName": "remarks",
                    "Notnull": "no",
                    "Sort": "1",
                    "DefaultValue": [
                        {
                            "Type": "default_value",
                            "Value": ""
                        }
                    ]
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "current_owner",
                    "Notnull": "yes",
                    "Sort": "2",
                    "DefaultValue": [
                        {
                            "Type": "record_value",
                            "DBModel": "Bug",
                            "Field": "current_owner"
                        }
                    ]
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "de",
                    "Notnull": "yes",
                    "Sort": "3",
                    "DefaultValue": [
                        {
                            "Type": "record_value",
                            "DBModel": "Bug",
                            "Field": "current_owner"
                        }
                    ]
                },
                {
                    "DBModel": "Bug",
                    "FieldName": "auditer",
                    "Notnull": "yes",
                    "Sort": "4",
                    "DefaultValue": [
                        {
                            "Type": "record_value",
                            "DBModel": "Bug",
                            "Field": "reporter"
                        }
                    ]
                }
            ]
        },
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">Name</td><td style="text-align:center;">状态（步骤）流转名称</td></tr><tr><td style="text-align:center;">StepPrevious</td><td style="text-align:center;">当前状态（步骤）</td></tr><tr><td style="text-align:center;">StepNext</td><td style="text-align:center;">目标状态（步骤）</td></tr><tr><td style="text-align:center;">Appendfield</td><td style="text-align:center;">状态（步骤）流转时需要补充的附加字段</td></tr><tr><td style="text-align:center;">AuthorizedUser</td><td style="text-align:center;">状态（步骤）流转权限设置</td></tr></tbody></table>

## appendfield字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">FieldName</td><td style="text-align:center;">字段名称，用于系统后台保存</td></tr><tr><td style="text-align:center;">Notnull</td><td style="text-align:center;">是否必填</td></tr><tr><td style="text-align:center;">Sort</td><td style="text-align:center;">显示排序号</td></tr><tr><td style="text-align:center;">default_value</td><td style="text-align:center;">默认值设置</td></tr></tbody></table>

## default\_value字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">Type</td><td style="text-align:center;">默认值类型</td></tr><tr><td style="text-align:center;">Field</td><td style="text-align:center;">对应系统字段</td></tr></tbody></table>

## 并行工作流说明

并行工作流中，节点操作只需要完成。StepPrevious 与 StepNext 相同的配置为完成操作需要补充的附加字段和流转权限配置
