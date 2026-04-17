# TAPD Open Platform API 完整参考文档

> **文档来源：** https://open.tapd.cn/document/api-doc/API文档/api_reference/
> **Base URL：** `https://api.tapd.cn`
> **认证方式：** HTTP Basic Auth（`api_user:api_password`）
> **数据格式：** JSON（默认），可通过 `?format=xml` 切换为 XML

---

## 目录

1. [需求 (Stories)](#1-需求-stories)
2. [缺陷 (Bugs)](#2-缺陷-bugs)
3. [任务 (Tasks)](#3-任务-tasks)
4. [迭代 (Iterations)](#4-迭代-iterations)
5. [发布计划 (Releases)](#5-发布计划-releases)
6. [Wiki](#6-wiki)
7. [评论 (Comments)](#7-评论-comments)
8. [测试 (Testing)](#8-测试-testing)
9. [项目/工作区 (Workspace)](#9-项目工作区-workspace)
10. [用户 (Users)](#10-用户-users)
11. [附件 (Attachments)](#11-附件-attachments)
12. [工时 (Worklogs)](#12-工时-worklogs)
13. [看板 (Kanban)](#13-看板-kanban)
14. [源码/提交 (Commits)](#14-源码提交-commits)
15. [Webhook](#15-webhook)
16. [通用查询规则](#通用查询规则)

---

## 1. 需求 (Stories)

### GET `/stories` — 获取需求列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 需求 ID（支持多 ID，逗号分隔） |
| `name` | ❌ | string | 标题（模糊匹配） |
| `status` | ❌ | string | 状态（枚举查询） |
| `owner` | ❌ | string | 处理人（模糊匹配） |
| `creator` | ❌ | string | 创建人（支持多人查询） |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `iteration_id` | ❌ | string | 迭代 ID |
| `category_id` | ❌ | integer | 需求分类 ID |
| `label` | ❌ | string | 标签（枚举查询） |
| `workitem_type_id` | ❌ | string | 需求类型 ID |
| `release_id` | ❌ | integer | 发布计划 ID |
| `parent_id` | ❌ | integer | 父需求 ID |
| `created` | ❌ | datetime | 创建时间（支持范围） |
| `modified` | ❌ | datetime | 最后修改时间 |
| `begin` | ❌ | date | 预计开始日期 |
| `due` | ❌ | date | 预计结束日期 |
| `effort` | ❌ | string | 预估工时 |
| `module` | ❌ | string | 模块 |
| `version` | ❌ | string | 版本 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |
| `limit` | ❌ | integer | 每页数量（默认 30，最大 200） |
| `page` | ❌ | integer | 页码（默认 1） |
| `order` | ❌ | string | 排序规则（如 `created desc`） |
| `fields` | ❌ | string | 返回字段，逗号分隔 |

---

### POST `/stories` — 创建需求

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ✅ | string | 需求标题 |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `business_value` | ❌ | integer | 业务价值 |
| `version` | ❌ | string | 版本 |
| `module` | ❌ | string | 模块 |
| `owner` | ❌ | string | 处理人 |
| `cc` | ❌ | string | 抄送人 |
| `creator` | ❌ | string | 创建人 |
| `developer` | ❌ | string | 开发人员 |
| `begin` | ❌ | date | 计划开始日期 |
| `due` | ❌ | date | 计划结束日期 |
| `iteration_id` | ❌ | string | 迭代 ID |
| `parent_id` | ❌ | integer | 父需求 ID |
| `effort` | ❌ | string | 预估工时 |
| `effort_completed` | ❌ | string | 完成工时 |
| `remain` | ❌ | float | 剩余工时 |
| `category_id` | ❌ | integer | 需求分类 ID |
| `workitem_type_id` | ❌ | integer | 需求类型 ID |
| `release_id` | ❌ | integer | 发布计划 ID |
| `source` / `type` / `feature` | ❌ | string | 来源 / 类型 / 特性 |
| `tech_risk` | ❌ | string | 技术风险 |
| `description` | ❌ | string | 详细描述 |
| `label` | ❌ | string | 标签（竖线分隔） |
| `templated_id` | ❌ | integer | 模板 ID |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### POST `/stories`（含 `id`）— 更新需求

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `id` | ✅ | integer | 需求 ID |
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ❌ | string | 标题 |
| `status` | ❌ | string | 状态 |
| `v_status` | ❌ | string | 状态（可用中文名） |
| `owner` | ❌ | string | 处理人 |
| `current_user` | ❌ | string | 修改人 |
| `iteration_id` | ❌ | string | 迭代 ID |
| `release_id` | ❌ | integer | 发布计划 ID |
| `description` | ❌ | string | 描述 |
| `label` | ❌ | string | 标签（竖线分隔） |
| `is_auto_close_task` | ❌ | integer | 自动关闭任务（1=是） |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### POST `/stories/batch_update_story` — 批量更新需求

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `workitems` | ✅ | array | 需求更新对象数组（最多 50 条） |
| ↳ `id` | ✅ | integer | 需求 ID（每条必填） |
| ↳ 其他字段 | ❌ | — | 同创建/更新需求字段 |

---

### GET `/stories/count` — 获取需求数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| （筛选参数）| ❌ | — | 与 GET /stories 相同 |

**返回：** `{ "data": { "count": N } }`

---

### GET `/story_changes` — 获取需求变更历史

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `story_id` | ✅* | integer | 需求 ID（*与 `created` 二选一） |
| `created` | ✅* | datetime | 变更日期（*与 `story_id` 二选一） |
| `id` | ❌ | integer | 变更记录 ID |
| `creator` | ❌ | string | 操作人 |
| `change_type` | ❌ | string | 变更类型 |
| `change_field` | ❌ | string | 变更字段过滤 |
| `need_parse_changes` | ❌ | integer | 返回 field_changes（默认 1） |
| `limit` | ❌ | integer | 最多 100 条 |
| `page` | ❌ | integer | 默认 1 |
| `order` / `fields` | ❌ | string | 排序/字段选择 |

---

### GET `/story_categories` — 获取需求分类

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 分类 ID |
| `name` | ❌ | string | 分类名（模糊） |
| `parent_id` | ❌ | integer | 父分类 ID |
| `created` / `modified` | ❌ | datetime | 时间过滤 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页与排序 |

---

### GET `/stories/get_link_stories` — 获取需求关联关系

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `story_id` | ✅ | integer | 需求 ID（19位） |

**返回：** `[ { type, id, story_id, linked_workspace_id, actas, created } ]`

---

### GET `/stories/get_fields_info` — 获取需求字段及候选值

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |

**返回：** 所有字段定义（标签、类型、选项值）

---

## 2. 缺陷 (Bugs)

### GET `/bugs` — 获取缺陷列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 缺陷 ID（支持多 ID） |
| `title` | ❌ | string | 标题（模糊匹配） |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `severity` | ❌ | string | 严重程度 |
| `status` | ❌ | string | 状态 |
| `iteration_id` | ❌ | string | 迭代 ID |
| `module` | ❌ | string | 模块 |
| `current_owner` | ❌ | string | 处理人（模糊） |
| `reporter` | ❌ | string | 报告人/创建人 |
| `label` | ❌ | string | 标签查询 |
| `release_id` | ❌ | integer | 发布计划 ID |
| `version_report` | ❌ | string | 发现版本 |
| `version_fix` | ❌ | string | 修复版本 |
| `cc` / `de` / `te` | ❌ | string | 抄送/开发/测试 |
| `created` | ❌ | datetime | 创建时间 |
| `modified` | ❌ | datetime | 最后修改时间 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |
| `limit` | ❌ | integer | 默认 30，最大 200 |
| `page` | ❌ | integer | 默认 1 |
| `order` | ❌ | string | 排序规则 |
| `fields` | ❌ | string | 返回字段 |

---

### POST `/bugs`（不含 `id`）— 创建缺陷

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `title` | ✅ | string | 缺陷标题 |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `severity` | ❌ | string | 严重程度 |
| `status` | ❌ | string | 状态 |
| `module` | ❌ | string | 模块 |
| `current_owner` | ❌ | string | 处理人 |
| `reporter` | ❌ | string | 报告人 |
| `cc` / `de` / `te` | ❌ | string | 抄送/开发/测试 |
| `iteration_id` | ❌ | string | 迭代 |
| `release_id` | ❌ | integer | 发布计划 |
| `label` | ❌ | string | 标签 |
| `description` | ❌ | string | 描述 |
| `estimate` | ❌ | integer | 预估修复时间 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### POST `/bugs`（含 `id`）— 更新缺陷

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `id` | ✅ | integer | 缺陷 ID |
| `workspace_id` | ✅ | integer | 项目 ID |
| `title` | ❌ | string | 标题 |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `severity` | ❌ | string | 严重程度 |
| `status` | ❌ | string | 状态 |
| `v_status` | ❌ | string | 状态（中文名） |
| `module` / `feature` | ❌ | string | 模块/特性 |
| `release_id` | ❌ | integer | 发布计划 |
| `current_owner` | ❌ | string | 处理人 |
| `cc` | ❌ | string | 抄送 |
| `reporter` | ❌ | string | 报告人 |
| `label` | ❌ | string | 标签 |
| `estimate` | ❌ | integer | 预估修复时间 |
| `description` | ❌ | string | 描述 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### POST `/bugs/copy_bug` — 复制缺陷

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 源项目 ID |
| `src_bug_id` | ✅ | integer | 源缺陷 ID |
| `dst_workspace_id` | ✅ | integer | 目标项目 ID |
| `sync_fields` | ❌ | string | 要复制的字段（逗号分隔）：title, status, description, attachment, current_owner, cc, module, iteration_id, priority, severity, version_*, label, custom_field 等 |

---

### POST `/bugs/batch_update_bugs` — 批量更新缺陷

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `workitems` | ✅ | array | 缺陷更新对象数组 |

---

### GET `/bugs/count` — 获取缺陷数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| （筛选参数）| ❌ | — | 与 GET /bugs 相同 |

**返回：** `{ "data": { "count": N } }`

---

### GET `/bug_changes` — 获取缺陷变更历史

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `bug_id` | ✅* | integer | 缺陷 ID（*与 `created` 二选一） |
| `created` | ✅* | datetime | 日期（*与 `bug_id` 二选一） |
| `id` | ❌ | integer | 变更记录 ID |
| `author` | ❌ | string | 变更人 |
| `field` | ❌ | string | 变更字段 |
| `old_value` / `new_value` | ❌ | string | 变更前/后值 |
| `include_add_bug` | ❌ | integer | 包含创建事件（1=是） |
| `limit` | ❌ | integer | 默认 30，最大 200 |
| `page` / `order` / `fields` | ❌ | — | 分页与排序 |

---

## 3. 任务 (Tasks)

### GET `/tasks` — 获取任务列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 任务 ID（支持多 ID） |
| `name` | ❌ | string | 标题（模糊） |
| `description` | ❌ | string | 描述 |
| `creator` | ❌ | string | 创建人 |
| `status` | ❌ | string | 状态（枚举） |
| `owner` | ❌ | string | 处理人（模糊） |
| `cc` | ❌ | string | 抄送人 |
| `story_id` | ❌ | integer | 关联需求 ID（支持多个） |
| `iteration_id` | ❌ | integer | 迭代 ID |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `progress` | ❌ | integer | 进度百分比 |
| `label` | ❌ | string | 标签（枚举） |
| `begin` / `due` | ❌ | date | 开始/结束日期 |
| `created` / `modified` / `completed` | ❌ | datetime | 时间过滤 |
| `effort` | ❌ | string | 预估工时 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

---

### POST `/tasks`（不含 `id`）— 创建任务

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ✅ | string | 任务标题 |
| `description` | ❌ | string | 描述 |
| `creator` | ❌ | string | 创建人 |
| `owner` | ❌ | string | 处理人 |
| `cc` | ❌ | string | 抄送人 |
| `begin` / `due` | ❌ | date | 开始/结束日期 |
| `story_id` | ❌ | integer | 关联需求 ID |
| `iteration_id` | ❌ | integer | 迭代 ID |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `effort` | ❌ | string | 预估工时 |
| `label` | ❌ | string | 标签 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### POST `/tasks`（含 `id`）— 更新任务

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `id` | ✅ | integer | 任务 ID |
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ❌ | string | 标题 |
| `description` | ❌ | string | 描述 |
| `status` | ❌ | string | 状态 |
| `owner` | ❌ | string | 处理人 |
| `current_user` | ❌ | string | 操作人 |
| `cc` | ❌ | string | 抄送人 |
| `begin` / `due` | ❌ | date | 开始/结束日期 |
| `story_id` | ❌ | integer | 关联需求 |
| `iteration_id` | ❌ | integer | 迭代 |
| `priority_label` | ❌ | string | 优先级（推荐使用） |
| `effort` | ❌ | string | 预估工时 |
| `auto_complete_effort` | ❌ | integer | 完成时自动填充工时 |
| `label` | ❌ | string | 标签 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### GET `/tasks/count` — 获取任务数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| （筛选参数）| ❌ | — | 与 GET /tasks 相同 |

**返回：** `{ "data": { "count": N } }`

---

### GET `/task_changes` — 获取任务变更历史

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `task_id` | ❌ | integer | 任务 ID（支持多 ID） |
| `created` | ❌ | datetime | 日期过滤 |
| `id` / `creator` | ❌ | — | 变更 ID / 操作人 |
| `change_summary` / `comment` | ❌ | string | 变更描述 |
| `need_parse_changes` | ❌ | integer | 返回 field_changes（默认 1） |
| `limit` | ❌ | integer | 默认 30，最大 100 |
| `page` / `order` / `fields` | ❌ | — | 分页 |

---

## 4. 迭代 (Iterations)

### GET `/iterations` — 获取迭代列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 迭代 ID（支持多 ID） |
| `name` | ❌ | string | 标题（模糊） |
| `description` | ❌ | string | 描述 |
| `status` | ❌ | string | 状态（open/done 或中文） |
| `startdate` / `enddate` | ❌ | date | 日期范围 |
| `workitem_type_id` | ❌ | integer | 迭代分类 |
| `plan_app_id` | ❌ | integer | 计划应用 ID |
| `creator` | ❌ | string | 创建人 |
| `created` / `modified` / `completed` | ❌ | datetime | 时间过滤 |
| `locker` | ❌ | string | 锁定人 |
| `custom_field_*` | ❌ | string/int | 自定义字段 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

---

### POST `/iterations` — 创建迭代

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ✅ | string | 标题 |
| `startdate` | ✅ | date | 开始日期 |
| `enddate` | ✅ | date | 结束日期 |
| `creator` | ✅ | string | 创建人 |
| `workitem_type_id` | ❌ | integer | 分类 ID |
| `plan_app_id` | ❌ | integer | 计划应用 ID |
| `entity_type` | ❌ | string | `iteration` 或 `release`（默认：iteration） |
| `parent_id` | ❌ | integer | 父计划 ID |
| `description` | ❌ | string | 描述 |
| `status` | ❌ | string | 状态 |
| `label` | ❌ | string | 标签（竖线分隔） |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### POST `/iterations`（含 `id`）— 更新迭代

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 迭代 ID |
| `current_user` | ✅ | string | 操作人 |
| `name` | ❌ | string | 标题 |
| `startdate` / `enddate` | ❌ | date | 日期范围 |
| `description` | ❌ | string | 描述 |
| `status` | ❌ | string | `open` 或 `done` |
| `custom_field_*` | ❌ | string/int | 自定义字段 |

---

### GET `/iterations/count` — 获取迭代数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| （筛选参数）| ❌ | — | 与 GET /iterations 相同 |

**返回：** `{ "data": { "count": N } }`

---

## 5. 发布计划 (Releases)

### GET `/releases` — 获取发布计划列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 发布 ID（支持多 ID） |
| `name` | ❌ | string | 标题（模糊） |
| `description` | ❌ | string | 描述 |
| `startdate` / `enddate` | ❌ | date | 日期范围 |
| `creator` | ❌ | string | 创建人 |
| `status` | ❌ | enum | `open` 或 `done` |
| `created` / `modified` | ❌ | datetime | 时间过滤 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

---

### POST `/releases` — 创建发布计划

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ✅ | string | 发布标题 |
| `startdate` | ✅ | date | 开始日期 |
| `enddate` | ✅ | date | 结束日期 |
| `description` | ❌ | string | 描述 |
| `creator` | ❌ | string | 创建人 |

---

### POST `/releases`（含 `id`）— 更新发布计划

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 发布 ID |
| `name` | ❌ | string | 标题 |
| `description` | ❌ | string | 描述 |
| `startdate` / `enddate` | ❌ | date | 日期范围 |
| `status` | ❌ | string | `open` 或 `done` |

---

## 6. Wiki

### GET `/wikis` — 获取 Wiki 列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | Wiki ID |
| `title` | ❌ | string | 标题（模糊） |
| `creator` | ❌ | string | 创建人 |
| `created` / `modified` | ❌ | datetime | 时间过滤 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

---

### POST `/wikis` — 创建 Wiki

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `title` | ✅ | string | Wiki 标题 |
| `description` | ❌ | string | Wiki 内容 |
| `creator` | ❌ | string | 创建人 |
| `parent_id` | ❌ | integer | 父 Wiki ID |
| `category_id` | ❌ | integer | 分类 |

---

### POST `/wikis`（含 `id`）— 更新 Wiki

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | Wiki ID |
| `title` | ❌ | string | 标题 |
| `description` | ❌ | string | 内容 |
| `operator` | ❌ | string | 操作人 |

---

### GET `/wikis/count` — 获取 Wiki 数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` / `id` | ❌ | integer | Wiki ID |

---

### GET `/wikis/get_drawio_data` — 获取流程图数据

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` | ✅ | integer | Wiki ID |

---

### GET `/wikis/get_wiki_followers` — 获取 Wiki 关注者列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` | ✅ | integer | Wiki ID |

---

### GET `/wikis/get_wiki_followers_count` — 获取 Wiki 关注者数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` | ✅ | integer | Wiki ID |

---

### GET `/wikis/get_wiki_permissions` — 获取 Wiki 权限

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` | ✅ | integer | Wiki ID |

---

### GET `/wikis/get_wiki_tags` — 获取 Wiki 标签

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` | ✅ | integer | Wiki ID |

---

### GET `/wikis/get_wiki_tags_count` — 获取 Wiki 标签数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` | ✅ | integer | Wiki ID |

---

### GET `/wikis/get_wiki_attachments_count` — 获取 Wiki 附件数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `wiki_id` | ✅ | integer | Wiki ID |

---

## 7. 评论 (Comments)

### POST `/comments` — 添加评论

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `description` | ✅ | string | 评论内容 |
| `author` | ✅ | string | 评论人用户名 |
| `entry_type` | ✅ | string | 对象类型：`bug`、`bug_remark`、`stories`、`tasks` |
| `entry_id` | ✅ | integer | 被评论对象的 ID |
| `root_id` | ❌ | integer | 根评论 ID（用于线程） |
| `reply_id` | ❌ | integer | 被回复的评论 ID |

---

### GET `/comments` — 获取评论列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 评论 ID（支持多 ID） |
| `author` | ❌ | string | 评论人 |
| `entry_type` | ❌ | string | 类型（竖线分隔多个） |
| `entry_id` | ❌ | integer | 关联对象 ID |
| `title` | ❌ | string | 标题 |
| `description` | ❌ | string | 内容 |
| `root_id` / `reply_id` | ❌ | integer | 线程 ID |
| `created` / `modified` | ❌ | datetime | 时间过滤 |
| `limit` | ❌ | integer | 默认 30，最大 200 |
| `page` / `order` / `fields` | ❌ | — | 分页 |

---

### GET `/comments/count` — 获取评论数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `entry_type` | ❌ | string | 评论类型过滤 |
| `entry_id` | ❌ | integer | 对象 ID 过滤 |
| （其他过滤参数）| ❌ | — | 与 GET /comments 相同 |

**返回：** `{ "data": { "count": N } }`

---

### POST `/comments`（含 `id`）— 更新评论

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 评论 ID |
| `description` | ✅ | string | 更新的内容 |
| `change_creator` | ❌ | string | 修改人姓名 |

---

## 8. 测试 (Testing)

### GET `/testcases` — 获取测试用例列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 用例 ID |
| `creator` | ❌ | string | 创建人 |
| `created` | ❌ | datetime | 创建时间 |
| `modified` | ❌ | datetime | 修改时间 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

---

### GET `/testcases/count` — 获取测试用例数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| （筛选参数）| ❌ | — | 与 GET /testcases 相同 |

---

### POST `/testcases` — 创建测试用例

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ✅ | string | 用例名称 |
| `creator` | ❌ | string | 创建人 |
| `description` | ❌ | string | 描述 |
| `priority` | ❌ | string | 优先级 |
| `category_id` | ❌ | integer | 用例分类 ID |

---

### POST `/testcases`（含 `id`）— 更新测试用例

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `id` | ✅ | integer | 用例 ID |
| `workspace_id` | ✅ | integer | 项目 ID |
| （其他字段）| ❌ | — | 同创建用例字段 |

---

### GET `/testcase_changes` — 获取测试用例变更历史

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `testcase_id` | ❌ | integer | 用例 ID |
| `created` | ❌ | datetime | 时间过滤 |

---

### GET `/testcases/get_fields_info` — 获取测试用例字段信息

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |

---

### GET `/testcases/custom_field` — 获取测试用例自定义字段配置

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |

---

### GET `/testplans` — 获取测试计划列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 计划 ID |
| `creator` | ❌ | string | 创建人 |
| `created` / `modified` | ❌ | datetime | 时间过滤 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

---

### POST `/testplans` — 创建测试计划

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `name` | ✅ | string | 计划名称 |
| `creator` | ❌ | string | 创建人 |
| `description` | ❌ | string | 描述 |

---

### POST `/testplans`（含 `id`）— 更新测试计划

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `id` | ✅ | integer | 计划 ID |
| `workspace_id` | ✅ | integer | 项目 ID |
| （其他字段）| ❌ | — | 同创建计划字段 |

---

### GET `/testplan_cases` — 获取测试计划中的用例

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `plan_id` | ✅ | integer | 计划 ID |
| `limit` / `page` / `fields` | ❌ | — | 分页 |

---

### POST `/testplan_cases` — 添加用例到测试计划

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `plan_id` | ✅ | integer | 计划 ID |
| `case_id` | ✅ | integer | 用例 ID |

---

### GET `/testruns` — 获取测试执行结果

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `plan_id` | ❌ | integer | 计划 ID |
| `case_id` | ❌ | integer | 用例 ID |
| `limit` / `page` / `fields` | ❌ | — | 分页 |

---

### POST `/testruns` — 更新测试执行结果

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 执行记录 ID |
| `status` | ❌ | string | 执行状态 |

---

### GET `/testcases/get_metrics` — 获取测试指标

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |

---

## 9. 项目/工作区 (Workspace)

### GET `/workspaces/get_workspace_info` — 获取项目信息

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |

**返回字段：** `id`, `name`, `pretty_name`, `category`, `status`（normal/closed/suspended）, `description`, `begin_date`, `end_date`, `creator`, `created`, `company_id`, `external_on`

---

### POST `/workspaces/add_workspace_member` — 添加项目成员

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `nick` | ✅ | string | 用户英文昵称 |
| `company_id` | ❌ | integer | 成员所属公司 ID |
| `role_ids` | ❌ | string | 角色组 ID（逗号分隔） |

*每次请求只添加一名成员；若成员已存在，则更新角色。*

---

### GET `/workspaces/sub_workspaces` — 获取子项目列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 父项目 ID |
| `template_id` | ❌ | integer | 项目模板 ID |

---

### GET `/workspaces/users` — 获取项目成员列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID 或公司 ID |
| `user` | ❌ | string | 昵称过滤（逗号分隔） |
| `fields` | ❌ | string | 返回字段：user, user_id, role_id, name, email, real_join_time |

---

### GET `/roles` — 获取用户组/角色 ID 映射

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |

**返回：** 角色 ID → 角色名称的映射对象

---

### GET `/workspaces/projects` — 获取公司项目列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `company_id` | ✅ | integer | 公司 ID |
| `category` | ❌ | string | `project`（协作）或 `mini_project`（轻量） |
| `with_extends` | ❌ | integer | `1` 包含自定义字段 |

---

### GET `/workspaces/user_participant_projects` — 获取用户参与的项目列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `nick` | ✅ | string | 成员昵称 |
| `company_id` | ✅ | integer | 公司 ID |

---

### GET `/workspaces/workspace_custom_field_settings` — 获取项目自定义字段配置

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 公司/工作区 ID |

**返回字段：** `id`, `workspace_id`, `entry_type`, `custom_field`, `type`, `name`, `options`, `extra_config`, `enabled`

---

### GET `/documents/get_workspace_documents` — 获取项目文档列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `limit` | ❌ | integer | 每页数量（默认 30，最大 200） |
| `page` | ❌ | integer | 页码（默认 1） |
| `fields` | ❌ | string | 返回字段（逗号分隔） |

---

### POST `/workspaces/update_workspace_info` — 更新项目信息

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `field` | ✅ | string | 要更新的字段：`description`, `begin_date`, `end_date`, `begin_end` |
| `value` | ✅ | string | 新值 |

---

### POST `/workspaces/set_custom_work_calendar` — 设置自定义工作日历

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `year` | ✅ | string | 日历年份 |
| `weekdays` | ❌ | array | 标准工作周天（1–7） |
| `holidays` | ❌ | array | 非工作日期 |
| `workdays` | ❌ | array | 额外工作日（覆盖节假日） |

---

### POST `/workspaces/enable_work_calendar` — 启用工作日历

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `type` | ✅ | string | `system` 或 `custom` |

---

### GET `/workspaces/get_custom_work_calendar` — 获取自定义工作日历详情

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `year` | ✅ | string | 日历年份 |

**返回字段：** `weekdays`, `holidays`, `workdays`

---

### GET `/workspaces/get_work_calendar_settings` — 获取工作日历设置列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |

**返回字段：** `name`, `type`, `enable`

---

### GET `/workspaces/get_workitems_long_id_by_short_ids` — 工作项短 ID ↔ 长 ID 转换

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `entity_type` | ✅ | string | `story`, `task` 或 `bug` |
| `short_ids` | ❌* | string | 短 ID（分号分隔） |
| `long_ids` | ❌* | string | 长 ID（分号分隔） |

*`short_ids` 与 `long_ids` 至少填一个。*

---

## 10. 用户 (Users)

### GET `/users/info` — 获取当前用户信息

*无需参数，使用认证用户上下文*

---

### GET `/users/get_personal_setting` — 获取用户个人配置

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 公司/工作区 ID |
| `nick` | ✅ | string | 用户标识（昵称） |

**返回字段：** `language`, `message_setting`

---

### GET `/user_oauth/get_user_view_list` — 获取用户需求视图列表

> ⚠️ 需要用户模式 OAuth Token，不支持 Basic Auth

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `type` | ❌ | string | 对象类型（目前仅支持 `story`） |

---

### GET `/user_oauth/get_user_todo_story` — 获取用户待处理需求

> ⚠️ 需要用户模式 OAuth Token

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `user` | ❌ | string | 用户标识 |
| `limit` | ❌ | integer | 每页数量（默认 30，最大 200） |
| `page` | ❌ | integer | 页码（默认 1） |
| `order` | ❌ | string | 排序（需 URL 编码） |
| `fields` | ❌ | string | 返回字段（逗号分隔） |

---

### GET `/user_oauth/get_user_todo_bug` — 获取用户待处理缺陷

> ⚠️ 需要用户模式 OAuth Token

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `user` | ❌ | string | 用户标识 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

---

### GET `/user_oauth/get_user_todo_task` — 获取用户待处理任务

> ⚠️ 需要用户模式 OAuth Token

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `user` | ❌ | string | 用户标识 |
| `limit` / `page` / `order` / `fields` | ❌ | — | 分页 |

**任务状态值：** `open` / `progressing` / `done`
**优先级值：** 4=高, 3=中, 2=低, 1=可选

---

### GET `/users/get_third_user_mapping` — 获取第三方用户 ID 映射

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 公司 ID |
| `user_id` | ✅ | integer | TAPD 用户 ID |

**返回字段：** `third_partys`（包含 `third_party_id`, `third_party_type` 的数组）

---

## 11. 附件 (Attachments)

### GET `/attachments` — 获取附件列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 附件 ID |
| `type` | ❌ | string | 附件分类 |
| `entry_id` | ❌ | integer | 关联对象 ID |
| `filename` | ❌ | string | 附件文件名 |
| `owner` | ❌ | string | 上传人 |
| `limit` / `page` | ❌ | integer | 分页 |

**返回字段：** `id`, `type`, `entry_id`, `filename`, `content_type`, `created`, `workspace_id`, `owner`

---

### GET `/attachments/down` — 获取单个附件下载链接

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 附件 ID |

**返回字段：** `id`, `type`, `entry_id`, `filename`, `description`, `content_type`, `created`, `workspace_id`, `owner`, `download_url`

*下载链接有效期 300 秒，每次请求返回一个附件。*

---

### GET `/files/get_image` — 获取单张图片下载链接

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `image_path` | ✅ | string | 图片路径（支持完整 URL，项目需匹配） |

**支持格式：** `png`, `gif`, `jpg`, `jpeg`, `bmp`，链接有效期 300 秒

---

### GET `/documents/down` — 获取单个文档下载链接

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 文档 ID |

**返回字段：** `id`, `workspace_id`, `name`, `type`, `folder_id`, `creator`, `modifier`, `created`, `modified`, `download_url`

---

## 12. 工时 (Worklogs)

### GET `/worklogs` — 获取工时记录列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 工时记录 ID |
| `owner` | ❌ | string | 记录人 |
| `entity_type` | ❌ | string | `story`, `bug`, `task` |
| `entity_id` | ❌ | integer | 关联对象 ID |
| `created` | ❌ | datetime | 创建时间过滤 |
| `spentdate` | ❌ | date | 工时日期 |
| `limit` | ❌ | integer | 每页数量（默认 30，最大 200） |
| `page` | ❌ | integer | 页码 |
| `fields` | ❌ | string | 返回字段 |

---

### POST `/worklogs` — 添加工时记录

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `owner` | ✅ | string | 记录人 |
| `entity_type` | ✅ | string | 关联对象类型 |
| `entity_id` | ✅ | integer | 关联对象 ID |
| `spentdate` | ✅ | date | 工时日期 |
| `spent` | ✅ | float | 工时数量（小时） |
| `description` | ❌ | string | 描述 |

---

### POST `/worklogs/update` — 更新工时记录

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 工时记录 ID |
| `spent` | ❌ | float | 更新工时数量 |
| `spentdate` | ❌ | date | 更新工时日期 |
| `description` | ❌ | string | 更新描述 |

---

### POST `/worklogs/delete` — 删除工时记录

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ✅ | integer | 工时记录 ID |

---

### GET `/worklogs/count` — 获取工时记录数量

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| （筛选参数）| ❌ | — | 与 GET /worklogs 相同 |

---

## 13. 看板 (Kanban)

### GET `/board_cards` — 获取看板卡片列表

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `id` | ❌ | integer | 工作项 ID（支持多个） |
| `b_board_id` | ❌ | integer | 看板 ID |
| `b_column_id` | ❌ | integer | 列/泳道 ID |
| `owner` | ❌ | string | 负责人 |
| `cc` | ❌ | string | 参与者 |
| `status` | ❌ | string | 状态 |
| `name` | ❌ | string | 工作项标题 |
| `created` | ❌ | datetime | 创建时间过滤 |
| `begin` | ❌ | datetime | 开始日期过滤 |
| `due` | ❌ | datetime | 截止日期过滤 |
| `b_label` | ❌ | integer | 标签 ID |
| `limit` | ❌ | integer | 每页数量（默认 30，最大 200） |
| `page` | ❌ | integer | 页码（默认 1） |
| `fields` | ❌ | string | 返回字段（逗号分隔） |

**返回字段：** `id`, `name`, `description`, `workspace_id`, `b_board_id`, `b_column_id`, `owner`, `cc`, `status`, `created`, `modified`, `begin`, `due`, `b_label`, `b_sort`

---

## 14. 源码/提交 (Commits)

### GET `/code_commit_infos` — 获取 Git 提交信息

| 参数 | 必填 | 类型 | 说明 |
|------|------|------|------|
| `workspace_id` | ✅ | integer | 项目 ID |
| `type` | ✅ | string | 对象类型：`story`, `bug`, 或 `task` |
| `object_id` | ✅ | integer | 关联对象 ID |
| `commit_time` | ❌ | string | 时间过滤：`>datetime`、`<datetime` 或 `date~date` |
| `related_type` | ❌ | string | 关联过滤：`all`/`branch`/`source_code`（默认：`all`） |
| `limit` | ❌ | integer | 每页数量（默认 30，最大 200） |
| `page` | ❌ | integer | 页码（默认 1） |

**返回字段：** `id`, `user_name`, `user_id`, `commit_id`, `message`, `commit_time`, `path`, `web_url`, `ref`, `git_env`, `file_commit`, `repo_id`, `branch_id`, `file_sort`

---

## 15. Webhook

Webhook 是 TAPD 的推送通知系统，通过开发者门户配置，无需 REST 调用。

### 支持的事件类型

| 事件 | 说明 |
|------|------|
| `story::create` | 需求创建 |
| `story::update` | 需求更新 |
| `story::delete` | 需求删除 |
| `bug::create` | 缺陷创建 |
| `bug::update` | 缺陷更新 |
| `bug::delete` | 缺陷删除 |
| `task::create` | 任务创建 |
| `task::update` | 任务更新 |
| `task::delete` | 任务删除 |
| `review::create` | 发布评审创建 |
| `review::update` | 发布评审更新 |
| `review::delete` | 发布评审删除 |
| 关联事件 | 绑定/解绑前后置对象 |

### Webhook 推送数据通用字段

| 字段 | 类型 | 说明 |
|------|------|------|
| `event` | string | 事件标识（如 `story::create`） |
| `event_from` | string | 触发来源：`web` 或 `api` |
| `workspace_id` | integer | 项目 ID |
| `current_user` | string | 操作人昵称 |
| `event_id` | integer | 唯一事件标识符 |
| `id` | integer | 对象 ID（19位格式） |
| `secret` | string | 可选验证密码 |
| `created` | datetime | 触发时间戳（Y-m-d H:i:s） |
| `old_*` | 多种 | 变更前字段值（仅更新事件） |
| `change_fields` | array | 变更字段列表（仅更新事件） |

**推送格式：** JSON **或** 表单编码（`key1=value1&key2=value2`）

---

## 通用查询规则

| 功能 | 语法示例 |
|------|------|
| **多 ID 查询** | `id=123,456,789` |
| **时间范围** | `created[]=2024-01-01&created[]=2024-12-31` |
| **不等式查询** | `status[ne]=done` 或 `status[in]=open,progressing` |
| **排序** | `order=created%20desc` |
| **字段选择** | `fields=id,name,status,owner` |
| **分页** | `limit=50&page=2`（每页最多 200） |
| **自定义字段** | `custom_field_1=value` |
| **模糊匹配** | `name`, `title`, `owner`, `reporter` 等字段支持 |

---

## API 分类汇总

| 分类 | 端点数量 | 说明 |
|------|----------|------|
| 需求 Stories | ~29 | CRUD、复制、分类、关联关系、自定义字段、模板 |
| 缺陷 Bugs | ~19 | CRUD、关联关系、自定义字段、模板 |
| 迭代 Iterations | ~8 | CRUD、自定义字段、锁定/解锁 |
| 任务 Tasks | ~10 | CRUD、批量操作 |
| 测试 Testing | ~32 | 测试用例、测试计划、执行记录、关联关系 |
| 发布 Release | ~11 | 计划、评审、附属对象 |
| 源码 Source Code | ~3 | 提交、关联、计数 |
| Wiki | ~10 | CRUD、关注者、权限、标签、附件 |
| 看板 Kanban | ~4 | 卡片和列管理 |
| 评论 Comments | ~4 | CRUD |
| 报表 Reports | ~1 | |
| 附件 Attachments | ~4 | 列表、下载链接 |
| 度量 Metrics | ~1 | |
| 工时 Worklogs | ~5 | CRUD + 计数 |
| 项目 Workspace | ~13 | 信息、成员、子项目、日历、自定义字段 |
| 项目集 Program | ~2 | |
| 工作流 Workflow | ~6 | |
| 配置 Settings | ~20 | 自定义字段、模块、版本、基线 |
| Webhook | — | 推送事件（需在开发者门户配置） |
| 用户 Users | ~8 | 信息、角色、待办、第三方映射 |
| **合计** | **~200+** | |

---

*文档整理日期：2026-04-17*
