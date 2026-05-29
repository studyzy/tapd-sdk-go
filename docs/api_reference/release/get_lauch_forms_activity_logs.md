# 获取发布评审日志

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_lauch_forms_activity_logs.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取发布评审日志](#获取发布评审日志)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

获取发布评审日志(评审过程的步骤和说明)

# url

`https://api.tapd.cn/launch_forms/get_activity_logs`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">form_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">发布评审ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取发布评审日志

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/launch_forms/get_activity_logs?workspace_id=10104801&form_id=1010104801079777231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "LaunchChange": {
                "id": "1010104801083909701",
                "workspace_id": "10104801",
                "type": "signing",
                "form_id": "1010104801079777231",
                "activity_form_id": null,
                "field": "sign",
                "old_value": "signing",
                "new_value": {
                    "comment": "同意 【通过企业微信快速审批】",
                    "result": "sign_agree"
                },
                "created_by": "v_xuanfang",
                "created": "2023-07-13 15:13:34",
                "operation": "sign_agree"
            }
        },
        {
            "LaunchChange": {
                "id": "1010104801083908923",
                "workspace_id": "10104801",
                "type": "auditing",
                "form_id": "1010104801079777231",
                "activity_form_id": "1010104801067750287",
                "field": "项目经理评审",
                "old_value": null,
                "new_value": {
                    "activity": {
                        "id": "1010104801067750287",
                        "comment": "2111",
                        "conclusion": "audit_absolutely",
                        "factors_total_value": 0,
                        "factors_pass_value": 0,
                        "factors_passed_rate": 100,
                        "audited": "2023-07-13  10:41:23",
                        "is_audited": true,
                        "audited_by": "v_xuanfang",
                        "workspace_id": "10104801"
                    },
                    "factor_result": []
                },
                "created_by": "v_xuanfang",
                "created": "2023-07-13 10:41:23",
                "operation": "audit_absolutely"
            }
        },
        {
            "LaunchChange": {
                "id": "1010104801083448609",
                "workspace_id": "10104801",
                "type": "initial",
                "form_id": "1010104801079777231",
                "activity_form_id": null,
                "field": "initialization",
                "old_value": null,
                "new_value": null,
                "created_by": "v_xuanfang",
                "created": "2022-09-08 17:27:34",
                "operation": "initialization"
            }
        }
    ],
    "info": "success"
}
```


# 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评审日志记录ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属项目ID</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">form_id</td><td style="text-align:center;">评审单ID</td></tr><tr><td style="text-align:center;">activity_form_id</td><td style="text-align:center;">发布评审活动ID</td></tr><tr><td style="text-align:center;">factor_result</td><td style="text-align:center;">评审要素</td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">字段名</td></tr><tr><td style="text-align:center;">old_value</td><td style="text-align:center;">变更前</td></tr><tr><td style="text-align:center;">new_value</td><td style="text-align:center;">变更后</td></tr><tr><td style="text-align:center;">created_by</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">operation</td><td style="text-align:center;">流程状态</td></tr></tbody></table>

# factor\_result字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">评审要素ID</td><td style="text-align:center;">评审结果取值（factor_y:是，factor_n:否，factor_inesstional:不涉及）</td></tr></tbody></table>
