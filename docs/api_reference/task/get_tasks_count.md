# 获取任务数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks_count.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的任务数量](#获取项目下的任务数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [任务状态(status)取值字段说明](#任务状态-status-取值字段说明)
-   [任务优先级(priority)取值字段说明](#任务优先级-priority-取值字段说明)

# 说明

计算符合查询条件的任务数量并返回

# url

`https://api.tapd.cn/tasks/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回任务数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询、模糊匹配</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">任务标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">任务详细描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签查询</td><td style="text-align:center;">支持枚举查询，5月30号上线</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">任务当前处理人</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">关联需求的ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">所属迭代的ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。为了兼容自定义优先级，<code>请使用 priority_label 字段</code>，详情参考：<a href="/document/api-doc/API文档/subject/custom_priority/" class="">如何兼容自定义优先级</a></td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。推荐使用这个字段</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">progress</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">进度</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">完成时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">effort_completed</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">完成工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">exceed</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">超出工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">remain</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">剩余工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预估工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/task/get_task_custom_fields_settings.html" class="">获取任务自定义字段配置</a> 获取</td><td style="text-align:center;">支持枚举查询</td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的任务数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tasks/count?workspace_id=10158231'`

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


## 任务状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">open</td><td style="text-align:center;">未开始</td></tr><tr><td style="text-align:center;">progressing</td><td style="text-align:center;">进行中</td></tr><tr><td style="text-align:center;">done</td><td style="text-align:center;">已完成</td></tr></tbody></table>

## 任务优先级(priority)取值字段说明

为了兼容自定义优先级，`请使用 priority_label 字段`，详情参考：[如何兼容自定义优先级](/document/api-doc/API文档/subject/custom_priority/) 。`下面取值将不再使用`。

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">4</td><td style="text-align:center;">High</td></tr><tr><td style="text-align:center;">3</td><td style="text-align:center;">Middle</td></tr><tr><td style="text-align:center;">2</td><td style="text-align:center;">Low</td></tr><tr><td style="text-align:center;">1</td><td style="text-align:center;">Nice To Have</td></tr></tbody></table>
