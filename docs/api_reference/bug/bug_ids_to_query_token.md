# 转换缺陷ID成列表queryToken

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/bug_ids_to_query_token.html

-   [说明](#说明)
-   [SDK 方法名](#sdk-方法名)
    -   [nodeJs](#nodejs)
    -   [python](#python)
    -   [golang](#golang)
-   [url](#url)
-   [支持格式](#支持格式)
-   [HTTP 请求方式](#http-请求方式)
-   [请求数限制](#请求数限制)
-   [请求参数](#请求参数)
-   [调用示例及返回结果](#调用示例及返回结果)
    -   [把缺陷ID 1010104801102653321,1010104801085527301 转换成 queryToken 及返回列表链接](#把缺陷id-1010104801102653321-1010104801085527301-转换成-querytoken-及返回列表链接)
-   [重要字段说明](#重要字段说明)
    -   [返回字段字段说明](#返回字段字段说明)

## 说明

把一批缺陷ID转换成页面能用的 QueryToken

## url

`https://api.tapd.cn/bugs/ids_to_query_token`

## 支持格式

JSON/XML（默认JSON格式）

## HTTP 请求方式

POST

## 请求数限制

为了保证页面显示效果，建议ID数量不超过500

## 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">workspace_id</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">integer</td><td style="text-align:center;">项目ID</td></tr><tr><td style="text-align:center;">ids</td><td style="text-align:center;"><code>是</code></td><td style="text-align:center;">string</td><td style="text-align:center;">缺陷ID，使用英文逗号 , 做分隔</td></tr></tbody></table>

## 调用示例及返回结果

### 把缺陷ID 1010104801102653321,1010104801085527301 转换成 queryToken 及返回列表链接

#### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'workspace_id=10104801&ids=1010104801102653321,1010104801085527301' 'https://api.tapd.cn/bugs/ids_to_query_token'`

#### 返回结果

```
{
    "status": 1,
    "data": {
        "queryToken": "71ab88eeb45d084d8fbc85686a0d2399",
        "href": "http://www.tapd.cn/tapd_fe/10104801/bug/list?page=1&queryToken=71ab88eeb45d084d8fbc85686a0d2399"
    },
    "info": "success"
}
```


## 重要字段说明

### 返回字段字段说明

<table><thead><tr><th style="text-align:center;">字段</th><th style="text-align:center;">说明</th></tr></thead><tbody><tr><td style="text-align:center;">queryToken</td><td style="text-align:center;">列表queryToken</td></tr><tr><td style="text-align:center;">href</td><td style="text-align:center;">对应的TAPD缺陷列表链接</td></tr></tbody></table>
