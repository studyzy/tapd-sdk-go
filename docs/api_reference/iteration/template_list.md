# 获取迭代模板列表

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/template_list.html

# 说明

获取迭代模板列表

TAPD迭代支持多类别/多模板, 每一个类别和模板是一一对应关系, 且对应同一套创建页预填写字段配置。此接口可以获取指定项目下所有的模板。

获取迭代模板后, 可以通过 获取迭代模板字段配置 接口取得该模板的创建页字段配置。

# url

`https://api.tapd.cn/iterations/template_list`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下迭代模板列表

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/iterations/template_list?workspace_id=20375553'`

### 返回结果

```json
{
    "status": 1,
    "data": [
        {
            "WorkitemTemplate": {
                "id": "1020375553000077579",
                "workspace_id": "20375553",
                "type": "iteration",
                "name": "迭代模板1",
                "creator": "SYSTEM",
                "created": "2021-07-06 10:55:29",
                "modified": "2021-07-06 10:55:29"
            }
        },
        {
            "WorkitemTemplate": {
                "id": "1020375553000091187",
                "workspace_id": "20375553",
                "type": "release",
                "name": "发布计划模板1",
                "creator": "TAPD system",
                "created": "2023-03-23 15:23:35",
                "modified": "2023-12-13 15:46:41"
            }
        },
        {
            "WorkitemTemplate": {
                "id": "1020375553001023175",
                "workspace_id": "20375553",
                "type": "release",
                "name": "发布计划模板2",
                "creator": "xinweihe",
                "created": "2023-11-15 22:16:27",
                "modified": "2023-12-13 09:59:49"
            }
        }
    ],
    "info": "success"
}
```
