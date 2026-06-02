# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/report/list_report.html

获取报告列表

# url

`https://api.tapd.cn/api/testx/report/v1/namespaces/{namespace}/reports`

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
| Search | 否 | string | 搜索关键词 |
| StartAt | 否 | string | 开始时间 |
| EndAt | 否 | string | 结束时间 |
| Creators | 否 | string of list | 创建者列表 |
| PlanUids | 否 | string of list | 计划唯一标识列表 |
| WithAssociated | 否 | bool | 是否需要查询多计划生成报告 |
| TemplateUid | 否 | string | 模板唯一标识 |
| Source | 否 | ReportSource | 报告创建来源（已弃用，会优先使用 Sources） |
| Sources | 否 | ReportSource of list | 报告创建来源 |

## ReportSource


| 参数名称 | 含义 |
| --- | --- |
| REPORTSOURCE_UNKNOW | 未知来源 |
| REPORTSOURCE_CREATE | 新建生成 |
| REPORTSOURCE_PLAN | 计划生成 |
| REPORTSOURCE_PERIOD | 定期报告 |

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
"Audit": {
"Creator": "734242230",
"Updater": "734242230",
"CreatedAt": "2025-07-02T21:40:13+08:00",
"UpdatedAt": "2025-07-02T21:40:13+08:00",
"Tenant": ""
},
"Uid": "report_uid",
"Title": "test",
"Namespace": "66136271",
"Summary": "",
"Plans": [
{
"Meta": {
"Uid": "plan_uid",
"Namespace": "",
"Audit": null,
"FolderUid": "",
"Name": "test",
"Description": "",
"DataSource": null,
"State": "NONE",
"Testers": [],
"FolderPath": "",
"Version": "V1",
"CasePath": null,
"Nid": "",
"Path": ""
},
"Spec": null
}
],
"Source": "REPORTSOURCE_PLAN",
"TemplateUid": "0",

"Tasks": [],
"Nid": "",
"Stat": null
}
],
"TotalCount": 1
}
```

# 测试用例仓库字段说明

## Report


| 参数名称 | 类型 | 含义 |
| --- | --- | --- |
| Uid | string | 报告唯一标识 |
| Title | string | 报告标题 |
| Namespace | string | 命名空间 |
| Summary | string | 报告摘要 |
| Plans | Plan of list | 计划列表 |
| Source | ReportSource | 报告来源 |
| TemplateUid | string | 模板唯一标识 |
| NotificationUid | string | 通知模版唯一标识 |
| Tasks | PeriodicTask of list | 定期任务列表 |
| Nid | string | 序号 |

## ReportSource


| 参数名称 | 含义 |
| --- | --- |
| REPORTSOURCE_UNKNOW | 未知来源 |
| REPORTSOURCE_CREATE | 新建生成 |
| REPORTSOURCE_PLAN | 计划生成 |
| REPORTSOURCE_PERIOD | 定期报告 |

## PeriodicTask


| 参数名称 | 类型 | 含义 |
| --- | --- | --- |
| Id | string | 定期任务 ID |
| StartedAt | string | 开始时间 |
| ExpiredAt | string | 截止时间 |
| Enabled | bool | 是否启用 |
| Cron | Cron | 定时任务表达式 |
| Namespace | string | 命名空间 |

## Cron


| 参数名称 | 类型 | 含义 |
| --- | --- | --- |
| Period | Period | 定时周期 |
| MonthlyDays | uint32 of list | 每月几号 |
| WeeklyDays | WeeklyDay of list | 周几 |
| TimeAt | string | 时:分，如 "13:30" |

## Cron.Period


| 参数名称 | 含义 |
| --- | --- |
| DAILY | 每天 |
| WORKDAY | 工作日 |
| WEEKLY | 每周 |
| BIWEEKLY | 每两周 |
| MONTHLY | 每月 |

## Cron.WeeklyDay


| 取值 | 含义 |
| --- | --- |
| SUNDAY | 周日 |
| MONDAY | 周一 |
| TUESDAY | 周二 |
| WEDNESDAY | 周三 |
| THURSDAY | 周四 |
| FRIDAY | 周五 |
| SATURDAY | 周六 |