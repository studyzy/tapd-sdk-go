# 解除测试用例关联并移出测试计划

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/delete_tcase_story_relation.html

# 说明

将测试用例移出测试计划，同时解除测试用例和需求的关联关系

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
deleteTcaseStoryRelation
```


# url

`https://api.tapd.cn/tcase_instance/delete_tcase_story_relation`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

支持批量移出测试用例

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试用例ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 将测试用例移出测试计划，同时解除测试用例和需求的关联关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10158231&story_id=1020357849500705291&tcase_id=1020357849077231363&test_plan_id=1020357849000015397' 'https://api.tapd.cn/tcase_instance/delete_tcase_story_relation'`

### 返回结果

```
{
    "status": 1,
    "data": true,
    "info": "success"
}
```

