# 获取发布计划数量接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_new_releases_count.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [SDK 方法名](#sdk-方法名)
-   [获取项目下发布计划数量](#获取项目下发布计划数量)
    -   [curl 调用示例](#curl-调用示例)
    -   [返回结果](#返回结果)

# 说明

返回符合查询条件的发布计划数量 [旧版发布计划数量接口](https://open.tapd.cn/document/api-doc/API文档/api_reference/release/get_releases_count.html)

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getNewReleasesCount
```


# url

`https://api.tapd.cn/new_releases/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">开始时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">结束时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">结束时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">timestamp</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">enum('done','open')</td><td style="text-align:center;">状态</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下发布计划数量

### curl 调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/new_releases/count?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 1
    },
    "info": "success"
}
```

