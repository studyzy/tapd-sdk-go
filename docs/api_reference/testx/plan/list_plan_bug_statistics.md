# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/list_plan_bug_statistics.html

批量查询计划关联缺陷统计数

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/bug-statistics`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| PlanUids | 是 | repeated string | 计划ID列表(不能为空) |
| Namespace | 是 | string | 命名空间 |

## pageinfo


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Offset | 是 | uint32 | 分页偏移位置 |
| Limit | 是 | uint32 | 每页数量 |

返回结果


```json
{
"Error": null,
"Data": [
{
"PlanUid": "plan_uid",
"BugCount": 0
}
]
}
```

# bug统计字段说明

## PlanBugStatistics


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| PlanUid | 是 | string | 所属计划ID |
| BugCount | 是 | uint32 | 缺陷总数 |