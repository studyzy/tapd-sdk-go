# 批量创建测试用例

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/batch_add_tcase.html

# 说明

批量新建测试用例，返回新建测试用例的数据

# url

`https://api.tapd.cn/tcases/batch_save`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

每次新增最大为两百

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">name</td><td style="text-align:center;">是</td><td style="text-align:center;">string</td><td style="text-align:center;">用例名称</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">steps</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例步骤</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">category_id</td><td style="text-align:center;">否</td><td style="text-align:center;">integer</td><td style="text-align:center;">用例目录</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">status</td><td style="text-align:center;">否</td><td style="text-align:center;">取值 updating, abandon, normal</td><td style="text-align:center;">用例状态</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">precondition</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">前置条件</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">expectation</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">预期结果</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">type</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例类型</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">priority</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">用例等级</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">creator</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">创建人</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">custom_field_*</td><td style="text-align:center;">否</td><td style="text-align:center;">string或者integer</td><td style="text-align:center;">自定义字段参数</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 批量创建测试用例

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d '[{"workspace_id": "69992160", "name": "简单用例1","creator":"XX1"}, {"workspace_id": "69992160", "name": "简单用例2","creator":"XX2"}]' 'https://api.tapd.cn/tcases/batch_save'`

### 返回结果

```json
{
    "status": 1,
    "data": [
        {
            "Tcase": {
                "id": "1069992160077456793",
                "workspace_id": "69992160",
                "category_id": "-1",
                "created": "2023-07-10 16:30:06",
                "modifier": "XX1",
                "modified": "2023-07-10 16:30:06",
                "creator": "XX1",
                "status": "normal",
                "name": "简单用例1",
                "precondition": null,
                "expectation": null,
                "type": "",
                "priority": ""
            }
        },
        {
            "Tcase": {
                "id": "1069992160077456795",
                "workspace_id": "69992160",
                "category_id": "-1",
                "created": "2023-07-10 16:30:06",
                "modifier": "XX2",
                "modified": "2023-07-10 16:30:06",
                "creator": "XX2",
                "status": "normal",
                "name": "简单用例2",
                "precondition": null,
                "expectation": null,
                "type": "",
                "priority": ""
            }
        }
    ],
    "info": "success"
}
```
