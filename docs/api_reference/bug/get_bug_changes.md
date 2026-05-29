# 获取缺陷变更历史

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_changes.html

-   [说明](#说明)
-   [提示](#提示)
-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [SDK 方法名](#sdk-方法名)
-   [url](#url)
-   [支持格式](#支持格式)
-   [HTTP请求方式](#http请求方式)
-   [请求数限制](#请求数限制)
-   [请求参数](#请求参数)
-   [调用示例及返回结果](#调用示例及返回结果)
    -   [获取项目下缺陷变更历史](#获取项目下缺陷变更历史)
-   [缺陷变更历史字段说明](#缺陷变更历史字段说明)
    -   [缺陷变更历史重要字段说明](#缺陷变更历史重要字段说明)
-   [最佳实践](#最佳实践)
    -   [做状态停留时长分析](#做状态停留时长分析)
    -   [全量同步变更历史数据](#全量同步变更历史数据)
    -   [增量同步变更历史数据](#增量同步变更历史数据)

## 说明

返回符合查询条件的所有缺陷变更历史（分页显示，默认一页30条）

## 提示

如果需要做状态停留时长分析，请直接使用 [获取状态流转时间接口](/document/api-doc/API文档/api_reference/measure/get_life_times.html)

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getBugChanges
```


## url

`https://api.tapd.cn/bug_changes`

## 支持格式

JSON/XML（默认JSON格式）

## HTTP请求方式

GET

## 请求数限制

-   `created` 参数与 `bug_id` 参数二选一必填。即 bug\_id、created 两个参数至少填写一个
-   其中 created 参数要填写成日期格式，比如 created=2022-02-22，则返回 2022-02-22 这一天的变更历史
-   默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

## 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持按天查询。比如 created=2022-02-22</td></tr><tr><td style="text-align:center;">bug_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">缺陷ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更字段</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">old_value</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更前</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">new_value</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更后</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">备注</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">include_add_bug</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回创建缺陷的记录（值传1）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

## 调用示例及返回结果

### 获取项目下缺陷变更历史

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bug_changes?workspace_id=10158231'`

#### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "BugChange": {
                "id": "10101582315000015921",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "severity",
                "old_value": "serious",
                "new_value": "normal",
                "memo": null,
                "created": "2019-06-26 20:48:52",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015919",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "severity",
                "old_value": "fatal",
                "new_value": "serious",
                "memo": null,
                "created": "2019-06-26 20:47:24",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015917",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "priority",
                "old_value": "urgent",
                "new_value": "high",
                "memo": null,
                "created": "2019-06-26 20:47:21",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015915",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "current_owner",
                "old_value": null,
                "new_value": "anyechen;",
                "memo": null,
                "created": "2019-06-26 20:47:18",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015913",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "severity",
                "old_value": "normal",
                "new_value": "fatal",
                "memo": null,
                "created": "2019-06-26 20:47:02",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015911",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "priority",
                "old_value": "medium",
                "new_value": "urgent",
                "memo": null,
                "created": "2019-06-26 20:46:59",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "1010158231606851987",
                "bug_id": "1010158231500628817",
                "author": "anyechen",
                "field": "severity",
                "old_value": "serious",
                "new_value": "prompt",
                "memo": null,
                "created": "2018-01-12 14:45:29",
                "workspace_id": "10158231"
            }
        }
    ],
    "info": "success"
}
```


获取缺陷ID为 1010158231500628815 的变更历史

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bug_changes?workspace_id=10158231&bug_id=1010158231500628815'`

#### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "BugChange": {
                "id": "10101582315000015921",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "severity",
                "old_value": "serious",
                "new_value": "normal",
                "memo": null,
                "created": "2019-06-26 20:48:52",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015919",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "severity",
                "old_value": "fatal",
                "new_value": "serious",
                "memo": null,
                "created": "2019-06-26 20:47:24",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015917",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "priority",
                "old_value": "urgent",
                "new_value": "high",
                "memo": null,
                "created": "2019-06-26 20:47:21",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015915",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "current_owner",
                "old_value": null,
                "new_value": "anyechen;",
                "memo": null,
                "created": "2019-06-26 20:47:18",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015913",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "severity",
                "old_value": "normal",
                "new_value": "fatal",
                "memo": null,
                "created": "2019-06-26 20:47:02",
                "workspace_id": "10158231"
            }
        },
        {
            "BugChange": {
                "id": "10101582315000015911",
                "bug_id": "1010158231500628815",
                "author": "anyechen",
                "field": "priority",
                "old_value": "medium",
                "new_value": "urgent",
                "memo": null,
                "created": "2019-06-26 20:46:59",
                "workspace_id": "10158231"
            }
        }
    ],
    "info": "success"
}
```


## 缺陷变更历史字段说明

### 缺陷变更历史重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">bug_id</td><td style="text-align:center;">缺陷ID</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">变更人</td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">变更字段</td></tr><tr><td style="text-align:center;">old_value</td><td style="text-align:center;">变更前</td></tr><tr><td style="text-align:center;">new_value</td><td style="text-align:center;">变更后</td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">备注</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr></tbody></table>

## 最佳实践

### 做状态停留时长分析

如果需要做状态停留时长分析，请直接使用 [获取状态流转时间接口](/document/api-doc/API文档/api_reference/measure/get_life_times.html)。

### 全量同步变更历史数据

如果需要同步项目全量状态变更历史数据，可以按照如下步骤：

1.  使用获取缺陷数据接口，分页获取缺陷的ID。比如 `https://api.tapd.cn/bugs?workspace_id=项目ID&limit=30&page=1&fields=id`
2.  根据步骤1获取到的缺陷ID，加到 bug\_id 参数里面，分页获取变更历史数据。比如 `https://api.tapd.cn/bug_changes?workspace_id=项目ID&bug_id=缺陷ID1,缺陷ID2,缺陷3&limit=30&page=1`
3.  重复步骤1、步骤2直到拿到所有缺陷的变更历史

### 增量同步变更历史数据

因为变更历史数据不会变，只增不减。所以可以传递 created 参数来增量获取变更历史数据。参考这样：

1.  传递 created=日期 。比如 `https://api.tapd.cn/bug_changes?workspace_id=项目ID&created=2022-02-22&limit=30&page=1` 来获取某天的变更历史数据。
