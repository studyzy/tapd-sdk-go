# 创建测试用例

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_tcase.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [创建个简单的测试用例](#创建个简单的测试用例)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [测试用例重要字段说明](#测试用例重要字段说明)
-   [测试用例类型(type)取值字段说明](#测试用例类型-type-取值字段说明)
-   [测试用例等级(priority)取值字段说明](#测试用例等级-priority-取值字段说明)
-   [测试用例状态(status)取值字段说明](#测试用例状态-status-取值字段说明)

# 说明

新建测试用例，返回新建测试用例的数据

# url

`https://api.tapd.cn/tcases`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">steps</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例步骤</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">用例目录</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">enum('updating','abandon','normal')</td><td style="text-align:center;">用例状态</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">用例名称</td></tr><tr><td style="text-align:center;">precondition</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">前置条件</td></tr><tr><td style="text-align:center;">expectation</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预期结果</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例类型</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例等级</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/tcase/get_tcase_custom_fields_settings.html" class="">获取测试用例自定义字段配置</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

## 创建个简单的测试用例

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'name=简单用例&workspace_id=10158231&cus_自定义字段的名称=custom_field_value' 'https://api.tapd.cn/tcases'`

### 返回结果

````
{
    "status": 1,
    "data": {
        "Tcase": {
            "id": "1010158231077224795",
            "steps": null,
            "workspace_id": "10158231",
            "category_id": "-1",
            "created": "2019-06-26 16:42:58",
            "modifier": "api_doc_oauth",
            "modified": "2019-06-26 16:42:58",
            "creator": "api_doc_oauth",
            "status": "normal",
            "name": "简单用例",
            "precondition": null,
            "expectation": null,
            "type": "",
            "priority": "",
            "cus_自定义字段的名称": "custom_field_value",
            "custom_field_1": null,
            "custom_field_2": null,
            "custom_field_3": null,
            "custom_field_4": null,
            "custom_field_5": null,
            "custom_field_6": null,
            "custom_field_7": null,
            "custom_field_8": null,
            "custom_field_9": null,
            "custom_field_10": null,
            "custom_field_11": null,
            "custom_field_12": null,
            "custom_field_13": null,
            "custom_field_14": null,
            "custom_field_15": null,
            "custom_field_16": null,
            "custom_field_17": null,
            "custom_field_18": null,
            "custom_field_19": null,
            "custom_field_20": null,
            "custom_field_21": "",
            "custom_field_22": "",
            "custom_field_23": "",
            "custom_field_24": "",
            "custom_field_25": "",
            "custom_field_26": "",
            "custom_field_27": "",
            "custom_field_28": "",
            "custom_field_29": "",
            "custom_field_30": "",
            "custom_field_31": null,
            "custom_field_32": null,
            "custom_field_33": null,
            "custom_field_34": null,
            "custom_field_35": null,
            "custom_field_36": null,
            "custom_field_37": null,
            "custom_field_38": null,
            "custom_field_39": null,
            "custom_field_40": null,
            "custom_field_41": null,
            "custom_field_42": null,
            "custom_field_43": null,
            "custom_field_44": null,
            "custom_field_45": null,
            "custom_field_46": null,
            "custom_field_47": null,
            "custom_field_48": null,
            "custom_field_49": null,
            "custom_field_50": null
        }
    },
    "info": "success"
}
```## 创建个稍微复杂的测试用例
### curl 使用 Basic Auth 鉴权调用示例
`curl -u 'api_user:api_password' -d 'type=其它&priority=高&status=待更新&steps=第一二三步&precondition=打开浏览器&expectation=无样式错误&creator=tapd&name=测试浏览器兼容性&workspace_id=10158231&cus_自定义字段的名称=custom_field_value' '{{ $page.apiHost }}/tcases'`


### 返回结果
```JSON
{
    "status": 1,
    "data": {
        "Tcase": {
            "id": "1010158231077224799",
            "steps": "第一二三步",
            "workspace_id": "10158231",
            "category_id": "-1",
            "created": "2019-06-26 16:42:59",
            "modifier": "tapd",
            "modified": "2019-06-26 16:42:59",
            "creator": "tapd",
            "status": "",
            "name": "测试浏览器兼容性",
            "precondition": "打开浏览器",
            "expectation": "无样式错误",
            "type": "其它",
            "priority": "高",
            "cus_自定义字段的名称": "custom_field_value",
            "custom_field_1": null,
            "custom_field_2": null,
            "custom_field_3": null,
            "custom_field_4": null,
            "custom_field_5": null,
            "custom_field_6": null,
            "custom_field_7": null,
            "custom_field_8": null,
            "custom_field_9": null,
            "custom_field_10": null,
            "custom_field_11": null,
            "custom_field_12": null,
            "custom_field_13": null,
            "custom_field_14": null,
            "custom_field_15": null,
            "custom_field_16": null,
            "custom_field_17": null,
            "custom_field_18": null,
            "custom_field_19": null,
            "custom_field_20": null,
            "custom_field_21": "",
            "custom_field_22": "",
            "custom_field_23": "",
            "custom_field_24": "",
            "custom_field_25": "",
            "custom_field_26": "",
            "custom_field_27": "",
            "custom_field_28": "",
            "custom_field_29": "",
            "custom_field_30": "",
            "custom_field_31": null,
            "custom_field_32": null,
            "custom_field_33": null,
            "custom_field_34": null,
            "custom_field_35": null,
            "custom_field_36": null,
            "custom_field_37": null,
            "custom_field_38": null,
            "custom_field_39": null,
            "custom_field_40": null,
            "custom_field_41": null,
            "custom_field_42": null,
            "custom_field_43": null,
            "custom_field_44": null,
            "custom_field_45": null,
            "custom_field_46": null,
            "custom_field_47": null,
            "custom_field_48": null,
            "custom_field_49": null,
            "custom_field_50": null
        }
    },
    "info": "success"
}
````


# 测试用例字段说明

## 测试用例重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">steps</td><td style="text-align:center;">用例步骤</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">用例目录</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">用例状态</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">用例名称</td></tr><tr><td style="text-align:center;">precondition</td><td style="text-align:center;">前置条件</td></tr><tr><td style="text-align:center;">expectation</td><td style="text-align:center;">预期结果</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">用例类型</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">用例等级</td></tr></tbody></table>

## 测试用例类型(type)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">功能测试</td><td style="text-align:center;">功能测试</td></tr><tr><td style="text-align:center;">性能测试</td><td style="text-align:center;">性能测试</td></tr><tr><td style="text-align:center;">安全性测试</td><td style="text-align:center;">安全性测试</td></tr><tr><td style="text-align:center;">其他</td><td style="text-align:center;">其他</td></tr></tbody></table>

## 测试用例等级(priority)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">高</td><td style="text-align:center;">高</td></tr><tr><td style="text-align:center;">中</td><td style="text-align:center;">中</td></tr><tr><td style="text-align:center;">低</td><td style="text-align:center;">低</td></tr></tbody></table>

## 测试用例状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">normal</td><td style="text-align:center;">正常</td></tr><tr><td style="text-align:center;">updating</td><td style="text-align:center;">待更新</td></tr><tr><td style="text-align:center;">abandon</td><td style="text-align:center;">已废弃</td></tr></tbody></table>
