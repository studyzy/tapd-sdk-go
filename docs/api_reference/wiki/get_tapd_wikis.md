# 获取 wiki

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis.html

# 说明

返回符合查询条件的所有Wiki（分页显示，默认一页30条）

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTapdWikis
```


# url

`https://api.tapd.cn/tapd_wikis`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">修改人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">note</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">备注</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">view_count</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">浏览量</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下的wiki

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tapd_wikis?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Wiki": {
                "id": "1210104801000043827",
                "name": "test888",
                "workspace_id": "10104801",
                "description": "",
                "markdown_description": "",
                "is_rich": "0",
                "parent_wiki_id": "0",
                "note": "",
                "view_count": "0",
                "created": "2020-08-25 11:24:44",
                "creator": "dev",
                "modified": "2020-08-25 11:24:44",
                "modifier": "dev"
            }
        },
        {
            "Wiki": {
                "id": "1210104801000043825",
                "name": "test888",
                "workspace_id": "10104801",
                "description": "",
                "markdown_description": "",
                "is_rich": "0",
                "parent_wiki_id": "0",
                "note": "",
                "view_count": "1",
                "created": "2020-08-25 11:22:24",
                "creator": "dev",
                "modified": "2020-08-25 11:22:24",
                "modifier": "dev"
            }
        },
        {
            "Wiki": {
                "id": "1210104801000043823",
                "name": "test888",
                "workspace_id": "10104801",
                "description": "",
                "markdown_description": "wfwwf4",
                "is_rich": "0",
                "parent_wiki_id": "0",
                "note": "",
                "view_count": "4",
                "created": "2020-08-25 11:22:20",
                "creator": "dev",
                "modified": "2020-08-25 15:39:05",
                "modifier": "dev"
            }
        }
    ],
    "info": "success"
}
```


# Wiki字段说明

## Wiki重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">富文本</td></tr><tr><td style="text-align:center;">markdown_description</td><td style="text-align:center;">Markdown</td></tr><tr><td style="text-align:center;">parent_wiki_id</td><td style="text-align:center;">父wiki ID</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">修改人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">note</td><td style="text-align:center;">备注</td></tr><tr><td style="text-align:center;">view_count</td><td style="text-align:center;">浏览量</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr></tbody></table>
