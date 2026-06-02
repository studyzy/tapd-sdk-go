# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/list_case_executions.html

获取用例执行记录

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{repo_version_uid}/cases/{case_uid}/executions`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求数限制

一次获取一批数据，不传默认获取10条。

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间/项目唯一标识 |
| CaseUid | 是 | string | 用例UID |
| PageInfo | 否 | Pageinfo | 分页信息 |
| Ordering | 否 | string | 排序字段 |
| RepoUid | 是 | string | 仓库UID |
| RepoVersionUid | 是 | string | 仓库版本UID |

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
"SourceName": "测试",
"SourceUid": "case-xx",
"Executor": "xx",
"StartExecuteTime": "",
"EndExecuteTime": "2025-07-01 17:02:59",
"ExecuteState": "FAIL",
"Message": "",
"BugCount": "0",
"LinkData": {
"DesignUid": "xxxx"
},
"Source": "DESIGN"
}
],
"TotalCount": 1
}
```

# 测试用例执行记录字段说明

## CaseExecution


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| SourceName | string | 执行源名称 |
| SourceUid | string | 执行源UID |
| Executor | string | 执行人 |
| StartExecuteTime | string | 开始执行时间 |
| EndExecuteTime | string | 结束执行时间 |
| ExecuteState | string | 执行状态 |
| Message | string | 执行消息 |
| BugCount | uint64 | 缺陷数量 |
| LinkData | google.protobuf.Struct | 关联数据 |
| Source | string | 执行来源 |