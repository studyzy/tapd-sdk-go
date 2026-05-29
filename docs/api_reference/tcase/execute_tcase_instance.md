# 执行测试用例

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/execute_tcase_instance.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取项目下测试用例](#获取项目下测试用例)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [结果状态(result\_status)取值字段说明](#结果状态-result-status-取值字段说明)

# 说明

执行测试用例，支持批量修改

# url

`https://api.tapd.cn/tcase_instance/execute`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

最大支持10条

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">用例ID</td><td style="text-align:center;">支持批量执行</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">result_status</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">执行结果</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">last_executor</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">执行人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">result_remark</td><td style="text-align:center;"><code>否</code></td><td style="text-align:center;">string</td><td style="text-align:center;">实际执行结果</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取项目下测试用例

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'test_plan_id=1010158231077224799&tcase_id=1020357849077231381&result_status=pass&workspace_id=10158231' 'https://api.tapd.cn/tcase_instance/execute'`

### 返回结果

```
{
    "status": 1,
    "data": [],
    "info": "success"
}
```


## 结果状态(result\_status)取值字段说明

|pass|通过| |no\_pass|不通过| |block|阻塞|
