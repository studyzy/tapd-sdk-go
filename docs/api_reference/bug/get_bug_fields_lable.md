# 获取缺陷所有字段的中英文

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bug_fields_lable.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下的缺陷字段](#获取项目下的缺陷字段)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

返回缺陷所有字段的中英文

# url

`https://api.tapd.cn/bugs/get_fields_lable`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的缺陷字段

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/bugs/get_fields_lable?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "id": "ID",
        "title": "标题",
        "priority": "优先级",
        "priority_label": "优先级",
        "severity": "严重程度",
        "status": "状态",
        "iteration_id": "迭代",
        "module": "模块",
        "release_id": "发布计划",
        "feature": "特性",
        "version_report": "发现版本",
        "version_test": "验证版本",
        "version_fix": "合入版本",
        "version_close": "关闭版本",
        "baseline_find": "发现基线",
        "baseline_join": "合入基线",
        "baseline_test": "验证基线 ",
        "baseline_close": "关闭基线",
        "current_owner": "处理人",
        "cc": "抄送人",
        "reporter": "创建人",
        "participator": "参与人",
        "te": "测试人员",
        "de": "开发人员",
        "auditer": "审核人",
        "confirmer": "验证人",
        "fixer": "修复人",
        "closer": "关闭人",
        "lastmodify": "最后修改人",
        "created": "创建时间",
        "reopen_time": "重新打开时间",
        "assigned_time": "分配时间",
        "in_progress_time": "接受处理时间 ",
        "resolved": "解决时间",
        "verify_time": "验证时间",
        "closed": "关闭时间",
        "reject_time": "拒绝时间",
        "modified": "最后修改时间",
        "begin": "预计开始",
        "due": "预计结束",
        "deadline": "解决期限",
        "os": "操作系统",
        "platform": "软件平台",
        "testmode": "测试方式",
        "testphase": "测试阶段",
        "testtype": "测试类型",
        "source": "缺陷根源",
        "bugtype": "缺陷类型",
        "frequency": "重现规律",
        "originphase": "发现阶段",
        "sourcephase": "引入阶段  ",
        "resolution": "解决方法",
        "estimate": "预计解决时间",
        "description": "详细描述",
        "custom_field_one": "测试下拉",
        "custom_field_two": "测试字段"
    },
    "info": "success"
}
```

