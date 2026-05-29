# 锁定迭代

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/lock_iteration.html

# 说明

锁定迭代

# url

`https://api.tapd.cn/iterations/lock_iteration`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次锁定一个迭代

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">iteration_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">迭代ID</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">lock_types</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">锁定对象。取值 __ALL_STORY__ / __ALL_BUG__ 。多个使用英文逗号分隔，比如 lock_types=__ALL_STORY__,__ALL_BUG__</td></tr></tbody></table>

# 调用示例及返回结果

## 在项目下锁定迭代

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&iteration_id=1010104801000723579' 'https://api.tapd.cn/iterations/lock_iteration'`

### 返回结果

```
{
    "status": 1,
    "data": "lock 1010104801000723579 successfully",
    "info": "success"
}
```

