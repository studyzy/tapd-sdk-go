# 获取回收站的缺陷

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_removed_bugs.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [回收站重要字段说明](#回收站重要字段说明)

# 说明

返回符合查询条件的所有回收站中的缺陷

# url

`https://api.tapd.cn/bugs/get_removed_bugs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">缺陷ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">删除时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">include_all</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">取 1 会返回所有删除的缺陷，包括 移动、合并、删除 的</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/get_removed_bugs?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "RemovedBug": {
                "id": "1100000755500695186",
                "title": "标题呀阿斯蒂芬你不是分别就开始放缓觉得斯芬克斯地方卡积分看到as艰苦地方阿克苏",
                "reporter": "gobichen",
                "created": "2021-04-22 21:29:41",
                "operation_user": "v_tingtdong",
                "modified": "2021-04-23 11:04:59",
                "removed_comment": "{\"action\":\"delete\"}",
                "type": "delete",
                "new_bug_url": ""
            }
        },
        {
            "RemovedBug": {
                "id": "1100000755500695184",
                "title": "标题呀",
                "reporter": "gobichen",
                "created": "2021-04-22 21:17:13",
                "operation_user": "v_tingtdong",
                "modified": "2021-04-23 10:25:05",
                "removed_comment":"{\"action\":\"merge\",\"comment\":\"\已\经\被\合\并\到 Bug#500695186, \点\击<a href=\\\"http:\\/\\/tiger.oa.com\\/755\\/bugtrace\\/bugs\\/view?bug_id=1100000755500695186\\\">\这\里<\\/a>\查\看\"}",
                "type": "merge",
                "new_bug_url": "http://tiger.oa.com/755/bugtrace/bugs/view?bug_id=1100000755500695186"
            }
        }
    ],
    "info": "success"
}
```


# 回收站字段说明

## 回收站重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">缺陷ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">reporter</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">operation_user</td><td style="text-align:center;">删除人</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">删除时间</td></tr><tr><td style="text-align:center;">removed_comment</td><td style="text-align:center;">删除附加信息</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">删除操作类型，取值：delete 删除、move 移动、merge 合并</td></tr><tr><td style="text-align:center;">new_bug_url</td><td style="text-align:center;">当 type 为 move、merge 时，这个字段会指向新的缺陷链接</td></tr></tbody></table>
