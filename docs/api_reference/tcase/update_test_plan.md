# 编辑测试计划

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/update_test_plan.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [在项目下更新测试计划](#在项目下更新测试计划)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [状态(status)取值字段说明](#状态-status-取值字段说明)

# 说明

编辑测试计划，返回编辑后测试计划的数据

# url

`https://api.tapd.cn/test_plans`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次插入一条数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">int</td><td style="text-align:center;">测试计划ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划详细描述</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">修改人</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划负责人</td></tr><tr><td style="text-align:center;">start_date</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本号</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">float</td><td style="text-align:center;">测试类型</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态，默认开启，值为open</td></tr><tr><td style="text-align:center;">template_id</td><td style="text-align:center;">否</td><td style="text-align:center;">int</td><td style="text-align:center;">模板ID</td></tr><tr><td style="text-align:center;">cus_{$自定义字段别名}</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">自定义字段值，参数名会由后台自动转义为custom_field_*，如：cus_自定义字段的名称</td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数，具体字段名通过接口 <a href="/document/api-doc/API文档/api_reference/tcase/get_test_plan_fields_info.html" class="">获取测试计划自定义字段配置</a> 获取</td></tr></tbody></table>

# 调用示例及返回结果

## 在项目下更新测试计划

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'id=1000000755000016443&name=test&workspace_id=10158231' 'https://api.tapd.cn/test_plans'`

### 返回结果

```

{
    "status": 1,
    "data": {
        "TestPlan": {
            "id": "1000000755000016443",
            "workspace_id": "755",
            "name": "test_plan_12",
            "description": "这不是一个测试",
            "version": "123456",
            "owner": "",
            "status": "open",
            "type": "",
            "start_date": null,
            "end_date": null,
            "creator": "dev",
            "created": "2020-01-09 21:11:52",
            "modified": "2020-01-09 21:11:52",
            "modifier": "dev",
            "created_from": "api",
            "cus_自定义字字段的名称": "custom_field_value",
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


```


## 状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">open</td><td style="text-align:center;">开启</td></tr><tr><td style="text-align:center;">close</td><td style="text-align:center;">关闭</td></tr></tbody></table>
