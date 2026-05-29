# 获取测试计划

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plans.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [测试计划重要字段说明](#测试计划重要字段说明)
-   [状态(status)取值字段说明](#状态-status-取值字段说明)

# 说明

返回符合查询条件的所有测试计划（分页显示，默认一页30条）

# url

`https://api.tapd.cn/test_plans`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">否</td><td style="text-align:center;">int</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划标题</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划详细描述</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">创建时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">修改人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">最后修改时间</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试计划负责人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">start_date</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计开始</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">end_date</td><td style="text-align:center;">否</td><td style="text-align:center;">date</td><td style="text-align:center;">预计结束</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">版本号</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">测试类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">状态，默认开启，值为open</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">limit</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">设置返回数量限制，默认为30</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">page</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">order</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td><td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr><tr><td style="text-align:center;">fields</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">设置获取的字段，多个字段间以','逗号隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/test_plans?workspace_id=10104801'`

### 返回结果

```
{
    "status": 1,
    "data": [
        {
            "TestPlan": {
                "id": "1010104801000045351",
                "workspace_id": "10104801",
                "name": "aada",
                "description": "<div>adfa</div>",
                "version": "",
                "owner": "v_xuanfang;",
                "status": "open",
                "type": "",
                "start_date": null,
                "end_date": null,
                "creator": "anyechen",
                "created": "2020-01-09 12:12:37",
                "modified": "2020-08-21 10:58:08",
                "modifier": "v_xuanfang",
                "created_from": "",
                "custom_field_1": "",
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


# 测试计划字段说明

## 测试计划重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">标题</td></tr><tr><td style="text-align:center;">description</td><td style="text-align:center;">详细描述</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">版本</td></tr><tr><td style="text-align:center;">owner</td><td style="text-align:center;">处理人</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">类型</td></tr><tr><td style="text-align:center;">startdate</td><td style="text-align:center;">开始时间</td></tr><tr><td style="text-align:center;">enddate</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">结束时间</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr></tbody></table>

## 状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">open</td><td style="text-align:center;">开启</td></tr><tr><td style="text-align:center;">close</td><td style="text-align:center;">关闭</td></tr></tbody></table>
