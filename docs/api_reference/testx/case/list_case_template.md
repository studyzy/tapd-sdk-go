# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/list_case_template.html

获取用例模版

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/case-templates`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

GET

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 命名空间/项目唯一标识 |

# 返回结果


```json
{
"Error": null,
"Data": [{
"Uid": "1587",
"Name": "默认模版",
"Fields": [{
"Uid": "24912",
"Field": {
"Widget": "STORY",
"Type": "STR",
"Uneditable": true,
"Name": "ISSUES",
"Label": "关联需求",
"Description": "",
"Help": "",
"Visible": false,
"Required": false,
"Default": null,
"Choices": [],
"Flag": "SYSTEM",
"Order": 0,
"EncryptOption": null,
"ValueSource": null,
"Properties": [],
"IsRichText": false
},
"Default": "",
"VisibleScopes": [],
"Required": false,
"Order": 0,
"Width": 0
}]
}]
}
```

# 测试用例模版字段说明

## CaseTemplate


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Uid | 是 | string | 模板唯一标识 |
| Name | 是 | string | 模板名称 |
| Fields | 是 | repeated CaseTemplateField | 模板字段列表 |

## CaseTemplateField


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Uid | 是 | string | 字段唯一标识(用于判断更新/新增) |
| Field | 是 | PropertyDefinition | 字段定义 |
| Default | 是 | google.protobuf.Value | 默认值 |
| VisibleScopes | 是 | repeated VisibleScope | 可见范围 |
| Required | 是 | bool | 是否必填 |
| Order | 是 | int32 | 排序序号 |
| Width | 是 | int32 | 宽度 |

## VisibleScope


| 枚举值 | 说明 |
| --- | --- |
| NONE | 不可见 |
| LIST | 列表中显示 |
| REVIEW | 评审中显示 |

## PropertyDefinition


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Widget | 是 | Widget | 输入框类型 |
| Type | 是 | Type | 字段类型 |
| Uneditable | 是 | bool | 是否可编辑 |
| Name | 是 | string | 字段名称 |
| Label | 是 | string | 显示名称 |
| Description | 是 | string | 描述 |
| Help | 是 | string | 帮助信息 |
| Visible | 是 | bool | 是否显示 |
| Required | 是 | bool | 是否必填 |
| Default | 是 | google.protobuf.Value | 默认值 |
| Choices | 是 | repeated ValueChoice | 可选值列表 |
| Flag | 是 | Flag | 类型标志 |
| Order | 是 | uint32 | 排序序号 |
| ValueSource | 否 | PropertyValueSource | 值来源 |
| Properties | 否 | repeated PropertyDefinition | 子属性 |
| IsRichText | 否 | bool | 是否富文本 |

## Widget


| 枚举值 | 说明 |
| --- | --- |
| TEXT | 文本输入框 |
| TEXTAREA | 文本域 |
| SELECT | 下拉选择 |
| MULTI_SELECT | 多选下拉 |
| USER | 用户选择 |
| MULTI_USER | 多用户选择 |
| RADIO | 单选按钮 |
| CHECKBOX | 复选框 |
| INTEGER | 整数输入 |
| FLOAT | 浮点数输入 |
| DATE | 日期选择 |
| DATETIME | 日期时间选择 |
| KEY_VALUE | 键值对 |
| MULTI_VALUE | 多值 |
| MULTI_KEY_VALUE | 多键值对 |
| TAPD | TAPD字段 |
| ITERATION | 迭代 |
| STORY | 需求 |
| BUG | 缺陷 |
| VERSION | 版本 |
| SWITCH | 开关 |

## Type


| 枚举值 | 说明 |
| --- | --- |
| STR | 字符串 |
| LONG | 长整型 |
| LIST_STR | 字符串列表 |
| MAP_STR_TO_STR | 字符串映射 |

## Flag


| 枚举值 | 说明 |
| --- | --- |
| NONE | 无标志 |
| SYSTEM | 系统属性 |
| CUSTOM | 自定义属性 |
| EXTRA_EDIT | 额外可编辑属性 |

## ValueChoice


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| ID | 是 | uint32 | 选项ID |
| Label | 是 | string | 显示标签 |
| Value | 是 | google.protobuf.Value | 选项值 |
| Order | 是 | uint32 | 排序序号 |