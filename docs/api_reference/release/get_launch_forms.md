# Get Launch Forms

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_launch_forms.html

- SDK 方法名nodeJspythongolang
- 获取发布评审curl 使用 Basic Auth 鉴权调用示例返回结果



 
# # 说明

 获取发布评审

  
# # url

 https://api.tapd.cn/launch_forms

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 GET

 
# # 请求数限制

 默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">integer</td> <td style="text-align:center;">项目ID</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">id</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">发布评审ID</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">creator</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">创建人</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">否</td> <td style="text-align:center;">datetime</td> <td style="text-align:center;">创建时间</td> <td style="text-align:center;">支持时间查询</td></tr> <tr><td style="text-align:center;">title</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">标题</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">状态</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">version_type</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">版本类型</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">baseline</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">基线</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">release_model</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布模块</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">roadmap_version</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">路标版本</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">release_type</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布类型</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">change_type</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">变更类型</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">signed_by</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">签发人</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">archived_by</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">发布确认人</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">cc</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">抄送人</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">change_notifier</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">变更通知人</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">limit</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">设置返回数量限制，默认为30</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">page</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">fields</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td> <td style="text-align:center;"></td></tr></tbody></table>

 
# # 调用示例及返回结果

 
## # 获取发布评审

 
### # curl 使用 Basic Auth 鉴权调用示例

 curl -u 'api_user:api_password' 'https://api.tapd.cn/launch_forms?workspace_id=10104801'

 
### # 返回结果

 ```
{
    "status": 1,
    "data": [
        {
            "LaunchForm": {
                "id": "1010104801079697767",
                "title": null,
                "name": "202101150008",
                "creator": "v_xuanfang",
                "created": "2021-01-15 19:47:37",
                "workspace_id": "10104801",
                "status": "LAUNCHFORM_STATUS_INITIAL",
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
                "flows": "LAUNCHFORM_STATUS_INITIAL",
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
        }
    ],
    "info": "success"
}
```

 123456789101112131415161718192021222324252627282930313233343536373839404142434445464748495051525354555657585960616263646566676869707172737475767778798081
# # 返回字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">评审ID</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">所属项目ID</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr> <tr><td style="text-align:center;">title</td> <td style="text-align:center;">标题</td></tr> <tr><td style="text-align:center;">status</td> <td style="text-align:center;">状态</td></tr> <tr><td style="text-align:center;">version_type</td> <td style="text-align:center;">版本类型</td></tr> <tr><td style="text-align:center;">baseline</td> <td style="text-align:center;">基线</td></tr> <tr><td style="text-align:center;">release_model</td> <td style="text-align:center;">发布模块</td></tr> <tr><td style="text-align:center;">roadmap_version</td> <td style="text-align:center;">路标版本</td></tr> <tr><td style="text-align:center;">release_type</td> <td style="text-align:center;">发布类型</td></tr> <tr><td style="text-align:center;">change_type</td> <td style="text-align:center;">变更类型</td></tr> <tr><td style="text-align:center;">signed_by</td> <td style="text-align:center;">签发人</td></tr> <tr><td style="text-align:center;">archived_by</td> <td style="text-align:center;">发布确认人</td></tr> <tr><td style="text-align:center;">cc</td> <td style="text-align:center;">抄送人</td></tr> <tr><td style="text-align:center;">change_notifier</td> <td style="text-align:center;">变更通知人</td></tr> <tr><td style="text-align:center;">signed</td> <td style="text-align:center;">签发时间</td></tr> <tr><td style="text-align:center;">archived</td> <td style="text-align:center;">归档时间</td></tr> <tr><td style="text-align:center;">signer_result</td> <td style="text-align:center;">签发结论</td></tr> <tr><td style="text-align:center;">signer_comment</td> <td style="text-align:center;">签发意见</td></tr> <tr><td style="text-align:center;">release_result</td> <td style="text-align:center;">发布结果</td></tr> <tr><td style="text-align:center;">release_comment</td> <td style="text-align:center;">发布意见</td></tr> <tr><td style="text-align:center;">test_path</td> <td style="text-align:center;">测试路径</td></tr> <tr><td style="text-align:center;">created_path</td> <td style="text-align:center;">归档路径</td></tr> <tr><td style="text-align:center;">remark</td> <td style="text-align:center;">备注</td></tr> <tr><td style="text-align:center;">participator</td> <td style="text-align:center;">参与人</td></tr> <tr><td style="text-align:center;">template_id</td> <td style="text-align:center;">模板ID</td></tr> <tr><td style="text-align:center;">iteration_id</td> <td style="text-align:center;">迭代ID</td></tr> <tr><td style="text-align:center;">release_id</td> <td style="text-align:center;">发布计划ID</td></tr> <tr><td style="text-align:center;">custom_field_*</td> <td style="text-align:center;">自定义字段</td></tr></tbody></table>
