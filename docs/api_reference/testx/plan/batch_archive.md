# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/plan/batch_archive.html

批量归档计划

# url

`https://api.tapd.cn/api/testx/plan/v1/namespaces/{namespace}/plans/batch-archive`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

PUT

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间 |
| Uids | 是 | repeated string | 要归档的计划UID列表 |
| ArchiveMode | 是 | ArchiveMode | 归档操作模式 |

## ArchiveMode


| 枚举值 | 说明 |
| --- | --- |
| ARCHIVEMODE_NONE | 获取全部 |
| ARCHIVEMODE_YES | 获取归档 |
| ARCHIVEMODE_NO | 获取非归档 |

# 入参示例


```json
{
"Uids": ["152323", "152322", "152321", "152319", "152318"],
"ArchiveMode": "ARCHIVEMODE_YES"
}
```

# 返回结果


```json
{
"RequestId": "",
"Error": null
}
```