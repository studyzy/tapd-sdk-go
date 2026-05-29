# 获取缺陷数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs_count.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的缺陷数量](#获取项目下的缺陷数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [获取项目下当前处理人为 anyechen 、高优先级、状态为新的缺陷数量](#获取项目下当前处理人为-anyechen-、高优先级、状态为新的缺陷数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

计算符合查询条件的缺陷数量并返回

# url

`https://api.tapd.cn/bugs/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回缺陷数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。为了兼容自定义优先级，<code>请使用 priority_label 字段</code>，详情参考：<a href="/document/api-doc/API文档/subject/custom_priority/" class="">如何兼容自定义优先级</a></td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。推荐使用这个字段</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">severity</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">严重程度</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;">支持不等于查询、枚举查询</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签查询</td><td style="text-align:center;">支持枚举查询，5月30号上线</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">module</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">模块</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">发布计划</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">version_report</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发现版本</td><td style="text-align:center;">枚举查询</td></tr><tr><td style="text-align:center;">version_test</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">验证版本</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">version_fix</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">合入版本</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">version_close</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">关闭版本</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">baseline_find</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发现基线</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">baseline_join</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">合入基线</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">baseline_test</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">验证基线</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">baseline_close</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">关闭基线</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">current_owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">reporter</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">participator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">参与人</td><td style="text-align:center;">支持多人员查询</td></tr><tr><td style="text-align:center;">te</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试人员</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">de</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">开发人员</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">auditer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">审核人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">confirmer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">验证人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fixer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">修复人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">closer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">关闭人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">lastmodify</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">最后修改人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">in_progress_time</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">接受处理时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">resolved</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">解决时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">verify_time</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">验证时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">closed</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">关闭时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">reject_time</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">拒绝时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">deadline</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">解决期限</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">os</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">操作系统</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">platform</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">软件平台</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">testmode</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试方式</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">testphase</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试阶段</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">testtype</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">source</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">缺陷根源</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">bugtype</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">缺陷类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">frequency</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">重现规律</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">originphase</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发现阶段</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">sourcephase</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">引入阶段</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">resolution</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">解决方法</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">estimate</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">预计解决时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">预估工时</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/bug/get_bug_custom_fields_settings.html" class="">获取缺陷自定义字段配置</a> 获取</td><td style="text-align:center;">支持枚举查询</td></tr><tr><td style="text-align:center;">custom_plan_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义计划应用参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/iteration/get_plan_apps.html" class="">获取自定义计划应用</a> 获取</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的缺陷数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/count?workspace_id=10158231'`

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


## 获取项目下当前处理人为 anyechen 、高优先级、状态为新的缺陷数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/count?workspace_id=10158231&current_owner=anyechen;&priority=high&status=new'`

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


# 缺陷字段说明

缺陷字段说明，请参考 [缺陷说明](/document/api-doc/API文档/api_reference/bug/bug.html)
