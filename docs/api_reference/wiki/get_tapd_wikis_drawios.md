# 获取wiki drawio数据

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_drawios.html

-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [获取wiki drawio数据](#获取wiki-drawio数据)
    -   [curl 使用 Basic Auth 鉴权调用示例](#curl-使用-basic-auth-鉴权调用示例)
    -   [返回结果](#返回结果)

# 说明

返回符合查询条件的Wiki drawio数据

# url

`https://api.tapd.cn/tapd_wikis_drawios`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

GET

# 请求数限制

返回1条记录

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">id</td><td style="text-align:center;">drawio数据的id，在wiki内容里</td></tr><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">token</td><td style="text-align:center;">否</td><td style="text-align:center;">string</td><td style="text-align:center;">验证用token</td><td style="text-align:center;"><code>如果wiki内容里有token，必须传，如果没有，则不用</code></td></tr></tbody></table>

# 调用示例及返回结果

## 获取wiki drawio数据

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' 'https://api.tapd.cn/tapd_wikis_drawios?workspace_id=10104801&id=1100000000000001102'`

### 返回结果

```
{
  "status": 1,
  "data": {
    "StaticData": {
      "id": "1100000000000001102",
      "values": "<mxGraphModel dx=\"1090\" dy=\"638\" grid=\"1\" gridSize=\"10\" guides=\"1\" tooltips=\"1\" connect=\"1\" arrows=\"1\" fold=\"1\" page=\"1\" pageScale=\"1\" pageWidth=\"827\" pageHeight=\"1169\"><root><mxCell id=\"0\"/><mxCell id=\"1\" parent=\"0\"/><mxCell id=\"2\" value=\"Text\" style=\"text;html=1;resizable=0;points=[];autosize=1;align=left;verticalAlign=top;spacingTop=-4;\" vertex=\"1\" parent=\"1\"><mxGeometry x=\"376\" y=\"238\" width=\"40\" height=\"20\" as=\"geometry\"/></mxCell></root></mxGraphModel>"
    }
  },
  "info": "success"
}
```


# 字段说明

-   values：drawio的xml数据
