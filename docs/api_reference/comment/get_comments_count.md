# 获取评论数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/get_comments_count.html

# 说明

计算符合查询条件的评论数量并返回

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getCommentsCount
```


# url

`https://api.tapd.cn/comments/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回评论数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">内容</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks 。多个类型间以竖线隔开）</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">entry_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">评论所依附的业务对象实体id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后更改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下评论数量

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments/count?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 61
    },
    "info": "success"
}
```


## 获取某条缺陷评论的数量

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/comments/count?entry_type=bug&amp;entry_id=1010104801074085199&amp;workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 1
    },
    "info": "success"
}
```

