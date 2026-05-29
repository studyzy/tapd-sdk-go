# Get Image

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/attachment/get_image.html

- 说明
- SDK 方法名nodeJspythongolang
- url
- 支持格式
- HTTP请求方式
- 请求数限制
- 请求参数
- 额外说明
- 调用示例及返回结果获取图片 /tfl/captures/2023-07/tapd10104801base641689686020146.png 的下载链接
- 字段说明返回字段说明



 
## # 说明

 获取单个图片下载链接

  
## # url

 https://api.tapd.cn/files/get_image

 
## # 支持格式

 JSON/XML（默认JSON格式）

 
## # HTTP请求方式

 GET

 
## # 请求数限制

 - 每次只能请求一张图片的下载链接，下载链接默认有效时间300s
- 文件名后缀仅限 png、gif、jpg、jpeg、bmp

 
## # 请求参数

 <table><thead><tr><th style="text-align:center;">字段名</th> <th style="text-align:center;">必选</th> <th style="text-align:center;">类型及范围</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">integer</td> <td style="text-align:center;">项目ID</td></tr> <tr><td style="text-align:center;">image_path</td> <td style="text-align:center;"><code>是</code></td> <td style="text-align:center;">string</td> <td style="text-align:center;">图片路径, 支持完整url地址, 图片所属项目必须和传入的项目id一致</td></tr></tbody></table>

  
## # 调用示例及返回结果

 
### # 获取图片 /tfl/captures/2023-07/tapd_10104801_base64_1689686020_146.png 的下载链接

 
#### # curl 使用 Basic Auth 鉴权调用示例

 curl -u 'api_user:api_password' 'https://api.tapd.cn/files/get_image?workspace_id=10104801&image_path=/tfl/captures/2023-07/tapd_10104801_base64_1689686020_146.png'

 
# #

 
#### # 返回结果

  ```
{
    "status": 1,
    "data": {
        "Attachment": {
            "type": "tfl_image",
            "value": "/tfl/captures/2023-07/tapd_10104801_base64_1689686020_146.png",
            "workspace_id": 10104801,
            "filename": "tapd_10104801_base64_1689686020_146.png",
            "download_url": "https://file.tapd.cn/attachments/tmp_download/tmp_wiki_attachments_down_c92481f2770c5611d1c7eafe7fb829bb?salt=73aa5cb432e749c68b85821503f4dec3&time=1689686364"

        }
    },
    "info": "success"
}
```

 1234567891011121314 
## # 字段说明

 
### # 返回字段说明

 <table><thead><tr><th style="text-align:center;">字段</th> <th style="text-align:center;">说明</th></tr></thead> <tbody><tr><td style="text-align:center;">workspace_id</td> <td style="text-align:center;">项目id</td></tr> <tr><td style="text-align:center;">filename</td> <td style="text-align:center;">图片文件名</td></tr> <tr><td style="text-align:center;">type</td> <td style="text-align:center;">文件类型</td></tr> <tr><td style="text-align:center;">value</td> <td style="text-align:center;">图片路径</td></tr> <tr><td style="text-align:center;">download_url</td> <td style="text-align:center;">单个图片下载地址</td></tr></tbody></table>
