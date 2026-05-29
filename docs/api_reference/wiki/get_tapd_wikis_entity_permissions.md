# 获取wiki可访问范围人员及用户组

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_entity_permissions.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [返回结果](#返回结果)

# 说明

返回符合查询条件的wiki可访问范围人员及用户组 注：仅当这个 wiki 的 is\_private 为 1 才需要调用此接口

# url

`https://api.tapd.cn/tapd_wikis_entity_permissions`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">wiki_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">wiki ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">target_type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">可访问的类型（role_id表示用户组，user_id表示用户）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">target_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">用户ID或用户组ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tapd_wikis_entity_permissions?workspace_id=10104801&wiki_id=1210104801001897607'`

## 返回结果

```
{
  "status": 1,
  "data": [
    {
      "EntityPermission": {
        "id": "1210158241000001519",
        "workspace_id": "10158241",
        "entry_type": "wiki",
        "target_type": "role_id",
        "target_id": "1000000000000000002",
        "wiki_id": "1210158241000048769"
      }
    },
    {
      "EntityPermission": {
        "id": "1210158241000001517",
        "workspace_id": "10158241",
        "entry_type": "wiki",
        "target_type": "role_id",
        "target_id": "1000000000000000009",
        "wiki_id": "1210158241000048769"
      }
    },
    {
      "EntityPermission": {
        "id": "1210158241000001515",
        "workspace_id": "10158241",
        "entry_type": "wiki",
        "target_type": "role_id",
        "target_id": "1000000000000000002",
        "wiki_id": "1210158241000048769"
      }
    },
    {
      "EntityPermission": {
        "id": "1210158241000001513",
        "workspace_id": "10158241",
        "entry_type": "wiki",
        "target_type": "nick",
        "target_id": "jmyan",
        "wiki_id": "1210158241000048769"
      }
    },
    {
      "EntityPermission": {
        "id": "1210158241000001511",
        "workspace_id": "10158241",
        "entry_type": "wiki",
        "target_type": "nick",
        "target_id": "austwayliu",
        "wiki_id": "1210158241000048769"
      }
    },
    {
      "EntityPermission": {
        "id": "1210158241000001509",
        "workspace_id": "10158241",
        "entry_type": "wiki",
        "target_type": "nick",
        "target_id": "austwayliu",
        "wiki_id": "1210158241000048769"
      }
    }
  ],
  "info": "success"
}
```


# 返回字段说明

##返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">记录ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">固定值 wiki</td></tr><tr><td style="text-align:center;">target_type</td><td style="text-align:center;">可访问的类型（role_id表示用户组，nick表示用户昵称）</td></tr><tr><td style="text-align:center;">target_id</td><td style="text-align:center;">用户昵称或用户组ID</td></tr><tr><td style="text-align:center;">wiki_id</td><td style="text-align:center;">wiki ID</td></tr></tbody></table>
