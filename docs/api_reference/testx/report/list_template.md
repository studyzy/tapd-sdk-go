# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/report/list_template.html

获取报告模版列表

# url

`https://api.tapd.cn/api/testx/report/v1/namespaces/{namespace}/templates`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间 |
| PageInfo | 否 | PageInfo | 分页信息 |

## pageinfo


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Offset | 是 | uint32 | 分页偏移位置 |
| Limit | 是 | uint32 | 每页数量 |

# 返回结果


```json
{
"RequestId": "",
"Error": null,
"Data": [
{
"Audit": null,
"Uid": "1742",
"Title": "系统报告模板",
"Namespace": "66136271",
"IsSystem": true,
"Desc": "系统报告模板",
"AsDefault": true,
"Order": 0,
"NotificationUid": "",
"NotificationName": "",
"Modules": [],
"Count": 17
}
],
"TotalCount": 1
}
```

# 测试报告模版字段说明

## ReportTemplate


| 参数名称 | 类型 | 含义 |
| --- | --- | --- |
| Uid | string | 模板唯一标识 |
| Title | string | 模板标题 |
| Namespace | string | 命名空间 |
| IsSystem | bool | 是否为系统模板 |
| Desc | string | 描述信息 |
| AsDefault | bool | 是否为默认模板 |
| Order | uint32 | 排序顺序 |
| NotificationUid | string | 通知唯一标识 |
| NotificationName | string | 通知名称 |
| Count | uint32 | 模块数量 |