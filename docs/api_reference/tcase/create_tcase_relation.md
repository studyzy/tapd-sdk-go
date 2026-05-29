# 创建测试计划和测试用例关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/create_tcase_relation.html

# 说明

创建测试计划和测试用例关联关系

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
createTcaseRelation
```


# url

`https://api.tapd.cn/test_plans/create_tcase_relation`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次最多保存10条需求关联关系

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">teet_plan_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">tcase_ids</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">要关联的测试用例ID,最多不超过10条，多个需求ID之间用,分割</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 创建测试计划和测试用例关联关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'test_plan_id=1010158231077224799&workspace_id=10158231&tcase_ids=1020357849077231603,1020357849077231393,1020357849077231605&creator=peter' 'https://api.tapd.cn/test_plans/create_tcase_relation'`

### 返回结果

```
{
    "status": 1,
    "data": [],
    "info": "create tcase relation success"
}
```

