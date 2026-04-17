# TAPD Mini API Reference

Base URL: `https://api.tapd.cn`

认证方式：
- Basic Auth: `curl -u 'api_user:api_password'`
- OAuth Access Token: `curl -H 'Authorization: Bearer ACCESS_TOKEN'`

标准响应格式: `{ "status": 1, "data": {...}, "info": "success" }`

---

## 一、工作项 (Mini Item)

### 1. 添加工作项

- **Method:** `POST`
- **Endpoint:** `/mini_items`
- **说明:** 每次只能添加一条记录

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| name | 是 | string | 标题 |
| priority | 否 | string | 优先级 |
| owner | 否 | string | 处理人 |
| creator | 否 | string | 创建人 |
| is_archived | 否 | boolean | 是否归档 |
| begin | 否 | date | 预计开始 |
| due | 否 | date | 预计结束 |
| parent_id | 否 | integer | 父工作项ID |
| category_id | 否 | integer | 分组 |
| description | 否 | string | 详细描述 |
| label | 否 | string | 标签，多个用`\|`分隔，不存在则自动创建 |
| custom_field_* | 否 | string/integer | 自定义字段 |

**响应字段 (data.MiniItem):** id, name, workspace_id, category_id, status, owner, begin, due, priority, label, description_type, description, markdown_description, ancestor_id, parent_id, children_id, level, creator, created, modifier, modified, completed, has_attachment, sort, is_archived, progress_manual, participator, custom_field_one ~ custom_field_100

---

### 2. 更新工作项

- **Method:** `POST`
- **Endpoint:** `/mini_items`
- **说明:** 每次只能更新一条记录

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| id | 是 | integer | 工作项ID |
| workspace_id | 是 | integer | 空间ID |
| name | 否 | string | 标题 |
| priority | 否 | string | 优先级 |
| status | 否 | string | 状态 |
| progress_manual | 否 | integer | 进度 |
| owner | 否 | string | 处理人 |
| begin | 否 | date | 预计开始 |
| due | 否 | date | 预计结束 |
| category_id | 否 | integer | 分组ID |
| description | 否 | string | 详细描述 |
| label | 否 | string | 标签，多个用`\|`分隔 |
| custom_field_* | 否 | string/integer | 自定义字段 |

**响应字段:** 同添加工作项

---

### 3. 获取工作项

- **Method:** `GET`
- **Endpoint:** `/mini_items`
- **分页:** 默认30条/页，最大200条

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 工作项ID，支持多ID查询 |
| name | 否 | string | 标题，支持模糊匹配 |
| priority | 否 | string | 优先级，支持枚举查询 |
| status | 否 | string | 状态，支持枚举查询 |
| label | 否 | string | 标签，支持枚举查询 |
| owner | 否 | string | 处理人，支持模糊匹配 |
| is_archived | 否 | boolean | 是否归档 |
| creator | 否 | string | 创建人，支持多人查询 |
| begin | 否 | date | 预计开始，支持时间范围查询 |
| due | 否 | date | 预计结束，支持时间范围查询 |
| created | 否 | datetime | 创建时间，支持时间范围查询 |
| modified | 否 | datetime | 最后修改时间，支持时间范围查询 |
| completed | 否 | datetime | 完成时间，支持时间范围查询 |
| progress_manual | 否 | integer | 进度 |
| category_id | 否 | integer | 分组，支持枚举查询 |
| parent_id | 否 | integer | 父工作项 |
| children_id | 否 | string | 子工作项 |
| description | 否 | string | 详细描述，支持模糊匹配 |
| custom_field_* | 否 | string/integer | 自定义字段 |
| limit | 否 | integer | 每页数量，默认30，最大200 |
| page | 否 | integer | 页码，默认1 |
| order | 否 | string | 排序规则，如 `created desc` |
| fields | 否 | string | 返回字段，逗号分隔 |

**响应字段:** 同添加工作项

---

### 4. 获取工作项数量

- **Method:** `GET`
- **Endpoint:** `/mini_items/count`
- **说明:** 默认包含已归档的工作项

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 工作项ID，支持多ID查询 |
| name | 否 | string | 标题，支持模糊匹配 |
| priority | 否 | string | 优先级，支持枚举查询 |
| status | 否 | string | 状态，支持枚举查询 |
| label | 否 | string | 标签，支持枚举查询 |
| owner | 否 | string | 处理人，支持模糊匹配 |
| creator | 否 | string | 创建人，支持模糊匹配 |
| begin | 否 | date | 预计开始，支持时间范围查询 |
| due | 否 | date | 预计结束，支持时间范围查询 |
| created | 否 | datetime | 创建时间，支持时间范围查询 |
| modified | 否 | datetime | 最后修改时间，支持时间范围查询 |
| completed | 否 | datetime | 完成时间，支持时间范围查询 |
| progress_manual | 否 | integer | 进度 |
| category_id | 否 | integer | 分组，支持枚举查询 |
| parent_id | 否 | integer | 父工作项 |
| children_id | 否 | string | 子工作项 |
| description | 否 | string | 详细描述，支持模糊匹配 |
| is_archived | 否 | boolean | 是否归档 |
| custom_field_* | 否 | string/integer | 自定义字段 |

**响应字段:** data.count (integer)

---

### 5. 添加分组

- **Method:** `POST`
- **Endpoint:** `/mini_item_categories`
- **说明:** 每次只能添加一条记录

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| name | 是 | string | 分组名称 |

**响应字段 (data.Category):** id, workspace_id, name, created, creator, modified, modifier

---

### 6. 更新分组

- **Method:** `POST`
- **Endpoint:** `/mini_item_categories`
- **说明:** 每次只能更新一条记录

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 是 | integer | 分组ID |
| name | 否 | string | 分组名称 |

**响应字段:** 同添加分组

---

### 7. 获取分组

- **Method:** `GET`
- **Endpoint:** `/mini_item_categories`
- **分页:** 默认30条/页，最大200条
- **说明:** "未分组"(id=-1)不可查询/修改/删除

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 分组ID，支持多ID查询 |
| name | 否 | string | 分组名称，支持模糊匹配 |
| created | 否 | datetime | 创建时间，支持时间查询 |
| modified | 否 | datetime | 最后修改时间，支持时间查询 |
| limit | 否 | integer | 每页数量，默认30，最大200 |
| page | 否 | integer | 页码，默认1 |
| order | 否 | string | 排序规则 |
| fields | 否 | string | 返回字段，逗号分隔 |

**响应字段 (data[].Category):** id, workspace_id, name, description, parent_id, created, creator, modified, modifier

---

### 8. 获取分组数量

- **Method:** `GET`
- **Endpoint:** `/mini_item_categories/count`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 分组ID，支持多ID查询 |
| name | 否 | string | 分组名称，支持模糊匹配 |
| created | 否 | datetime | 创建时间，支持时间查询 |
| modified | 否 | datetime | 最后修改时间，支持时间查询 |

**响应字段:** data.count (integer)

---

### 9. 获取工作项动态

- **Method:** `GET`
- **Endpoint:** `/mini_item_changes`
- **分页:** 默认30条/页，最大100条
- **说明:** created 和 mini_item_id 至少提供一个

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 动态ID，支持多ID查询 |
| mini_item_id | 否 | integer | 工作项ID，支持多ID查询 |
| creator | 否 | string | 创建人/操作人 |
| created | 否 | datetime | 创建时间/变更时间，支持时间查询 |
| change_type | 否 | string | 变更类型：`api`, `manual_update` |
| change_summary | 否 | string | 变更描述 |
| comment | 否 | string | 评论 |
| change_field | 否 | string | 指定变更字段 |
| need_parse_changes | 否 | integer | 是否返回field_changes，默认1 |
| limit | 否 | integer | 每页数量，默认30，最大100 |
| page | 否 | integer | 页码，默认1 |
| order | 否 | string | 排序规则 |
| fields | 否 | string | 返回字段，逗号分隔 |

**响应字段 (data[].WorkitemChange):** id, workspace_id, workitem_type_id, creator, created, change_summary, comment, changes, entity_type, change_type, change_type_text, field_changes, mini_item_id

**field_changes 子字段:** field, value_before, value_after, field_name, field_type, value_before_parsed, value_after_parsed, field_label

---

### 10. 获取工作项动态数量

- **Method:** `GET`
- **Endpoint:** `/mini_item_changes/count`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 动态ID |
| mini_item_id | 否 | integer | 工作项ID |
| creator | 否 | string | 创建人/操作人 |
| created | 否 | datetime | 创建时间/变更时间，支持时间查询 |
| change_summary | 否 | string | 变更描述 |
| comment | 否 | string | 评论 |
| changes | 否 | string | 变更详细记录 |

**响应字段:** data.count (integer)

---

### 11. 获取工作项自定义字段配置

- **Method:** `GET`
- **Endpoint:** `/mini_items/custom_fields_settings`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |

**响应字段 (data[].CustomFieldConfig):** id, workspace_id, entry_type, custom_field, type, name, options, enabled, sort, freeze, extra_config, memo

---

### 12. 获取工作项所有字段的中英文

- **Method:** `GET`
- **Endpoint:** `/mini_items/get_fields_label`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |

**响应字段 (data):** 返回中文标签到英文字段名的映射对象，如 `{ "标题": "name", "状态": "status", ... }`

---

### 13. 添加工作项与其他业务对象的关联关系

- **Method:** `POST`
- **Endpoint:** `/mini_items/create_mini_item_relation`
- **说明:** 每次只能添加一条记录

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| target_type | 是 | string | 目标对象类型：`story`, `bug`, `mini_item` |
| source_id | 是 | integer | 源对象ID |
| target_id | 是 | integer | 目标对象ID |

**响应字段:** data.success (boolean)

---

### 14. 获取关联需求

- **Method:** `GET`
- **Endpoint:** `/mini_items/get_link_stories`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| mini_item_id | 是 | integer | 工作项ID |

**响应字段 (data[]):** type, id, workspace_id, linked_workspace_id, actas, created, modified

---

### 15. 获取关联缺陷

- **Method:** `GET`
- **Endpoint:** `/mini_items/get_related_bugs`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| mini_item_id | 是 | integer | 工作项ID |

**响应字段 (data[]):** type, id, workspace_id, linked_workspace_id, actas, created, modified

---

### 16. 解除工作项与其他业务对象的关联关系

- **Method:** `POST`
- **Endpoint:** `/mini_items/remove_mini_item_relation`
- **说明:** 每次只能解除一条关联

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| mini_item_id | 是 | integer | 工作项ID |
| target_type | 是 | string | 业务对象类型，如 `bug` |
| target_id | 是 | integer | 业务对象ID |

**响应字段:** data.success (boolean)

---

### 17. 获取回收站内的工作项

- **Method:** `GET`
- **Endpoint:** `/mini_items/get_removed_mini_items`
- **分页:** 默认30条/页，最大200条

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 工作项ID |
| creator | 否 | string | 创建人 |
| created | 否 | date | 创建时间 |
| deleted | 否 | date | 删除时间 |
| limit | 否 | integer | 每页数量，默认30 |
| page | 否 | integer | 页码，默认1 |

**响应字段 (data[].RemovedMiniItem):** id, name, creator, created, operation_user, deleted

---

## 二、空间 (Workspace)

### 1. 获取空间信息

- **Method:** `GET`
- **Endpoint:** `/workspaces/get_workspace_info`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |

**响应字段 (data.Workspace):** id, name, category, status, description, product_type, begin_date, end_date, creator, created

---

### 2. 添加空间成员

- **Method:** `POST`
- **Endpoint:** `/workspaces/add_workspace_member`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| nick | 是 | string | 用户英文昵称 |
| company_id | 否 | integer | 公司ID（云版/SaaS版必填） |
| role_ids | 否 | string | 角色组ID，多个逗号分隔，取最高权限 |

**响应字段:** data.success (boolean)

---

### 3. 获取项目成员列表

- **Method:** `GET`
- **Endpoint:** `/workspaces/users`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID 或 公司ID |
| fields | 否 | string | 查询字段：user, role_id, email，逗号分隔 |

**响应字段 (data[].UserWorkspace):** user, role_id, email, name

**角色ID参考：**
| 角色ID | 角色名称 |
|--------|----------|
| 1000000000000000012 | 项目拥有者 |
| 1000000000000000002 | 空间管理员 |
| 1000000000000000089 | 普通成员 |
| 1000000000000000006 | 只读成员 |

---

### 4. 新建空间

- **Method:** `POST`
- **Endpoint:** `/workspaces/create_mini_project`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| company_id | 是 | integer | 公司ID |
| name | 是 | string | 空间名称 |
| creator | 是 | string | 创建人（必须属于当前公司） |
| description | 否 | string | 空间描述 |
| template_id | 否 | integer | 模板ID |

**响应字段:** data.workspace_id (string), data.workspace_url (string)

---

### 5. 获取用户所有参与的空间

- **Method:** `GET`
- **Endpoint:** `/workspaces/get_mini_project_list_with_permission`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| user | 是 | string | 用户名 |
| company_id | 是 | string | 公司ID |
| can_edit | 否 | boolean | 是否过滤有编辑权限的空间，默认false |

**响应字段 (data[]):** id, name, status, creator, created

---

## 三、评论 (Comment)

### 1. 添加评论

- **Method:** `POST`
- **Endpoint:** `/comments`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| description | 是 | string | 评论内容（需用`<p>`标签包裹） |
| author | 是 | string | 评论作者 |
| entry_id | 是 | integer | 工作项ID |
| root_id | 否 | integer | 根评论ID（用于嵌套评论） |
| reply_id | 否 | integer | 被回复的评论ID |

**响应字段 (data.Comment):** id, title, description, author, entry_type, entry_id, reply_id, root_id, created, modified, workspace_id

---

### 2. 更新评论

- **Method:** `POST`
- **Endpoint:** `/comments`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 是 | integer | 评论ID |
| description | 是 | string | 评论内容 |
| change_creator | 否 | string | 变更人 |

**响应字段:** 同添加评论

---

### 3. 获取评论

- **Method:** `GET`
- **Endpoint:** `/comments`
- **分页:** 默认30条/页，最大200条

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 评论ID，支持多ID查询 |
| title | 否 | string | 标题 |
| description | 否 | string | 内容 |
| author | 否 | string | 评论作者 |
| entry_id | 否 | integer | 工作项ID |
| created | 否 | datetime | 创建时间，支持时间范围查询 |
| modified | 否 | datetime | 最后修改时间，支持时间范围查询 |
| root_id | 否 | integer | 根评论ID |
| reply_id | 否 | integer | 被回复评论ID |
| limit | 否 | integer | 每页数量，默认30，最大200 |
| page | 否 | integer | 页码，默认1 |
| order | 否 | string | 排序规则 |
| fields | 否 | string | 返回字段，逗号分隔 |

**响应字段:** 同添加评论

---

### 4. 获取评论数量

- **Method:** `GET`
- **Endpoint:** `/comments/count`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 评论ID，支持多ID查询 |
| title | 否 | string | 标题 |
| description | 否 | string | 内容 |
| author | 否 | string | 评论作者 |
| entry_id | 否 | integer | 工作项ID |
| created | 否 | datetime | 创建时间，支持时间范围查询 |
| modified | 否 | datetime | 最后修改时间，支持时间范围查询 |

**响应字段:** data.count (integer)

---

## 四、附件 (Attachment)

### 1. 附件上传

- **Method:** `POST`
- **Endpoint:** `/files/upload_attachment`
- **说明:** 每次只能上传一个文件，最大150MB

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| file | 是 | file | 上传的文件 |
| type | 是 | string | 固定值：`story_custom_field` |
| custom_field | 是 | string | 自定义字段英文名 |
| entry_id | 是 | integer | 工作项ID |
| owner | 否 | string | 附件创建者 |

**响应字段 (data.Attachment):** id, type, entry_id, filename, description, content_type, created, workspace_id, owner

---

### 2. 上传base64图片

- **Method:** `POST`
- **Endpoint:** `/files/upload_image_base64`
- **说明:** 每次只能上传一张图片，最大15MB

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| base64_data | 是 | string | base64格式的图片数据 |
| type | 是 | string | 固定值：`story_custom_field` |
| custom_field | 是 | string | 自定义字段英文名 |
| entry_id | 是 | integer | 工作项ID |
| owner | 否 | string | 附件创建者 |

**响应字段 (data.Attachment):** id, type, entry_id, filename, description, content_type, created

---

### 3. 获取单个附件下载链接

- **Method:** `GET`
- **Endpoint:** `/attachments/down`

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 是 | integer | 附件ID |

**响应字段 (data.Attachment):** id, type, entry_id, filename, description, content_type, created, workspace_id, owner, download_url (临时下载链接)

---

### 4. 获取附件

- **Method:** `GET`
- **Endpoint:** `/attachments`
- **分页:** 默认30条/页，最大200条

| 参数 | 必选 | 类型 | 说明 |
|------|------|------|------|
| workspace_id | 是 | integer | 空间ID |
| id | 否 | integer | 附件ID |
| entry_id | 否 | integer | 工作项ID |
| filename | 否 | string | 附件文件名 |
| owner | 否 | string | 上传者 |

**响应字段 (data[].Attachment):** id, type, entry_id, filename, description, content_type, created, workspace_id, owner

---

## API 汇总

| 分类 | API数量 |
|------|---------|
| 工作项 (Mini Item) | 17 |
| 空间 (Workspace) | 5 |
| 评论 (Comment) | 4 |
| 附件 (Attachment) | 4 |
| **合计** | **30** |
