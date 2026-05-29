# 更新系统字段

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/update_system_select_field_options.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [修改bugtype的选项为test和test111](#修改bugtype的选项为test和test111)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [重要字段说明](#重要字段说明)

# 说明

修改缺陷系统字段选项,会覆盖掉原有选项

# url

`https://api.tapd.cn/bugs/update_system_select_field_options`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次返回所有符合条件的值

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">字段 目前支持: bugtype（缺陷类型）</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">array</td><td style="text-align:center;">选项列表</td><td style="text-align:center;">无</td></tr><tr><td style="text-align:center;">value</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">选项对应value</td><td style="text-align:center;">无</td></tr></tbody></table>

# 调用示例及返回结果

## 修改bugtype的选项为test和test111

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/update_system_select_field_options --data '{"workspace_id": 62187798,"field": "bugtype","options": [{"value": "test"},{"value": "test111"}]}'`

### 返回结果

```
{
    "status": 1,
    "data": true,
    "info": "success"
}
```


# 字段说明

## 重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody></tbody></table>
