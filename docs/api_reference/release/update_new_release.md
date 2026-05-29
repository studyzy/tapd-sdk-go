# 更新发布计划接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/update_new_release.html

# 说明

更新发布计划，返回更新发布计划的数据 [旧版发布计划更新接口](https://open.tapd.cn/document/api-doc/API文档/api_reference/release/update_release.html)

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
updateNewRelease
```


# url

`https://api.tapd.cn/new_releases`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">发布计划ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态('open’或'done')</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&id=1010104801100003081&description=内容被更新' 'https://api.tapd.cn/new_releases'`

## 返回结果

```
{
    "status": 1,
    "data": {
        "Release": {
            "id": "1010104801100003081",
            "workspace_id": "10104801",
            "name": "test2",
            "description": "内容被更新",
            "startdate": "2020-10-20",
            "enddate": "2020-11-20",
            "creator": "dev",
            "created": "2020-10-27 11:09:14",
            "modified": "2020-10-27 11:24:48",
            "status": "open"
        }
    },
    "info": "success"
}
```


# 版本字段说明

## 版本字段重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">状态</td></tr></tbody></table>
