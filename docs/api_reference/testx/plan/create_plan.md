# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/create_plan.html

创建计划

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求数限制

一次创建一条计划

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Plan | 是 | Plan | 计划信息 |

## Plan


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Meta | 是 | PlanMeta | 计划元信息 |
| Spec | 否 | PlanSpec | 计划规格 |

## PlanMeta


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Uid | 是 | string | 计划唯一标识 |
| Namespace | 是 | string | 命名空间 |
| FolderUid | 是 | string | 目录ID |
| Name | 是 | string | 接口库标识名 |
| Description | 否 | string | 描述信息 |
| DataSource | 否 | PlanDataSource | 数据源 |
| State | 否 | State | 计划状态 |
| Testers | 否 | repeated string | 测试人员列表 |
| FolderPath | 否 | string | 所属目录绝对路径 |
| Version | 否 | Version | 计划API版本 |
| CasePath | 否 | CasePath | 用例存储路径 |
| Nid | 否 | string | 序号 |
| Path | 否 | string | 路径 |

## PlanSpec


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Field | PlanFieldSpec | 计划字段 |
| CustomFields | repeated Property | 自定义属性 |
| Stories | repeated Issue | 需求列表 |
| Bugs | repeated Issue | 缺陷列表 |
| Iterations | repeated Issue | 迭代列表 |
| Statistic | PlanStatistic | 统计信息 |
| Properties | repeated Property | 计划扩展属性 |
| SystemFields | repeated Property | 系统字段 |
| Scope | PlanScope | 计划范围 |
| Target | PlanTarget | 计划目标 |
| Versions | repeated Issue | 关联版本列表 |

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

## PlanStatistic


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| SucceedCaseCount | uint32 | 成功用例数 |
| FailedCaseCount | uint32 | 失败用例数 |
| BlockCaseCount | uint32 | 阻塞用例数 |
| ErrorCaseCount | uint32 | 错误用例数 |
| RetryCaseCount | uint32 | 重试用例数 |
| IgnoreCaseCount | uint32 | 忽略用例数 |
| TodoCaseCount | uint32 | 未测试用例数 |
| ManualSucceedRate | string | 手工用例成功率 |
| TestedCaseCount | uint32 | 已测试用例数 |
| TotalCaseCount | uint32 | 用例总数 |
| CaseSucceedRate | string | 用例成功率 |
| StoryPassedRate | string | 需求通过率 |
| FailedTaskCount | uint32 | 未通过任务数 |
| CaseCoverageRate | string | 用例覆盖率 |
| StoryCount | uint32 | 功能需求数 |

## PlanFieldSpec


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Developers | repeated string | 开发人员列表 |
| ProjectManagers | repeated string | 项目经理列表 |
| ProductManagers | repeated string | 产品经理列表 |
| Summary | string | 总结 |
| Version | string | 版本 |
| Type | string | 类型 |
| ExtranetPublish | string | 外网发布 |
| PublishReview | string | 发布评审 |
| EstimateTestAt | string | 预计提测时间 |
| ActualTestAt | string | 实际提测时间 |
| EstimatePublishAt | string | 预计发布时间 |
| ActualPublishAt | string | 实际发布时间 |
| EstimateStartedAt | string | 测试预计开始时间 |
| EstimateEndedAt | string | 测试预计完成时间 |
| Env | string | 测试环境 |

## Property


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Name | string | 属性名称 |
| Label | string | 显示值 |
| Value | google.protobuf.Value | 实际值 |
| Url | string | 超链接地址 |
| flag | string | 属性标志 |

## PlanScope


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Stories | repeated Issue | 需求列表 |
| StoryCaseStrategy | StoryCaseStrategy | 用例关联策略 |

## StoryCaseStrategy


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Mode | Mode | 关联模式 |
| Priorities | repeated Priority | 用例优先级列表 |

## Priority


| 枚举值 | 说明 |
| --- | --- |
| Unknown | 未知优先级 |
| P0 | 最高优先级 |
| P1 | 高优先级 |
| P2 | 中优先级 |
| P3 | 低优先级 |

## StoryCaseStrategyMode


| 枚举值 | 说明 |
| --- | --- |
| UNKNOWN | 未知 |
| ALL | 关联全部用例 |
| PATCH | 按条件关联部分用例 |
| NOT | 不关联用例 |
| SELECT_CASE | 取需求用例和入参用例的交集 |

## PlanTarget


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Description | string | 目标描述 |
| Feature | string | 涉及功能 |
| Deliverables | string | 交付成果 |
| SuccessConditions | repeated string | 通过标准 |

## State


| 枚举值 | 说明 |
| --- | --- |
| NONE | 无状态 |
| WAITING | 等待中 |
| RUNNING | 运行中 |
| DONE | 已完成 |
| ARCHIVE | 已归档 |

## Version


| 枚举值 | 说明 |
| --- | --- |
| V1 | 版本1 |
| V2 | 版本2 |

## CasePath


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| RepoUid | string | 用例库版本ID |
| RepoVersionUid | string | 用例库版本ID |
| FolderUid | string | 父用例目录ID |

## PlanDataSource


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Mode | Mode | 数据源模式 |
| CaseUids | repeated string | 用例UID列表 |
| CasesetUids | repeated string | 用例集UID列表 |
| CaseSuiteUid | string | 测试套件UID |

## Mode


| 枚举值 | 说明 |
| --- | --- |
| UNKNOWN | 未知模式 |
| CASE_REPO | 用例库模式 |
| CASE_SET | 用例集模式 |

# 入参示例


```json
// 简单创建
{
"Plan": {
"Meta": {
"FolderUid": "xxx", //计划目录id
"DataSource": {},
"Testers": [],
"Name": "测试", // 计划名称
"Description": ""
}
}
}
```

# 返回结果


```json
{
"Error": null,
"Data": {
"Meta": {
"Uid": "xx",
"Namespace": "xxx",
"Audit": {
"Creator": "xxx",
"CreatedAt": "2025-07-14T14:18:56+08:00",
"StartedAt": "",
"EndedAt": "",
"Starter": "",
"Terminator": "xx",
"UpdatedAt": "2025-07-14T14:18:56+08:00",
"Updater": ""
},
"FolderUid": "xx",
"Name": "测试",
"Description": "",
"DataSource": {
"Mode": "CASE_REPO",
"CaseUids": [],
"CasesetUids": [],
"CaseSuiteUid": "xx"
},
"State": "WAITING",
"Testers": [],
"FolderPath": "/test/test子目录",
"Version": "V2",
"CasePath": {
"RepoUid": "",
"RepoVersionUid": "",
"FolderUid": ""
},
"Nid": "",
"Path": ".152316.152317."
},
"Spec": {
"Field": {
"Developers": [],
"ProjectManagers": [],
"ProductManagers": [],
"Summary": "",
"Version": "",
"Type": "",
"ExtranetPublish": "",
"PublishReview": "",
"EstimateTestAt": "",
"ActualTestAt": "",
"EstimatePublishAt": "",
"ActualPublishAt": "",
"EstimateStartedAt": "",
"EstimateEndedAt": "",
"Env": ""
},
"CustomFields": [],
"Stories": [],
"Bugs": [],
"Iterations": [],
"Statistic": {
"SucceedCaseCount": 0,
"FailedCaseCount": 0,
"BlockCaseCount": 0,
"ErrorCaseCount": 0,
"RetryCaseCount": 0,
"IgnoreCaseCount": 0,
"TodoCaseCount": 0,
"ManualSucceedRate": "",
"TestedCaseCount": 0,
"TotalCaseCount": 0,
"CaseSucceedRate": "",
"StoryPassedRate": "",
"FailedTaskCount": 0,
"CaseCoverageRate": "",
"StoryCount": 0
},
"Properties": [],
"SystemFields": [],
"Scope": {
"Stories": [],
"StoryCaseStrategy": {
"Mode": "UNKNOWN",
"Priorities": []
}
},
"Target": {
"Description": "",
"Feature": "",
"Deliverables": "",
"SuccessConditions": [
"",
""
]
},
"Versions": []
}
}
}
```