# 分配测试用例

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/assign_tcase_instance.html

# 说明

-   支持修改测试用例执行人、负责人
-   支持通过tcase\_id批量修改
-   支持通过用例目录批量修改

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
assignTcaseInstance
```


# url

`https://api.tapd.cn/tcase_instance/assign`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

最大支持10条

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">用例ID</td><td style="text-align:center;">支持批量执行，用,分割，如1020357849077231381,1020357849077231382</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">用例目录ID</td><td style="text-align:center;">tcase_id和category_id不能同时为空</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">executor</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">执行人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">assignee</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">负责人</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 分配测试用例

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'test_plan_id=1010158231077224799&tcase_id=1020357849077231381&executor=peter&workspace_id=10158231' 'https://api.tapd.cn/tcase_instance/assign'`

### 返回结果

```
{
    "status": 1,
    "data": [],
    "info": "success"
}
```

