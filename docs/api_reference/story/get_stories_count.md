# 获取需求数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories_count.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下需求的数量](#获取项目下需求的数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取项目下优先级为 High 需求的数量](#获取项目下优先级为-high-需求的数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

计算符合查询条件的需求数量并返回

# url

`https://api.tapd.cn/stories/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回需求数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。为了兼容自定义优先级，<code>请使用 priority_label 字段</code>，详情参考：<a href="/document/api-doc/API文档/subject/custom_priority/" class="">如何兼容自定义优先级</a></td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。推荐使用这个字段</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">business_value</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">业务价值</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">v_status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态(支持传入中文状态名称)</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">with_v_status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">值=1可以返回中文状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签查询</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求类别ID</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">module</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">模块</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">feature</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">特性</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">test_focus</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试重点</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">size</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">规模</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">tech_risk</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">技术风险</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">business_value</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">业务价值</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;">支持多人员查询</td></tr><tr><td style="text-align:center;">developer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">开发人员</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">完成时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">迭代ID</td><td style="text-align:center;">支持不等于查询或枚举查询</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预估工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">effort_completed</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">完成工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">remain</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">剩余工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">exceed</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">超出工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求分类</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">发布计划</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">source</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求来源</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">ancestor_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">祖先需求，查询指定需求下所有子需求</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">父需求</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">children_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">子需求</td><td style="text-align:center;">为空查询传：丨</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/story/get_story_custom_fields_settings.html" class="">获取需求自定义字段配置</a> 获取</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">custom_plan_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义计划应用参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/iteration/get_plan_apps.html" class="">获取自定义计划应用</a> 获取</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">include_sub_category</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">是否包含子分类</td><td style="text-align:center;">取值 0或者1，默认取 0</td></tr><tr><td style="text-align:center;">include_sub_iteration</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">是否包含子迭代</td><td style="text-align:center;">取值 0或者1，默认取 0</td></tr><tr><td style="text-align:center;">include_leaf_stories</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">是否包含子需求</td><td style="text-align:center;">取值 0或者1，默认取 0</td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下需求的数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/count?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 7
    },
    "info": "success"
}
```


## 获取项目下优先级为 High 需求的数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/count?workspace_id=10158231&priority=4'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 4
    },
    "info": "success"
}
```

