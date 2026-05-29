# 获取任务所有字段及候选值

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_fields_info.html

# 说明

获取任务所有字段及候选值

# url

`https://api.tapd.cn/tasks/get_fields_info`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取任务所有字段及候选值

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tasks/get_fields_info?workspace_id=10104801'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "id": {
            "name": "id",
            "html_type": "input",
            "label": "ID",
            "options": [],
            "color_options": [],
            "pure_options": []
        },
        "name": {
            "name": "name",
            "label": "标题",
            "options": [],
            "color_options": [],
            "pure_options": [],
            "html_type": "input"
        },
        "status": {
            "name": "status",
            "html_type": "select",
            "label": "状态",
            "options": {
                "open": "未开始",
                "progressing": "进行中",
                "done": "已完成"
            },
            "color_options": [],
            "pure_options": []
        },
        "description": {
            "name": "description",
            "label": "详细描述",
            "options": [],
            "color_options": [],
            "pure_options": [],
            "html_type": "rich_edit"
        },
        "owner": {
            "name": "owner",
            "html_type": "user_chooser",
            "label": "处理人",
            "options": [],
            "color_options": [],
            "pure_options": []
        },
        "creator": {
            "name": "creator",
            "html_type": "user_chooser",
            "label": "创建人",
            "options": [],
            "color_options": [],
            "pure_options": []
        },
        "created": {
            "name": "created",
            "label": "创建时间",
            "options": [],
            "color_options": [],
            "pure_options": [],
            "html_type": "datetime"
        },
        "effort": {
            "name": "effort",
            "html_type": "float",
            "label": "预估工时",
            "options": [],
            "color_options": [],
            "pure_options": []
        },
        "effort_completed": {
            "name": "effort_completed",
            "label": "完成工时",
            "options": [],
            "color_options": [],
            "pure_options": [],
            "html_type": "float"
        },
        "remain": {
            "name": "remain",
            "html_type": "float",
            "label": "剩余工时",
            "options": [],
            "color_options": [],
            "pure_options": []
        },
        "exceed": {
            "name": "exceed",
            "html_type": "float",
            "label": "超出工时",
            "options": [],
            "color_options": [],
            "pure_options": []
        },
        "progress": {
            "name": "progress",
            "html_type": "input",
            "label": "进度",
            "options": [],
            "color_options": [],
            "pure_options": []
        },
        "modified": {
            "name": "modified",
            "label": "最后修改时间",
            "options": [],
            "color_options": [],
            "pure_options": [],
            "html_type": "datetime"
        },
        "priority": {
            "name": "priority",
            "html_type": "select",
            "label": "优先级",
            "options": {
                "4": "High",
                "3": "Middle",
                "2": "Low",
                "1": "Nice To Have"
            },
            "color_options": [],
            "pure_options": []
        }
    },
    "info": "success"
}
```
