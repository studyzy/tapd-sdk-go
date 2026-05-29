# 获取工作项动态

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/mini_api_reference/mini_item/get_mini_item_changes.html

-   [说明](#说明)
-   [url](#url)
-   [支持格式](#支持格式)
-   [HTTP请求方式](#http请求方式)
-   [请求数限制](#请求数限制)
-   [请求参数](#请求参数)
-   [调用示例及返回结果](#调用示例及返回结果)
    -   [获取空间下工作项动态](#获取空间下工作项动态)
    -   [以创建时间倒序获取工作项ID为 1069993260856110917 的工作项 于2023-07-18日创建的动态](#以创建时间倒序获取工作项id为-1069993260856110917-的工作项-于2023-07-18日创建的动态)
-   [工作项动态字段说明](#工作项动态字段说明)
    -   [工作项动态重要字段说明](#工作项动态重要字段说明)
-   [附录1](#附录1)
-   [最佳实践](#最佳实践)
    -   [全量同步动态数据](#全量同步动态数据)
    -   [增量同步动态数据](#增量同步动态数据)

## 说明

返回符合查询条件的所有工作项动态（分页显示，默认一页30条）

## url

`https://api.tapd.cn/mini_item_changes`

## 支持格式

JSON/XML（默认JSON格式）

## HTTP请求方式

GET

## 请求数限制

-   `created` 参数与 `mini_item_id` 参数二选一必填。即 mini\_item\_id、created 两个参数至少填写一个
-   其中 created 参数要填写成日期格式，比如 created=2023-07-18，则返回 2023-07-18 这一天的动态
-   默认返回 30 条。可通过传 limit 参数设置，最大取 100。也可以传 page 参数翻页

## 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">动态id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">mini_item_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">工作项id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">空间ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人（操作人）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间（变更时间）</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">change_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更类型</td><td style="text-align:center;">值范围见文档下方附录1</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">工作项变更描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">change_field</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取变更字段如（status）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">need_parse_changes</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置field_changes字段是否返回（默认取 1。取 0 则不返回）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

## 调用示例及返回结果

### 获取空间下工作项动态

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_item_changes?workspace_id=69993260'`

#### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_item_changes?workspace_id=69993260'`

#### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemChange": {
                "id": "1069993260086183599",
                "workspace_id": "69993260",
                "workitem_type_id": "0",
                "creator": "orangecyang",
                "created": "2023-07-18 16:02:56",
                "change_summary": "update_mini_item",
                "comment": null,
                "changes":"[{\"field\":\"progress_manual\",\"value_before\":\"0\",\"value_after\":100,\"field_name\":\"\进\度\",\"field_type\":\"integer\"}",
                "entity_type": "MiniItem",
                "change_type": "manual_update",
                "change_type_text": "手动变更",
                "field_changes": [
                    {
                        "field": "progress_manual",
                        "value_before": "0",
                        "value_after": 100,
                        "field_name": "进度",
                        "field_type": "integer",
                        "value_before_parsed": "0",
                        "value_after_parsed": "100",
                        "field_label": "进度"
                    }
                ],
                "mini_item_id": "1069993260856110919"
            }
        },
        {
            "WorkitemChange": {
                "id": "1069993260086183603",
                "workspace_id": "69993260",
                "workitem_type_id": "0",
                "creator": "orangecyang",
                "created": "2023-07-18 16:03:19",
                "change_summary": "update_mini_item",
                "comment": null,
                "changes":"[{\"field\":\"owner\",\"value_before\":\"\",\"value_after\":\"orangecyang;\",\"field_name\":\"\处\理\人\",\"field_type\":\"user_chooser\"}",
                "entity_type": "MiniItem",
                "change_type": "manual_update",
                "change_type_text": "手动变更",
                "field_changes": [
                    {
                        "field": "owner",
                        "value_before": "",
                        "value_after": "orangecyang;",
                        "field_name": "处理人",
                        "field_type": "user_chooser",
                        "value_before_parsed": "--",
                        "value_after_parsed": "orangecyang;",
                        "field_label": "处理人"
                    }
                ],
                "mini_item_id": "1069993260856110919"
            }
        }
    ],
    "info": "success"
}
```


### 以创建时间倒序获取工作项ID为 1069993260856110917 的工作项 于2023-07-18日创建的动态

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/mini_item_changes?workspace_id=69993260&created=2023-07-18&mini_item_id=1069993260856110917&order=created desc'`

#### curl 使用 OAuth Access Token 鉴权调用示例

`curl -H 'Authorization: Bearer ACCESS_TOKEN' 'https://api.tapd.cn/mini_item_changes?workspace_id=69993260&created=2023-07-18&mini_item_id=1069993260856110917&order=created desc'`

#### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemChange": {
                "id": "1069993260086183749",
                "workspace_id": "69993260",
                "workitem_type_id": "0",
                "creator": "orangecyang",
                "created": "2023-07-18 16:40:55",
                "change_summary": "update_mini_item",
                "comment": null,
                "changes":"[{\"field\":\"name\",\"value_before\":\"mini_item_created_by_api\",\"value_after\":\"mini_item_created_by_api1\",\"field_name\":\"\标\题\",\"field_type\":\"input\"}",
                "entity_type": "MiniItem",
                "change_type": "manual_update",
                "change_type_text": "手动变更",
                "field_changes": [
                    {
                        "field": "name",
                        "value_before": "mini_item_created_by_api",
                        "value_after": "mini_item_created_by_api1",
                        "field_name": "标题",
                        "field_type": "input",
                        "value_before_parsed": "mini_item_created_by_api",
                        "value_after_parsed": "mini_item_created_by_api1",
                        "field_label": "标题"
                    }
                ],
                "mini_item_id": "1069993260856110917"
            }
        },
        {
            "WorkitemChange": {
                "id": "1069993260086183745",
                "workspace_id": "69993260",
                "workitem_type_id": "0",
                "creator": "orangecyang",
                "created": "2023-07-18 16:36:44",
                "change_summary": "update_mini_item",
                "comment": null,
                "changes":"[{\"field\":\"name\",\"value_before\":\"mini_item_created_by_api1\",\"value_after\":\"mini_item_created_by_api\",\"field_name\":\"\标\题\",\"field_type\":\"input\"}",
                "entity_type": "MiniItem",
                "change_type": "manual_update",
                "change_type_text": "手动变更",
                "field_changes": [
                    {
                        "field": "name",
                        "value_before": "mini_item_created_by_api1",
                        "value_after": "mini_item_created_by_api",
                        "field_name": "标题",
                        "field_type": "input",
                        "value_before_parsed": "mini_item_created_by_api1",
                        "value_after_parsed": "mini_item_created_by_api",
                        "field_label": "标题"
                    }
                ],
                "mini_item_id": "1069993260856110917"
            }
        }
    ],
    "info": "success"
}
```


## 工作项动态字段说明

### 工作项动态重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">空间ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人（操作人）</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间（变更时间）</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">工作项变更描述</td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">评论</td></tr><tr><td style="text-align:center;">changes</td><td style="text-align:center;">变更详细记录</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">变更的对象类型</td></tr><tr><td style="text-align:center;">change_type</td><td style="text-align:center;">变更类型</td></tr><tr><td style="text-align:center;">change_type_text</td><td style="text-align:center;">变更类型结果中文</td></tr><tr><td style="text-align:center;">field_changes</td><td style="text-align:center;">变更详细记录</td></tr></tbody></table>

## 附录1

<table><thead><tr><th style="text-align:center;">参数值</th><th style="text-align:center;">参数含义</th></tr></thead><tbody><tr><td style="text-align:center;">api</td><td style="text-align:center;">API变更</td></tr><tr><td style="text-align:center;">manual_update</td><td style="text-align:center;">手动变更</td></tr></tbody></table>

## 最佳实践

### 全量同步动态数据

如果需要同步空间全量状态动态数据，可以按照如下步骤：

1.  使用获取工作项数据接口，分页获取工作项的ID。比如 `https://api.tapd.cn/mini_items?workspace_id=空间ID&limit=30&page=1&fields=id`
2.  根据步骤1获取到的工作项ID，加到 mini\_item\_id 参数里面，分页获取动态数据。比如 `https://api.tapd.cn/mini_item_changes?workspace_id=空间ID&mini_item_id=工作项ID1,工作项ID2,工作项3&limit=30&page=1`
3.  重复步骤1、步骤2直到拿到所有工作项的动态

### 增量同步动态数据

因为动态数据不会变，只增不减。所以可以传递 created 参数来增量获取动态数据。参考这样：

1.  传递 created=日期 。比如 `https://api.tapd.cn/mini_item_changes?workspace_id=空间ID&created=2022-02-22&limit=30&page=1` 来获取某天的动态数据
