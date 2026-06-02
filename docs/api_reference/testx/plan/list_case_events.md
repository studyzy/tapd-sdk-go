# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/list_case_events.html

获取计划下用例的事件列表

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/{plan_uid}/cases/{case_uid}/events`

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
| Namespace | 是 | string | 命名空间/项目唯一标识 |
| PlanUid | 是 | string | 计划UID |
| CaseUid | 是 | string | 用例UID |
| PageInfo | 否 | PageInfo | 分页信息(缺省获取全部) |

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
"Audit": {
"Creator": "xxx",
"Updater": "xxx",
"CreatedAt": "2025-06-20T11:07:27+08:00",
"UpdatedAt": "2025-06-20T11:07:27+08:00",
"Tenant": ""
},
"Type": "UPDATE_RESULT",
"Detail": "{\"type\":\"update_result\",\"status\":\"fail\"}",
"Source": "MANUAL",
"Attachments": []
},
{
"Audit": {
"Creator": "xxx",
"Updater": "xxx",
"CreatedAt": "2025-06-20T11:07:25+08:00",
"UpdatedAt": "2025-06-20T11:07:25+08:00",
"Tenant": ""
},
"Type": "UPDATE_RESULT",
"Detail": "{\"type\":\"update_result\",\"status\":\"succeed\"}",
"Source": "MANUAL",
"Attachments": []
},
{
"Audit": {
"Creator": "xxx",
"Updater": "xxx",
"CreatedAt": "2025-06-20T11:07:17+08:00",
"UpdatedAt": "2025-06-20T11:07:17+08:00",
"Tenant": ""
},
"Type": "UPDATE_TESTER",
"Detail": "{\"type\":\"update_tester\",\"tester\":\"xxx\"}",
"Source": "MANUAL",
"Attachments": []
}
],
"TotalCount": 3
}
```

# 测试计划字段说明

## Data


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Type | 是 | Type | 事件类型 |
| Detail | 是 | string | 事件详细信息 |
| Source | 是 | Source | 结果来源 |

## Type


| 值 | 说明 |
| --- | --- |
| UNKNOWN | 未知类型 |
| UPDATE_RESULT | 用例执行结果 |
| POST_COMMENT | 发表评论 |
| RELATE_ISSUE | 关联需求 |
| UPDATE_TESTER | 指定测试人员 |
| CREATE_REVIEW | 创建评审 |

## Source


| 值 | 说明 |
| --- | --- |
| NONE | 无来源 |
| MANUAL | 手动 |