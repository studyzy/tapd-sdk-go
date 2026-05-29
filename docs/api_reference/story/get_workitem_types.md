# 获取需求类别

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_workitem_types.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下需求类别](#获取项目下需求类别)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [需求类别重要字段说明](#需求类别重要字段说明)

# 说明

返回符合查询条件的所有需求类别（分页显示，默认一页30条）

# url

`https://api.tapd.cn/workitem_types`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求类别名称</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">类别别名</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">english_name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">英文名称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workflow_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">工作流ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modified_by</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下需求类别

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/workitem_types?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "WorkitemType": {
                "id": "1010104801000066691",
                "workspace_id": "10104801",
                "app_id": "1",
                "entity_type": "story",
                "name": "bbbee",
                "english_name": "be",
                "status": "3",
                "color": "#3582fb",
                "workflow_id": "1010104801000050043",
                "children_ids": "|",
                "parent_ids": "",
                "icon": "10104801/icon/1010104801503447605",
                "icon_small": "10104801/icon/1010104801503447607",
                "creator": "anyechen",
                "created": "2021-01-26 16:42:16",
                "modified_by": "anyechen",
                "modified": "2024-01-17 16:03:35",
                "icon_viper": "https://viper.wolf.woa.com/icon/files/10104801/icon/1010104801503447605.png?token=6c47847284f7c6b85cc484c32c07c152f4f2d5afd0229b9a27ad0b83c0cdb01b&version=1705478615",
                "icon_small_viper": "https://viper.wolf.woa.com/icon/files/10104801/icon/1010104801503447607.png?token=3369d6880a360a83ff09377aeea128b2b11804184b7d964f504c49bf75125791&version=1705478615"
            }
        },
        {
            "WorkitemType": {
                "id": "1010104801000037383",
                "workspace_id": "10104801",
                "app_id": "1",
                "entity_type": "story",
                "name": "需求",
                "english_name": "story",
                "status": "3",
                "color": "#3582fb",
                "workflow_id": "1010104801000128627",
                "children_ids": "",
                "parent_ids": "",
                "icon": "",
                "icon_small": "",
                "creator": "TAPD system",
                "created": "2020-11-13 17:02:52",
                "modified_by": "",
                "modified": "2024-01-03 19:56:58",
                "icon_viper": "https://wolf.woa.com//img/workitem_type/default_icon/@2/story.png",
                "icon_small_viper": "https://wolf.woa.com//img/workitem_type/default_icon/@2/story_small.png"
            }
        }
    ],
    "info": "success"
}
```


# 需求类别字段说明

## 需求类别重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">需求类别名称</td></tr><tr><td style="text-align:center;">entity_type</td><td style="text-align:center;">类别别名</td></tr><tr><td style="text-align:center;">english_name</td><td style="text-align:center;">英文名称</td></tr><tr><td style="text-align:center;">workflow_id</td><td style="text-align:center;">工作流ID</td></tr><tr><td style="text-align:center;">children_ids</td><td style="text-align:center;">允许的子需求类别。为空是允许创建任何类别子需求；为 | 是不允许创建子需求；其它则为指定类别子需求。</td></tr><tr><td style="text-align:center;">parent_ids</td><td style="text-align:center;">允许的父需求类别。为空是没限制；其它为必须选择指定类别父需求。</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">modified_by</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">color</td><td style="text-align:center;">颜色</td></tr><tr><td style="text-align:center;">icon</td><td style="text-align:center;">图标</td></tr><tr><td style="text-align:center;">icon_small</td><td style="text-align:center;">大图标地址</td></tr><tr><td style="text-align:center;">icon_viper</td><td style="text-align:center;">图标地址</td></tr><tr><td style="text-align:center;">icon_small_viper</td><td style="text-align:center;">小图标地址</td></tr></tbody></table>

# 状态说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">1</td><td style="text-align:center;">未完成</td></tr><tr><td style="text-align:center;">2</td><td style="text-align:center;">未启用</td></tr><tr><td style="text-align:center;">3</td><td style="text-align:center;">已启用</td></tr></tbody></table>
