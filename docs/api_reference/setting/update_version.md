# 更新版本接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/update_version.html

# 说明

更新版本，返回版本更新后的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
updateVersion
```


# url

`https://api.tapd.cn/versions`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">Version ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">当前处理人</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">提交人</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">负责人</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u "api_user:password" -d "name=uptest3&id=1164851079001000271&workspace_id=64851079&modifier=tapd_api" "https://api.tapd.cn/versions"`

## 返回结果

```
{
    "status": 1,
    "data": {
        "Version": {
            "id": "1164851079001000271",
            "name": "uptest3",
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
            "modifier": "tapd_api",
            "modified_time": null,
            "workspace_id": "64851079",
            "creator": "tapd_api"
        }
    },
    "info": "success"
}
```


# 版本字段说明

## 版本重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">Version ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">当前处理人</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">提交人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">上次修改时间</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">负责人</td></tr></tbody></table>
