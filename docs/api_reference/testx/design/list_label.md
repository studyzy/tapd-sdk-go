# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/design/list_label.html

查询测试设计标签列表

# url

`https://api.tapd.cn/api/testx/design/v2/namespaces/{namespace}/labels`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求数限制

根据uid获取

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| namespace | 是 | string | 项目标识 |
| design_uid | 否 | string | 测试设计标识 |
| kind | 否 | Kind | 标签所在的节点类型 |
| name | 否 | string | 标签名称(支持模糊匹配) |

## Kind


| 取值 | 说明 |
| --- | --- |
| NONE | 未知节点 |
| CUSTOM | 自由节点 |
| GROUP | 分组 |
| DESIGN | 测试设计 |
| STORY | 需求 |
| FEATURE | 功能点 |
| SCENE | 场景 |
| TEST_POINT | 测试点 |
| CASE | 用例 |
| CASE_PRE_CONDITION | 用例前置条件节点 |
| CASE_STEP | 用例步骤 |
| CASE_EXPECTION_RESULT | 用例预期结果 |
| CASE_GROUP | 目录类型节点 |
| BUG | 缺陷 |

# 返回结果


```json
{
"Error": null,
"Data": [
{
"Name": "label_name",
"Value": "",
"Tag": "",
"Color": "#EE7C89",
"Uneditable": false,
"DisplayName": "test1",
"Source": "MANUAL"
},
{
"Name": "label_name",
"Value": "",
"Tag": "",
"Color": "#fe6fd4",
"Uneditable": false,
"DisplayName": "test2",
"Source": "DATASET"
}
],
"TotalCount": 0
}
```

# 测试设计标签字段说明

## Label


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| name | string | 标签名称(在某一范围是唯一，如namespace下) |
| value | string | 标签值 |
| tag | string | 扩展标识(如case里面关联的key) |
| color | string | 标签颜色 |
| uneditable | bool | 是否不可编辑 |
| display_name | string | 前端展示的中文名 |
| source | string | 标签来源 |