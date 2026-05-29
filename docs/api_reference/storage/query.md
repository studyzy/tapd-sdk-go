# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/storage/query.html

查询数据

 

 
# # url

 https://api.tapd.cn/open_app_storage/query

 
# # 支持格式

 JSON

 
# # HTTP请求方式

 POST

 
# # 请求数限制

 无

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">document</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">文档名</td> <td style="text-align:center;">无</td></tr> <tr><td style="text-align:center;">condition</td> <td style="text-align:center;"><code>否</code></td> <td style="text-align:center;">object</td> <td style="text-align:center;">查询条件</td> <td style="text-align:center;">参考<a href="/document/api-doc/API文档/api_reference/storage/condition.html" class="">条件语法</a></td></tr> <tr><td style="text-align:center;">order</td> <td style="text-align:center;"><code>否</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode，多字段排序，使用,分隔</td> <td style="text-align:center;">无</td></tr> <tr><td style="text-align:center;">limit</td> <td style="text-align:center;"><code>否</code></td> <td style="text-align:center;">int</td> <td style="text-align:center;">条数限制</td> <td style="text-align:center;">无</td></tr> <tr><td style="text-align:center;">offset</td> <td style="text-align:center;"><code>否</code></td> <td style="text-align:center;">int</td> <td style="text-align:center;">偏移量</td> <td style="text-align:center;">无</td></tr> <tr><td style="text-align:center;">fields</td> <td style="text-align:center;"><code>否</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">返回字段</td> <td style="text-align:center;">无</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
### # 返回结果

 ```
{
  "status": 1,
  "data": [
    {
      "id": "1",
      "context_data": "{\\"field1\\": \\"value1\\"}",
      "field1": "value1"
    },
    {
      "id": "2",
      "context_data": "{\\"field1\\": \\"value1\\"}",
      "field1": "value1"
    }
  ],
  "info": "success",
  "total": "2"
}
```

 1234567891011121314151617
