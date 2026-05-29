# 获取版本接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/get_versions.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [返回结果](#返回结果)
-   [版本重要字段说明](#版本重要字段说明)

# 说明

返回符合查询条件的所有版本（分页显示，默认一页30条）

# url

`https://api.tapd.cn/versions`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">版本ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">负责人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">提交人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">enum(0,1)</td><td style="text-align:center;">状态</td><td style="text-align:center;">0为未关闭（Closed），1为已关闭（Unclosed）</td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

`curl -u 'api_user:api_password' 'https://api.tapd.cn/versions?workspace_id=64851079'`

## 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Version": {
                "id": "1164851079001000271",
                "name": "test3",
                "description": null,
                "created": "2019-05-15 11:17:52",
                "owner": null,
                "due": null,
                "completed": "0",
                "default": "0",
                "parent_id": null,
                "path": null,
                "module_id": null,
                "start": null,
                "realend": null,
                "testtime": null,
                "realbegin": null,
                "releasetime": null,
                "business_module": null,
                "version_type": null,
                "modifier": null,
                "modified_time": null,
				            "status": "Unclosed"
                "workspace_id": "64851079",
                "creator": "tapd_api"
            }
        }
    ],
    "info": "success"
}
```


# 版本字段说明

## 版本重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">版本ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">负责人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">提交人</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">版本标题</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态（Unclosed, Closed）</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr></tbody></table>
