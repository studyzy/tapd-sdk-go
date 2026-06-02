# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/list_case_bugs.html

获取用例关联的bug记录

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{repo_version_uid}/cases/{case_uid}/bugs`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间 |
| RepoUid | 是 | string | 仓库UID |
| RepoVersionUid | 是 | string | 仓库版本UID |
| CaseUid | 是 | string | 用例UID |
| Status | 否 | Status | 缺陷状态 |
| Priority | 否 | string | 缺陷优先级 |
| Handler | 否 | string | 缺陷处理人 |
| PageInfo | 否 | PageInfo | 分页信息 |
| Name | 否 | string | 缺陷标题 |
| SourceTypes | 否 | repeated Source | 缺陷来源类型列表 |

## Source


| 枚举值 | 说明 |
| --- | --- |
| UNKNOWN_SOURCE | 未知来源 |
| DESIGN | 来源于测试执行 |
| CASE | 来源于用例 |
| PLAN | 来源于计划 |
| DESIGN_RELATED | 来源于测试设计 |

## Status


| 枚举值 | 说明 |
| --- | --- |
| STATUS_ALL | 全部状态 |
| NEW | 新 |
| POSTPONED | 延期 |
| REJECTED | 已拒绝 |
| REOPENED | 重新打开 |
| RESOLVED | 已解决 |
| SUSPENDED | 挂起 |
| UNCONFIRMED | 待确认 |
| VERIFIED | 已验证 |
| IN_PROGRESS | 接受/处理中 |
| ASSIGNED | 已分配 |
| CLOSED | 已关闭 |

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
"Bug": {
"WorkspaceId": "xxx",
"IterationId": "0",
"Id": "xxxx",
"Summary": "test",
"Creator": "xxx",
"Priority": {
"Label": "P2",
"Value": "P2"
},
"Created": "2025-03-18 15:30:37",
"Status": {
"Label": "新",
"Value": "NEW"
},
"LastUpdater": "ivanzbzhang",
"UpdateTime": "2025-06-12 10:11:25",
"Severity": {
"Label": "",
"Value": "SEVERITY_INIT"
},
"Owners": "",
"Begin": "",
"Due": "",
"NamespaceId": "",
"Description": "【问题描述】

fsf

【截图】

【链接地址】

",
"Url": "",
"Developers": "",
"Testers": "",
"Module": "",
"Attachments": [],
"Tags": [],
"RelatedType": "",
"IconType": "",
"Version": [],
"ReleaseId": "",
"VersionReport": "",
"VersionTest": "",
"VersionFix": "",
"VersionClose": "",
"BaselineFind": "",
"BaselineJoin": "",
"BaselineClose": "",
"BaselineTest": "",
"CC": "",
"Participator": "",
"Auditer": "",
"Confirmer": "",
"Fixer": "",
"Closer": "",
"ReopenTime": "",
"InProgressTime": "",
"Resolved": "",
"VerifyTime": "",
"Closed": "",
"RejectTime": "",
"Modified": "2025-06-12 10:11:25",
"Deadline": "",
"Os": "",
"Platform": "",
"Testmode": "",
"Testphase": "",
"Testtype": "",
"Source": "",
"Frequency": "",
"Originphase": "",
"Sourcephase": "",
"Resolution": {
"Label": "",
"Value": "RESOLUTION_INIT"
},
"Estimate": "",
"Lastmodify": "",
"CustomFields": {},
"TemplateId": "",
"IsApplyTemplateDefaultValue": false,
"CustomPlanFields": {},
"Feature": "",
"Effort": "",
"Bugtype": "",
"Size": ""
},
"Sources": [
{
"Uid": "xxx",
"Type": "CASE",
"Name": "评估环节退出确认-用户主动退出",
"Label": "用例关联",
"Url": ""
}
]
}
]    
"TotalCount": 1
}
```

# 测试用例bug字段说明

## Bug


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| WorkspaceId | string | 工作空间ID |
| IterationId | string | 迭代ID |
| Id | string | 缺陷ID |
| Summary | string | 标题 |
| Creator | string | 创建人 |
| Priority | BugPriority | 优先级 |
| Created | string | 创建时间 |
| Status | BugStatus | 状态 |
| LastUpdater | string | 最后更新人 |
| UpdateTime | string | 更新时间 |
| Severity | BugSeverity | 严重程度 |
| Owners | string | 处理人 |
| Begin | string | 预计开始时间 |
| Due | string | 结束时间 |
| NamespaceId | string | 命名空间ID |
| Description | string | 描述 |
| Url | string | URL |
| Developers | string | 开发人员 |
| Testers | string | 测试人员 |
| Module | string | 模块 |
| Attachments | repeated Attachment | 附件列表 |
| Tags | repeated string | 标签 |
| RelatedType | string | 关联类型 |
| IconType | string | 图标类型 |
| Version | repeated string | 版本 |
| ReleaseId | string | 发布计划ID |
| VersionReport | string | 发现版本 |
| VersionTest | string | 验证版本 |
| VersionFix | string | 合入版本 |
| VersionClose | string | 关闭版本 |
| BaselineFind | string | 发现基线 |
| BaselineJoin | string | 合入基线 |
| BaselineClose | string | 关闭基线 |
| BaselineTest | string | 验证基线 |
| CC | string | 抄送人 |
| Participator | string | 参与人 |
| Auditer | string | 审核人 |
| Confirmer | string | 验证人 |
| Fixer | string | 修复人 |
| Closer | string | 关闭人 |
| ReopenTime | string | 重新打开时间 |
| InProgressTime | string | 接受处理时间 |
| Resolved | string | 解决时间 |
| VerifyTime | string | 验证时间 |
| Closed | string | 关闭时间 |
| RejectTime | string | 拒绝时间 |
| Modified | string | 最后修改时间 |
| Deadline | string | 解决期限 |
| Os | string | 操作系统 |
| Platform | string | 软件平台 |
| Testmode | string | 测试方式 |
| Testphase | string | 测试阶段 |
| Testtype | string | 测试类型 |
| Source | string | 缺陷根源 |
| Frequency | string | 重现规律 |
| Originphase | string | 发现阶段 |
| Sourcephase | string | 引入阶段 |
| Resolution | BugResolution | 解决方法 |
| Estimate | string | 预计工时 |
| Lastmodify | string | 最后修改人 |
| CustomFields | map<string, string> | 自定义字段 |
| TemplateId | string | 模板ID |
| IsApplyTemplateDefaultValue | bool | 是否从模板继承默认值 |
| CustomPlanFields | map<string, string> | 自定义计划字段 |
| Feature | string | 特性 |
| Effort | string | 预计工时 |
| Bugtype | string | 缺陷类型 |
| Size | string | 规模 |

## BugStatus


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Label | string | 状态标签 |
| Value | State | 状态值 |

## BugPriority


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Label | string | 优先级标签 |
| Value | string | 优先级值 |

## BugSeverity


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Label | string | 严重程度标签 |
| Value | Severity | 严重程度值 |

## BugResolution


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Label | string | 解决方法标签 |
| Value | Resolution | 解决方法值 |

## Status


| 枚举值 | 说明 |
| --- | --- |
| STATUS_ALL | 全部状态 |
| NEW | 新 |
| POSTPONED | 延期 |
| REJECTED | 已拒绝 |
| REOPENED | 重新打开 |
| RESOLVED | 已解决 |
| SUSPENDED | 挂起 |
| UNCONFIRMED | 待确认 |
| VERIFIED | 已验证 |
| IN_PROGRESS | 接受/处理中 |
| ASSIGNED | 已分配 |
| CLOSED | 已关闭 |

## Severity


| 值 | 说明 |
| --- | --- |
| SEVERITY_ALL | 所有严重程度 |
| FATAL | 致命 |
| NORMAL | 一般 |
| PROMPT | 提示 |
| SERIOUS | 严重 |
| ADVICE | 建议 |
| SEVERITY_INIT | 未设置 |

## Resolution


| 值 | 说明 |
| --- | --- |
| RESOLUTION_ALL | 所有解决方法 |
| IGNORE | 无需解决 |
| FIXED | 已修改 |
| FIX_LATER | 延期解决 |
| FAILED_TO_RECUR | 无法重现 |
| EXTERNAL_REASON | 外部原因 |
| DUPLICATED | 重复 |
| INTENTIONAL_DESIGN | 设计如此 |
| UNCLEAR_DESCRIPTION | 问题描述不准确 |
| FEATURE_CHANGE | 需求变更 |
| TRANSFERRED_TO_STORY | 已转需求 |
| HOLD | 挂起 |
| RESOLUTION_INIT | 未设置 |

## Attachment


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Id | string | 附件ID |
| DownloadUrl | string | 下载URL |
| Type | string | 类型 |
| EntryId | string | 条目ID |
| Filename | string | 文件名 |
| Description | string | 描述 |
| ContentType | string | 内容类型 |
| Created | string | 创建时间 |
| WorkspaceId | string | 工作空间ID |
| Owner | string | 所有者 |