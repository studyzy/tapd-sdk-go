# 复制需求

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/copy_story.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [复制需求到另外项目](#复制需求到另外项目)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [同步复制需求到另外项目，同时设置状态、详细字段为同步字段](#同步复制需求到另外项目-同时设置状态、详细字段为同步字段)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [需求重要字段说明](#需求重要字段说明)
-   [需求优先级(priority)取值字段说明](#需求优先级-priority-取值字段说明)

# 说明

复制需求，返回新建需求的数据

# url

`https://api.tapd.cn/stories/copy_story`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP 请求方式

POST

# 请求数限制

-   一次复制一条需求
-   同步复制支持的字段：name(标题)、 status(状态)、 description(详细描述)、 attachment(附件)、begin\_due(预计开始结束时间)、 module(模块)、 feature(特性)、 priority(优先级)、 owner(处理人)、 developer(开发人员)、 business\_value(业务价值)、 size(规模)、 effort(Estimated effort)、 cc(抄送人)、 test\_focus(测试重点)、 version(版本)、 label(标签)、 tech\_risk(技术风险)、 iteration\_id(迭代)、 comments(评论)、 custom\_field::字段中文名

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源项目ID</td></tr><tr><td style="text-align:center;">src_story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">源需求ID</td></tr><tr><td style="text-align:center;">dst_workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">目标项目ID</td></tr><tr><td style="text-align:center;">sync_fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需要同步的字段。多写使用 , 分隔</td></tr><tr><td style="text-align:center;">dst_workitem_type_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">目标需求类别ID</td></tr><tr><td style="text-align:center;">new_creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">新需求创建人</td></tr><tr><td style="text-align:center;">new_status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">新需求状态</td></tr></tbody></table>

# 调用示例及返回结果

## 复制需求到另外项目

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&src_story_id=1010104801854843773&dst_workspace_id=755' 'https://api.tapd.cn/stories/copy_story'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Story": {
            "id": "1000000755854845111",
            "workitem_type_id": "1000000755000000003",
            "name": "bbbbbbbb",
            "description": "<p><b><span style=\"color: #ff0000;\">作为</span></b>&nbsp;</p>\n<div>&nbsp;\n<div>\n<div><b><span style=\"color: #ff0000;\">我希望</span></b>&nbsp;</div>\n<div></div>\n<div><b><span style=\"color: #ff0000;\">以便</span></b> ADFADFADF</div>\n<div></div>\n<div>【验收标准】</div>\n<div>1、</div>\n<div>2、</div>\n<div>3、</div>\n<div>ADFADFDFDAADFFADS</div>\n</div>\n</div>",
            "workspace_id": "755",
            "creator": "anyechen",
            "created": "2020-12-09 17:00:09",
            "modified": "2020-12-09 17:00:10",
            "status": "planning",
            "owner": "",
            "cc": "",
            "begin": null,
            "due": null,
            "size": "0",
            "priority": "",
            "developer": "",
            "iteration_id": "0",
            "test_focus": "",
            "type": "",
            "source": "",
            "module": "",
            "version": "",
            "completed": null,
            "category_id": "-1",
            "path": "1000000755854845111",
            "parent_id": "0",
            "children_id": "|",
            "ancestor_id": "1000000755854845111",
            "business_value": null,
            "effort": "0",
            "effort_completed": "0",
            "exceed": "0",
            "remain": "0",
            "release_id": "0",
            "custom_field_one": null,
            "custom_field_two": null,
            "custom_field_three": null,
            "custom_field_four": null,
            "custom_field_five": null,
            "custom_field_six": null,
            "custom_field_seven": null,
            "custom_field_eight": null,
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
            "custom_field_100": ""
        }
    },
    "info": "success"
}
```


## 同步复制需求到另外项目，同时设置状态、详细字段为同步字段

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&src_story_id=1010104801854843773&dst_workspace_id=755&sync_fields=description,status,owner' 'https://api.tapd.cn/stories/copy_story'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Story": {
            "id": "1000000755854845109",
            "workitem_type_id": "1000000755000000003",
            "name": "bbbbbbbb",
            "description": "<p><b><span style=\"color: #ff0000;\">作为</span></b>&nbsp;</p>\n<div>&nbsp;\n<div>\n<div><b><span style=\"color: #ff0000;\">我希望</span></b>&nbsp;</div>\n<div></div>\n<div><b><span style=\"color: #ff0000;\">以便</span></b> ADFADFADF</div>\n<div></div>\n<div>【验收标准】</div>\n<div>1、</div>\n<div>2、</div>\n<div>3、</div>\n<div>ADFADFDFDAADFFADS</div>\n</div>\n</div>",
            "workspace_id": "755",
            "creator": "anyechen",
            "created": "2020-12-09 16:49:47",
            "modified": "2020-12-09 16:49:47",
            "status": "planning",
            "owner": "",
            "cc": "",
            "begin": null,
            "due": null,
            "size": "0",
            "priority": "",
            "developer": "",
            "iteration_id": "0",
            "test_focus": "",
            "type": "",
            "source": "",
            "module": "",
            "version": "",
            "completed": null,
            "category_id": "-1",
            "path": "1000000755854845109",
            "parent_id": "0",
            "children_id": "|",
            "ancestor_id": "1000000755854845109",
            "business_value": null,
            "effort": "0",
            "effort_completed": "0",
            "exceed": "0",
            "remain": "0",
            "release_id": "0",
            "custom_field_one": null,
            "custom_field_two": null,
            "custom_field_three": null,
            "custom_field_four": null,
            "custom_field_five": null,
            "custom_field_six": null,
            "custom_field_seven": null,
            "custom_field_eight": null,
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
            "custom_field_100": ""
        }
    },
    "info": "success"
}
```


# 需求字段说明

## 需求重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">优先级</td></tr><tr><td style="text-align:center;">business_value</td><td style="text-align:center;">业务价值</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">版本</td></tr><tr><td style="text-align:center;">module</td><td style="text-align:center;">模块</td></tr><tr><td style="text-align:center;">test_focus</td><td style="text-align:center;">测试重点</td></tr><tr><td style="text-align:center;">size</td><td style="text-align:center;">规模</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">developer</td><td style="text-align:center;">开发人员</td></tr><tr><td style="text-align:center;">lastmodify</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">完成时间</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">迭代</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">预估工时</td></tr><tr><td style="text-align:center;">effort_completed</td><td style="text-align:center;">完成工时</td></tr><tr><td style="text-align:center;">remain</td><td style="text-align:center;">剩余工时</td></tr><tr><td style="text-align:center;">exceed</td><td style="text-align:center;">超出工时</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">需求分类</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">发布计划</td></tr><tr><td style="text-align:center;">source</td><td style="text-align:center;">来源</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">父需求</td></tr><tr><td style="text-align:center;">children_id</td><td style="text-align:center;">子需求</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">需求类别</td></tr></tbody></table>

## 需求优先级(priority)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">4</td><td style="text-align:center;">High</td></tr><tr><td style="text-align:center;">3</td><td style="text-align:center;">Middle</td></tr><tr><td style="text-align:center;">2</td><td style="text-align:center;">Low</td></tr><tr><td style="text-align:center;">1</td><td style="text-align:center;">Nice To Have</td></tr></tbody></table>

需求字段说明，请参考：[需求字段说明](/document/api-doc/API文档/api_reference/story/story.html)
