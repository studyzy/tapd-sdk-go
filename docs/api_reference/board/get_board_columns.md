# 获取看板板块

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/get_board_columns.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [返回结果](#返回结果)
-   [板块重要字段说明](#板块重要字段说明)

# 说明

返回符合查询条件的所有板块（分页显示，默认一页30条）

# url

`https://api.tapd.cn/board_columns`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">板块ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">板块名称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">board_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">看板ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' 'https://api.tapd.cn/board_columns?workspace_id=10104801'`

## 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Column": {
                "id": "1010104801000032321",
                "name": "已上线",
                "board_id": "1010104801000007781",
                "status": "open",
                "sort": "7",
                "created": "2018-07-19 16:46:04",
                "creator": "TAPD",
                "workspace_id": "10104801"
            }
        },
        {
            "Column": {
                "id": "1010104801000032319",
                "name": "开发中",
                "board_id": "1010104801000007781",
                "status": "open",
                "sort": "6",
                "created": "2018-07-19 16:46:04",
                "creator": "TAPD",
                "workspace_id": "10104801"
            }
        },
        {
            "Column": {
                "id": "1010104801000032317",
                "name": "设计中",
                "board_id": "1010104801000007781",
                "status": "open",
                "sort": "5",
                "created": "2018-07-19 16:46:04",
                "creator": "TAPD",
                "workspace_id": "10104801"
            }
        },
        {
            "Column": {
                "id": "1010104801000032315",
                "name": "需求排期",
                "board_id": "1010104801000007781",
                "status": "open",
                "sort": "4",
                "created": "2018-07-19 16:46:04",
                "creator": "TAPD",
                "workspace_id": "10104801"
            }
        },
        {
            "Column": {
                "id": "1010104801000032313",
                "name": "需求评审",
                "board_id": "1010104801000007781",
                "status": "open",
                "sort": "3",
                "created": "2018-07-19 16:46:04",
                "creator": "TAPD",
                "workspace_id": "10104801"
            }
        }
    ],
    "info": "success"
}
```


# 板块字段说明

## 板块重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">板块ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">板块名称</td></tr><tr><td style="text-align:center;">board_id</td><td style="text-align:center;">看板ID</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">计数</td></tr></tbody></table>
