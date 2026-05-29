# # 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/query_label.html

返回符合查询条件的所有自定义标签（分页显示，默认一页30条）

  
# # url

 https://api.tapd.cn/label

 
# # 支持格式

 JSON/XML（默认JSON格式）

 
# # HTTP请求方式

 GET

 
# # 请求数限制

 默认返回 30 条。可通过传 limit 参数设置，最大取 200。也可以传 page 参数翻页

 
# # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th> <th style="text-align:center;">特殊规则</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">integer</td> <td style="text-align:center;">项目ID</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">id</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">id</td> <td style="text-align:center;">支持多ID查询</td></tr> <tr><td style="text-align:center;">name</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">标签名称</td> <td style="text-align:center;">支持模糊匹配</td></tr> <tr><td style="text-align:center;">creator</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">创建人</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">否</td> <td style="text-align:center;">datetime</td> <td style="text-align:center;">创建时间</td> <td style="text-align:center;">支持时间查询</td></tr> <tr><td style="text-align:center;">limit</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">设置返回数量限制，默认为30</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">page</td> <td style="text-align:center;">否</td> <td style="text-align:center;">integer</td> <td style="text-align:center;">返回当前数量限制下第N页的数据，默认为1（第一页）</td> <td style="text-align:center;"></td></tr> <tr><td style="text-align:center;">order</td> <td style="text-align:center;">否</td> <td style="text-align:center;">string</td> <td style="text-align:center;">排序规则，规则：字段名 ASC或者DESC，然后 urlencode</td> <td style="text-align:center;">如按创建时间逆序：order=created%20desc</td></tr></tbody></table>

 
# # 调用示例及返回结果

 
## # 获取项目下的标签信息

 
### # curl 使用 Basic Auth 鉴权调用示例

 curl -u 'api_user:api_password' 'https://api.tapd.cn/label?workspace_id=10158231'

 curl 'https://api.tapd.cn/label?workspace_id=10158231&access_token=ACCESS_TOKEN'

 
### # 返回结果

 ```
{
    "status": 1,
    "data": [
      {
        "LabelPool": {
          "id": "1220358527000001547",
          "workspace_id": "20358527",
          "name": "啦啦",
          "color": "4",
          "creator": "huanjinxie",
          "modifier": "",
          "created": "2022-09-25 21:37:44",
          "modified": "2022-09-25 21:37:44",
          "color_value": "#748096"
        }
      }
    ],
    "info": "success"
}
```

 12345678910111213141516171819
# # 标签字段说明

 
## # 标签重要字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">id</td> <td style="text-align:center;">id</td></tr> <tr><td style="text-align:center;">name</td> <td style="text-align:center;">标签名称</td></tr> <tr><td style="text-align:center;">color</td> <td style="text-align:center;">颜色ID，枚举值，可选[1,2,3,4]</td></tr> <tr><td style="text-align:center;">color_value</td> <td style="text-align:center;">实际颜色值 [1 =&gt; #2FB280, 2 =&gt; #F85E5E,3 =&gt; #FFAF21 ,4  =&gt; #748096]</td></tr> <tr><td style="text-align:center;">created</td> <td style="text-align:center;">创建时间</td></tr> <tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">项目ID</td></tr></tbody></table>
