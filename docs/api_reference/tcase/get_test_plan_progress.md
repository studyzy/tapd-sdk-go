# 获取测试计划执行进度

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_progress.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取测试计划执行进度](#获取测试计划执行进度)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

-   获取测试计划整体执行进度
-   获取指定执行人的测试计划执行进度

# url

`https://api.tapd.cn/test_plans/progress`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

一次获取一条测试计划执行进度

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">测试计划ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 获取测试计划执行进度

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/test_plans/progress?id=1010158231077224799&workspace_id=10158231'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "story_count": 1,
        "tcase_count": 10,
        "status_counter": {
            "pass": "5",
            "no_pass": "0",
            "block": "0",
            "unexecuted": 5
        },
        "executed_rate": "50%"
    },
    "info": "success"
}
```

