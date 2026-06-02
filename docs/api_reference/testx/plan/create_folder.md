# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/create_folder.html

创建计划目录，返回计划目录

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/folders`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间 |
| ParentUid | 是 | string | 父目录ID |
| Name | 是 | string | 接口库标识名 |
| Description | 否 | string | 描述信息 |

# 返回结果


```json
{
"Error": null,
"Data": {
"Uid": "xxx",
"Namespace": "xxx",
"Audit": {
"Creator": "xxx",
"Updater": "xxx",
"CreatedAt": "2025-07-14T14:10:58+08:00",
"UpdatedAt": "2025-07-14T14:10:58+08:00",
"Tenant": ""
},
"ParentUid": "0",
"Name": "test",
"Description": "test desc",
"PlanCount": 0,
"ArchiveAuto": false,
"Folders": [],
"Plans": [],
"Path": "."
}
}
```

# 测试计划目录重要字段说明

## Folder


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| Uid | string | 目录唯一标识 |
| Namespace | string | 命名空间 |
| Audit | testx.common.Audit | 审计信息 |
| ParentUid | string | 父目录ID |
| Name | string | 接口库标识名 |
| Description | string | 描述信息 |
| PlanCount | uint32 | 计划数量 |
| ArchiveAuto | bool | 已完成计划自动归档 |
| Folders | repeated Folder | 子目录基本信息 |
| Plans | repeated PlanMeta | 计划基本信息 |
| Path | string | 路径 |