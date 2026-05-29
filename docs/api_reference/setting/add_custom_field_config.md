# 创建自定义字段（需求及缺陷）

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/add_custom_field_config.html

-   [说明](#说明)
-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [url](#url)
-   [支持格式](#支持格式)
-   [HTTP请求方式](#http请求方式)
-   [请求数限制](#请求数限制)
-   [请求参数](#请求参数)
    -   [请求参数 option 详细说明](#请求参数-option-详细说明)
-   [调用示例及返回结果](#调用示例及返回结果)
    -   [创建需求自定义字段配置](#创建需求自定义字段配置)
    -   [创建需求选类型checkbox自定义字段配置](#创建需求选类型checkbox自定义字段配置)
    -   [创建缺陷级联多选字段](#创建缺陷级联多选字段)
-   [返回字段说明](#返回字段说明)

## 说明

创建自定义字段（需求及缺陷）

## url

`https://api.tapd.cn/custom_field_configs`

## 支持格式

JSON/XML（默认JSON格式）

## HTTP请求方式

POST

## 请求数限制

-   不能超过自定义字段的最大个数
-   选项个数不能超过 500

## 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">新增字段分类</td><td style="text-align:center;">story(需求) bug(缺陷) task(任务) tcase(测试用例)</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">字段名称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">字段类型</td><td style="text-align:center;">select 单选下拉列表、multi_select 多选下拉框、text 单行文本框、checkbox 复选框 、radio 单选框 、textarea 多行文本框 、 user_chooser 人名输入框 、dateinput 日期 、datetime 时分类型 、float 数值型、integer 整型、 cascade_checkbox 级联复选、 cascade_radio 级联单选</td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">string</td><td style="text-align:center;">字段备注</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">option</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">string</td><td style="text-align:center;">字段值选项</td><td style="text-align:center;">详细使用说明见下面</td></tr></tbody></table>

### 请求参数 option 详细说明

1.  字段类型为 select、 multi\_select、radio、 checkbox、 cascade\_checkbox、cascade\_radio 类型时为必填
2.  字段类型为 select、 multi\_select、radio、 checkbox时，选项参数格式如：AA|BB|CC 格式，使用“|”隔开
3.  字段类型为 cascade\_checkbox、cascade\_radio 级联类型时，选项参数格式必须是 json 结构，name 属性表示选项名， children 表示子选项。示例：`[{"name":"1 ","children":[{"name":"11","children":[{"name":"111"},{"name":"112"}]}]},{"name":" 2","children":[{"name":"21","children":[{"name":"211"},{"name":"212"}]}]}]` ，展开如下：

```
[
    {
        "name": "1  ",
        "children": [
            {
                "name": "11",
                "children": [
                    {
                        "name": "111"
                    },
                    {
                        "name": "112"
                    }
                ]
            }
        ]
    },
    {
        "name": " 2",
        "children": [
            {
                "name": "21",
                "children": [
                    {
                        "name": "211"
                    },
                    {
                        "name": "212"
                    }
                ]
            }
        ]
    }
]
```


## 调用示例及返回结果

### 创建需求自定义字段配置

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'entry_type=story&name=add_field&type=text&workspace_id=10158231' 'https://api.tapd.cn/custom_field_configs'`

#### 返回结果

```
{
    "status": 1,
    "data": {
        "CustomFieldConfig": {
            "id": "1010158231215016293",
            "workspace_id": "10158231",
            "entry_type": "story",
            "custom_field": "custom_field_three",
            "type": "text",
            "name": "add_field",
            "options": null,
            "enabled": "1",
            "sort": null,
            "memo": ""
        }
    },
    "info": "success"
}
```


### 创建需求选类型checkbox自定义字段配置

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'entry_type=story&name=add_field2&memo=add_field_detail&type=checkbox&option=item1|item2|item3&workspace_id=10158231' 'https://api.tapd.cn/custom_field_configs'`

#### 返回结果

```
{
    "status": 1,
    "data": {
        "CustomFieldConfig": {
            "id": "1010158231215016295",
            "workspace_id": "10158231",
            "entry_type": "story",
            "custom_field": "custom_field_four",
            "type": "checkbox",
            "name": "add_field2",
            "options": "{\"1\":\"item1\",\"2\":\"item2\",\"3\":\"item3\"}",
            "enabled": "1",
            "sort": null,
            "memo": "add_field_detail"
        }
    },
    "info": "success"
}
```


### 创建缺陷级联多选字段

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=51240826&entry_type=bug&option=%5B%7B%22name%22%3A%221%20%20%22%2C%22children%22%3A%5B%7B%22name%22%3A%2211%22%2C%22children%22%3A%5B%7B%22name%22%3A%22111%22%7D%2C%7B%22name%22%3A%22112%22%7D%5D%7D%5D%7D%2C%7B%22name%22%3A%22%202%22%2C%22children%22%3A%5B%7B%22name%22%3A%2221%22%2C%22children%22%3A%5B%7B%22name%22%3A%22211%22%7D%2C%7B%22name%22%3A%22212%22%7D%5D%7D%5D%7D%5D&type=cascade_checkbox&name=%E7%BC%BA%E9%99%B7%E7%BA%A7%E8%81%94tt' 'https://api.tapd.cn/custom_field_configs'`

#### 返回结果

```
{
    "status": 1,
    "data": {
        "CustomFieldConfig": {
            "id": "1151240826001007121",
            "workspace_id": "51240826",
            "entry_type": "bug",
            "custom_field": "custom_field_three",
            "type": "cascade_checkbox",
            "name": "缺陷级联tt",
            "options": "[{\"name\":\"1\",\"children\":[{\"name\":\"11\",\"children\":[{\"name\":\"111\"},{\"name\":\"112\"}]}]},{\"name\":\"2\",\"children\":[{\"name\":\"21\",\"children\":[{\"name\":\"211\"},{\"name\":\"212\"}]}]}]",
            "extra_config": null,
            "enabled": "1",
            "sort": "3",
            "memo": null
        }
    },
    "info": "success"
}
```


## 返回字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">自定义字段配置的ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">所属项目ID</td></tr><tr><td style="text-align:center;">entry_type</td><td style="text-align:center;">所属实体对象(story 、bug)</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">自定义字段名称</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">输入字段类型</td></tr><tr><td style="text-align:center;">created_from</td><td style="text-align:center;">请求来源</td></tr><tr><td style="text-align:center;">enabled</td><td style="text-align:center;">是否启用</td></tr><tr><td style="text-align:center;">memo</td><td style="text-align:center;">自定义字段显示备注</td></tr><tr><td style="text-align:center;">custom_field</td><td style="text-align:center;">自定义字段预留个数名称</td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;">自定义字段可选值</td></tr><tr><td style="text-align:center;">enabled</td><td style="text-align:center;">是否启用</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">显示时排序系数</td></tr></tbody></table>
