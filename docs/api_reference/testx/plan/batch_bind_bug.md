# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/batch_bind_bug.html

计划用例批量添加缺陷

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/{plan_uid}/cases/bugs`

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
| CaseUids | 是 | repeated string | 用例UID列表 |
| BindBugs | 是 | repeated Issue | 绑定的bug信息列表 |

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
"CaseUids": ["case_uid", "case_uid_1"],
"BindBugs": [{
"Type": "BUG",
"Source": "TAPD",
"IssueUid": "Bug_Id",
"IssueName": "test",
"Namespace": "",
"WorkspaceUid": "WorkspaceUid"
}, {
"Type": "BUG",
"Source": "TAPD",
"IssueUid": "Bug_Id",
"IssueName": "dsf",
"Namespace": "",
"WorkspaceUid": "WorkspaceUid"
}],
"StoryIds": []
}
```

# 返回结果


```json
{
"RequestId": "",
"Error": null
}
```