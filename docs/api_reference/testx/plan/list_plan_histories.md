# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/list_plan_histories.html

获取计划变更历史

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/{plan_uid}/histories`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| PlanUid | 是 | string | 计划ID |
| Namespace | 是 | string | 命名空间 |
| PageInfo | 否 | PageInfo | 分页信息 |

## pageinfo


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Offset | 是 | uint32 | 分页偏移位置 |
| Limit | 是 | uint32 | 每页数量 |

# 返回结果


```json
{
"Error": null,
"Data": [
{
"Uid": "Uid",
"Creator": "Creator",
"CreateAt": "2025-07-16T17:12:10+08:00",
"PlanUid": "PlanUid",
"Fields": [
{
"Uid": "Uid",
"Name": "测试字段一",
"PreValue": "test-pre",
"PostValue": "test"
},
{
"Uid": "Uid",
"Name": "测试字段二",
"PreValue": "",
"PostValue": "21"
}
]
}
],
"TotalCount": 1
}
```

# 测试计划字段说明

## PlanHistory


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Uid | 是 | string | 计划变更历史ID |
| Creator | 是 | string | 记录人 |
| CreateAt | 是 | string | 记录时间 |
| PlanUid | 是 | string | 所属计划ID |
| Fields | 是 | repeated PlanHistoryField | 变更字段详情 |

## PlanHistoryField


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Uid | 是 | string | 字段变更ID |
| Name | 是 | string | 变更字段名称 |
| PreValue | 是 | string | 变更前值 |
| PostValue | 是 | string | 变更后值 |