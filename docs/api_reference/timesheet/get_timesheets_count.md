# 获取工时花费的数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/get_timesheets_count.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下花费工时的数量](#获取项目下花费工时的数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

计算符合查询条件的花费工时数量并返回

# url

`https://api.tapd.cn/timesheets/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回花费工时数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">对象类型，如story、task、bug等</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">对象ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">timespent</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">spentdate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">花费日期</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">is_delete</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否已删除。默认取 0，不返回已删除的工时记录。取 1 可以返回已删除的记录</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下花费工时的数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/timesheets/count?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 14
    },
    "info": "success"
}
```

