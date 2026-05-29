# 更新缺陷

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/update_bug.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [更新缺陷 1010158231500628817 的当前处理人给 anyechen](#更新缺陷-1010158231500628817-的当前处理人给-anyechen)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

用于更新缺陷的字段 或者 流转缺陷状态，返回缺陷更新后的数据

# url

`https://api.tapd.cn/bugs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。为了兼容自定义优先级，<code>请使用 priority_label 字段</code>，详情参考：<a href="/document/api-doc/API文档/subject/custom_priority/" class="">如何兼容自定义优先级</a></td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。推荐使用这个字段</td></tr><tr><td style="text-align:center;">severity</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">严重程度</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">v_status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态(支持传入中文状态名称)</td></tr><tr><td style="text-align:center;">module</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">模块</td></tr><tr><td style="text-align:center;">feature</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">特性</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">发布计划</td></tr><tr><td style="text-align:center;">version_report</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发现版本</td></tr><tr><td style="text-align:center;">version_test</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">验证版本</td></tr><tr><td style="text-align:center;">version_fix</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">合入版本</td></tr><tr><td style="text-align:center;">version_close</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">关闭版本</td></tr><tr><td style="text-align:center;">baseline_find</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发现基线</td></tr><tr><td style="text-align:center;">baseline_join</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">合入基线</td></tr><tr><td style="text-align:center;">baseline_test</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">验证基线</td></tr><tr><td style="text-align:center;">baseline_close</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">关闭基线</td></tr><tr><td style="text-align:center;">current_owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">reporter</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">current_user</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更人</td></tr><tr><td style="text-align:center;">participator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">参与人</td></tr><tr><td style="text-align:center;">te</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试人员</td></tr><tr><td style="text-align:center;">de</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">开发人员</td></tr><tr><td style="text-align:center;">auditer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">审核人</td></tr><tr><td style="text-align:center;">confirmer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">验证人</td></tr><tr><td style="text-align:center;">fixer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">修复人</td></tr><tr><td style="text-align:center;">closer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">关闭人</td></tr><tr><td style="text-align:center;">lastmodify</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">in_progress_time</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">接受处理时间</td></tr><tr><td style="text-align:center;">verify_time</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">验证时间</td></tr><tr><td style="text-align:center;">reject_time</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">拒绝时间</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">deadline</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">解决期限</td></tr><tr><td style="text-align:center;">os</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">操作系统</td></tr><tr><td style="text-align:center;">size</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">规模</td></tr><tr><td style="text-align:center;">platform</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">软件平台</td></tr><tr><td style="text-align:center;">testmode</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试方式</td></tr><tr><td style="text-align:center;">testphase</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试阶段</td></tr><tr><td style="text-align:center;">testtype</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试类型</td></tr><tr><td style="text-align:center;">source</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">缺陷根源</td></tr><tr><td style="text-align:center;">bugtype</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">缺陷类型</td></tr><tr><td style="text-align:center;">frequency</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">重现规律</td></tr><tr><td style="text-align:center;">originphase</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发现阶段</td></tr><tr><td style="text-align:center;">sourcephase</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">引入阶段</td></tr><tr><td style="text-align:center;">resolution</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">解决方法</td></tr><tr><td style="text-align:center;">estimate</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">预计解决时间</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签，标签不存在时将自动创建，多个以英文坚线分格</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">预估工时</td></tr><tr><td style="text-align:center;">keep_owner</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否保存处理人。取 1 则是保留</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/bug/get_bug_custom_fields_settings.html" class="">获取缺陷自定义字段配置</a> 获取</td></tr><tr><td style="text-align:center;">custom_plan_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义计划应用参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/iteration/get_plan_apps.html" class="">获取自定义计划应用</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

## 更新缺陷 1010158231500628817 的当前处理人给 anyechen

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'current_owner=anyechen;&id=1010158231500628817&workspace_id=10158231' 'https://api.tapd.cn/bugs'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Bug": {
            "id": "1010158231500628817",
            "title": "【示例】新官网Chrome浏览器兼容性bug",
            "description": null,
            "priority": "high",
            "severity": "prompt",
            "module": null,
            "status": "in_progress",
            "reporter": "anyechen",
            "deadline": null,
            "created": "2017-06-20 16:49:19",
            "bugtype": "",
            "resolved": null,
            "closed": null,
            "modified": "2019-06-27 14:29:03",
            "lastmodify": "anyechen",
            "auditer": null,
            "de": null,
            "fixer": null,
            "version_test": "",
            "version_report": "版本1",
            "version_close": "",
            "version_fix": "",
            "baseline_find": "",
            "baseline_join": "",
            "baseline_close": "",
            "baseline_test": "",
            "sourcephase": "",
            "te": null,
            "current_owner": "anyechen;",
            "iteration_id": "0",
            "resolution": "",
            "source": "",
            "originphase": "",
            "confirmer": null,
            "milestone": null,
            "participator": null,
            "closer": null,
            "platform": "",
            "os": "",
            "testtype": "",
            "testphase": "",
            "frequency": "",
            "cc": null,
            "regression_number": "0",
            "flows": "new",
            "feature": null,
            "testmode": "",
            "estimate": null,
            "issue_id": null,
            "created_from": null,
            "in_progress_time": null,
            "verify_time": null,
            "reject_time": null,
            "reopen_time": null,
            "audit_time": null,
            "suspend_time": null,
            "due": null,
            "begin": null,
            "release_id": null,
            "label": "阻塞|延期",
            "cus_自定义字段的名称": "custom_field_value",
            "custom_field_one": "",
            "custom_field_two": "",
            "custom_field_three": "",
            "custom_field_four": "",
            "custom_field_five": "",
            "custom_field_6": "",
            "custom_field_7": "",
            "custom_field_8": "",
            "custom_field_9": "",
            "custom_field_10": "",
            "custom_field_11": "",
            "custom_field_12": "",
            "custom_field_13": "",
            "custom_field_14": "",
            "custom_field_15": "",
            "custom_field_16": "",
            "custom_field_17": "",
            "custom_field_18": "",
            "custom_field_19": "",
            "custom_field_20": "",
            "custom_field_21": "",
            "custom_field_22": "",
            "custom_field_23": "",
            "custom_field_24": "",
            "custom_field_25": "",
            "custom_field_26": "",
            "custom_field_27": "",
            "custom_field_28": "",
            "custom_field_29": "",
            "custom_field_30": "",
            "custom_field_31": "",
            "custom_field_32": "",
            "custom_field_33": "",
            "custom_field_34": "",
            "custom_field_35": "",
            "custom_field_36": "",
            "custom_field_37": "",
            "custom_field_38": "",
            "custom_field_39": "",
            "custom_field_40": "",
            "custom_field_41": "",
            "custom_field_42": "",
            "custom_field_43": "",
            "custom_field_44": "",
            "custom_field_45": "",
            "custom_field_46": "",
            "custom_field_47": "",
            "custom_field_48": "",
            "custom_field_49": "",
            "custom_field_50": "",
            "workspace_id": "10158231"
        }
    },
    "info": "success"
}
```


# 缺陷字段说明

缺陷字段说明，请参考 [缺陷说明](/document/api-doc/API文档/api_reference/bug/bug.html)
