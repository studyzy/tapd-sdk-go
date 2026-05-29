# 创建测试用例目录

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_tcase_category.html

# 说明

创建测试用例目录，返回创建测试用例目录的数据

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
addTcaseCategory
```


# url

`https://api.tapd.cn/tcase_categories`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">目录名称</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录描述</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">父目录ID</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr></tbody></table>

# 调用示例及返回结果

## 创建测试用例目录

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=20355782&name=test' 'https://api.tapd.cn/tcase_categories'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "TcaseCategory": {
            "id": "1020355782075922101",
            "workspace_id": "20355782",
            "name": "test",
            "description": null,
            "parent_id": "0",
            "modified": "2020-05-26 15:04:19",
            "created": "2020-05-26 15:04:19",
            "creator": "v_xuanfang",
            "modifier": "v_xuanfang",
            "sorting": "0"
        }
    },
    "info": "success"
}
```


# 测试用例目录字段说明

## 测试用例目录重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">目录名称</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">目录描述</td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">父目录ID</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">目录创建人</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">目录最后修改人</td></tr><tr><td style="text-align:center;">sorting</td><td style="text-align:center;">目录排序序号</td></tr></tbody></table>
