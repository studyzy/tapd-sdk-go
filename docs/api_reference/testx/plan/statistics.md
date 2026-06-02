# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/statistics.html

获取计划统计信息

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/statistics`

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
| Uids | 是 | repeated string | 计划ID列表 |
| Namespace | 是 | string | 命名空间 |
| PageInfo | 否 | PageInfo | 分页信息(缺省获取全部) |

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
"Uid": "xx",
"Statistic": {
"SucceedCaseCount": 1,
"FailedCaseCount": 1,
"BlockCaseCount": 0,
"ErrorCaseCount": 0,
"RetryCaseCount": 0,
"IgnoreCaseCount": 0,
"TodoCaseCount": 2,
"ManualSucceedRate": "25%",
"TestedCaseCount": 2,
"TotalCaseCount": 4,
"CaseSucceedRate": "25%",
"StoryPassedRate": "0%",
"FailedTaskCount": 1,
"CaseCoverageRate": "0%",
"StoryCount": 0
},
"CaseRepos": [
{
"Audit": {
"Creator": "xx",
"Updater": "xxx",
"CreatedAt": "2025-03-17T17:22:00+08:00",
"UpdatedAt": "2025-07-12T09:44:53+08:00",
"Tenant": "xx"
},
"Namespace": "xxx",
"Uid": "17114",
"Name": "32日2人",
"Description": "csacaca",
"Versions": [
{
"Audit": {
"Creator": "xx",
"Updater": "xxx",
"CreatedAt": "2025-03-17T17:22:00+08:00",
"UpdatedAt": "2025-03-17T17:22:00+08:00",
"Tenant": "xx"
},
"Uid": "18182",
"Repo": {
"Audit": null,
"Namespace": "",
"Uid": "17114",
"Name": "",
"Description": "",
"Versions": [],
"Type": "DEFALT"
},
"Name": "master",
"Description": "稳定版本(默认)",
}
],
"Type": "DEFALT"
}
]
}
],
"TotalCount": 1
}
```

# 测试计划字段说明

## 测试计划用例统计重要字段说明


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| SucceedCaseCount | uint32 | 成功用例数 |
| FailedCaseCount | uint32 | 失败用例数 |
| BlockCaseCount | uint32 | 阻塞用例数 |
| ErrorCaseCount | uint32 | 错误用例数 |
| RetryCaseCount | uint32 | 重试用例数 |
| IgnoreCaseCount | uint32 | 忽略用例数 |
| TodoCaseCount | uint32 | 未测试用例数 |
| ManualSucceedRate | string | 手工用例成功率 |
| TestedCaseCount | uint32 | 已测试用例数 |
| TotalCaseCount | uint32 | 用例总数 |
| CaseSucceedRate | string | 用例成功率 |
| StoryPassedRate | string | 需求通过率 |
| FailedTaskCount | uint32 | 未通过任务数 |
| CaseCoverageRate | string | 用例覆盖率 |
| StoryCount | uint32 | 功能需求数 |

## 测试用例仓库重要字段说明


| 字段 | 说明 |
| --- | --- |
| Namespace | 项目ID |
| Uid | 仓库uid |
| Name | 仓库名称 |
| Description | 仓库描述 |
| Versions | 用例库版本信息 |
| type | 用例类型 |

## 测试用例仓库版本字段说明


| 取值 | 字面值 |
| --- | --- |
| Uid | 用例库版本 uid |
| Repo | 所属用例库 |
| Name | 用例库版本名称 |
| Description | 用例库版本描述 |

## 测试用例仓库(type)取值字段说明


| 取值 | 字面值 |
| --- | --- |
| DEFALT | 缺省 |
| SYSTEM | 系统默认 |