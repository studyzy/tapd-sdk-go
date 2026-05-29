# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/storage/save.html

保存数据

 

 
# # url

 https://api.tapd.cn/open_app_storage/save

 
# # 支持格式

 JSON

 
# # HTTP请求方式

 POST

 
# # 请求数限制

 无

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">document</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">文档名，若不存在将被创建</td> <td style="text-align:center;">无</td></tr> <tr><td style="text-align:center;">data</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">array</td> <td style="text-align:center;">要保存的数据，支持单条和批量，格式为{a:1,b:2}或[{a:1,b:2},{a:2,b:3}]</td> <td style="text-align:center;">无</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
### # 返回结果

 ```
{
  "status": 1,
  "data": [],
  "info": "success"
}
```

 12345
