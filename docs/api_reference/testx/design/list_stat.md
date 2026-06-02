# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/design/list_stat.html

查询测试设计列表统计信息

# url

`https://api.tapd.cn/api/testx/design/v2/namespaces/{namespace}/stat-list`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求数限制

根据uid获取

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| namespace | 是 | string | 命名空间 |
| design_uids | 否 | repeated string | 设计唯一标识列表 |

# 返回结果


```json
{
"Error": null,
"Data": [
{
"DesignUid": "design_uid",
"ExtendFeatures": {
"UsedAI": false,
"Archive": {
"Status": 0,
"Count": 0,
"CaseNum": 1
},
"Review": {
"State": "UNREVIEW",
"ReviewedIssueCount": 0,
"TotalIssueCount": 2,
"ReviewedCaseCount": 0,
"TotalCaseCount": 1,
"AIReviewedCaseCount": 0,
"TotalBugCount": 0
}
}
}
]
}
```

# 测试设计状态字段说明

## DesignExtendFeatures


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| design_uid | string | 设计标识 |
| extend_features | ExtendFeatures | 扩展信息 |

## ExtendFeatures


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| used_ai | bool | 是否使用过AI |
| archive | ArchiveState | 归档数据 |
| review | ReviewState | 评审数据 |

## ArchiveState


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| status | int32 | 状态(0未归档/1已归档) |
| count | int32 | 已归档用例数量 |
| case_num | int32 | 所有用例数 |

## ReviewState


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| state | State | 评审状态 |
| reviewed_issue_count | uint32 | 已评审需求数 |
| total_issue_count | uint32 | 所有需求数 |
| reviewed_case_count | uint32 | 已评审用例数 |
| total_case_count | uint32 | 所有用例数 |
| ai_reviewed_case_count | uint32 | AI已评审用例数 |
| total_bug_count | uint32 | 所有Bug数 |

## State


| 取值 | 说明 |
| --- | --- |
| UNREVIEW | 未评审 |
| PROCESSING | 评审中 |
| FINISHED | 评审完成 |
| PASS | 通过 |