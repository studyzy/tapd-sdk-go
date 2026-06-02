# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/list_plan_cases.html

获取计划下用例列表

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/{uid}/cases-search`

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
| Uid | 是 | string | 计划UID |
| ShowMode | 否 | ShowMode | 用例展示模式 |
| PageInfo | 否 | PageInfo | 分页信息 |
| Testers | 否 | string list | 分配owner过滤 |
| States | 否 | State list | 用例状态过滤 |
| Sources | 否 | Source list | 用例来源过滤 |
| CaseFilter | 否 | Filter | 用例过滤条件 |
| WithAncestor | 否 | bool | 是否返回父目录 |
| RepoVersionUid | 否 | string | 用例库版本ID |
| CaseSelectFields | 否 | string list | 用例字段过滤 |
| SelectFields | 否 | SelectField list | 关联字段 |
| Version | 否 | Version | 接口版本 |
| WithFullPath | 否 | bool | 用例是否包含全路径 |
| CaseExtendFields | 否 | repeated string | 用例扩展字段 |
| StoryShowMode | 否 | StoryShowMode | 需求展示模式 |
| WithCaseReviewState | 否 | bool | 是否返回用例评审状态 |

## ShowMode


| 枚举值 | 说明 |
| --- | --- |
| FLAT | 平铺展示 |
| TREE | 树形展示 |

## PlanCaseInternalState


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

## PlanCaseSource


| 枚举值 | 说明 |
| --- | --- |
| NULL | 无来源 |
| MANUAL | 手动执行 |

## ListPlanCasesRequestSelectField


| 枚举值 | 说明 |
| --- | --- |
| NONE | 不获取关联字段 |
| STATISTIC | 获取统计信息 |
| EVENT | 获取事件 |
| BUG | 获取缺陷 |

## PlanMetaVersion


| 枚举值 | 说明 |
| --- | --- |
| V1 | v1版本 |
| V2 | v2版本 |

## Filter


| 字段名 | 类型 | 说明 |
| --- | --- | --- |
| Name | string | 模糊匹配用例名称 |
| Uuid | string | 精确匹配 UUID |
| Priorities | Priority list | 匹配用例优先级 |
| Description | string | 模糊匹配用例描述 |
| Labels | string list | 匹配用例标签 |
| ReviewStates | StepType list | 匹配用例审核状态 |
| Owners | string list | 精确匹配用例责任人 |
| CustomFields | Property list | 匹配自定义字段 |
| Issues | string list | 精确匹配用例关联的 issue |
| ItemType | ItemType | 是否仅返回用例/目录 |
| Creators | string list | 精确匹配用例创建人 |

## Priority


| 枚举值 | 说明 |
| --- | --- |
| Unknown | 未知优先级 |
| P0 | 最高优先级 |
| P1 | 高优先级 |
| P2 | 中优先级 |
| P3 | 低优先级 |

## StepType


| 枚举值 | 说明 |
| --- | --- |
| STEP | 默认类型，即测试用例步骤为非文本类型 |
| TEXT | 文本类型，即测试用例步骤为文本类型 |

## Property


| 字段名 | 类型 | 说明 |
| --- | --- | --- |
| Name | string | 属性名称 |
| Label | string | 显示值 |
| Value | google.protobuf.Value | 实际值 |

## ItemType


| 枚举值 | 说明 |
| --- | --- |
| ALL | 文件夹和用例 |
| FOLDER | 仅展示文件夹 |
| CASE | 仅展示用例 |

## pageinfo


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Offset | 是 | uint32 | 分页偏移位置 |
| Limit | 是 | uint32 | 每页数量 |

# 入参示例


```json
{
"ShowMode": "TREE",
"Version": "V2",
"StoryShowMode": "TREE",
"WithCaseReviewState": true,
"CaseFilter": {}
}
```

# 返回结果


```json
{
"Error": null,
"Data": {
"CaseUidToDetail": {
"testcase": {
"State": "BLOCK",
"Tester": "",
"FinalTester": "FinalTester",
"EndedAt": "2025-07-18T17:05:26+08:00",
"Source": "MANUAL",
"PlanUid": "testplan",
"CaseUid": "testcase",
"CaseNid": "",
"Bugs": [{
"IssueUid": "IssueUid",
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
"IssueUid": "IssueUid",
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
"Events": [],
"Statistic": null,
"Uid": "Uid",
"CaseName": "",
"PlanName": "",
"PlanFolderUid": "",
"RunTimes": 2,
"ReviewState": "PENDING_REVIEW"
}
},
"PlanCasesInfo": {
"StoryGroup": [{
"Story": {
"IssueUid": "IssueUid",
"Namespace": "",
"WorkspaceUid": "WorkspaceUid",
"IssueUrl": "",
"Type": "STORY",
"Source": "NONE",
"Detail": {
"CategoryId": "-1",
"Completed": "0",
"Created": "2025-03-10 14:47:18",
"Creator": "TAPD",
"CustomFields": {
"custom_field_10": "",
"custom_field_100": "",
"custom_field_101": "",
"custom_field_102": "",
"custom_field_103": "",
"custom_field_104": "",
"custom_field_105": "",
"custom_field_106": "",
"custom_field_107": "",
"custom_field_108": "",
"custom_field_109": "",
"custom_field_11": "",
"custom_field_110": "",
"custom_field_111": "",
"custom_field_112": "",
"custom_field_113": "",
"custom_field_114": "",
"custom_field_115": "",
"custom_field_116": "",
"custom_field_117": "",
"custom_field_118": "",
"custom_field_119": "",
"custom_field_12": "",
"custom_field_120": "",
"custom_field_121": "",
"custom_field_122": "",
"custom_field_123": "",
"custom_field_124": "",
"custom_field_125": "",
"custom_field_126": "",
"custom_field_127": "",
"custom_field_128": "",
"custom_field_129": "",
"custom_field_13": "",
"custom_field_130": "",
"custom_field_131": "",
"custom_field_132": "",
"custom_field_133": "",
"custom_field_134": "",
"custom_field_135": "",
"custom_field_136": "",
"custom_field_137": "",
"custom_field_138": "",
"custom_field_139": "",
"custom_field_14": "",
"custom_field_140": "",
"custom_field_141": "",
"custom_field_142": "",
"custom_field_143": "",
"custom_field_144": "",
"custom_field_145": "",
"custom_field_146": "",
"custom_field_147": "",
"custom_field_148": "",
"custom_field_149": "",
"custom_field_15": "",
"custom_field_150": "",
"custom_field_151": "",
"custom_field_152": "",
"custom_field_153": "",
"custom_field_154": "",
"custom_field_155": "",
"custom_field_156": "",
"custom_field_157": "",
"custom_field_158": "",
"custom_field_159": "",
"custom_field_16": "",
"custom_field_160": "",
"custom_field_161": "",
"custom_field_162": "",
"custom_field_163": "",
"custom_field_164": "",
"custom_field_165": "",
"custom_field_166": "",
"custom_field_167": "",
"custom_field_168": "",
"custom_field_169": "",
"custom_field_17": "",
"custom_field_170": "",
"custom_field_171": "",
"custom_field_172": "",
"custom_field_173": "",
"custom_field_174": "",
"custom_field_175": "",
"custom_field_176": "",
"custom_field_177": "",
"custom_field_178": "",
"custom_field_179": "",
"custom_field_18": "",
"custom_field_180": "",
"custom_field_181": "",
"custom_field_182": "",
"custom_field_183": "",
"custom_field_184": "",
"custom_field_185": "",
"custom_field_186": "",
"custom_field_187": "",
"custom_field_188": "",
"custom_field_189": "",
"custom_field_19": "",
"custom_field_190": "",
"custom_field_191": "",
"custom_field_192": "",
"custom_field_193": "",
"custom_field_194": "",
"custom_field_195": "",
"custom_field_196": "",
"custom_field_197": "",
"custom_field_198": "",
"custom_field_199": "",
"custom_field_20": "",
"custom_field_200": "",
"custom_field_21": "",
"custom_field_22": "",
"custom_field_23": "",
"custom_field_24": "",
"custom_field_25": "",
"custom_field_26": "",
"custom_field_27": "",
"custom_field_28": "",
"custom_field_29": "",
"custom_field_30": "",
"custom_field_31": "",
"custom_field_32": "",
"custom_field_33": "",
"custom_field_34": "",
"custom_field_35": "",
"custom_field_36": "",
"custom_field_37": "",
"custom_field_38": "",
"custom_field_39": "",
"custom_field_40": "",
"custom_field_41": "",
"custom_field_42": "",
"custom_field_43": "",
"custom_field_44": "",
"custom_field_45": "",
"custom_field_46": "",
"custom_field_47": "",
"custom_field_48": "",
"custom_field_49": "",
"custom_field_50": "",
"custom_field_51": "",
"custom_field_52": "",
"custom_field_53": "",
"custom_field_54": "",
"custom_field_55": "",
"custom_field_56": "",
"custom_field_57": "",
"custom_field_58": "",
"custom_field_59": "",
"custom_field_60": "",
"custom_field_61": "",
"custom_field_62": "",
"custom_field_63": "",
"custom_field_64": "",
"custom_field_65": "",
"custom_field_66": "",
"custom_field_67": "",
"custom_field_68": "",
"custom_field_69": "",
"custom_field_70": "",
"custom_field_71": "",
"custom_field_72": "",
"custom_field_73": "",
"custom_field_74": "",
"custom_field_75": "",
"custom_field_76": "",
"custom_field_77": "",
"custom_field_78": "",
"custom_field_79": "",
"custom_field_80": "",
"custom_field_81": "",
"custom_field_82": "",
"custom_field_83": "",
"custom_field_84": "",
"custom_field_85": "",
"custom_field_86": "",
"custom_field_87": "",
"custom_field_88": "",
"custom_field_89": "",
"custom_field_9": "",
"custom_field_90": "",
"custom_field_91": "",
"custom_field_92": "",
"custom_field_93": "",
"custom_field_94": "",
"custom_field_95": "",
"custom_field_96": "",
"custom_field_97": "",
"custom_field_98": "",
"custom_field_99": "",
"custom_field_eight": "",
"custom_field_five": "",
"custom_field_four": "",
"custom_field_one": "",
"custom_field_seven": "",
"custom_field_six": "",
"custom_field_three": "",
"custom_field_two": "",
"custom_plan_field_1": "1166136271001000361",
"custom_plan_field_10": "0",
"custom_plan_field_2": "0",
"custom_plan_field_3": "0",
"custom_plan_field_4": "0",
"custom_plan_field_5": "0",
"custom_plan_field_6": "0",
"custom_plan_field_7": "0",
"custom_plan_field_8": "0",
"custom_plan_field_9": "0"
},
"Description": "",
"Effort": "3",
"EndDate": "2025-03-09",
"Exceed": "0",
"Id": "issueid",
"IterationId": "0",
"Modified": "2025-05-29 20:57:09",
"ParentId": "0",
"Priority": {
"Label": "High",
"Value": "High"
},
"ReleaseID": "0",
"Remain": "3",
"StartDate": "2025-03-08",
"Status": {
"Label": "待处理",
"Value": "UNKNOWN"
},
"Summary": "测试总结",
"TemplatedId": "模版id",
"WorkitemTypeId": "工作项id",
"WorkspaceId": "WorkspaceId"
},
"IssueName": "",
"IsDeleted": false,
"Uid": ""
},
"Cases": [],
"Folders": [],
"ExecCaseCount": 0
}],
"RepoVersionGroup": [{
"RepoVersion": {
"Audit": {
"Creator": "Creator",
"Updater": "",
"CreatedAt": "2025-03-10T14:48:01+08:00",
"UpdatedAt": "2025-03-10T14:48:01+08:00",
"Tenant": "Tenant"
},
"Uid": "Uid",
"Repo": {
"Audit": null,
"Namespace": "",
"Uid": "Uid",
"Name": "默认用例库",
"Description": "",
"CodeRepo": null,
"RefAutomationRepos": [],
"Versions": [],
"AssociatedResources": [],
"Nid": "",
"Type": "DEFALT"
},
"ParentVersionUid": "0",
"Name": "master",
"Description": "稳定版本(默认)",
"SourceCount": 0,
"TargetCount": 0,
"TargetFolderCount": 0,
"TimeLeft": 0,
"Status": "COMPLETED",
"FolderOnly": false,
"CaseUids": [],
"CaseCount": 0,
"Nid": ""
},
"Folders": [{
"Uid": "Uid",
"RepoUid": "RepoUid",
"RepoVersionUid": "RepoVersionUid",
"FolderUid": "",
"FullPath": "",
"Name": "Name",
"Owners": [],
"Description": "",
"CaseCount": 1,
"Folders": [{
"Uid": "Uid",
"RepoUid": "RepoUid",
"RepoVersionUid": "RepoVersionUid",
"FolderUid": "FolderUid",
"FullPath": "",
"Name": "test2",
"Owners": [],
"Description": "",
"CaseCount": 2,
"Folders": [],
"Cases": [{
"Audit": {
"Creator": "Creator",
"Updater": "Updater",
"CreatedAt": "2025-07-15T17:04:31+08:00",
"UpdatedAt": "2025-07-15T17:06:50+08:00",
"Tenant": "Tenant"
},
"Uid": "Uid",
"RepoUid": "RepoUid",
"RepoVersionUid": "RepoVersionUid",
"FolderUid": "FolderUid",
"FullPath": "",
"UUID": "af575027-67bb-439b-a07f-b5eda8acd91b",
"Name": "2-test1",
"Description": "",
"Priority": "P1",
"PreConditions": "",
"Type": "",
"StepType": "STEP",
"Steps": [],
"StepText": null,
"Attachments": [],
"CustomFields": [],
"Labels": [],
"Source": "TESTX",
"ReviewAt": "",
"ReviewState": "DRAFT",
"AutomatedState": "",
"IsManualRelation": false,
"AutomationCases": [],
"Issues": [],
"RelatedSuiteUids": [],
"RelatedSetUids": [],
"Owners": [],
"ManHourEstimated": "",
"Nid": "",
"Path": ".path1.path2.",
"BugCount": "0",
"Bugs": [],
"RunTimes": "0",
"Executions": [],
"Reviews": [],
"CaseBug": null,
"FolderExtensions": null,
"IsFolder": false,
"SetUids": [],
"IsBaseline": false,
"ExecutionPlatforms": ""
}],
"UUID": "aa8847b2-1240-480c-ae98-2eade3ce69b3",
"Nid": "",
"Path": ".path1.",
"ArchiveUid": "",
"Audit": {
"Creator": "Creator",
"Updater": "Updater",
"CreatedAt": "2025-07-15T17:04:20+08:00",
"UpdatedAt": "2025-07-15T17:04:20+08:00",
"Tenant": "Tenant"
},
"Issues": [],
"CopyCaseUids": [],
"IncludeAncestors": false,
"FolderExtensions": {
"FolderType": "GENERAL",
"AllowActions": [],
"CanImport": false,
"DisplayMonitor": false,
"MergingCount": 0
},
"Reviews": [],
"IsBaseline": false
}]
}]
}]
}
},
"TotalCount": 1
}
```

# 测试计划用例字段说明

RepoVersion字段[参考](/document/api-doc/API文档/api_reference/testx/case/create_repo.html)
Folders字段[参考](/document/api-doc/API文档/api_reference/testx/case/search_case.html)