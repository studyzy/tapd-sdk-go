# 获取测试用例执行结果

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_result.html

# 说明

获取测试用例的执行结果，及关联bug的情况

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTcaseResult
```


# url

`https://api.tapd.cn/tcase_instance/result`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次获取一条测试用例执行结果

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">用例ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 分配测试用例

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tcase_instance/result?test_plan_id=1010158231077224799&tcase_id=1020357849077231381&workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "1020357849000703565": {
            "id": "1020357849000703565",
            "executed_at": "2020-03-06 17:46:13",
            "executor": "jeffjffang",
            "result_status": "pass",
            "result_remark": null,
            "bug_id": [],
            "Bug": []
        },
        "1020357849000703563": {
            "id": "1020357849000703563",
            "executed_at": "2020-03-06 17:33:49",
            "executor": "jeffjffang",
            "result_status": "no_pass",
            "result_remark": null,
            "bug_id": [],
            "Bug": []
        },
        "1020357849000703543": {
            "id": "1020357849000703543",
            "executed_at": "2020-03-05 16:05:21",
            "executor": "jeffjffang",
            "result_status": "pass",
            "result_remark": null,
            "bug_id": [],
            "Bug": []
        },
        "1020357849000703535": {
            "id": "1020357849000703535",
            "executed_at": "2020-03-05 14:50:28",
            "executor": "jeffjffang",
            "result_status": "no_pass",
            "result_remark": null,
            "bug_id": [],
            "Bug": []
        },
        "1020357849000703491": {
            "id": "1020357849000703491",
            "executed_at": "2020-03-05 10:14:19",
            "executor": "jeffjffang",
            "result_status": "block",
            "result_remark": null,
            "bug_id": [],
            "Bug": []
        },
        "1020357849000703489": {
            "id": "1020357849000703489",
            "executed_at": "2020-03-05 10:12:03",
            "executor": "jeffjffang",
            "result_status": "no_pass",
            "result_remark": null,
            "bug_id": [],
            "Bug": []
        },
        "1020357849000703483": {
            "id": "1020357849000703483",
            "executed_at": "2020-03-04 17:29:29",
            "executor": "jeffjffang",
            "result_status": "pass",
            "result_remark": null,
            "bug_id": [
                "1020357849500655643",
                "1020357849500655855"
            ],
            "Bug": [
                {
                    "id": "1020357849500655643",
                    "title": "1231",
                    "severity": "",
                    "priority": "",
                    "status": "新"
                },
                {
                    "id": "1020357849500655855",
                    "title": "用例失败bug关联",
                    "severity": "",
                    "priority": "",
                    "status": "新"
                }
            ]
        }
    },
    "info": "success"
}
```


# 返回结果重要字段说明

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">executed_at</td><td style="text-align:center;">执行时间</td></tr><tr><td style="text-align:center;">executor</td><td style="text-align:center;">执行人</td></tr><tr><td style="text-align:center;">result_status</td><td style="text-align:center;">执行结果</td></tr><tr><td style="text-align:center;">result_remark</td><td style="text-align:center;">执行结果备注</td></tr><tr><td style="text-align:center;">bug_id</td><td style="text-align:center;">关联缺陷ID</td></tr><tr><td style="text-align:center;">Bug</td><td style="text-align:center;">关联缺陷详情</td></tr></tbody></table>
