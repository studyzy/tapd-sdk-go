# 创建发布评审接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/add_launch_form.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
-   [返回结果](#返回结果)
-   [版本字段重要字段说明](#版本字段重要字段说明)

# 说明

创建发布评审单，返回创建发布评审单的数据

# url

`https://api.tapd.cn/launch_forms`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">version_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本类型</td></tr><tr><td style="text-align:center;">baseline</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">基线</td></tr><tr><td style="text-align:center;">release_model</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发布模块</td></tr><tr><td style="text-align:center;">roadmap_version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">路标版本</td></tr><tr><td style="text-align:center;">release_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发布类型</td></tr><tr><td style="text-align:center;">signed_by</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">签发人</td></tr><tr><td style="text-align:center;">archived_by</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">发布确认人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/release/get_launch_forms_custom_fields_settings.html" class="">获取发布评审自定义字段配置</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&template_id=1010104801065798331&creator=tapd_api&cus_自定义字段的名称=custom_field_value' 'https://api.tapd.cn/launch_forms'`

## 返回结果

```
{
    "status": 1,
    "data": {
        "LaunchForm": {
            "id": "1010104801079724013",
            "title": null,
            "name": "202107210009",
            "creator": "v_xuanfang",
            "created": "2021-07-21 14:27:36",
            "workspace_id": "10104801",
            "status": "initial",
            "version_type": null,
            "baseline": null,
            "release_model": null,
            "roadmap_version": null,
            "release_type": "正常发布",
            "change_type": null,
            "signed_by": null,
            "archived_by": null,
            "cc": null,
            "change_notifier": null,
            "signed": null,
            "archived": null,
            "signer_result": null,
            "signer_comment": null,
            "release_result": null,
            "release_comment": null,
            "test_path": null,
            "created_path": null,
            "remark": null,
            "participator": ";v_xuanfang;",
            "template_id": "1010104801065798331",
            "iteration_id": null,
            "release_id": null,
            "flows": "initial",
            "cus_自定义字段的名称": "custom_field_value",
            "custom_field_one": null,
            "custom_field_two": null,
            "custom_field_three": null,
            "custom_field_four": null,
            "custom_field_five": null,
            "custom_field_six": null,
            "custom_field_seven": null,
            "custom_field_eight": null,
            "custom_field_9": null,
            "custom_field_10": null,
            "custom_field_11": null,
            "custom_field_12": null,
            "custom_field_13": null,
            "custom_field_14": null,
            "custom_field_15": null,
            "custom_field_16": null,
            "custom_field_17": null,
            "custom_field_18": null,
            "custom_field_19": null,
            "custom_field_20": null,
            "custom_field_21": null,
            "custom_field_22": null,
            "custom_field_23": null,
            "custom_field_24": null,
            "custom_field_25": null,
            "custom_field_26": null,
            "custom_field_27": null,
            "custom_field_28": null,
            "custom_field_29": null,
            "custom_field_30": null,
            "custom_field_31": null,
            "custom_field_32": null,
            "custom_field_33": null,
            "custom_field_34": null,
            "custom_field_35": null,
            "custom_field_36": null,
            "custom_field_37": null,
            "custom_field_38": null,
            "custom_field_39": null,
            "custom_field_40": null
        }
    },
    "info": "success"
}
```


# 版本字段说明

## 版本字段重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">评审ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属项目ID</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">title</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">version_type</td><td style="text-align:center;">版本类型</td></tr><tr><td style="text-align:center;">baseline</td><td style="text-align:center;">基线</td></tr><tr><td style="text-align:center;">release_model</td><td style="text-align:center;">发布模块</td></tr><tr><td style="text-align:center;">roadmap_version</td><td style="text-align:center;">路标版本</td></tr><tr><td style="text-align:center;">release_type</td><td style="text-align:center;">发布类型</td></tr><tr><td style="text-align:center;">change_type</td><td style="text-align:center;">变更类型</td></tr><tr><td style="text-align:center;">signed_by</td><td style="text-align:center;">签发人</td></tr><tr><td style="text-align:center;">archived_by</td><td style="text-align:center;">发布确认人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">change_notifier</td><td style="text-align:center;">变更通知人</td></tr><tr><td style="text-align:center;">signed</td><td style="text-align:center;">签发时间</td></tr><tr><td style="text-align:center;">archived</td><td style="text-align:center;">归档时间</td></tr><tr><td style="text-align:center;">signer_result</td><td style="text-align:center;">签发结论</td></tr><tr><td style="text-align:center;">signer_comment</td><td style="text-align:center;">签发意见</td></tr><tr><td style="text-align:center;">release_result</td><td style="text-align:center;">发布结果</td></tr><tr><td style="text-align:center;">release_comment</td><td style="text-align:center;">发布意见</td></tr><tr><td style="text-align:center;">test_path</td><td style="text-align:center;">测试路径</td></tr><tr><td style="text-align:center;">created_path</td><td style="text-align:center;">归档路径</td></tr><tr><td style="text-align:center;">remark</td><td style="text-align:center;">备注</td></tr><tr><td style="text-align:center;">participator</td><td style="text-align:center;">参与人</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">发布计划ID</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">自定义字段</td></tr></tbody></table>
