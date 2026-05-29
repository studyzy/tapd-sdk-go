# 获取状态流转时间

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/measure/get_life_times.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下状态流转时间](#获取项目下状态流转时间)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [状态流转时间重要字段说明](#状态流转时间重要字段说明)

# 说明

返回符合查询条件的所有状态流转时间（分页显示，默认一页30条）

# url

`https://api.tapd.cn/life_times`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">entity_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">业务对象ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">业务对象类型</td><td style="text-align:center;">目前type可选值：task,story,bug</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下状态流转时间

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/life_times?entity_id=1000000755854844301&entity_type=story&workspace_id=755'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "LifeTime": {
                "id": "1000000755025642521",
                "workspace_id": "755",
                "entity_type": "story",
                "entity_id": "1000000755854844301",
                "status": "status_23",
                "owner": "peerxu",
                "is_repeated": "0",
                "begin_date": "2020-12-07 11:29:27",
                "end_date": "2020-12-07 11:29:27",
                "time_cost": "0",
                "created": "2020-12-07 11:29:27",
                "operator": null
            }
        },
        {
            "LifeTime": {
                "id": "1000000755025642515",
                "workspace_id": "755",
                "entity_type": "story",
                "entity_id": "1000000755854844301",
                "status": "planning",
                "owner": "peerxu",
                "is_repeated": "0",
                "begin_date": "2020-12-07 11:29:20",
                "end_date": "2020-12-07 11:30:10",
                "time_cost": "1",
                "created": "2020-12-07 11:29:20",
                "operator": "barryhu"
            }
        },
        {
            "LifeTime": {
                "id": "1000000755025642237",
                "workspace_id": "755",
                "entity_type": "story",
                "entity_id": "1000000755854844301",
                "status": "status_23",
                "owner": "peerxu",
                "is_repeated": "0",
                "begin_date": "2020-12-07 10:18:59",
                "end_date": "2020-12-07 11:30:03",
                "time_cost": "2",
                "created": "2020-12-07 10:18:59",
                "operator": "barryhu"
            }
        }
    ],
    "info": "success"
}
```


# 状态流转时间字段说明

## 状态流转时间重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">业务对象类型</td></tr><tr><td style="text-align:center;">entity_id</td><td style="text-align:center;">业务对象ID</td></tr><tr><td style="text-align:center;">begin_date</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间（变更时间）</td></tr><tr><td style="text-align:center;">operator</td><td style="text-align:center;">操作人</td></tr><tr><td style="text-align:center;">is_repeated</td><td style="text-align:center;">是否重复</td></tr><tr><td style="text-align:center;">time_cost</td><td style="text-align:center;">停留时长，单位：小时</td></tr><tr><td style="text-align:center;">change_from</td><td style="text-align:center;">变更来源（status或者owner）</td></tr></tbody></table>
