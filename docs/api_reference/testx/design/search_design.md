# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/design/search_design.html

查询测试设计列表

# url

`https://api.tapd.cn/api/testx/design/v2/namespaces/{namespace}/designs/search`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

POST

# 请求参数

## request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| namespace | 是 | string | 命名空间 |
| filter | 否 | Filter | 过滤条件 |
| page_info | 否 | PageInfo | 分页信息 |

## Filter


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| uids | 否 | repeated string | 设计唯一标识列表 |
| search | 否 | string | 搜索关键词 |
| creators | 否 | repeated string | 创建者列表 |
| updaters | 否 | repeated string | 更新者列表 |
| time_filters | 否 | repeated TimeFilter | 时间过滤条件 |
| orders | 否 | repeated Order | 排序条件 |
| period | 否 | Period | 时间范围 |
| ref_uid | 否 | string | 用例管理中主键ID |
| states | 否 | repeated State | 用例管理中主键ID |
| labels | 否 | repeated string | 标签过滤 |
| story_ids | 否 | repeated string | 需求过滤 |
| unscoped | 否 | bool | 是否查询全部数据 |
| custom_fields | 否 | map<string, string> | 用户自定义查询条件 |

## Period


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| begin_time | 是 | string | 开始时间 |
| end_time | 是 | string | 截止时间 |
| type | 是 | Type | 时间类型 |

## Type


| 取值 | 说明 |
| --- | --- |
| NONE | 接口时间不确定动作 |
| CREATE | 创建时间 |
| UPDATE | 更新时间 |

## TimeFilter


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| start_at | 否 | string | 开始时间 |
| end_at | 否 | string | 结束时间 |
| time_type | 否 | TimeType | 时间类型 |

## TimeType


| 取值 | 说明 |
| --- | --- |
| CREATE_AT | 创建时间 |
| UPDATE_AT | 更新时间 |

## Order


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| order_field | 否 | OrderFiled | 排序字段 |
| order_type | 否 | OrderType | 排序类型 |

## OrderType


| 取值 | 说明 |
| --- | --- |
| ASC | 升序 |
| DESC | 降序 |

## OrderFiled


| 取值 | 说明 |
| --- | --- |
| CREATE_AT | 创建时间 |
| UPDATE_AT | 更新时间 |

## State


| 取值 | 说明 |
| --- | --- |
| UNREVIEW | 未评审 |
| PROCESSING | 评审中 |
| FINISHED | 评审完成 |
| PASS | 通过 |

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
"Meta": {
"Audit": {
"Creator": "user_id",
"Updater": "user_id",
"CreatedAt": "2025-07-09T14:42:57+08:00",
"UpdatedAt": "2025-07-14T16:51:31+08:00",
"Tenant": ""
},
"Namespace": "namespace",
"Uid": "design_uid",
"ParentUid": "",
"Name": "test",
"Description": "",
"Thumbnail": "",
"Source": "MANUALLY",
"UseAI": false,
"TemplateUid": ""
},
"Spec": {
"Theme": "",
"Labels": [],
"Stories": [
{
"IssueUid": "issue_uid",
"Namespace": "namespace",
"WorkspaceUid": "workspace_uid",
"IssueUrl": "",
"Type": "STORY",
"Source": "TAPD",
"Detail": null,
"IssueName": "name",
"IsDeleted": false,
"Uid": "story_uid"
}
],
"Bugs": [],
"CustomFields": {}
},
"Review": {
"Audit": {
"Creator": "user_id",
"Updater": "user_id",
"CreatedAt": "2025-07-09T14:42:59+08:00",
"UpdatedAt": "2025-07-09T14:42:59+08:00",
"Tenant": ""
},
"Namespace": "namespace",
"DesignUid": "",
"Uid": "design_review_uid",
"State": "UNREVIEW",
"Comments": [],
"Reviewers": [],
"Stat": null
},
"Nodes": [],
"Stat": {
"CustomCount": 0,
"StoryCount": 2,
"SceneCount": 0,
"TestPointCount": 1,
"CaseCount": 1,
"FeatureCount": 1,
"BugCount": 0,
"TotalCount": 5,
"NodeUid": "",
"ExecCount": 0,
"ExecPassCount": 0,
"AiCaseCount": 0
}
}
],
"TotalCount": 1
}
```

# 测试设计字段说明

## Design


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| meta | Meta | 元数据 |
| spec | Spec | 规格信息 |
| review | Review | 总评信息 |
| stat | Statistics | 节点数据统计 |

## Meta


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| namespace | string | 命名空间 |
| uid | string | 唯一标识 |
| parent_uid | string | 父级唯一标识 |
| name | string | 名称 |
| description | string | 描述 |
| thumbnail | string | 缩略图地址 |
| source | Source | 来源 |
| use_ai | bool | 是否使用过AI |
| template_uid | string | 模版UID |

## Source


| 取值 | 说明 |
| --- | --- |
| MANUALLY | 手动创建 |
| IMPORT | 导入 |
| TEMPLATE | 模板创建 |

## Spec


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| theme | string | 画布主题 |
| labels | repeated Label | 关联标签 |
| stories | repeated Issue | 关联需求 |
| bugs | repeated Issue | 关联缺陷 |
| custom_fields | map<string, string> | 用户自定义属性值 |

## Statistics


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| custom_count | uint32 | 自定义计数 |
| story_count | uint32 | 需求计数 |
| scene_count | uint32 | 场景计数 |
| test_point_count | uint32 | 测试点计数 |
| case_count | uint32 | 用例计数 |
| feature_count | uint32 | 特性计数 |
| bug_count | uint32 | 缺陷计数 |
| total_count | uint32 | 节点总数 |
| node_uid | string | 节点Uid |
| exec_count | uint32 | 执行次数 |
| exec_pass_count | uint32 | 执行通过次数 |
| ai_case_count | uint32 | AI用例数量 |

## Review


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| namespace | string | 命名空间 |
| design_uid | string | 测试设计Uid |
| uid | string | 总评Uid |
| state | State | 总评状态 |
| comments | repeated Comment | 总评论信息列表 |
| reviewers | repeated string | 画布评审人 |

## State


| 取值 | 说明 |
| --- | --- |
| UNREVIEW | 未评审 |
| PROCESSING | 评审中 |
| FINISHED | 评审完成 |
| PASS | 通过 |

## Comment


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| uid | string | Uid |
| message | string | 评论信息 |
| reply_uid | string | 回复模块Uid |
| review_uid | string | 评审Uid |

## Issue


| 字段名 | 类型及范围 | 说明 |
| --- | --- | --- |
| issue_uid | string | Issue唯一标识 |
| namespace | string | 命名空间 |
| workspace_uid | string | 工作空间唯一标识 |
| issue_url | string | Issue链接 |
| type | Type | Issue类型 |
| source | Source | Issue来源 |
| detail | google.protobuf.Struct | Issue降级数据(需求删除时，前端可以降级显示) |
| issue_name | string | Issue名称 |
| is_deleted | bool | 是否已删除 |
| uid | string | 业务存储Issue的UID |

## IssueType


| 参数名称 | 说明 |
| --- | --- |
| UNKNOWN | 未知类型 |
| ITERATION | 迭代 |
| VERSION | 版本 |
| STORY | 需求 |
| TASK | 子任务 |
| BUG | Bug |
| TODO | 事项 |

## IssueSource


| 参数名称 | 说明 |
| --- | --- |
| NONE | 无来源 |
| TAPD | TAPD需求 |
| ZHIYAN | 智研需求 |
| LOCAL_WORD | 3本地Word |