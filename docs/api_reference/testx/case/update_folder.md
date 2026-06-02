# 说明

> Source: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/testx/case/update_folder.html

更新用例目录，返回用例目录

# url

`https://api.tapd.cn/api/testx/case/v1/namespaces/{namespace}/repos/{repo_uid}/versions/{version_uid}/folders/{folder_uid}`

# 支持格式

JSON

# 调用方式

参考 [授权-用户态](/document/api-doc/API文档/授权凭证/用户态.html)

# HTTP请求方式

PUT

# 请求数限制

一次更新一条数据

# 请求参数

## Request


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Namespace | 是 | string | 项目命名空间 |
| Folder | 是 | Folder | 目录信息 |

## Folder


| 字段名 | 必选 | 类型及范围 | 说明 |
| --- | --- | --- | --- |
| Uid | 是 | string | 唯一标识 |
| RepoUid | 是 | string | 仓库唯一标识 |
| RepoVersionUid | 是 | string | 仓库版本唯一标识 |
| FolderUid | 是 | string | 目录唯一标识 |
| Name | 是 | string | 目录名称 |
| Owners | 否 | repeated string | 负责人列表 |
| Description | 否 | string | 描述信息 |

# 返回结果


```json
{
"Error": null,
"Data": {
"Uid": "12572731",
"RepoUid": "",
"RepoVersionUid": "",
"FolderUid": "",
"FullPath": "",
"Name": "test",
"Owners": [
"xxxx"
],
"Description": "test",
"CaseCount": 0,
"UUID": "40340dec-cea7-4fe6-9a96-ee07eaef2bde",
"Path": ".",
"Audit": {
"Creator": "xx",
"Updater": "xx",
"CreatedAt": "2025-07-12T10:26:55+08:00",
"UpdatedAt": "2025-07-12T10:26:55+08:00",
"Tenant": "xx"
},
}
}
```

# 测试用例目录字段说明

## 测试用例重要字段说明


| 字段 | 说明 |
| --- | --- |
| Uid | 目录uid |
| UUID | 目录uuid |
| RepoUid | 仓库uid |
| RepoVersionUid | 仓库版本uid |
| Description | 仓库描述 |
| CaseCount | 目录用例数 |
| FullPath | 目录中文路径 |
| Path | 目录路径 |
| Owners | 目录负责人列表 |