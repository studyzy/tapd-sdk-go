# 批量更新任务

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/batch_update_task.html

# 说明

批量更新任务，返回任务更新后的数据

# url

`https://api.tapd.cn/tasks/batch_update_task`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次最多允许批量更新五十条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">workitems</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">array</td><td style="text-align:center;">要更新对象数组 （下面所有属性数组）</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">任务标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">任务详细描述</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">任务当前处理人</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">操作人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">关联需求的ID</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">所属迭代的ID</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。为了兼容自定义优先级，<code>请使用 priority_label 字段</code>，详情参考：<a href="/document/api-doc/API文档/subject/custom_priority/" class="">如何兼容自定义优先级</a></td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。推荐使用这个字段</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预估工时</td></tr><tr><td style="text-align:center;">auto_complete_effort</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否自动补齐工时，取1时，并且状态流转到 done，就补齐</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签，标签不存在时将自动创建，多个以英文坚线分格</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_这是一个自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/task/get_task_custom_fields_settings.html" class="">获取任务自定义字段配置</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

## 更新任务 1010158231500600385 的状态到 已完成（done）

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10158231&workitems=[{"id"=123123,"name"="测试“}]' 'https://api.tapd.cn/tasks/batch_update_task'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "msg": "batch update success"
    },
    "info": "success"
}
```


# 任务字段说明

## 任务重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">任务标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">任务详细描述</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">任务当前处理人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">关联需求的ID</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">所属迭代的ID</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">优先级</td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">优先级</td></tr><tr><td style="text-align:center;">progress</td><td style="text-align:center;">进度</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">完成时间</td></tr><tr><td style="text-align:center;">effort_completed</td><td style="text-align:center;">完成工时</td></tr><tr><td style="text-align:center;">exceed</td><td style="text-align:center;">超出工时</td></tr><tr><td style="text-align:center;">remain</td><td style="text-align:center;">剩余工时</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">预估工时</td></tr></tbody></table>

## 任务状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">open</td><td style="text-align:center;">未开始</td></tr><tr><td style="text-align:center;">progressing</td><td style="text-align:center;">进行中</td></tr><tr><td style="text-align:center;">done</td><td style="text-align:center;">已完成</td></tr></tbody></table>

## 任务优先级(priority)取值字段说明

为了兼容自定义优先级，`请使用 priority_label 字段`，详情参考：[如何兼容自定义优先级](/document/api-doc/API文档/subject/custom_priority/) 。`下面取值将不再使用`。

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">4</td><td style="text-align:center;">High</td></tr><tr><td style="text-align:center;">3</td><td style="text-align:center;">Middle</td></tr><tr><td style="text-align:center;">2</td><td style="text-align:center;">Low</td></tr><tr><td style="text-align:center;">1</td><td style="text-align:center;">Nice To Have</td></tr></tbody></table>
