# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/list_repo.html

获取用例仓库列表，返回用例仓库列表信息

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求数限制

一次获取一批数据，不传默认一次获取20条

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 项目 |
| PageInfo | 是 | PageInfo PageInfo | 分页信息 |
| Search | 否 | string | 搜索关键字 |
| Uids | 否 | repeated string | 用户ID列表 |
| Reverse | 否 | bool | 是否根据id倒排序 |

## PageInfo


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
"CreatedAt": "2025-07-11T11:12:20+08:00",
"UpdatedAt": "2025-07-11T11:12:20+08:00",
"Tenant": "xx"
},
"Namespace": "xx",
"Uid": "17131",
"Name": "默认用例库",
"Description": "",
"Versions": [
{
"Audit": {
"Creator": "xx",
"Updater": "",
"CreatedAt": "2025-07-11T11:12:20+08:00",
"UpdatedAt": "2025-07-11T11:12:20+08:00",
"Tenant": "xx"
},
"Uid": "18199",
"Name": "master",
"Description": "稳定版本(默认)",
}
],
"Type": "SYSTEM"
}
],
"TotalCount": 1
}
```

# 测试用例仓库字段说明

## 测试用例仓库重要字段说明


| 字段 | 说明 |
| --- | --- |
| Namespace | 项目ID |
| Uid | 仓库uid |
| Name | 仓库名称 |
| Description | 仓库描述 |
| Versions | 用例库版本信息 |
| type | 用例库类型 |

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