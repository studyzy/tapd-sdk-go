# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/batch_update_case.html

批量更新用例，返回用例

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{version_uid}/cases/batch-update`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求数限制

编辑用例属性不能超过5个

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 项目命名空间 |
| RepoUid | 是 | string | 仓库唯一标识 |
| RepoVersionUid | 是 | string | 仓库版本唯一标识 |
| CaseUids | 是 | repeated string | 用例唯一标识列表 |
| UpdateInfos | 是 | repeated UpdateInfo | 更新信息列表 |

## UpdateInfo


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| FieldName | 是 | string | 字段名称(用例对象的字段名即可) |
| FieldValue | 是 | google.protobuf.Value | 字段值 |
| Action | 是 | Action | 操作类型 |

## Action


| 取值 | 字面值 |
| --- | --- |
| UNKNOWN | 未知操作 |
| APPEND | 附加 |
| OVERRIDE | 覆盖 |
| DELETE | 删除 |

# 入参示例


```json
{
"CaseUids": [
"12571828",
"xxxxx",
"xxxxx"
],
"UpdateInfos": [
{
"FieldValue": "44907",
"FieldName": "Type",
"Action": "OVERRIDE"
},
{
"FieldValue": "34456",
"FieldName": "Labels",
"Action": "APPEND"
}
]
}
```

# 返回结果


```json
{
"RequestId": "",
"Error": null
}
```

# 测试用例字段说明

## FieldName

取值可参考页面批量编辑


| 取值 | 含义 |
| --- | --- |
| Priority | 优先级 |
| PreConditions | 前置条件 |
| Owners | 负责人 |
| ... | ... |