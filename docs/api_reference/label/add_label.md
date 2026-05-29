# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/add_label.html

新建标签，返回新建标签的数据

  
# # url

 https://api.tapd.cn/label

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 POST

 
# # 请求数限制

 一次插入一条数据

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">是</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">项目ID</td></tr> <tr><td style="text-align:center;">name</td> <td style="text-align:center;">是</td> <td style="text-align:center;">string</td> <td style="text-align:center;">标签名称,不能包括英文坚线</td></tr> <tr><td style="text-align:center;">color</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">颜色标识，枚举值，可选[1,2,3,4]</td></tr> <tr><td style="text-align:center;">creator</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">创建人</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
### # curl 使用 Basic Auth 鉴权调用示例

 curl -u 'api_user:api_password' -d 'workspace_id=10104801&name=创建标签&creator=tapd_api' 'https://api.tapd.cn/label'

 
## # 返回结果

 ```
{
  "status": 1,
  "data": {
    "LabelPool": {
      "id": "1220358527000001577",
      "workspace_id": "20358527",
      "name": "创建标签",
      "color": "2",
      "creator": "",
      "modifier": "",
      "created": "2022-09-26 20:25:02",
      "modified": "2022-09-26 20:25:02"
    }
  },
  "info": "success"
}
```

 12345678910111213141516
# # 标签字段说明

 
## # 标签字段重要字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">ID</td></tr> <tr><td style="text-align:center;">name</td> <td style="text-align:center;">标签名称</td></tr> <tr><td style="text-align:center;">color</td> <td style="text-align:center;">详细描述</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr></tbody></table>
