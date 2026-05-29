# 获取迭代变更历史

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iteration_changes.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下迭代变更历史](#获取项目下迭代变更历史)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [迭代变更历史重要字段说明](#迭代变更历史重要字段说明)

# 说明

返回符合查询条件的所有迭代变更历史（分页显示，默认一页30条）

# url

`https://api.tapd.cn/iteration_changes`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">ID</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">old_value</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更前</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">new_value</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">变更后</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">字段名称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下迭代变更历史

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/iteration_changes?workspace_id=20355782'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "IterationChange": {
                "id": "1020355782015033213",
                "iteration_id": "1020355782000700291",
                "author": "v_xinyucao",
                "field": "name",
                "old_value": null,
                "new_value": "对方的身份",
                "memo": null,
                "created": "2020-04-29 10:42:02",
                "modifyversion": "1588128122",
                "operater_type": "add",
                "workspace_id": "20355782"
            }
        }
    ],
    "info": "success"
}
```


# 迭代字段说明

## 迭代变更历史重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">ID</td></tr><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">field</td><td style="text-align:center;">字段名称</td></tr><tr><td style="text-align:center;">old_value</td><td style="text-align:center;">变更前</td></tr><tr><td style="text-align:center;">new_value</td><td style="text-align:center;">变更后</td></tr><tr><td style="text-align:center;">operater_type</td><td style="text-align:center;">变更类型</td></tr><tr><td style="text-align:center;">author</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr></tbody></table>
