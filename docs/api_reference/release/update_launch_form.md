# Update Launch Form

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/update_launch_form.html

- SDK 方法名nodeJspythongolangcurl 使用 Basic Auth 鉴权调用示例
- 返回结果
- 版本字段重要字段说明



 
# # 说明

 更新发布评审单，返回更新发布评审单的数据

  
# # url

 https://api.tapd.cn/launch_forms

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 POST

 
# # 请求数限制

 一次更新一条数据

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布评审ID</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">integer</td> <td style="text-align:center;">项目ID</td></tr> <tr><td style="text-align:center;">title</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">标题</td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布评审状态，可选值为["initial","auditing","signing","sign_completed","finished"]，分别表示初始化、评审中、待签发、签发结束、发布结束</td></tr> <tr><td style="text-align:center;">version_type</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">版本类型，可选值为["Alpha版本","Beta版本","正式版本","Preview版本"]</td></tr> <tr><td style="text-align:center;">baseline</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">基线</td></tr> <tr><td style="text-align:center;">release_model</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布模块，可选值为["Client","Web","Server"]</td></tr> <tr><td style="text-align:center;">roadmap_version</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">路标版本</td></tr> <tr><td style="text-align:center;">release_type</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布类型，可选值为["全发布","正常发布","紧急发布","灰度发布","简便发布"]</td></tr> <tr><td style="text-align:center;">change_type</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">变更类型，可选值为["特性变更","缺陷变更","行业标准或政策性调整","首次发布","运维需求变更"]</td></tr> <tr><td style="text-align:center;">signed_by</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">签发人，多个用户使用“;”分割，例如：a;b</td></tr> <tr><td style="text-align:center;">archived_by</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布确认人，多个用户使用“;”分割，例如：a;b</td></tr> <tr><td style="text-align:center;">cc</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">抄送人，多个用户使用“;”分割，例如：a;b</td></tr> <tr><td style="text-align:center;">change_notifier</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">变更通知人，多个用户使用“;”分割，例如：a;b</td></tr> <tr><td style="text-align:center;">signer_comment</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">签发意见，status从signing流转到initial或者auditing时，签发意见是必填项</td></tr> <tr><td style="text-align:center;">release_result</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布结果，可选值为["release_success", "release_fail"]，当status变更为finished时，release_result字段必填</td></tr> <tr><td style="text-align:center;">release_comment</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布意见，当status变更为finished时，并且release_result为release_fail， release_comment必填</td></tr> <tr><td style="text-align:center;">remark</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">备注</td></tr> <tr><td style="text-align:center;">cus_{$自定义字段别名}</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td></tr> <tr><td style="text-align:center;">custom_field_*</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string或者integer</td> <td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/release/get_launch_forms_custom_fields_settings.html" class="">获取发布评审自定义字段配置</a> 获取</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
### # curl 使用 Basic Auth 鉴权调用示例

 curl -u 'api_user:api_password' -d 'workspace_id=10104801&id=1010104801065798331&creator=tapd_api&cus_自定义字段的名称=custom_field_value' 'https://api.tapd.cn/launch_forms'

 
## # 返回结果

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

 1234567891011121314151617181920212223242526272829303132333435363738394041424344454647484950515253545556575859606162636465666768697071727374757677787980
# # 版本字段说明

 
## # 版本字段重要字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">评审ID</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">所属项目ID</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr> <tr><td style="text-align:center;">title</td> <td style="text-align:center;">标题</td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">状态</td></tr> <tr><td style="text-align:center;">version_type</td> <td style="text-align:center;">版本类型</td></tr> <tr><td style="text-align:center;">baseline</td> <td style="text-align:center;">基线</td></tr> <tr><td style="text-align:center;">release_model</td> <td style="text-align:center;">发布模块</td></tr> <tr><td style="text-align:center;">roadmap_version</td> <td style="text-align:center;">路标版本</td></tr> <tr><td style="text-align:center;">release_type</td> <td style="text-align:center;">发布类型</td></tr> <tr><td style="text-align:center;">change_type</td> <td style="text-align:center;">变更类型</td></tr> <tr><td style="text-align:center;">signed_by</td> <td style="text-align:center;">签发人</td></tr> <tr><td style="text-align:center;">archived_by</td> <td style="text-align:center;">发布确认人</td></tr> <tr><td style="text-align:center;">cc</td> <td style="text-align:center;">抄送人</td></tr> <tr><td style="text-align:center;">change_notifier</td> <td style="text-align:center;">变更通知人</td></tr> <tr><td style="text-align:center;">signed</td> <td style="text-align:center;">签发时间</td></tr> <tr><td style="text-align:center;">archived</td> <td style="text-align:center;">归档时间</td></tr> <tr><td style="text-align:center;">signer_result</td> <td style="text-align:center;">签发结论</td></tr> <tr><td style="text-align:center;">signer_comment</td> <td style="text-align:center;">签发意见</td></tr> <tr><td style="text-align:center;">release_result</td> <td style="text-align:center;">发布结果</td></tr> <tr><td style="text-align:center;">release_comment</td> <td style="text-align:center;">发布意见</td></tr> <tr><td style="text-align:center;">test_path</td> <td style="text-align:center;">测试路径</td></tr> <tr><td style="text-align:center;">created_path</td> <td style="text-align:center;">归档路径</td></tr> <tr><td style="text-align:center;">remark</td> <td style="text-align:center;">备注</td></tr> <tr><td style="text-align:center;">participator</td> <td style="text-align:center;">参与人</td></tr> <tr><td style="text-align:center;">template_id</td> <td style="text-align:center;">模板ID</td></tr> <tr><td style="text-align:center;">iteration_id</td> <td style="text-align:center;">迭代ID</td></tr> <tr><td style="text-align:center;">release_id</td> <td style="text-align:center;">发布计划ID</td></tr> <tr><td style="text-align:center;">custom_field_*</td> <td style="text-align:center;">自定义字段</td></tr></tbody></table>
