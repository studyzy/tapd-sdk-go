# 获取测试用例所有字段及候选值

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_fields_info.html

# 说明

返回测试用例所有字段及候选值

# url

`https://api.tapd.cn/tcases/get_fields_info`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的测试用例字段

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tcases/get_fields_info?workspace_id=10104801'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "id": {
            "html_type": "input",
            "label": "ID",
            "options": [],
            "name": "id"
        },
        "steps": {
            "html_type": "qmeditor",
            "label": "用例步骤",
            "options": [],
            "name": "steps"
        },
        "workspace_id": {
            "html_type": "input",
            "label": "Workspace",
            "options": [],
            "name": "workspace_id"
        },
        "category_id": {
            "html_type": "dialog",
            "label": "用例目录",
            "options": {
                "-1": "未规划目录"
            },
            "name": "category_id"
        },
        "created": {
            "html_type": "dateinput",
            "label": "创建时间",
            "options": [],
            "name": "created"
        },
        "modifier": {
            "html_type": "user_chooser",
            "label": "最后修改人",
            "options": [],
            "name": "modifier"
        },
        "modified": {
            "html_type": "dateinput",
            "label": "最后修改时间",
            "options": [],
            "name": "modified"
        },
        "creator": {
            "html_type": "user_chooser",
            "label": "创建人",
            "options": [],
            "name": "creator"
        },
        "status": {
            "html_type": "select",
            "label": "用例状态",
            "options": {
                "normal": "正常",
                "updating": "待更新",
                "abandon": "已废弃"
            },
            "name": "status",
            "default": "normal"
        },
        "name": {
            "html_type": "input",
            "label": "用例名称",
            "options": [],
            "name": "name"
        },
        "precondition": {
            "html_type": "qmeditor",
            "label": "前置条件",
            "options": [],
            "name": "precondition"
        },
        "expectation": {
            "html_type": "qmeditor",
            "label": "预期结果",
            "options": [],
            "name": "expectation"
        },
        "type": {
            "html_type": "select",
            "label": "用例类型",
            "options": {
                "功能测试": "功能测试",
                "性能测试": "性能测试",
                "安全性测试": "安全性测试",
                "其他": "其他"
            },
            "name": "type"
        },
        "priority": {
            "html_type": "select",
            "label": "用例等级",
            "options": {
                "高": "高",
                "中": "中",
                "低": "低"
            },
            "name": "priority"
        },
        "is_automated": {
            "html_type": "select",
            "label": "是否实现自动化",
            "options": {
                "1": "是",
                "2": "否"
            },
            "name": "is_automated"
        },
        "automation_type": {
            "html_type": "multi_select",
            "label": "自动化测试类型",
            "options": {
                "UI自动化": "UI自动化",
                "接口自动化": "接口自动化",
                "未分析": "未分析",
                "无法自动化": "无法自动化",
                "待自动化": "待自动化"
            },
            "name": "automation_type"
        },
        "automation_platform": {
            "html_type": "multi_select",
            "label": "自动化测试平台",
            "options": {},
            "name": "automation_platform"
        }
    },
    "info": "success"
}
```
