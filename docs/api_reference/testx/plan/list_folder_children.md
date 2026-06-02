# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/list_folder_children.html

获取计划目录子信息，返回目录下的子目录和计划信息

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/folders/children`

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
| Uid | 是 | string | 计划目录ID |
| Namespace | 是 | string | 命名空间 |
| WithDescendant | 否 | bool | 是否级联获取所有子目录 |
| WithAncestor | 否 | bool | 是否获取所有祖先目录 |
| Name | 否 | string | 名称模糊查询 |
| PlanArchive | 否 | ArchiveMode | 计划归档状态 |
| PlanStates | 否 | repeated State | 计划状态过滤 |
| ItemType | 否 | ItemType | 返回内容类型 |

## ItemType


| 枚举值 | 说明 |
| --- | --- |
| ALL | 返回目录和计划 |
| FOLDER | 仅返回目录 |

## State


| 枚举值 | 说明 |
| --- | --- |
| NONE | 无状态 |
| WAITING | 等待中 |
| RUNNING | 运行中 |
| DONE | 已完成 |
| ARCHIVE | 已归档 |

## ArchiveMode


| 枚举值 | 说明 |
| --- | --- |
| ARCHIVEMODE_NONE | 获取全部 |
| ARCHIVEMODE_YES | 获取归档 |
| ARCHIVEMODE_NO | 获取非归档 |

# 返回结果


```json
{
"Error": null,
"Data": {
"Folders": [
{
"Uid": "xxx",
"Namespace": "xxx",
"Audit": {
"Creator": "xxx",
"Updater": "xxx",
"CreatedAt": "2025-07-14T14:18:40+08:00",
"UpdatedAt": "2025-07-14T14:18:40+08:00",
"Tenant": ""
},
"ParentUid": "xxx",
"Name": "test子目录",
"Description": "",
"PlanCount": 1,
"ArchiveAuto": false,
"Folders": [],
"Plans": [],
"Nid": "",
"Path": ".152316."
}
],
"Plans": [
{
"Uid": "xx",
"Namespace": "xx",
"Audit": {
"Creator": "xx",
"CreatedAt": "2025-07-14T14:31:08+08:00",
"StartedAt": "",
"EndedAt": "",
"Starter": "",
"Terminator": "xx",
"UpdatedAt": "2025-07-14T14:31:08+08:00",
"Updater": ""
},
"FolderUid": "xx",
"Name": "test计划2",
"Description": "",
"DataSource": {
"Mode": "CASE_REPO",
"CaseUids": [],
"CasesetUids": [],
"CaseSuiteUid": "1788319"
},
"State": "WAITING",
"Testers": [],
"FolderPath": "",
"Version": "V2",
"CasePath": {
"RepoUid": "",
"RepoVersionUid": "",
"FolderUid": ""
},
"Nid": "",
"Path": ".152316."
}
]
}
}
```

# 测试计划目录字段说明

## Data


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Folders | repeated Folder | 目录列表 |
| Plans | repeated PlanMeta | 计划列表 |

## Folder


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Uid | string | 目录唯一标识 |
| Namespace | string | 命名空间 |
| Audit | testx.common.Audit | 审计信息 |
| ParentUid | string | 父目录ID |
| Name | string | 接口库标识名 |
| Description | string | 描述信息 |
| PlanCount | uint32 | 计划数量 |
| ArchiveAuto | bool | 已完成计划自动归档 |
| Folders | repeated Folder | 子目录基本信息 |
| Plans | repeated PlanMeta | 计划基本信息 |
| Path | string | 路径 |

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