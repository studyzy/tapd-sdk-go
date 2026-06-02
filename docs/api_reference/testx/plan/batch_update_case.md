# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/batch_update_case.html

批量更新计划用例的分配人、执行结果、关联缺陷、注释

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/{plan_uid}/cases/batch-update`

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
| Namespace | 是 | string | 命名空间 |
| PlanUid | 是 | string | 计划UID |
| CaseInfos | 是 | repeated PlanCase | 批量更新的用例信息 |
| Events | 是 | repeated PlanCaseEvent | 计划事件列表 |

## PlanCase


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| State | 是 | State | 用例执行状态 |
| Tester | 是 | string | 测试者 |
| CaseUid | 是 | string | 用例UID |
| Bugs | 是 | repeated Issue | 关联的缺陷列表 |

## State


| 枚举值 | 说明 |
| --- | --- |
| NONE | 无状态 |
| SUCCEED | 通过 |
| FAIL | 失败 |
| BLOCK | 受阻 |
| ERROR | 错误 |
| RETRY | 重测 |
| NOT_EXECUTE | 无法执行 |
| IGNORE | 忽略 |
| NOT_TEST | 未测试 |

## PlanCaseEvent


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Type | 是 | Type | 事件类型 |
| Detail | 是 | string | 事件详情 |
| Source | 是 | Source | 结果来源 |

## Type


| 枚举值 | 说明 |
| --- | --- |
| UNKNOWN | 未知类型 |
| UPDATE_RESULT | 更新用例执行结果 |
| POST_COMMENT | 发表评论 |
| RELATE_ISSUE | 关联需求 |
| UPDATE_TESTER | 指定测试人员 |
| CREATE_REVIEW | 创建评审 |

## Source


| 枚举值 | 说明 |
| --- | --- |
| NONE | 无来源 |
| MANUAL | 手工测试 |

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

# 入参示例


```json
{
"CaseInfos": [{
"CaseUid": "xx",
"Tester": "734242230",
"Bugs": [{
"IssueUid": "xx",
"Namespace": "xx",
"WorkspaceUid": "xx",
"IssueUrl": "",
"Type": "BUG",
"Source": "TAPD",
"IssueName": "test"
}, {
"IssueUid": "xx",
"Namespace": "xx",
"WorkspaceUid": "xx",
"IssueUrl": "",
"Type": "BUG",
"Source": "TAPD",
"IssueName": "dsf"
}],
"State": "SUCCEED"
}, {
"CaseUid": "xx",
"Tester": "734242230",
"Bugs": [{
"IssueUid": "xx",
"Namespace": "xx",
"WorkspaceUid": "xx",
"IssueUrl": "",
"Type": "BUG",
"Source": "TAPD",
"IssueName": "test"
}, {
"IssueUid": "xx",
"Namespace": "xx",
"WorkspaceUid": "xx",
"IssueUrl": "",
"Type": "BUG",
"Source": "TAPD",
"IssueName": "dsf"
}],
"State": "SUCCEED"
}],
"Events": [{
"Type": "POST_COMMENT",
"Source": "MANUAL",
"Detail": "{\"type\":\"comment\",\"comment\":\"xxxxasa\"}"
}]
}
```

# 返回结果


```json
{
"RequestId": "",
"Error": null
}
```