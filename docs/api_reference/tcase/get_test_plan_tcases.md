# 获取测试计划与测试用例关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_tcases.html

# 说明

返回符合查询条件的所有测试用例&需求

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTestPlanTcases
```


# url

`https://api.tapd.cn/test_plans/get_test_plan_tcase`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

返回全部数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下测试用例

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/test_plans/get_test_plan_tcase?test_plan_id=1000000755077233617&workspace_id=755'`

### 返回结果

```

{
    "status": 1,
    "data": [
        {
            "TestPlanStoryTcaseRelation": {
                "id": "1000000755002248699",
                "workspace_id": "755",
                "test_plan_id": "1000000755077233617",
                "story_id": "0",
                "tcase_id": "1000000755000026804",
                "sort": "0",
                "creator": "v_xuanfang",
                "created": "0000-00-00 00:00:00"
            }
        },
        {
            "TestPlanStoryTcaseRelation": {
                "id": "1000000755002248753",
                "workspace_id": "755",
                "test_plan_id": "1000000755077233617",
                "story_id": "0",
                "tcase_id": "1000000755075912019",
                "sort": "0",
                "creator": "v_xuanfang",
                "created": "0000-00-00 00:00:00"
            }
        },
    ],
"info": "success"
}

```


# 测试用例字段说明

## 测试用例重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">关系id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;">测试计划ID</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">需求ID</td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;">测试用例ID</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">关系创建时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr></tbody></table>
