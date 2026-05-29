# 获取任务变更次数

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes_count.html

# 说明

计算符合查询条件的任务变更历史数量并返回

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTaskChangesCount
```


# url

`https://api.tapd.cn/task_changes/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回任务变更历史数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">task_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">任务ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人（操作人）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间（变更时间）</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">change_summary</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求变更描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">comment</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">评论</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">changes</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更详细记录</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更的对象类型</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下任务变更次数

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/task_changes/count?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 42
    },
    "info": "success"
}
```


获取任务ID为 1010158231500600411 的任务变更次数

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/task_changes/count?workspace_id=10158231&task_id=1010158231500600411'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 2
    },
    "info": "success"
}
```

