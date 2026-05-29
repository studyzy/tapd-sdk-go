# 批量更新需求

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/batch_update_story.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [更新需求 1010104801125341253 的优先级为高，处理人为 anyechen](#更新需求-1010104801125341253-的优先级为高-处理人为-anyechen)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

批量更新需求，返回需求更新后的数据

# url

`https://api.tapd.cn/stories/batch_update_story`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次最多允许批量更新五十条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">workitems</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">array</td><td style="text-align:center;">要更新对象数组 （下面所有属性数组）</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。为了兼容自定义优先级，<code>请使用 priority_label 字段</code>，详情参考：<a href="/document/api-doc/API文档/subject/custom_priority/" class="">如何兼容自定义优先级</a></td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。推荐使用这个字段</td></tr><tr><td style="text-align:center;">business_value</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">业务价值</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态。需求当前使用并行工作流时，将按状态重置来更新节点，进行中节点变更参考<a href="/document/api-doc/API文档/api_reference/story/update_story_step_status.html" class="">节点完成接口</a></td></tr><tr><td style="text-align:center;">v_status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态(支持传入中文状态名称)</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本</td></tr><tr><td style="text-align:center;">module</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">模块</td></tr><tr><td style="text-align:center;">test_focus</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试重点</td></tr><tr><td style="text-align:center;">size</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">规模</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">developer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">开发人员</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预估工时</td></tr><tr><td style="text-align:center;">effort_completed</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">完成工时</td></tr><tr><td style="text-align:center;">remain</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">剩余工时</td></tr><tr><td style="text-align:center;">exceed</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">超出工时</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求分类</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">发布计划</td></tr><tr><td style="text-align:center;">source</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">来源</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">is_auto_close_task</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求流转到结束状态时，是否自动关闭关联的任务。为 1 时会自动关闭；默认取 0，不关闭</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签，标签不存在时将自动创建，多个以英文坚线分格</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_这是一个自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/story/get_story_custom_fields_settings.html" class="">获取需求自定义字段配置</a> 获取</td></tr><tr><td style="text-align:center;">custom_plan_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义计划应用参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/iteration/get_plan_apps.html" class="">获取自定义计划应用</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

## 更新需求 1010104801125341253 的优先级为高，处理人为 anyechen

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10158231&workitems=[{"id"=123123,"name"="测试“}]' 'https://api.tapd.cn/stories/batch_update_story'`

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


# 需求字段说明

需求字段说明，请参考：[需求字段说明](/document/api-doc/API文档/api_reference/story/story.html)
