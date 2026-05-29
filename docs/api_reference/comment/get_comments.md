# 获取评论

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/get_comments.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下评论](#获取项目下评论)
    -   [curl 调用示例](#curl-调用示例)
    -   [返回结果](#返回结果)
-   [获取某条缺陷的评论](#获取某条缺陷的评论)
    -   [curl 调用示例](#curl-调用示例)
    -   [返回结果](#返回结果)
-   [评论重要字段说明](#评论重要字段说明)

# 说明

返回符合查询条件的所有评论（分页显示，默认一页30条）

# url

`https://api.tapd.cn/comments`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">内容</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks 。多个类型间以竖线隔开）</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论所依附的业务对象实体id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后更改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">root_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">根评论ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">reply_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论回复的ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下评论

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Comment": {
                "id": "1010104801048492751",
                "title": "流转状态从 [规划中] 到 [实现中] 添加",
                "description": "<b><i><u>adfasd</u></i></b>",
                "author": "v_xuanfang",
                "entry_type": "stories",
                "entry_id": "1010104801858505231",
                "created": "2020-06-09 10:51:06",
                "modified": "2020-06-09 10:51:06",
                "workspace_id": "10104801"
            }
        },
        {
            "Comment": {
                "id": "1010104801047343797",
                "title": "流转状态从 [规划中] 到 [实现中] 添加",
                "description": "<b><i><u>adfasd</u></i></b>",
                "author": "v_xuanfang",
                "entry_type": "stories",
                "entry_id": "1010104801856464295",
                "created": "2020-04-26 10:44:29",
                "modified": "2020-04-26 10:44:29",
                "workspace_id": "10104801"
            }
        }
    ],
    "info": "success"
}
```


## 获取某条缺陷的评论

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments?entry_type=bug|bug_remark&entry_id=1010104801074085199&limit=2&workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Comment": {
                "id": "1010104801042258591",
                "title": "在状态 [新] 添加",
                "description": "<div>&amp;nbsp;撒大大爱上大</div>",
                "author": "v_xuanfang",
                "entry_type": "bug",
                "entry_id": "1010104801074085199",
                "created": "2019-08-29 14:24:44",
                "modified": "2019-08-29 14:24:44",
                "workspace_id": "10104801"
            }
        }
    ],
    "info": "success"
}
```


# 评论字段说明

## 评论重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评论ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">内容</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">评论人</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks、wiki 。多个类型间以竖线隔开）</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">评论所依附的业务对象实体id</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后更改时间</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">root_id</td><td style="text-align:center;">根评论ID</td></tr><tr><td style="text-align:center;">reply_id</td><td style="text-align:center;">评论回复的ID</td></tr></tbody></table>
