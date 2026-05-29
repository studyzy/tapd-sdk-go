# 获取需求变更历史

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes.html

-   [说明](#说明)
-   [提示](#提示)
-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [url](#url)
-   [支持格式](#支持格式)
-   [HTTP请求方式](#http请求方式)
-   [请求数限制](#请求数限制)
-   [请求参数](#请求参数)
-   [调用示例及返回结果](#调用示例及返回结果)
    -   [获取项目下需求变更历史](#获取项目下需求变更历史)
    -   [获取需求ID为 1010104801056751735 的需求变更历史](#获取需求id为-1010104801056751735-的需求变更历史)
-   [需求变更历史字段说明](#需求变更历史字段说明)
    -   [需求变更历史重要字段说明](#需求变更历史重要字段说明)
-   [最佳实践](#最佳实践)
    -   [做状态停留时长分析](#做状态停留时长分析)
    -   [全量同步变更历史数据](#全量同步变更历史数据)
    -   [增量同步变更历史数据](#增量同步变更历史数据)

## 说明

返回符合查询条件的所有需求变更历史（分页显示，默认一页30条）

## 提示

如果需要做状态停留时长分析，请直接使用 [获取状态流转时间接口](/document/api-doc/API文档/api_reference/measure/get_life_times.html)

## url

`https://api.tapd.cn/story_changes`

## 支持格式

JSON/XML（默认JSON格式）

## HTTP请求方式

GET

## 请求数限制

-   `created` 参数与 `story_id` 参数二选一必填。即 story\_id、created 两个参数至少填写一个
-   其中 created 参数要填写成日期格式，比如 created=2022-02-22，则返回 2022-02-22 这一天的变更历史
-   默认返回 30 条。可通过传 limit 参数设置，最大取 100。也可以传 page 参数翻页

## 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">变更历史id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人（操作人）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间（变更时间）</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">change_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更类型</td><td style="text-align:center;">值范围见文档下方附录1</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求变更描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更的对象类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">change_field</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取变更字段如（status）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">need_parse_changes</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置field_changes字段是否返回（默认取 1。取 0 则不返回）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30，最大取 100</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

## 调用示例及返回结果

### 获取项目下需求变更历史

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/story_changes?workspace_id=10104801'`

#### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemChange": {
                "id": "1010104801027730979",
                "workspace_id": "10104801",
                "app_id": "1",
                "workitem_type_id": "0",
                "creator": "anyechen",
                "created": "2015-06-30 14:28:53",
                "change_summary": "planning",
                "comment": null,
                "changes": "[{\"field\":\"parent_id\",\"value_before\":\"0\",\"value_after\":\"1010104801056751739\"}]",
                "entity_type": "Story",
                "change_type": "",
                "change_type_detail": "",
                "updated": "2024-09-05 22:35:46",
                "change_type_text": "",
                "field_changes": [
                    {
                        "field": "parent_id",
                        "value_before": "0",
                        "value_after": "1010104801056751739",
                        "value_before_parsed": "0",
                        "value_after_parsed": "工具调研",
                        "field_label": "父需求"
                    }
                ],
                "story_id": "1010104801056751735"
            }
        },
        {
            "WorkitemChange": {
                "id": "1010104801027731105",
                "workspace_id": "10104801",
                "app_id": "1",
                "workitem_type_id": "0",
                "creator": "anyechen",
                "created": "2015-06-30 14:31:26",
                "change_summary": "planning",
                "comment": null,
                "changes": "[{\"field\":\"parent_id\",\"value_before\":\"0\",\"value_after\":\"1010104801056751739\"}]",
                "entity_type": "Story",
                "change_type": "",
                "change_type_detail": "",
                "updated": "2024-09-05 22:35:46",
                "change_type_text": "",
                "field_changes": [
                    {
                        "field": "parent_id",
                        "value_before": "0",
                        "value_after": "1010104801056751739",
                        "value_before_parsed": "0",
                        "value_after_parsed": "工具调研",
                        "field_label": "父需求"
                    }
                ],
                "story_id": "1010104801056751727"
            }
        }
    ],
    "info": "success"
}
```


### 获取需求ID为 1010104801056751735 的需求变更历史

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/story_changes?workspace_id=10104801&story_id=1010104801056751735'`

#### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemChange": {
                "id": "1010104801027730979",
                "workspace_id": "10104801",
                "app_id": "1",
                "workitem_type_id": "0",
                "creator": "anyechen",
                "created": "2015-06-30 14:28:53",
                "change_summary": "planning",
                "comment": null,
                "changes": "[{\"field\":\"parent_id\",\"value_before\":\"0\",\"value_after\":\"1010104801056751739\"}]",
                "entity_type": "Story",
                "change_type": "",
                "change_type_detail": "",
                "updated": "2024-09-05 22:35:46",
                "change_type_text": "",
                "field_changes": [
                    {
                        "field": "parent_id",
                        "value_before": "0",
                        "value_after": "1010104801056751739",
                        "value_before_parsed": "0",
                        "value_after_parsed": "工具调研",
                        "field_label": "父需求"
                    }
                ],
                "story_id": "1010104801056751735"
            }
        }
    ],
    "info": "success"
}
```


## 需求变更历史字段说明

### 需求变更历史重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">需求类别ID</td></tr><tr><td style="text-align:center;">updated</td><td style="text-align:center;">更新时间</td></tr><tr><td style="text-align:center;">app_id</td><td style="text-align:center;">检查项</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人（操作人）</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间（变更时间）</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">需求变更描述</td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">评论</td></tr><tr><td style="text-align:center;">changes</td><td style="text-align:center;">变更详细记录</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">变更的对象类型</td></tr><tr><td style="text-align:center;">change_type</td><td style="text-align:center;">变更类型</td></tr><tr><td style="text-align:center;">change_type_text</td><td style="text-align:center;">变更类型结果中文</td></tr><tr><td style="text-align:center;">change_type_detail</td><td style="text-align:center;">api账号</td></tr><tr><td style="text-align:center;">field_changes</td><td style="text-align:center;">变更详细记录</td></tr></tbody></table>

##附录1

<table><thead><tr><th style="text-align:center;">参数值</th><th style="text-align:center;">参数含义</th></tr></thead><tbody><tr><td style="text-align:center;">sync_copy</td><td style="text-align:center;">同步复制联动</td></tr><tr><td style="text-align:center;">story_status_relation</td><td style="text-align:center;">父子需求联动</td></tr><tr><td style="text-align:center;">story_task_relation</td><td style="text-align:center;">需求任务联动</td></tr><tr><td style="text-align:center;">api</td><td style="text-align:center;">API变更</td></tr><tr><td style="text-align:center;">smart_commit</td><td style="text-align:center;">Smart Commit触发</td></tr><tr><td style="text-align:center;">auto_task</td><td style="text-align:center;">自动化任务触发</td></tr><tr><td style="text-align:center;">auto_workflow</td><td style="text-align:center;">自动化工作流触发</td></tr><tr><td style="text-align:center;">manual_update</td><td style="text-align:center;">手动变更</td></tr><tr><td style="text-align:center;">import_update</td><td style="text-align:center;">导入更新</td></tr><tr><td style="text-align:center;">code_change</td><td style="text-align:center;">代码变更</td></tr><tr><td style="text-align:center;">status_delete</td><td style="text-align:center;">状态删除</td></tr><tr><td style="text-align:center;">exit_workspace</td><td style="text-align:center;">退出项目触发</td></tr><tr><td style="text-align:center;">link_update</td><td style="text-align:center;">更新关联</td></tr><tr><td style="text-align:center;">link_create</td><td style="text-align:center;">创建关联</td></tr><tr><td style="text-align:center;">link_delete</td><td style="text-align:center;">删除关联</td></tr><tr><td style="text-align:center;">create_story_from_copy</td><td style="text-align:center;">复制创建</td></tr></tbody></table>

## 最佳实践

### 做状态停留时长分析

如果需要做状态停留时长分析，请直接使用 [获取状态流转时间接口](/document/api-doc/API文档/api_reference/measure/get_life_times.html)。

### 全量同步变更历史数据

如果需要同步项目全量状态变更历史数据，可以按照如下步骤：

1.  使用获取需求数据接口，分页获取需求的ID。比如 `https://api.tapd.cn/stories?workspace_id=项目ID&limit=30&page=1&fields=id`
2.  根据步骤1获取到的需求ID，加到 story\_id 参数里面，分页获取变更历史数据。比如 `https://api.tapd.cn/story_changes?workspace_id=项目ID&story_id=需求ID1,需求ID2,需求3&limit=30&page=1`
3.  重复步骤1、步骤2直到拿到所有需求的变更历史

### 增量同步变更历史数据

因为变更历史数据不会变，只增不减。所以可以传递 created 参数来增量获取变更历史数据。参考这样：

1.  传递 created=日期 。比如 `https://api.tapd.cn/story_changes?workspace_id=项目ID&created=2022-02-22&limit=30&page=1` 来获取某天的变更历史数据
