# 获取测试计划关联bug

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_bugs.html

# 说明

获取测试计划整体bug情况

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTestPlanBugs
```


# url

`https://api.tapd.cn/test_plans/result_relation_bugs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次获取一条测试计划的关联bug情况

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr></tbody></table>

# 调用示例及返回结果

## 获取测试计划关联bug

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/test_plans/result_relation_bugs?id=1010158231077224799&workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "id": 1020357849077231365,
            "name": "用例2",
            "tcase_result_relate_bugs": {
                "1020357849000703497": {
                    "id": "1020357849000703497",
                    "executed_at": "2020-03-05 10:15:22",
                    "executor": "jeffjffang",
                    "result_status": "no_pass",
                    "result_remark": null,
                    "bug_id": [
                        "1020357849500646067"
                    ],
                    "Bug": [
                        {
                            "id": "1020357849500646067",
                            "title": "【示例】缺陷1",
                            "severity": "serious",
                            "priority": "high",
                            "status": "接受/处理"
                        }
                    ]
                }
            }
        },
        {
            "id": 1020357849077231363,
            "name": "用例1",
            "tcase_result_relate_bugs": {
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
            }
        }
    ],
    "info": "success"
}
```

