# 获取指定commit关联的TAPD对象（需求、任务、缺陷）

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/source/get_commit_objects.html

# 说明

获取指定commit关联的业务对象

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getCodeCommitObjects
```


# url

`https://api.tapd.cn/code_commit_objects/workitems`

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">commit_id</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">提交ID</td><td style="text-align:center;">支持多个commit，以逗号(,)隔开</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">业务对象类型</td><td style="text-align:center;">story-需求，bug-缺陷，task-任务</td></tr><tr><td style="text-align:center;">scm_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">来源类型默认"p4"候选值"tgit,云端有gitlab，github，tgit"</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -H 'Authorization: Bearer Access Token' -X GET 'https://api.tapd.cn/code_commit_objects/workitems?commit_id=7b0645c6a467a502fe1d3b678fea8bdf2890aa8d&entity_type=story&workspace_id=20355782'`

### 返回结果

```
{
  "status": 1,
  "data": [
    {
      "Task": {
        "id": "1020355782500602947",
        "name": "666",
        "description": null,
        "workspace_id": "20355782",
        "creator": "v_tingtdong",
        "created": "2019-10-30 16:26:45",
        "modified": "2019-10-30 16:26:45",
        "status": "open",
        "owner": "",
        "cc": "",
        "begin": null,
        "due": null,
        "story_id": "1020355782500697089",
        "iteration_id": "1020355782000390769",
        "priority": "",
        "progress": "0",
        "completed": null,
        "effort_completed": "0",
        "exceed": "0",
        "remain": "0",
        "effort": "0",
        "has_attachment": "0",
        "release_id": "1020355782500697089",
        "label": null,
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
        "custom_field_50": ""
      }
    }
  ],
  "info": "success"
}
```


# 需求字段说明

请参考 [获取需求](/document/api-doc/API文档/api_reference/story/get_stories.html)

# 缺陷字段说明

请参考 [获取缺陷](/document/api-doc/API文档/api_reference/bug/get_bugs.html)

# 任务字段说明

请参考 [获取任务](/document/api-doc/API文档/api_reference/task/get_tasks.html)
