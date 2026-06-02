# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/batch_unbind_bug.html

批量解绑bug

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{repo_version_uid}/cases/{case_uid}/bugs/batch-unbind`

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
| BugUids | 是 | repeated string | 绑定的bug ids |

# 入参示例


```json
{
"BugUids": ["bug_uid"]
}
```

# 返回结果


```json
{
"RequestId": "",
"Error": null
}
```