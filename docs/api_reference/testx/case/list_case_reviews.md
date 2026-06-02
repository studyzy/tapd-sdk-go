# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/list_case_reviews.html

获取用例评审记录

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{repo_version_uid}/cases/{case_uid}/reviews`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求数限制

一次获取一批数据，不传默认获取全部

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| PageInfo | 否 | Pageinfo | 分页信息 |
| Namespace | 是 | string | 命名空间/项目唯一标识 |
| CaseUid | 是 | string | 用例UID |
| RepoUid | 是 | string | 仓库UID |
| RepoVersionUid | 是 | string | 仓库版本UID |
| Source | 否 | string | 来源(待改为枚举) |
| MainUid | 否 | string | 来源主体资源的UID |
| SourceKind | 否 | string | 来源主体资源的具体资源类型 |
| SourceUid | 否 | string | 来源主体资源的具体资源UID |
| MainUids | 否 | repeated string | 同时查询计划和设计下的评审信息 |
| IsLastReview | 否 | bool | 仅展示最后一次评审 |
| CaseUids | 否 | repeated string | 用例UID列表 |
| SourceUids | 否 | repeated string | 来源UID列表 |

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
"SourceName": "测试复现",
"SourceUid": "xx.xxx",
"Reviewer": "734242230",
"ReviewTime": "2025-07-12 16:28:06",
"ReviewState": "REVIEW_RESULT_AGREE",
"Message": "11",
"LinkData": {
"DesignUid": "xxx"
},
"Source": "DESIGN",
"Uid": "xxx",
"MainUid": "xxx",
"SourceKind": "CASE",
"Total": 0,
"CaseUid": "xxx"
}
],
"TotalCount": 1
}
```

# 测试用例评审字段说明

## CaseReview


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| SourceName | string | 来源名称 |
| SourceUid | string | 来源主体资源的具体资源UID |
| Reviewer | string | 评审人 |
| ReviewTime | string | 评审时间 |
| ReviewState | string | 评审状态 |
| Message | string | 评审消息 |
| LinkData | google.protobuf.Struct | 关联数据 |
| Source | string | 来源(建议改为枚举) |
| Uid | string | 唯一标识 |
| MainUid | string | 来源主体资源的UID |
| SourceKind | string | 来源主体资源的具体资源类型 |
| Total | uint32 | 当前来源产生的评审次数 |
| CaseUid | string | 实际用例UID(可能为空) |