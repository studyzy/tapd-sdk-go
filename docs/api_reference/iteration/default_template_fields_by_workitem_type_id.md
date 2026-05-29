# 获取迭代类别默认模板字段配置

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/default_template_fields_by_workitem_type_id.html

# 说明

获取迭代类别默认模板字段配置。

TAPD迭代支持多类别/多模板，每一个类别和模板是一一对应关系，且对应同一套创建页预填写字段配置。此接口可以根据项目ID和迭代类别ID获取对应的默认模板的创建页字段配置。

除了通过迭代类别ID `workitem_type_id` 获取配置，还可通过模板ID `template_id` 获取，详见 获取迭代模板字段配置 接口。

# url

`https://api.tapd.cn/iterations/default_template_fields_by_workitem_type_id`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代类别ID</td></tr></tbody></table>

# 调用示例及返回结果

## 获取迭代类别默认模板字段配置

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/iterations/default_template_fields_by_workitem_type_id?workspace_id=20375552&workitem_type_id=1020375553000070695'`

### 返回结果

```json
{
    "status": 1,
    "data": [
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067379",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "description",
                "value": "",
                "required": "1",
                "sort": "0"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067381",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "name",
                "value": "",
                "required": "1",
                "sort": "0"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067397",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "crucial_moment",
                "value": "[{\"key\":\"StartDate\",\"name\":\"Start Date\",\"period\":\"\"},{\"key\":\"CustomMoment1\",\"name\":\"发布评审\",\"period\":\"2\"},{\"key\":\"CustomMoment2\",\"name\":\"预发布\",\"period\":\"4\"},{\"key\":\"EndDate\",\"name\":\"Start Date\",\"period\":\"6\"}]",
                "required": "0",
                "sort": "0"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067397",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "jump_holiday",
                "value": "1",
                "required": "0",
                "sort": "0"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067387",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "custom_field_1",
                "value": "",
                "required": "0",
                "sort": "1"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067389",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "custom_field_2",
                "value": "",
                "required": "0",
                "sort": "2"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067391",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "custom_field_3",
                "value": "",
                "required": "0",
                "sort": "3"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067393",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "custom_field_4",
                "value": "",
                "required": "0",
                "sort": "4"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067395",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "custom_field_8",
                "value": "",
                "required": "0",
                "sort": "5"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067383",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "startdate",
                "value": "",
                "required": "1",
                "sort": "6"
            }
        },
        {
            "WorkitemTemplateField": {
                "id": "1020375553001067385",
                "workspace_id": "20375553",
                "type": "iteration",
                "template_id": "1020375553000077579",
                "field": "enddate",
                "value": "",
                "required": "1",
                "sort": "7"
            }
        }
    ],
    "info": "success"
}
```
