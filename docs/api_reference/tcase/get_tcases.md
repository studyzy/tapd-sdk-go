# 获取测试用例

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcases.html

# 说明

返回符合查询条件的所有测试用例（分页显示，默认一页30条）

## SDK 方法名

nodeJs:

[SDK使用方式 (opens new window)](https://open.tapd.cn/document/api-doc/SDK/node-README.html)

[插件中使用SDK](/document/plugin-doc/learning/api-and-security/index.html)

方法名::

```
getTcases
```


# url

`https://api.tapd.cn/tcases`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">支持多ID查询</td></tr><tr><td style="text-align:center;">steps</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例步骤</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">用例目录</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">创建时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">最后修改人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">datetime</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;">支持时间查询</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">enum('updating','abandon','normal')</td><td style="text-align:center;">用例状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例名称</td><td style="text-align:center;">支持模糊匹配</td></tr><tr><td style="text-align:center;">precondition</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">前置条件</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">expectation</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预期结果</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例等级</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 获取</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下测试用例

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tcases?workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "Tcase": {
                "id": "1120003271001000049",
                "steps": null,
                "workspace_id": "10158231",
                "category_id": "-1",
                "created": "2019-06-26 16:42:59",
                "modifier": "api_doc_oauth",
                "modified": "2019-06-26 16:43:00",
                "creator": "",
                "status": "abandon",
                "name": "",
                "precondition": null,
                "expectation": null,
                "type": "",
                "priority": "",
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
                "custom_field_21": null,
                "custom_field_22": null,
                "custom_field_23": null,
                "custom_field_24": null,
                "custom_field_25": null,
                "custom_field_26": null,
                "custom_field_27": null,
                "custom_field_28": null,
                "custom_field_29": null,
                "custom_field_30": null,
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
        {
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
        }
    ],
    "info": "success"
}
```


# 测试用例字段说明

## 测试用例重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">steps</td><td style="text-align:center;">用例步骤</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">用例目录</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">用例状态</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">用例名称</td></tr><tr><td style="text-align:center;">precondition</td><td style="text-align:center;">前置条件</td></tr><tr><td style="text-align:center;">expectation</td><td style="text-align:center;">预期结果</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">用例类型</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">用例等级</td></tr></tbody></table>

## 测试用例类型(type)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">功能测试</td><td style="text-align:center;">功能测试</td></tr><tr><td style="text-align:center;">性能测试</td><td style="text-align:center;">性能测试</td></tr><tr><td style="text-align:center;">安全性测试</td><td style="text-align:center;">安全性测试</td></tr><tr><td style="text-align:center;">其他</td><td style="text-align:center;">其他</td></tr></tbody></table>

## 测试用例等级(priority)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">高</td><td style="text-align:center;">高</td></tr><tr><td style="text-align:center;">中</td><td style="text-align:center;">中</td></tr><tr><td style="text-align:center;">低</td><td style="text-align:center;">低</td></tr></tbody></table>

## 测试用例状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">normal</td><td style="text-align:center;">正常</td></tr><tr><td style="text-align:center;">updating</td><td style="text-align:center;">待更新</td></tr><tr><td style="text-align:center;">abandon</td><td style="text-align:center;">已废弃</td></tr></tbody></table>
