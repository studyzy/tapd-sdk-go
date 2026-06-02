# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/create_repo.html

创建用例仓库，返回用例仓库信息

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 项目 |
| Name | 是 | string | 用例库名称 |
| Description | 否 | string | 用例库描述 |

# 入参示例


```json
{
"Name":"测试创建",
"Description":"测试描述"
}
```

# 返回结果


```json
{
"Error": null,
"Data": {
"Audit": {
"Creator": "xx",
"Updater": "xx",
"CreatedAt": "2025-07-08T17:24:07+08:00",
"UpdatedAt": "2025-07-08T17:24:07+08:00",
"Tenant": "xx"
},
"Namespace": "xxx",
"Uid": "17130",
"Name": "测试用例库",
"Description": "测试",
"Versions": [
{
"Audit": {
"Creator": "xx",
"Updater": "xx",
"CreatedAt": "2025-07-08T17:24:07+08:00",
"UpdatedAt": "2025-07-08T17:24:07+08:00",
"Tenant": "xx"
},
"Uid": "18198",
"Repo": {
"Audit": null,
"Namespace": "",
"Uid": "17130",
"Name": "",
"Description": "",
"Versions": [],
"Nid": "",
"Type": "DEFALT"
},
"Name": "master",
"Description": "稳定版本(默认)",
}
],
"Type": "DEFALT"
}
}
```

# 测试用例仓库字段说明

## 测试用例仓库


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