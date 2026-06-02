# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/batch_bind_bug.html

批量关联bug

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{repo_version_uid}/cases/{case_uid}/bugs/batch-bind`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间 |
| RepoUid | 是 | string | 仓库UID |
| RepoVersionUid | 是 | string | 仓库版本UID |
| CaseUid | 是 | string | 用例UID |
| BindBugs | 是 | repeated Bug | 绑定的bug信息列表 |

## bug


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| IssueUid | 是 | string | bug唯一标识 |
| WorkspaceUid | 是 | string | 工作空间唯一标识 |

# 入参示例


```json
{
"BindBugs": [{
"IssueUid": "bug_uid",
"WorkspaceUid": "workspace_Uid"
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