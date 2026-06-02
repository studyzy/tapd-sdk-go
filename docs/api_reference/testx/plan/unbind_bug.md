# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/unbind_bug.html

移除计划用例关联的缺陷

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/{plan_uid}/cases/{case_uid}/issues/{issue_uid}`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

DELETE

# 请求数限制

一次移除一条

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间/项目唯一标识 |
| PlanUid | 是 | string | 计划UID |
| CaseUid | 是 | string | 用例UID |
| IssueUid | 是 | string | 问题UID(TAPD ID) |

# 返回结果


```json
{
"RequestId": "",
"Error": null
}
```