# 获取迭代下测试计划

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_by_iteration_id.html

# 说明

返回迭代下的测试计划

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getByIterationId
```


# url

`https://api.tapd.cn/test_plans/get_by_iteration_id`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr></tbody></table>

# 调用示例及返回结果

## 获取迭代下测试计划

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/test_plans/get_by_iteration_id?iteration_id=1151650666001000111&workspace_id=51650666'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "workspace_id": "51650666",
            "iteration_id": "1151650666001000111",
            "test_plan_id": "1151650666001000019"
        },
        {
            "workspace_id": "51650666",
            "iteration_id": "1151650666001000111",
            "test_plan_id": "1151650666001000018"
        },
        {
            "workspace_id": "51650666",
            "iteration_id": "1151650666001000111",
            "test_plan_id": "1151650666001000017"
        }
    ],
    "info": "success"
}
```


## 字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目id</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">迭代id</td></tr><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;">测试计划id</td></tr></tbody></table>
