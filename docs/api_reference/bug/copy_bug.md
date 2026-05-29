# 复制缺陷

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/copy_bug.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [复制缺陷到另外项目](#复制缺陷到另外项目)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [同步复制缺陷到另外项目，同时设置状态、详细字段为同步字段](#同步复制缺陷到另外项目-同时设置状态、详细字段为同步字段)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

复制缺陷，返回复制缺陷的数据

# url

`https://api.tapd.cn/bugs/copy_bug`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP 请求方式

POST

# 请求数限制

-   一次复制一条缺陷
-   同步字段仅支持 title(标题)、 status(状态)、 description(详细描述)、 attachment(附件)、 begin\_due(预计开始结束时间)、 current\_owner(处理人)、 cc(抄送人)、 de(开发人员)、 module(模块)、 iteration\_id(迭代)、 priority(优先级)、 severity(严重程度)、 version\_report(发现版本)、 version\_test(验证版本)、 version\_fix(合入版本)、 version\_close(关闭版本)、 baseline\_find(发现基线)、 baseline\_join(合入基线)、 baseline\_close(关闭基线)、 baseline\_test(验证基线 )、 os(操作系统)、 platform(软件平台)、 testmode(测试方式)、 testphase(测试阶段)、 testtype(测试类型)、 source(缺陷根源)、 bugtype(缺陷类型)、 frequency(重现规律)、 originphase(发现阶段)、 sourcephase(引入阶段 )、 resolution(解决方法)、 deadline(解决期限)、 release\_id(发布计划)、 label(标签)、 size(规模)、 estimate(预计解决时间)、 comments(评论)、 custom\_field::字段中文名

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源项目ID</td></tr><tr><td style="text-align:center;">src_bug_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源缺陷ID</td></tr><tr><td style="text-align:center;">dst_workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">目标项目ID</td></tr><tr><td style="text-align:center;">sync_fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需要同步的字段。多写使用 , 分隔</td></tr></tbody></table>

# 调用示例及返回结果

## 复制缺陷到另外项目

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10158231&src_bug_id=1010158231500664575&dst_workspace_id=10104801' 'https://api.tapd.cn/bugs/copy_bug'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Bug": {
            "id": "1010104801500664637",
            "title": "同步复制-源",
            "description": null,
            "priority": "",
            "severity": "",
            "module": "",
            "status": "rejected",
            "reporter": "anyechen",
            "created": "2020-11-26 15:22:34",
            "bugtype": "",
            "resolved": null,
            "closed": null,
            "modified": "2020-11-26 15:22:34",
            "lastmodify": "TAPD_SYSTEM",
            "auditer": "",
            "de": "",
            "fixer": "",
            "version_test": "",
            "version_report": "",
            "version_close": "",
            "version_fix": "",
            "baseline_find": "",
            "baseline_join": "",
            "baseline_close": "",
            "baseline_test": "",
            "sourcephase": "",
            "te": "",
            "current_owner": "anyechen",
            "iteration_id": "0",
            "resolution": "",
            "source": "",
            "originphase": "",
            "confirmer": "",
            "milestone": "",
            "participator": "",
            "closer": "",
            "platform": "",
            "os": "",
            "testtype": "",
            "testphase": "",
            "frequency": "",
            "cc": "",
            "regression_number": "0",
            "flows": "",
            "feature": "",
            "testmode": "",
            "estimate": null,
            "issue_id": null,
            "created_from": "",
            "release_id": null,
            "verify_time": null,
            "reject_time": null,
            "reopen_time": null,
            "audit_time": null,
            "suspend_time": null,
            "due": null,
            "begin": null,
            "deadline": null,
            "in_progress_time": null,
            "assigned_time": null,
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
            "custom_field_51": "",
            "custom_field_52": "",
            "custom_field_53": "",
            "custom_field_54": "",
            "custom_field_55": "",
            "custom_field_56": "",
            "custom_field_57": "",
            "custom_field_58": "",
            "custom_field_59": "",
            "custom_field_60": "",
            "custom_field_61": "",
            "custom_field_62": "",
            "custom_field_63": "",
            "custom_field_64": "",
            "custom_field_65": "",
            "custom_field_66": "",
            "custom_field_67": "",
            "custom_field_68": "",
            "custom_field_69": "",
            "custom_field_70": "",
            "custom_field_71": "",
            "custom_field_72": "",
            "custom_field_73": "",
            "custom_field_74": "",
            "custom_field_75": "",
            "custom_field_76": "",
            "custom_field_77": "",
            "custom_field_78": "",
            "custom_field_79": "",
            "custom_field_80": "",
            "custom_field_81": "",
            "custom_field_82": "",
            "custom_field_83": "",
            "custom_field_84": "",
            "custom_field_85": "",
            "custom_field_86": "",
            "custom_field_87": "",
            "custom_field_88": "",
            "custom_field_89": "",
            "custom_field_90": "",
            "custom_field_91": "",
            "custom_field_92": "",
            "custom_field_93": "",
            "custom_field_94": "",
            "custom_field_95": "",
            "custom_field_96": "",
            "custom_field_97": "",
            "custom_field_98": "",
            "custom_field_99": "",
            "custom_field_100": "",
            "workspace_id": "10104801"
        }
    },
    "info": "success"
}
```


## 同步复制缺陷到另外项目，同时设置状态、详细字段为同步字段

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10158231&src_bug_id=1010158231500664575&dst_workspace_id=755&sync_fields=description,status' 'https://api.tapd.cn/bugs/copy_bug'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Bug": {
            "id": "1000000755500664641",
            "title": "同步复制-源",
            "description": null,
            "priority": "",
            "severity": "",
            "module": "",
            "status": "rejected",
            "reporter": "anyechen",
            "created": "2020-11-26 15:41:03",
            "bugtype": "",
            "resolved": null,
            "closed": null,
            "modified": "2020-11-26 15:41:03",
            "lastmodify": "TAPD_SYSTEM",
            "auditer": "",
            "de": "",
            "fixer": "",
            "version_test": "",
            "version_report": "",
            "version_close": "",
            "version_fix": "",
            "baseline_find": "",
            "baseline_join": "",
            "baseline_close": "",
            "baseline_test": "",
            "sourcephase": "",
            "te": "",
            "current_owner": "anyechen",
            "iteration_id": "0",
            "resolution": "",
            "source": "",
            "originphase": "",
            "confirmer": "",
            "milestone": "",
            "participator": "",
            "closer": "",
            "platform": "",
            "os": "",
            "testtype": "",
            "testphase": "",
            "frequency": "",
            "cc": "",
            "regression_number": "0",
            "flows": "",
            "feature": "",
            "testmode": "",
            "estimate": null,
            "issue_id": null,
            "created_from": "",
            "release_id": null,
            "verify_time": null,
            "reject_time": null,
            "reopen_time": null,
            "audit_time": null,
            "suspend_time": null,
            "due": null,
            "begin": null,
            "deadline": null,
            "in_progress_time": null,
            "assigned_time": null,
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
            "custom_field_51": "",
            "custom_field_52": "",
            "custom_field_53": "",
            "custom_field_54": "",
            "custom_field_55": "",
            "custom_field_56": "",
            "custom_field_57": "",
            "custom_field_58": "",
            "custom_field_59": "",
            "custom_field_60": "",
            "custom_field_61": "",
            "custom_field_62": "",
            "custom_field_63": "",
            "custom_field_64": "",
            "custom_field_65": "",
            "custom_field_66": "",
            "custom_field_67": "",
            "custom_field_68": "",
            "custom_field_69": "",
            "custom_field_70": "",
            "custom_field_71": "",
            "custom_field_72": "",
            "custom_field_73": "",
            "custom_field_74": "",
            "custom_field_75": "",
            "custom_field_76": "",
            "custom_field_77": "",
            "custom_field_78": "",
            "custom_field_79": "",
            "custom_field_80": "",
            "custom_field_81": "",
            "custom_field_82": "",
            "custom_field_83": "",
            "custom_field_84": "",
            "custom_field_85": "",
            "custom_field_86": "",
            "custom_field_87": "",
            "custom_field_88": "",
            "custom_field_89": "",
            "custom_field_90": "",
            "custom_field_91": "",
            "custom_field_92": "",
            "custom_field_93": "",
            "custom_field_94": "",
            "custom_field_95": "",
            "custom_field_96": "",
            "custom_field_97": "",
            "custom_field_98": "",
            "custom_field_99": "",
            "custom_field_100": "",
            "workspace_id": "755"
        }
    },
    "info": "success"
}
```


# 缺陷字段说明

缺陷字段说明，请参考 [缺陷说明](/document/api-doc/API文档/api_reference/bug/bug.html)
