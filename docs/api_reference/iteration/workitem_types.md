# 获取迭代类别列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/workitem_types.html

# 说明

获取迭代类别列表

TAPD迭代支持多类别/多模板, 每一个类别和模板是一一对应关系, 且对应同一套创建页预填写字段配置。此接口可以获取指定项目下所有的迭代类别。

获取迭代类别后, 可以:

1. 通过 **获取迭代类别默认模板字段配置** 接口取得该类别迭代的创建页字段配置;
2. 在 **创建迭代** 时指定并传递 `entity_type` 和 `workitem_type_id` 来预填迭代字段和自动推导计算迭代关键日期。

# url

`https://api.tapd.cn/iterations/workitem_types`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下迭代类别列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/iterations/workitem_types?workspace_id=20375553'`

### 返回结果

```json
{
    "status": 1,
    "data": [
        {
            "WorkitemType": {
                "id": "1020375553000072217",
                "workspace_id": "20375553",
                "entity_type": "release",
                "name": "发布计划类别1",
                "creator": "TAPD system",
                "created": "2023-03-23 15:23:35",
                "modified": "2023-12-13 15:46:41"
            }
        },
        {
            "WorkitemType": {
                "id": "1020375553000070695",
                "workspace_id": "20375553",
                "entity_type": "iteration",
                "name": "迭代类别1",
                "creator": "TAPD system",
                "created": "2022-12-13 15:06:20",
                "modified": "2022-12-13 15:06:20"
            }
        },
        {
            "WorkitemType": {
                "id": "1020375553860932091",
                "workspace_id": "20375553",
                "entity_type": "release",
                "name": "发布计划类别2",
                "creator": "xinweihe",
                "created": "2023-11-15 22:16:27",
                "modified": "2023-12-13 09:59:49"
            }
        }
    ],
    "info": "success"
}
```
