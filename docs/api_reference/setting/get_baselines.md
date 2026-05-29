# 获取基线接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/get_baselines.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下模块的数量](#获取项目下模块的数量)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [基线重要字段说明](#基线重要字段说明)

# 说明

计算符合查询条件的基线并返回

# url

`https://api.tapd.cn/baselines`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">标题</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">version_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">版本</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">详细描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">处理人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下模块的数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/baselines?workspace_id=10104801'`

`curl 'https://api.tapd.cn/baselines?workspace_id=10104801&access_token=ACCESS_TOKEN'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Baseline": {
                "id": "1010104801101308343",
                "name": "创建1",
                "description": "更新基线1",
                "created": "2020-11-30 17:09:52",
                "owner": "",
                "due": null,
                "completed": "0",
                "default": "0",
                "version_id": "1010104801001473663",
                "module_id": "0",
                "svn_tag": "0",
                "svn_project_id": "0",
                "svn_path_id": "0",
                "svn_sync_type": "",
                "workspace_id": "10104801"
            }
        },
        {
            "Baseline": {
                "id": "1010104801101007975",
                "name": "b1",
                "description": "",
                "created": "2017-09-11 22:47:02",
                "owner": "",
                "due": null,
                "completed": "0",
                "default": "0",
                "version_id": "0",
                "module_id": "0",
                "svn_tag": "0",
                "svn_project_id": "0",
                "svn_path_id": "0",
                "svn_sync_type": "",
                "workspace_id": "10104801"
            }
        }
    ],
    "info": "success"
}
```


# 基线字段说明

## 基线重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">version_id</td><td style="text-align:center;">版本ID</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">completed</td><td style="text-align:center;">0（未完成）1（完成状态）</td></tr><tr><td style="text-align:center;">due</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr></tbody></table>
