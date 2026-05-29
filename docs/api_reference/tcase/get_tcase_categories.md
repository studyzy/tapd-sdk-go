# 获取测试用例目录

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_categories.html

# 说明

返回符合查询条件的所有测试用例目录（分页显示，默认一页30条）

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTcaseCategories
```


# url

`https://api.tapd.cn/tcase_categories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录名称</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">父目录ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录最后修改人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">sorting</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">目录排序序号</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下测试用例目录

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tcase_categories?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "TcaseCategory": {
                "id": "1010158231075917759",
                "workspace_id": "10158231",
                "name": "None Category",
                "description": "未规划目录",
                "parent_id": "0",
                "modified": "2019-06-26 16:42:50",
                "created": "2019-06-26 16:42:50",
                "creator": null,
                "modifier": null,
                "sorting": "0"
            }
        },
        {
            "TcaseCategory": {
                "id": "1010158231000082521",
                "workspace_id": "10158231",
                "name": "用例目录3",
                "description": null,
                "parent_id": "0",
                "modified": "2017-06-20 16:48:37",
                "created": "2017-06-20 16:48:37",
                "creator": "system",
                "modifier": "system",
                "sorting": null
            }
        },
        {
            "TcaseCategory": {
                "id": "1010158231000082519",
                "workspace_id": "10158231",
                "name": "用例目录2",
                "description": null,
                "parent_id": "0",
                "modified": "2017-06-20 16:48:37",
                "created": "2017-06-20 16:48:37",
                "creator": "system",
                "modifier": "system",
                "sorting": null
            }
        },
        {
            "TcaseCategory": {
                "id": "1010158231000082517",
                "workspace_id": "10158231",
                "name": "用例目录1",
                "description": null,
                "parent_id": "0",
                "modified": "2017-06-20 16:48:37",
                "created": "2017-06-20 16:48:37",
                "creator": "system",
                "modifier": "system",
                "sorting": null
            }
        }
    ],
    "info": "success"
}
```


# 测试用例目录字段说明

## 测试用例目录重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">目录名称</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">目录描述</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">父目录ID</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">目录创建人</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">目录最后修改人</td></tr><tr><td style="text-align:center;">sorting</td><td style="text-align:center;">目录排序序号</td></tr></tbody></table>
