# 创建需求

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/add_story.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [SDK 方法名](#sdk-方法名)
-   [在项目下创建需求](#在项目下创建需求)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

新建需求，返回新建需求的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
addStory
```


# url

`https://api.tapd.cn/stories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。为了兼容自定义优先级，<code>请使用 priority_label 字段</code>，详情参考：<a href="/document/api-doc/API文档/subject/custom_priority/" class="">如何兼容自定义优先级</a></td></tr><tr><td style="text-align:center;">priority_label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">优先级。推荐使用这个字段</td></tr><tr><td style="text-align:center;">business_value</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">业务价值</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本</td></tr><tr><td style="text-align:center;">module</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">模块</td></tr><tr><td style="text-align:center;">test_focus</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试重点</td></tr><tr><td style="text-align:center;">size</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">规模</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">cc</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">抄送人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">developer</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">开发人员</td></tr><tr><td style="text-align:center;">begin</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">templated_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">父需求ID</td></tr><tr><td style="text-align:center;">effort</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预估工时</td></tr><tr><td style="text-align:center;">effort_completed</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">完成工时</td></tr><tr><td style="text-align:center;">remain</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">剩余工时</td></tr><tr><td style="text-align:center;">exceed</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">超出工时</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求分类</td></tr><tr><td style="text-align:center;">workitem_type_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">需求类别</td></tr><tr><td style="text-align:center;">release_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">发布计划</td></tr><tr><td style="text-align:center;">source</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">来源</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">feature</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">特性</td></tr><tr><td style="text-align:center;">tech_risk</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">技术风险</td></tr><tr><td style="text-align:center;">business_value</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">业务价值</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">label</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标签，标签不存在时将自动创建，多个以英文坚线分格</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/story/get_story_custom_fields_settings.html" class="">获取需求自定义字段配置</a> 获取</td></tr><tr><td style="text-align:center;">custom_plan_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义计划应用参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/iteration/get_plan_apps.html" class="">获取自定义计划应用</a> 获取</td></tr><tr><td style="text-align:center;">is_apply_template_default_value</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否从模板继承默认值、保密设置（传值=1继承模板默认值）</td></tr><tr><td style="text-align:center;">apply_template</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">模版选项,支持多个选项传入,使用','隔开 如: "option1,option2" 当前支持参数:preset_stories(支持创建需求模板预设子需求),preset_tasks(支持创建需求模板预设子任务)</td></tr></tbody></table>

# 调用示例及返回结果

## 在项目下创建需求

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'name=story_created_by_api&workspace_id=10158231&cus_自定义字段的名称=custom_field_value' 'https://api.tapd.cn/stories'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "Story": {
            "id": "1010104801124922063",
            "workitem_type_id": "1010104801000022091",
            "name": "story_created_by_api",
            "description": null,
            "workspace_id": "10104801",
            "creator": "v_xuanfang",
            "created": "2025-06-16 14:42:59",
            "modified": "2025-06-16 14:42:59",
            "status": "planning",
            "step": "",
            "owner": "",
            "cc": "",
            "begin": null,
            "due": null,
            "size": null,
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
            "path": "1010104801124922063:",
            "parent_id": "0",
            "children_id": "|",
            "ancestor_id": "1010104801124922063",
            "level": "0",
            "business_value": null,
            "effort": null,
            "effort_completed": "0",
            "exceed": "0",
            "remain": "0",
            "release_id": "0",
            "bug_id": null,
            "templated_id": null,
            "created_from": "api",
            "feature": "",
            "label": "",
            "progress": "0",
            "is_archived": "0",
            "tech_risk": null,
            "flows": null,
            "custom_field_one": "",
            "custom_field_two": "",
            "custom_field_three": "",
            "custom_field_four": "",
            "custom_field_five": "",
            "custom_field_six": "",
            "custom_field_seven": "",
            "custom_field_eight": "",
            "secret_root_id": "0",
            "progress_manual": "0",
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
            "custom_field_101": "",
            "custom_field_102": "",
            "custom_field_103": "",
            "custom_field_104": "",
            "custom_field_105": "",
            "custom_field_106": "",
            "custom_field_107": "",
            "custom_field_108": "",
            "custom_field_109": "",
            "custom_field_110": "",
            "custom_field_111": "",
            "custom_field_112": "",
            "custom_field_113": "",
            "custom_field_114": "",
            "custom_field_115": "",
            "custom_field_116": "",
            "custom_field_117": "",
            "custom_field_118": "",
            "custom_field_119": "",
            "custom_field_120": "",
            "custom_field_121": "",
            "custom_field_122": "",
            "custom_field_123": "",
            "custom_field_124": "",
            "custom_field_125": "",
            "custom_field_126": "",
            "custom_field_127": "",
            "custom_field_128": "",
            "custom_field_129": "",
            "custom_field_130": "",
            "custom_field_131": "",
            "custom_field_132": "",
            "custom_field_133": "",
            "custom_field_134": "",
            "custom_field_135": "",
            "custom_field_136": "",
            "custom_field_137": "",
            "custom_field_138": "",
            "custom_field_139": "",
            "custom_field_140": "",
            "custom_field_141": "",
            "custom_field_142": "",
            "custom_field_143": "",
            "custom_field_144": "",
            "custom_field_145": "",
            "custom_field_146": "",
            "custom_field_147": "",
            "custom_field_148": "",
            "custom_field_149": "",
            "custom_field_150": "",
            "custom_field_151": "",
            "custom_field_152": "",
            "custom_field_153": "",
            "custom_field_154": "",
            "custom_field_155": "",
            "custom_field_156": "",
            "custom_field_157": "",
            "custom_field_158": "",
            "custom_field_159": "",
            "custom_field_160": "",
            "custom_field_161": "",
            "custom_field_162": "",
            "custom_field_163": "",
            "custom_field_164": "",
            "custom_field_165": "",
            "custom_field_166": "",
            "custom_field_167": "",
            "custom_field_168": "",
            "custom_field_169": "",
            "custom_field_170": "",
            "custom_field_171": "",
            "custom_field_172": "",
            "custom_field_173": "",
            "custom_field_174": "",
            "custom_field_175": "",
            "custom_field_176": "",
            "custom_field_177": "",
            "custom_field_178": "",
            "custom_field_179": "",
            "custom_field_180": "",
            "custom_field_181": "",
            "custom_field_182": "",
            "custom_field_183": "",
            "custom_field_184": "",
            "custom_field_185": "",
            "custom_field_186": "",
            "custom_field_187": "",
            "custom_field_188": "",
            "custom_field_189": "",
            "custom_field_190": "",
            "custom_field_191": "",
            "custom_field_192": "",
            "custom_field_193": "",
            "custom_field_194": "",
            "custom_field_195": "",
            "custom_field_196": "",
            "custom_field_197": "",
            "custom_field_198": "",
            "custom_field_199": "",
            "custom_field_200": "",
            "custom_plan_field_1": "0",
            "custom_plan_field_2": "0",
            "custom_plan_field_3": "0",
            "custom_plan_field_4": "0",
            "custom_plan_field_5": "0",
            "custom_plan_field_6": "0",
            "custom_plan_field_7": "0",
            "custom_plan_field_8": "0",
            "custom_plan_field_9": "0",
            "custom_plan_field_10": "0",
            "priority_label": ""
        }
    },
    "info": "success"
}
```


需求字段说明，请参考：[需求字段说明](/document/api-doc/API文档/api_reference/story/story.html)
