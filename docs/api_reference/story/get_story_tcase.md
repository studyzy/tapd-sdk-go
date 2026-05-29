# 获取需求与测试用例关联关系

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_tcase.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)
-   [测试用例重要字段说明](#测试用例重要字段说明)

# 说明

返回符合查询条件的所有测试用例&测试计划

# url

`https://api.tapd.cn/stories/get_story_tcase`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

返回全部数据

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">需求id</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">include_test_plan</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">是否包含测试计划</td><td style="text-align:center;">取值为1或0，默认为1</td></tr></tbody></table>

# 调用示例及返回结果

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/stories/get_story_tcase?workspace_id=10104801&story_id=1010104801866191641'`

### 返回结果

```

{
    "status": 1,
    "data": [
        {
            "TestPlanStoryTcaseRelation": {
                "id": "1010104801021215005",
                "workspace_id": "10104801",
                "test_plan_id": "0",
                "story_id": "1010104801866191641",
                "tcase_id": "1010104801076110789",
                "sort": "0",
                "creator": "v_xuanfang",
                "created": "2021-08-06 12:35:01"
            }
        }
    ],
    "info": "success"
}

```


# 测试用例字段说明

## 测试用例重要字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;">关系id</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">test_plan_id</td><td style="text-align:center;">测试计划ID</td></tr><tr><td style="text-align:center;">story_id</td><td style="text-align:center;">需求ID</td></tr><tr><td style="text-align:center;">tcase_id</td><td style="text-align:center;">测试用例ID</td></tr><tr><td style="text-align:center;">created</td><td style="text-align:center;">关系创建时间</td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">创建人</td></tr><tr><td style="text-align:center;">sort</td><td style="text-align:center;">显示时排序系数</td></tr></tbody></table>
