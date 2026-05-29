# 更新下拉类型自定义字段候选值

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/update_select_field_options.html

# 说明

更新下拉类型自定义字段候选值

# url

`https://api.tapd.cn/custom_field_configs/update_select_field_options`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次只允许更新一个字段，并且是全量更新这个字段的所有候选值。支持对象：需求、缺陷、任务、迭代、测试用例

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">19位长度的自定义字段配置ID，可以调用对应获取自定义对象接口取到</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">options</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">候选值。以英文竖线隔开</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 更新需求自定义配置ID为 1010158231214910583 的候选值为 开发，测试，产品，运营

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'options=开发|测试|产品|运营&id=1010158231214910583&workspace_id=10158231' 'https://api.tapd.cn/custom_field_configs/update_select_field_options'`

### 返回结果

```json
{
    "status": 1,
    "data": {
        "status": 1
    },
    "info": "success"
}
```
