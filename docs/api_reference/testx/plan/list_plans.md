# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/list_plans.html

获取目录下计划列表

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/folders/{folder_uid}/plans-list`

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
| FolderUid | 是 | uint64 | 目录ID(0表示根目录) |
| WithDetail | 否 | bool | 是否包含详情 |
| PageInfo | 否 | pageinfo | 分页信息 |
| Ordering | 否 | string | 排序字段 |
| States | 否 | repeated State | 状态过滤 |
| Testers | 否 | repeated string | 测试者过滤 |
| Uids | 否 | repeated string | 计划ID列表 |
| Archive | 否 | ArchiveMode | 归档状态 |
| Name | 否 | string | 名称模糊查询 |
| Description | 否 | string | 描述模糊查询 |
| Assigned | 否 | NullBool | 是否分配 |
| IterationUids | 否 | repeated string | 迭代UID列表 |
| IgnoreFolderUid | 否 | bool | 是否忽略目录ID条件 |
| StoryUids | 否 | repeated string | 需求UID列表 |
| StartAt | 否 | string | 开始时间 |
| EndAt | 否 | string | 结束时间 |
| TimeIntervalType | 否 | TimeIntervalType | 时间区间类型 |
| Fields | 否 | repeated string | 返回字段列表 |

## ArchiveMode


| 枚举值 | 说明 |
| --- | --- |
| ARCHIVEMODE_NONE | 获取全部 |
| ARCHIVEMODE_YES | 获取归档 |
| ARCHIVEMODE_NO | 获取非归档 |

## TimeIntervalType


| 枚举值 | 说明 |
| --- | --- |
| CREATE_TIME | 创建时间 |
| UPDATE_TIME | 更新时间 |

## State


| 枚举值 | 说明 |
| --- | --- |
| NONE | 无状态 |
| WAITING | 等待中 |
| RUNNING | 运行中 |
| DONE | 已完成 |
| ARCHIVE | 已归档 |

## NullBool


| 枚举值 | 说明 |
| --- | --- |
| NULL_BOOl_NONE | 默认值(全部) |
| NULL_BOOl_FALSE | 未分配 |
| NULL_BOOl_TRUE | 已分配 |

## pageinfo


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Offset | 是 | uint32 | 分页偏移位置 |
| Limit | 是 | uint32 | 每页数量 |

# 返回结果


```json
{
"Error": null,
"Data": [{
"Meta": {
"Uid": "Uid",
"Namespace": "Namespace",
"Audit": {
"Creator": "Creator",
"CreatedAt": "2025-07-14T14:38:42+08:00",
"StartedAt": "",
"EndedAt": "",
"Starter": "",
"Terminator": "Terminator",
"UpdatedAt": "2025-07-14T14:38:42+08:00",
"Updater": ""
},
"FolderUid": "FolderUid",
"Name": "测试",
"Description": "",
"DataSource": {
"Mode": "CASE_REPO",
"CaseUids": [],
"CasesetUids": [],
"CaseSuiteUid": "CaseSuiteUid"
},
"State": "WAITING",
"Testers": [
"tester1",
"tester2"
],
"FolderPath": "",
"Version": "V2",
"CasePath": {
"RepoUid": "",
"RepoVersionUid": "",
"FolderUid": ""
},
"Nid": "",
"Path": ".xxxx."
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
"AutoTasks": [],
"Statistic": null,
"Properties": [],
"SystemFields": [],
"Scope": null,
"Target": null,
"Versions": []
}
}
],
"TotalCount": 1
}
```

# 测试计划字段说明

## Data


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Meta | PlanMeta | 计划元信息 |
| Spec | PlanSpec | 计划规格 |

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

## PlanMeta


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Uid | string | 计划唯一标识 |
| Namespace | string | 命名空间 |
| FolderUid | string | 目录ID |
| Name | string | 接口库标识名 |
| Description | string | 描述信息 |
| DataSource | PlanDataSource | 数据源 |
| State | State | 计划状态 |
| Testers | repeated string | 测试人员列表 |
| FolderPath | string | 所属目录绝对路径 |
| Version | Version | 计划API版本 |
| CasePath | CasePath | 用例存储路径 |
| Nid | string | 序号 |
| Path | string | 路径 |

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