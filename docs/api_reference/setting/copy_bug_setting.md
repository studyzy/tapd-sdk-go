# 复制缺陷配置接口

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/copy_bug_setting.html

# 说明

复制缺陷配置到其他项目

使用该接口将复制如下信息到目标项目中:
(1)缺陷字段值设置，如测试方式、操作系统等
(2)缺陷自定义字段设置
(3)缺陷模板设置
(4)默认缺陷工作流
(5)缺陷其他配置项

目标项目中的现有缺陷字段和自定义字段配置将会被覆盖，缺陷模板和缺陷工作流会在原有基础上新增，且默认启用，此操作有可能会影响目标项目的历史缺陷，请谨慎操作！

# url

`https://api.tapd.cn/bugs/copy_settings`

# 支持格式

JSON/XML（默认JSON格式）

# HTTP请求方式

POST

# 请求数限制

一次只能复制到一个项目

# 请求参数

<table><thead><tr><th style="text-align:center;">字段名</th><th style="text-align:center;">必选</th><th style="text-align:center;">类型及范围</th><th style="text-align:center;">说明</th><th style="text-align:center;">特殊规则</th></tr></thead><tbody><tr><td style="text-align:center;">src_workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">源项目ID</td><td style="text-align:center;"></td></tr><tr><td style="text-align:center;">target_workspace_id</td><td style="text-align:center;">是</td><td style="text-align:center;">integer</td><td style="text-align:center;">目标项目ID</td><td style="text-align:center;"></td></tr></tbody></table>

# 调用示例及返回结果

## 复制缺陷配置到目标项目

### curl 使用 Basic Auth 鉴权调用示例

`curl -u 'api_user:api_password' -d 'src_workspace_id=10104801&target_workspace_id=755' 'https://api.tapd.cn/bugs/copy_settings'`

### 返回结果

```json
{
    "status": 1,
    "info": "success"
}
```
