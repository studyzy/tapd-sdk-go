# webhook文档

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/webhook/webhook_document.html

-   [新建](#新建)
-   [更新](#更新)
-   [删除](#删除)

目前 TAPD 支持 Webhook 了。现在支持需求、缺陷、任务、发布评审在创建、变更、删除后，往指定 URL POST 相关信息。

# 如何接入

目前可以在 左下角头像 -> `问题反馈` 上提交申请。提供要监听的`项目ID`或者`公司ID`，`事件`（需求创建、需求更新、需求状态变更、需求删除、缺陷创建、缺陷更新、缺陷状态变更、缺陷删除、任务创建、任务更新、任务状态变更、任务删除、发布评审创建、发布评审更新，前后置对象绑定，前后置对象解绑 ，可多选），`接收数据的 URL`，验证密码（非必选，给接入方验证请求是否来自 TAPD），数据格式（非必选，可选有 json 、form，默认为 form 形式，下面有说明）。

# 数据格式及结构

数据格式支持 json 及 form（x-www-form-urlencoded格式，即 key1=value1&key2=value2&key3=value 的形式，key及value都进行过 urlencode）。

## 新建

新建类事件有需求创建（story::create）、缺陷创建（bug::create）、任务创建（task::create）、发布评审创建（launchform::create）。当事件发生时， POST 过去的数据如下：

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">类型</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">event</td><td style="text-align:center;">string</td><td style="text-align:center;">事件名</td></tr><tr><td style="text-align:center;">event_from</td><td style="text-align:center;">string</td><td style="text-align:center;">事件触发地方。有web、api</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">string</td><td style="text-align:center;">操作人昵称</td></tr><tr><td style="text-align:center;">event_id</td><td style="text-align:center;">integer</td><td style="text-align:center;">事件ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">integer</td><td style="text-align:center;">新对象ID。需求、缺陷、任务、发布评审的都是19位长度的长ID</td></tr><tr><td style="text-align:center;">secret</td><td style="text-align:center;">string</td><td style="text-align:center;">双方约定的验证密码</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">datetime</td><td style="text-align:center;">事件触发时的时间。格式：Y-m-d H:i:s，如 2011-11-11 11:11:11</td></tr></tbody></table>

例子：

```
(
    [event] => story::create
    [event_from] => web
    [workspace_id] => 10094331
    [id] => 1010094331500623503
    [secret] => 
    [created] => 2017-04-12 09:04:29
)
```


## 更新

更新类的事件有需求更新（story::update）、缺陷更新（bug::update）、任务更新（task::update）、发布评审更新（launchform::update）。当事件发生时， POST 过去的数据如下：

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">类型</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">event</td><td style="text-align:center;">string</td><td style="text-align:center;">事件名</td></tr><tr><td style="text-align:center;">event_from</td><td style="text-align:center;">string</td><td style="text-align:center;">事件触发地方。有web、api</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">string</td><td style="text-align:center;">操作人昵称</td></tr><tr><td style="text-align:center;">event_id</td><td style="text-align:center;">integer</td><td style="text-align:center;">事件ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">integer</td><td style="text-align:center;">新对象ID。需求、缺陷、任务、发布评审的都是19位长度的长ID</td></tr><tr><td style="text-align:center;">old_*</td><td style="text-align:center;">string</td><td style="text-align:center;">会有一系列 old_字段名 的字段，这个是发生变更前的数据。如果需要拉取最新数据，建议配合 tapdapi 使用</td></tr><tr><td style="text-align:center;">change_fields</td><td style="text-align:center;">string</td><td style="text-align:center;">发生变更的字段，以 , 分隔</td></tr><tr><td style="text-align:center;">secret</td><td style="text-align:center;">string</td><td style="text-align:center;">双方约定的验证密码</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">datetime</td><td style="text-align:center;">事件触发时的时间。格式：Y-m-d H:i:s，如 2011-11-11 11:11:11</td></tr></tbody></table>

例子：

```
(
    [event] => task::update
    [event_from] => web
    [workspace_id] => 755
    [id] => 1000000755500588953
    [old_id] => 1000000755500588953
    [old_name] => 批量添加_33
    [old_description_type] => 1
    [old_creator] => jemerychen
    [old_created] => 2017-04-07 16:49:46
    [old_modified] => 2017-04-14 16:20:49
    [old_parent_id] => 0
    [old_children_id] => |
    [old_ancestor_id] => 0
    [old_workspace_id] => 755
    [old_status] => open
    [old_priority] => 
    [old_owner] => darohu;
    [old_begin] => 2017-04-04
    [old_story_id] => 1000000755500623195
    [old_iteration_id] => 0
    [old_effort] => 0
    [old_effort_completed] => 0
    [old_exceed] => 0
    [old_remain] => 0
    [old_progress] => 0
    [old_entity_type] => Task
    [old_attachment_count] => 0
    [change_fields] => due
    [secret] => 
    [created] => 2017-04-14 16:20:52
)
```


## 删除

删除类的事件有需求删除（story::delete）、缺陷删除（bug::delete）、任务删除（task::delete）。当事件发生时， POST 过去的数据如下：

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">类型</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">event</td><td style="text-align:center;">string</td><td style="text-align:center;">事件名</td></tr><tr><td style="text-align:center;">event_from</td><td style="text-align:center;">string</td><td style="text-align:center;">事件触发地方。有web、api</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">string</td><td style="text-align:center;">操作人昵称</td></tr><tr><td style="text-align:center;">event_id</td><td style="text-align:center;">integer</td><td style="text-align:center;">事件ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">integer</td><td style="text-align:center;">新对象ID。需求、缺陷、任务、发布评审的都是19位长度的长ID</td></tr><tr><td style="text-align:center;">secret</td><td style="text-align:center;">string</td><td style="text-align:center;">双方约定的验证密码</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">datetime</td><td style="text-align:center;">事件触发时的时间。格式：Y-m-d H:i:s，如 2011-11-11 11:11:11</td></tr></tbody></table>

例子：

```
(
    [event] => story::delete
    [event_from] => web
    [workspace_id] => 755
    [id] => 1000000755500618867
    [secret] => 
    [created] => 2017-04-13 15:52:58
)
```

