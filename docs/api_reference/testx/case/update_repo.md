# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/update_repo.html

更新用例仓库，返回用例仓库信息

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

PUT

# 请求数限制

一次更新一条数据

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 项目 |
| Name | 是 | string | 用例库名称 |
| Description | 否 | string | 用例库描述 |

# 返回结果


```json
{
"Error": null,
"Data": {
"Audit": {
"Creator": "xx",
"Updater": "xx",
"CreatedAt": "",
"UpdatedAt": "",
"Tenant": "xxx"
},
"Namespace": "xxx",
"Uid": "17114",
"Name": "xxx",
"Description": "x",
"Type": "DEFALT"
}
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
| type | 用例库类型 |

## 测试用例仓库(type)取值字段说明


| 取值 | 字面值 |
| --- | --- |
| DEFALT | 缺省 |
| SYSTEM | 系统默认 |