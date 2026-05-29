# 获取测试计划测试结果

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_details.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [测试重要字段说明](#测试重要字段说明)
-   [测试结果(TcaseResult)中字段说明](#测试结果-tcaseresult-中字段说明)
-   [结果状态(result\_status)取值字段说明](#结果状态-result-status-取值字段说明)
-   [用例状态(status)取值字段说明](#用例状态-status-取值字段说明)

# 说明

获取测试计划测试结果

# url

`https://api.tapd.cn/test_plans/details`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

默认返回所有数据。

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td></tr><tr><td style="text-align:center;">last_executor</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">最后执行人</td></tr><tr><td style="text-align:center;">include_repeat</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">include_repeat=1 可以获取到所有数据</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/test_plans/details?workspace_id=10158231&id=1010158231000005241'`

### 返回结果

```

{
	"status": 1,
	"data": [{
		"Tcase": {
			"id": "1010158231075919347",
			"mid": "1010158231075919347",
			"steps": null,
			"workspace_id": "10158231",
			"category_id": "-1",
			"version": "0",
			"created": "2017-06-20 16:49:29",
			"modifier": "anyechen",
			"modified": "2017-06-20 16:49:29",
			"creator": "anyechen",
			"status": "normal",
			"name": "Firefox\u6d4f\u89c8\u5668\u517c\u5bb9\u6027\u6d4b\u8bd5",
			"precondition": null,
			"expectation": null,
			"sort": "0",
			"indexcode": "",
			"type": "",
			"priority": "",
			"template_id": "0",
			"created_from": "",
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
			"custom_field_50": null,
			"TcaseResult": {
				"id": 0,
				"tcase_id": 1010158231075919347,
				"created": "",
				"test_plan_id": 0,
				"result_status": "unexecuted",
				"result_remark": "",
				"workspace_id": 10158231,
				"status": "",
				"executor": "",
				"executed_at": ""
			}
		}
	}, {
		"Tcase": {
			"id": "1010158231075919345",
			"mid": "1010158231075919345",
			"steps": null,
			"workspace_id": "10158231",
			"category_id": "-1",
			"version": "0",
			"created": "2017-06-20 16:49:28",
			"modifier": "anyechen",
			"modified": "2017-06-20 16:49:29",
			"creator": "anyechen",
			"status": "normal",
			"name": "Chrome\u6d4f\u89c8\u5668\u517c\u5bb9\u6027\u6d4b\u8bd5",
			"precondition": null,
			"expectation": null,
			"sort": "0",
			"indexcode": "",
			"type": "",
			"priority": "",
			"template_id": "0",
			"created_from": "",
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
			"custom_field_50": null,
			"TcaseResult": {
				"id": "1010158231000703323",
				"tcase_id": "1010158231077230667",
				"created": "2020-02-12 23:42:56",
				"test_plan_id": "1010158231000005241",
				"result_status": "block",
				"result_remark": "xxx",
				"workspace_id": "10158231",
				"status": "normal",
				"executor": "anyechen",
				"executed_at": "2020-02-12 23:42:56"
			}
		}
	}, {
		"Tcase": {
			"id": "1010158231075919341",
			"mid": "1010158231075919341",
			"steps": null,
			"workspace_id": "10158231",
			"category_id": "-1",
			"version": "0",
			"created": "2017-06-20 16:49:26",
			"modifier": "anyechen",
			"modified": "2017-06-20 16:49:26",
			"creator": "anyechen",
			"status": "normal",
			"name": "IE\u6d4f\u89c8\u5668\u517c\u5bb9\u6027\u6d4b\u8bd5",
			"precondition": null,
			"expectation": null,
			"sort": "0",
			"indexcode": "",
			"type": "",
			"priority": "",
			"template_id": "0",
			"created_from": "",
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
			"custom_field_50": null,
			"TcaseResult": {
				"id": "1010158231000018041",
				"tcase_id": "1010158231075919343",
				"created": "2017-06-20 16:49:28",
				"test_plan_id": "1010158231000005241",
				"result_status": "pass",
				"result_remark": null,
				"workspace_id": "10158231",
				"status": "normal",
				"executor": "ruirayli",
				"executed_at": "2017-06-20 16:49:28"
			}
		}
	}],
	"info": "success"
}


```


# 测试字段说明

## 测试重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">用例ID</td></tr><tr><td style="text-align:center;">steps</td><td style="text-align:center;">用例步骤</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">用例目录</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">modifier</td><td style="text-align:center;">最后修改人</td></tr><tr><td style="text-align:center;">modified</td><td style="text-align:center;">最后修改时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">version</td><td style="text-align:center;">版本</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">用例状态</td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">用例名称</td></tr><tr><td style="text-align:center;">precondition</td><td style="text-align:center;">前置条件</td></tr><tr><td style="text-align:center;">expectation</td><td style="text-align:center;">预期结果</td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">用例类型</td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">用例等级</td></tr></tbody></table>

## 测试结果(TcaseResult)中字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">测试结果id</td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;">测试用例id</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">创建时间</td></tr><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;">测试计划id</td></tr><tr><td style="text-align:center;">result_status</td><td style="text-align:center;">结果状态</td></tr><tr><td style="text-align:center;">result_remark</td><td style="text-align:center;">结果备注</td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">状态</td></tr><tr><td style="text-align:center;">executor</td><td style="text-align:center;">最后执行人</td></tr><tr><td style="text-align:center;">executed_at</td><td style="text-align:center;">执行时间</td></tr></tbody></table>

## 结果状态(result\_status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">pass</td><td style="text-align:center;">通过</td></tr><tr><td style="text-align:center;">no_pass</td><td style="text-align:center;">不通过</td></tr><tr><td style="text-align:center;">block</td><td style="text-align:center;">阻塞</td></tr><tr><td style="text-align:center;">unexecuted</td><td style="text-align:center;">未执行</td></tr></tbody></table>

## 用例状态(status)取值字段说明

<table><thead><tr><th style="text-align:center;">取值</th><th style="text-align:center;">字面值</th></tr></thead><tbody><tr><td style="text-align:center;">normal</td><td style="text-align:center;">正常</td></tr><tr><td style="text-align:center;">updating</td><td style="text-align:center;">待更新</td></tr><tr><td style="text-align:center;">abandon</td><td style="text-align:center;">已废弃</td></tr></tbody></table>
