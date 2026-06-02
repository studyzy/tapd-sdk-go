# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/search_case.html

查询用例列表，获取用例

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{version_uid}/cases/search`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求数限制

每页最大200条

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| PageInfo | 否 | PageInfo | 用例列表搜索的分页信息 |
| Namespace | 是 | string | 搜索范围-命名空间 |
| RepoUid | 是 | string | 搜索范围-仓库 UID |
| RepoVersionUid | 是 | string | 搜索范围-仓库版本 UID |
| FolderUid | 否 | string | 搜索范围-目录 UID |
| CaseUids | 否 | repeated string | 搜索范围-用例 UID 列表 |
| ExcludeCaseUids | 否 | repeated string | 排除的用例 UID 列表 |
| Filter | 是 | Filter | 搜索条件 |
| ShowMode | 否 | ShowMode | 返回结果的显示模式 |
| ExtendFields | 否 | repeated string | 需要在基础用例字段上返回的关联字段列表 |
| IncludeDescendants | 否 | bool | 是否返回所有子孙用例 |
| IncludeAncestors | 否 | bool | 是否返回父目录 |
| SelectFields | 否 | repeated string | 需要返回的字段列表 |

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

## ShowMode


| 枚举值 | 说明 |
| --- | --- |
| FLAT | 平铺展示 |
| TREE | 树形展示 |

## pageinfo


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Offset | 是 | uint32 | 分页偏移位置 |
| Limit | 是 | uint32 | 每页数量 |

# 返回结果


```json
{
"Error": null,
"Data": {
"Folders": [
{
"Uid": "12571827",
"RepoUid": "17099",
"RepoVersionUid": "18167",
"FolderUid": "",
"FullPath": "",
"Name": "test",
"Owners": [],
"Description": "",
"CaseCount": 23,
"Folders": [],
"Cases": [],
"UUID": "cec41f49-7f5a-462a-a76c-d762e8632e03",
"Path": ".",
"ArchiveUid": "",
"Audit": {
"Creator": "734242230",
"Updater": "734242230",
"CreatedAt": "2025-03-19T11:12:23+08:00",
"UpdatedAt": "2025-03-19T11:12:23+08:00",
"Tenant": "43300792"
}
}
],
"Cases": [
{
"Audit": {
"Creator": "734242230",
"Updater": "734242230",
"CreatedAt": "2025-03-19T11:12:48+08:00",
"UpdatedAt": "2025-07-12T14:44:28+08:00",
"Tenant": "43300792"
},
"Uid": "12571845",
"RepoUid": "17099",
"RepoVersionUid": "18167",
"FolderUid": "12571834",
"FullPath": "test/55/示例测试用例/注册与登录/登录",
"UUID": "31332127-75df-4cf9-980e-568463240eb6",
"Name": "验证登录时记住用户名功能",
"Description": "验证在登录时勾选‘记住用户名’后，退出登录并重新进入登录页面时，用户名是否自动填充为之前记住的用户名。",
"Priority": "P0",
"PreConditions": "1. 进入平台首页\n2. 已成功注册账户，用户名为testxadmin，密码为admintestx",
"Type": "44907",
"StepType": "STEP",
"Steps": [],
"StepText": null,
"Attachments": [],
"CustomFields": [],
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
"ReviewAt": "",
"ReviewState": "DRAFT",
"Issues": [],
"Owners": [],
"ManHourEstimated": "",
"Nid": "",
"Path": ".12571827.12571828.12571829.12571830.12571834.",
"BugCount": "0",
"Bugs": [],
"RunTimes": "0",
"Executions": [],
"Reviews": [],
"CaseBug": null,
"IsFolder": false,
},           
],
"Repos": [
{
"Audit": {
"Creator": "xx",
"Updater": "",
"CreatedAt": "2025-07-11T11:12:20+08:00",
"UpdatedAt": "2025-07-11T11:12:20+08:00",
"Tenant": "xx"
},
"Namespace": "xx",
"Uid": "17131",
"Name": "默认用例库",
"Description": "",
"Versions": [
{
"Audit": {
"Creator": "xx",
"Updater": "",
"CreatedAt": "2025-07-11T11:12:20+08:00",
"UpdatedAt": "2025-07-11T11:12:20+08:00",
"Tenant": "xx"
},
"Uid": "18199",
"Name": "master",
"Description": "稳定版本(默认)",
}
],
"Type": "SYSTEM"
}
]
},
"TotalCount": 1
}
```

# 测试用例字段说明

## 测试用例


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