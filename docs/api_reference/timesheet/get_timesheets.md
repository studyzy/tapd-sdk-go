# 获取工时花费

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/get_timesheets.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的花费工时信息](#获取项目下的花费工时信息)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取项目下 2020年5月5日 的花费工时信息](#获取项目下-2020年5月5日-的花费工时信息)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取需求ID为 1120003271001000198 的花费工时信息](#获取需求id为-1120003271001000198-的花费工时信息)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取项目下某成员的花费工时时间](#获取项目下某成员的花费工时时间)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [花费工时重要字段说明](#花费工时重要字段说明)

# 说明

返回符合查询条件的所有花费工时（分页显示，默认一页30条）

# url

`https://api.tapd.cn/timesheets`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">对象类型，如story、task、bug等</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">对象ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">timespent</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">spentdate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">花费日期</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">include_parent_story_timesheet</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">值=0不返回父需求的花费</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">花费描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">is_delete</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否已删除。默认取 0，不返回已删除的工时记录。取 1 可以返回已删除的记录</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workflow_step</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">工作流节点原名</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">relation_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">关联类型，entity或workflow</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的花费工时信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/timesheets?workspace_id=10158231&limit=3'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Timesheet": {
                "id": "1010158231001168997",
                "entity_type": "story",
                "entity_id": "1010158231500709717",
                "timespent": "8",
                "spentdate": "2020-05-05",
                "owner": "anyechen",
                "created": "2020-05-06 19:32:35",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        },
        {
            "Timesheet": {
                "id": "1010158231001168995",
                "entity_type": "story",
                "entity_id": "1010158231500709717",
                "timespent": "2",
                "spentdate": "2020-05-06",
                "owner": "anyechen",
                "created": "2020-05-06 19:32:05",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        },
        {
            "Timesheet": {
                "id": "1010158231001168993",
                "entity_type": "task",
                "entity_id": "1010158231500606309",
                "timespent": "3",
                "spentdate": "2020-05-06",
                "owner": "anyechen",
                "created": "2020-05-06 19:22:07",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        }
    ],
    "info": "success"
}
```


## 获取项目下 2020年5月5日 的花费工时信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/timesheets?workspace_id=10158231&spentdate=2020-05-05'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Timesheet": {
                "id": "1010158231001168997",
                "entity_type": "story",
                "entity_id": "1010158231500709717",
                "timespent": "8",
                "spentdate": "2020-05-05",
                "owner": "anyechen",
                "created": "2020-05-06 19:32:35",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        },
        {
            "Timesheet": {
                "id": "1010158231001168987",
                "entity_type": "task",
                "entity_id": "1010158231500606309",
                "timespent": "2",
                "spentdate": "2020-05-05",
                "owner": "anyechen",
                "created": "2020-05-06 18:33:31",
                "workspace_id": "10158231",
                "memo": null
            }
        },
        {
            "Timesheet": {
                "id": "1010158231001168971",
                "entity_type": "task",
                "entity_id": "1010158231500606305",
                "timespent": "1",
                "spentdate": "2020-05-05",
                "owner": "anyechen",
                "created": "2020-05-06 12:57:08",
                "workspace_id": "10158231",
                "memo": "xxx"
            }
        }
    ],
    "info": "success"
}
```


## 获取需求ID为 1120003271001000198 的花费工时信息

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/timesheets?workspace_id=10158231&entity_type=story&entity_id=1010158231500709717'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Timesheet": {
                "id": "1010158231001168997",
                "entity_type": "story",
                "entity_id": "1010158231500709717",
                "timespent": "8",
                "spentdate": "2020-05-05",
                "owner": "anyechen",
                "created": "2020-05-06 19:32:35",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        },
        {
            "Timesheet": {
                "id": "1010158231001168995",
                "entity_type": "story",
                "entity_id": "1010158231500709717",
                "timespent": "2",
                "spentdate": "2020-05-06",
                "owner": "anyechen",
                "created": "2020-05-06 19:32:05",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        }
    ],
    "info": "success"
}
```


## 获取项目下某成员的花费工时时间

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/timesheets?workspace_id=10158231&owner=anyechen&limit=3'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Timesheet": {
                "id": "1010158231001168997",
                "entity_type": "story",
                "entity_id": "1010158231500709717",
                "timespent": "8",
                "spentdate": "2020-05-05",
                "owner": "anyechen",
                "created": "2020-05-06 19:32:35",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        },
        {
            "Timesheet": {
                "id": "1010158231001168995",
                "entity_type": "story",
                "entity_id": "1010158231500709717",
                "timespent": "2",
                "spentdate": "2020-05-06",
                "owner": "anyechen",
                "created": "2020-05-06 19:32:05",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        },
        {
            "Timesheet": {
                "id": "1010158231001168993",
                "entity_type": "task",
                "entity_id": "1010158231500606309",
                "timespent": "3",
                "spentdate": "2020-05-06",
                "owner": "anyechen",
                "created": "2020-05-06 19:22:07",
                "workspace_id": "10158231",
                "memo": "hey"
            }
        }
    ],
    "info": "success"
}
```


# 花费工时字段说明

## 花费工时重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">对象类型，如story、task、bug等</td></tr><tr><td style="text-align:center;">entity_id</td><td style="text-align:center;">对象ID</td></tr><tr><td style="text-align:center;">timespent</td><td style="text-align:center;">花费工时</td></tr><tr><td style="text-align:center;">spentdate</td><td style="text-align:center;">花费日期</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">花费创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">花费描述</td></tr><tr><td style="text-align:center;">is_delete</td><td style="text-align:center;">是否已删除</td></tr><tr><td style="text-align:center;">workflow_step</td><td style="text-align:center;">工作流节点原名</td></tr><tr><td style="text-align:center;">relation_type</td><td style="text-align:center;">关联类型，entity（直接关联业务对象）或workflow（关联工作流节点）</td></tr></tbody></table>
