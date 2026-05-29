# 获取任务变更历史

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes.html

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
    -   [获取项目下任务变更历史](#获取项目下任务变更历史)
-   [获取任务ID为 1010158231500600411 的任务变更历史](#获取任务id为-1010158231500600411-的任务变更历史)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [任务变更历史字段说明](#任务变更历史字段说明)
    -   [任务变更历史重要字段说明](#任务变更历史重要字段说明)
-   [最佳实践](#最佳实践)
    -   [做状态停留时长分析](#做状态停留时长分析)
    -   [全量同步变更历史数据](#全量同步变更历史数据)
    -   [增量同步变更历史数据](#增量同步变更历史数据)

## 说明

返回符合查询条件的所有任务变更历史（分页显示，默认一页30条）

## 提示

如果需要做状态停留时长分析，请直接使用 [获取状态流转时间接口](/document/api-doc/API文档/api_reference/measure/get_life_times.html)

## url

`https://api.tapd.cn/task_changes`

## 支持格式

JSON/XML（默认JSON格式）

## HTTP请求方式

GET

## 请求数限制

-   `created` 参数与 `task_id` 参数二选一必填。即 task\_id、created 两个参数至少填写一个
-   其中 created 参数要填写成日期格式，比如 created=2022-02-22，则返回 2022-02-22 这一天的变更历史
-   默认返回 30 条。可通过传 limit 参数设置，最大取 100。也可以传 page 参数翻页

## 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">task_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">任务ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人（操作人）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间（变更时间）</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求变更描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">changes</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更详细记录</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更的对象类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">need_parse_changes</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置field_changes字段是否返回（默认取 1。取 0 则不返回）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

## 调用示例及返回结果

### 获取项目下任务变更历史

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/task_changes?workspace_id=10158231&limit=2'`

#### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemChange": {
                "id": "1010158231046789917",
                "workspace_id": "10158231",
                "creator": "anyechen",
                "created": "2020-02-28 00:55:03",
                "change_summary": "open",
                "comment": null,
                "changes": "[{\"field\":\"effort\",\"value_before\":\"0\",\"value_after\":\"10\"}]",
                "entity_type": "Task",
                "task_id": "1010158231500600411"
            }
        },
        {
            "WorkitemChange": {
                "id": "1010158231046789915",
                "workspace_id": "10158231",
                "creator": "anyechen",
                "created": "2020-02-28 00:54:56",
                "change_summary": null,
                "comment": null,
                "changes": "[{\"field\":\"status\",\"value_before\":null,\"value_after\":\"open\"}]",
                "entity_type": "Task",
                "task_id": "1010158231500600411"
            }
        }
    ],
    "info": "success"
}
```


## 获取任务ID为 1010158231500600411 的任务变更历史

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/task_changes?workspace_id=10158231&task_id=1010158231500600411'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemChange": {
                "id": "1010158231046789917",
                "workspace_id": "10158231",
                "creator": "anyechen",
                "created": "2020-02-28 00:55:03",
                "change_summary": "open",
                "comment": null,
                "changes": "[{\"field\":\"effort\",\"value_before\":\"0\",\"value_after\":\"10\"}]",
                "entity_type": "Task",
                "task_id": "1010158231500600411"
            }
        },
        {
            "WorkitemChange": {
                "id": "1010158231046789915",
                "workspace_id": "10158231",
                "creator": "anyechen",
                "created": "2020-02-28 00:54:56",
                "change_summary": null,
                "comment": null,
                "changes": "[{\"field\":\"status\",\"value_before\":null,\"value_after\":\"open\"}]",
                "entity_type": "Task",
                "task_id": "1010158231500600411"
            }
        }
    ],
    "info": "success"
}
```


## 任务变更历史字段说明

### 任务变更历史重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人（操作人）</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间（变更时间）</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">需求变更描述</td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">评论</td></tr><tr><td style="text-align:center;">changes</td><td style="text-align:center;">变更详细记录</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">变更的对象类型</td></tr></tbody></table>

## 最佳实践

### 做状态停留时长分析

如果需要做状态停留时长分析，请直接使用 [获取状态流转时间接口](/document/api-doc/API文档/api_reference/measure/get_life_times.html)。

### 全量同步变更历史数据

如果需要同步项目全量状态变更历史数据，可以按照如下步骤：

1.  使用获取任务数据接口，分页获取任务的ID。比如 `https://api.tapd.cn/tasks?workspace_id=项目ID&limit=30&page=1&fields=id`
2.  根据步骤1获取到的任务ID，加到 task\_id 参数里面，分页获取变更历史数据。比如 `https://api.tapd.cn/task_changes?workspace_id=项目ID&task_id=任务ID1,任务ID2,任务3&limit=30&page=1`
3.  重复步骤1、步骤2直到拿到所有任务的变更历史

### 增量同步变更历史数据

因为变更历史数据不会变，只增不减。所以可以传递 created 参数来增量获取变更历史数据。参考这样：

1.  传递 created=日期 。比如 `https://api.tapd.cn/task_changes?workspace_id=项目ID&created=2022-02-22&limit=30&page=1` 来获取某天的变更历史数据
