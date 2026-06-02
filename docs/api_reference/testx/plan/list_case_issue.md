# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/list_case_issue.html

获取计划下用例关联的缺陷列表

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/{plan_uid}/cases/{case_uid}/issues`

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
| Type | 是 | IssueType | 问题类型 |
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
"IssueUid": "issue_uid",
"Namespace": "Namespace",
"WorkspaceUid": "WorkspaceUid",
"IssueUrl": "",
"Type": "BUG",
"Source": "NONE",
"Detail": null,
"IssueName": "",
"IsDeleted": false,
"Uid": ""
},
{
"IssueUid": "issue_uid_2",
"Namespace": "Namespace",
"WorkspaceUid": "WorkspaceUid",
"IssueUrl": "",
"Type": "BUG",
"Source": "NONE",
"Detail": null,
"IssueName": "",
"IsDeleted": false,
"Uid": ""
}
],
"TotalCount": 2
}
```

# 测试计划Issue说明

## Issue


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| IssueUid | string | 问题唯一标识 |
| Namespace | string | 命名空间 |
| WorkspaceUid | string | 工作空间ID |
| Type | Type | 问题类型 |
| Source | Source | 问题来源 |
| Detail | google.protobuf.Struct | 问题详情数据 |
| IssueName | string | 问题名称 |
| IsDeleted | bool | 是否已删除 |

## IssueType


| 枚举值 | 说明 |
| --- | --- |
| UNKNOWN | 未知类型 |
| ITERATION | 迭代 |
| VERSION | 版本 |
| STORY | 需求 |
| TASK | 子任务 |
| BUG | 缺陷 |

## IssueSource


| 枚举值 | 说明 |
| --- | --- |
| NONE | 无来源 |
| TAPD | TAPD |