# 获取测试计划数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plans_count.html

# 说明

计算符合查询条件的测试计划数量并返回

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTestPlansCount
```


# url

`https://api.tapd.cn/test_plans/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回评论数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划详细描述</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">修改人</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划负责人</td></tr><tr><td style="text-align:center;">start_date</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">关联迭代的ID</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本号</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">测试类型</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态，默认开启，值为open</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/test_plans/count?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 4
    },
    "info": "success"
}
```

