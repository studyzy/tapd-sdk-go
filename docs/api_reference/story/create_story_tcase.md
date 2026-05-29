# 创建需求与测试用例关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/create_story_tcase.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [创建需求与测试用例关联关系](#创建需求与测试用例关联关系)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

创建需求与测试用例关联关系

# url

`https://api.tapd.cn/stories/add_story_tcase`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

-   测试用例支持传多ID，一次不超过20个

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">测试用例ID</td><td style="text-align:center;">支持传多ID，使用英文逗号 , 分隔，不超过20个</td></tr></tbody></table>

# 调用示例及返回结果

## 创建需求与测试用例关联关系

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&story_id=1010104801854919725&tcase_id=1010104801077291609' 'https://api.tapd.cn/stories/add_story_tcase'`

### 返回结果

```
{
    "status": 1,
    "data": {
        "success_id": [
            "1010104801077291609"
        ]
    },
    "info": "success"
}
```

