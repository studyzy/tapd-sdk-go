# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/list_case_historys.html

获取用例变更历史

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{repo_version_uid}/cases/{case_uid}/history`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求数限制

一次获取一批数据，不传默认获取20条。

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| PageInfo | 否 | PageInfo | 分页信息 |
| Namespace | 是 | string | 命名空间 |
| RepoUid | 是 | string | 仓库 UID |
| RepoVersionUid | 是 | string | 仓库版本 UID |
| CaseUid | 是 | string | 用例 UID |

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
"Audit": {
"Creator": "xx",
"Updater": "",
"CreatedAt": "2025-07-12T14:44:28+08:00",
"UpdatedAt": "",
"Tenant": ""
},
"Uid": "xxxx",
"ChangeType": "BATCH_UPDATE",
"ChangeContents": [
{
"FieldName": "Type",
"PreValue": "",
"PostValue": "接口测试"
},
{
"FieldName": "Labels",
"PreValue": "",
"PostValue": "[{\"Color\":\"#7ad94e\",\"DisplayName\":\"121\",\"Name\":\"34456\"}]"
}
]
},
{
"Audit": {
"Creator": "xxx",
"Updater": "",
"CreatedAt": "2025-03-19T11:12:48+08:00",
"UpdatedAt": "",
"Tenant": ""
},
"Uid": "xxx",
"ChangeType": "CREATE",
"ChangeContents": [
{
"FieldName": "Name",
"PreValue": "",
"PostValue": "验证登录时记住用户名功能"
},
{
"FieldName": "FolderUid",
"PreValue": "",
"PostValue": "{\"Name\":\"登录\",\"Uid\":\"12571834\"}"
},
{
"FieldName": "Description",
"PreValue": "",
"PostValue": "验证在登录时勾选‘记住用户名’后，退出登录并重新进入登录页面时，用户名是否自动填充为之前记住的用户名。"
},
{
"FieldName": "PreConditions",
"PreValue": "",
"PostValue": "1. 进入平台首页\n2. 已成功注册账户，用户名为testxadmin，密码为admintestx"
},
{
"FieldName": "Priority",
"PreValue": "",
"PostValue": "P0"
}
]
}
],
"TotalCount": 2
}
```

# 测试用例历史字段说明

## CaseHistory


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Audit | common.Audit | 审计信息 |
| Uid | string | 测试用例 UID |
| ChangeType | ChangeType | 变更类型 |
| ChangeContents | repeated ChangeContent | 变更内容列表 |

## ChangeContent


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| FieldName | string | 变更字段名称 |
| PreValue | string | 变更前的值 |
| PostValue | string | 变更后的值 |

## ChangeType


| 枚举值 | 说明 |
| --- | --- |
| CREATE | 创建 |
| MANUAL_UPDATE | 手动变更 |
| BATCH_UPDATE | 批量变更 |
| CODE_UPDATE | 代码触发 |