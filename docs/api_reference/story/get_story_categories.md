# 获取需求分类

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_categories.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [SDK 方法名](#sdk-方法名)
-   [获取项目下需求分类](#获取项目下需求分类)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [需求分类重要字段说明](#需求分类重要字段说明)

# 说明

返回符合查询条件的所有需求分类（分页显示，默认一页30条）

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getStoryCategories
```


# url

`https://api.tapd.cn/story_categories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求分类名称</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">需求分类描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">父分类ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下需求分类

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/story_categories?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Category": {
                "id": "-1",
                "workspace_id": "10158231",
                "name": "未分类",
                "description": "未分类",
                "parent_id": "0",
                "modified": "2017-06-20 14:05:53",
                "created": "2017-06-20 14:05:53",
                "creator": null,
                "modifier": "sunyoungsun"
            }
        },
        {
            "Category": {
                "id": "1010158231000072437",
                "workspace_id": "10158231",
                "name": "优化需求",
                "description": "优化需求",
                "parent_id": "1010158231000072431",
                "modified": "2017-06-20 14:05:13",
                "created": "2017-06-20 14:05:13",
                "creator": null,
                "modifier": "ruirayli"
            }
        },
        {
            "Category": {
                "id": "1010158231000072435",
                "workspace_id": "10158231",
                "name": "新功能",
                "description": "新功能",
                "parent_id": "1010158231000072431",
                "modified": "2017-06-20 14:05:13",
                "created": "2017-06-20 14:05:13",
                "creator": null,
                "modifier": "ruirayli"
            }
        },
        {
            "Category": {
                "id": "1010158231000072433",
                "workspace_id": "10158231",
                "name": "技术需求",
                "description": "技术需求",
                "parent_id": "0",
                "modified": "2017-06-20 14:05:13",
                "created": "2017-06-20 14:05:13",
                "creator": null,
                "modifier": "ruirayli"
            }
        },
        {
            "Category": {
                "id": "1010158231000072431",
                "workspace_id": "10158231",
                "name": "产品需求",
                "description": "产品需求",
                "parent_id": "0",
                "modified": "2017-06-20 14:05:13",
                "created": "2017-06-20 14:05:13",
                "creator": null,
                "modifier": "ruirayli"
            }
        }
    ],
    "info": "success"
}
```


# 需求分类字段说明

## 需求分类重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">需求分类名称</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">需求分类描述</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">父分类ID</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr></tbody></table>
