# 获取测试用例目录数量

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_categories_count.html

# 说明

计算符合查询条件的测试计划数量并返回

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTcaseCategoriesCount
```


# url

`https://api.tapd.cn/tcase_categories/count`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

只返回测试计划数量

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录名称</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">parent_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">父目录ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">目录最后修改人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">sorting</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">目录排序序号</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下用例目录数量

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tcase_categories/count?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "count": 4
    },
    "info": "success"
}
```

