# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/update_case.html

更新用例，返回用例

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{version_uid}/cases/{case_uid}`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

PUT

# 请求数限制

一次更新一条数据

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 项目命名空间 |
| RepoUid | 是 | string | 仓库唯一标识 |
| RepoVersionUid | 是 | string | 仓库版本唯一标识 |
| Case | 是 | Case | 用例信息 |

## Case


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Uid | 是 | string | 唯一标识 |
| RepoUid | 是 | string | 仓库唯一标识 |
| RepoVersionUid | 是 | string | 仓库版本唯一标识 |
| FolderUid | 是 | string | 目录唯一标识 |
| Name | 是 | string | 用例名称 |
| Description | 否 | string | 描述信息 |
| Priority | 是 | Priority | 优先级（枚举值：Unknown, P0, P1, P2, P3） |
| PreConditions | 否 | string | 前置条件 |
| Type | 否 | string | 用例类型 |
| StepType | 否 | StepType | 步骤类型（枚举值：STEP, TEXT） |
| Steps | 否 | repeated CaseStep | 步骤列表 |
| StepText | 否 | CaseStep | 文本步骤 |
| Attachments | 否 | repeated CaseAttachment | 附件列表 |
| CustomFields | 否 | repeated Property | 自定义字段 |
| Labels | 否 | repeated Label | 标签列表 |
| Issues | 否 | repeated Issue | 关联问题列表 |
| Owners | 否 | repeated string | 负责人列表 |
| ManHourEstimated | 否 | string | 评估工时 |
| RunTimes | 否 | string | 执行次数 |

## CaseStep


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Content | 否 | string | 步骤描述 |
| ExpectedResult | 否 | string | 预期结果 |

## CaseAttachment


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Key | 是 | string | 附件Key |
| FileName | 是 | string | 文件名 |
| Size | 否 | uint64 | 文件大小 |

## Property


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Name | 是 | string | 属性名称 |
| Label | 是 | string | 显示值 |
| Value | 是 | google.protobuf.Value | 实际值 |
| Url | 否 | string | 超链接地址 |
| flag | 否 | string | 属性标志 |

## Label


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Name | 是 | string | 标签名称 |
| Value | 是 | string | 标签值 |
| Color | 否 | string | 颜色 |
| Uneditable | 否 | bool | 是否可编辑 |
| DisplayName | 否 | string | 前端展示名 |

## Issue


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| IssueUid | 是 | string | 需求唯一标识 |
| IssueName | 是 | string | 需求名称 |
| IssueUrl | 是 | string | 需求链接 |
| WorkspaceUid | 是 | string | 工作空间唯一标识 |
| Type | 是 | Type | 需求类型（枚举值：UNKNOWN, ITERATION, VERSION, STORY, TASK, BUG, TODO, PLAN） |
| Source | 是 | Source | 需求来源（枚举值：NONE, TAPD, LOCAL_WORD） |

# 返回结果


```json
{
"Error": null,
"Data": {
"Audit": {
"Creator": "xxx",
"Updater": "xxx",
"CreatedAt": "2025-07-12T10:58:42+08:00",
"UpdatedAt": "2025-07-12T10:58:42+08:00",
"Tenant": "xx"
},
"Uid": "12572733",
"RepoUid": "17099",
"RepoVersionUid": "18167",
"FolderUid": "12571827",
"FullPath": "/test/",
"UUID": "2f275e19-c053-4af9-b682-07d8c5ed13a6",
"Name": "测试",
"Description": "test desc",
"Priority": "P1",
"PreConditions": "test pre cond",
"Type": "44907",
"StepType": "STEP",
"Steps": [
{
"Id": "16068067",
"Content": "1",
"ExpectedResult": "2",
"NID": "cstp-ezyXrQVLFy"
}
],
"StepText": null,
"Attachments": [
{
"Id": "xxxx",
"Key": "/testimage.png",
"FileName": "testimage.png",
"Size": "223596"
}
],
"CustomFields": [
{
"Name": "ceshizidingyi",
"Label": "测试自定义",
"Value": "45297",
"Url": "",
"flag": ""
}
],
"Labels": [
{
"Name": "34456",
"Value": "",
"Tag": "",
"Color": "#7ad94e",
"Uneditable": false,
"DisplayName": "121",
"Module": ""
}
],
"Source": "TESTX",
"IsManualRelation": false,
"Issues": [
{
"IssueUid": "1166136271001007531",
"Namespace": "",
"WorkspaceUid": "66136271",
"IssueUrl": "66136271",
"Type": "STORY",
"Source": "TAPD",
"Detail": null,
"IssueName": "",
"IsDeleted": false,
"Uid": ""
}
],
"Owners": [
"734242230"
],
"ManHourEstimated": "1",
"Path": ".12571827.",
"RunTimes": "0",
"IsFolder": false
}
}
```

# 测试用例字段说明

## Source


| 取值 | 含义 |
| --- | --- |
| TESTX | 来源于 CT（Continuous Testing），指持续测试 |

## Priority


| 取值 | 含义 |
| --- | --- |
| Unknown | 未知优先级 |
| P0 | 最高优先级 |
| P1 | 高优先级 |
| P2 | 中优先级 |
| P3 | 低优先级 |

## StepType


| 取值 | 含义 |
| --- | --- |
| STEP | 默认类型，即测试用例步骤为非文本类型 |
| TEXT | 文本类型，即测试用例步骤为文本类型 |

## IssueType


| 取值 | 含义 |
| --- | --- |
| UNKNOWN | 未知 |
| ITERATION | 迭代 |
| VERSION | 版本 |
| STORY | 需求 |
| TASK | 子任务 |
| BUG | bug |

## IssueSource


| 取值 | 含义 |
| --- | --- |
| NONE | 无 |
| TAPD | TAPD |
| LOCAL_WORD | word本地文 |