# tapd-sdk-go

[![Go](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)

TAPD（腾讯敏捷产品研发平台）Open API 的 Go SDK，零外部依赖，仅使用标准库实现。

## 安装

```bash
go get github.com/studyzy/tapd-sdk-go
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"
    "log"

    tapd "github.com/studyzy/tapd-sdk-go"
    "github.com/studyzy/tapd-sdk-go/model"
)

func main() {
    // 方式一：使用 Access Token（推荐）
    client := tapd.NewClient("your-access-token", "", "")

    // 方式二：使用 Basic Auth
    // client := tapd.NewClient("", "api_user", "api_password")

    ctx := context.Background()

    // 查询需求列表
    stories, err := client.ListStories(ctx, &model.ListStoriesRequest{
        WorkspaceID: "12345678",
        Status:      "open",
        Limit:       10,
    })
    if err != nil {
        log.Fatal(err)
    }
    for _, s := range stories {
        fmt.Printf("需求: %s [%s]\n", s.Name, s.Status)
    }

    // 创建缺陷
    bug, err := client.CreateBug(ctx, &model.CreateBugRequest{
        WorkspaceID: "12345678",
        Title:       "登录页面崩溃",
        Severity:    "serious",
        Reporter:    "zhangsan",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("创建缺陷: %s (ID: %s)\n", bug.Title, bug.ID)
}
```

## 认证方式

| 方式 | 说明 |
|------|------|
| Bearer Token（推荐） | `tapd.NewClient("your-access-token", "", "")` |
| Basic Auth | `tapd.NewClient("", "api_user", "api_password")` |

当同时提供两种凭据时，Bearer Token 优先。

### 自定义站点地址

连接非 `tapd.cn` 的 TAPD 部署：

```go
client := tapd.NewClientWithBaseURL(
    "https://api.my-tapd.com",   // API 地址（空字符串使用默认值 https://api.tapd.cn）
    "https://www.my-tapd.com",   // 前端页面地址（空字符串使用默认值 https://www.tapd.cn）
    "your-access-token", "", "",
)
```

## 功能覆盖

SDK 覆盖 TAPD 全部 24 个资源模块，提供 **310 个 API 方法**，100% 覆盖官方文档所有公开接口。

| 资源 | 方法数 | 说明 |
|------|--------|------|
| **Story（需求）** | 35 | CRUD、计数、变更历史、分类、链接关系、时间关系、模板、复制、批量更新、视图查询、保密需求、QueryToken、工作项类型转换、并行工作流节点 |
| **Bug（缺陷）** | 19 | CRUD、计数、变更历史、链接缺陷、关联需求、复制、批量更新、模板、视图查询、已删除缺陷、QueryToken、系统字段选项更新 |
| **Task（任务）** | 10 | CRUD、计数、批量更新、变更历史、已删除任务、视图查询、字段信息、自定义字段配置 |
| **Iteration（迭代）** | 15 | CRUD、计数、变更历史、锁定/解锁、仪表盘卡片、模板、迭代类别、计划应用 |
| **MiniItem（轻量工作项）** | 18 | CRUD、计数、分类、变更历史、关联关系、已删除工作项、自定义字段、字段标签 |
| **TCase（测试用例）** | 11 | CRUD、批量创建、分类、关联需求、自定义字段、字段信息 |
| **TestPlan（测试计划）** | 14 | CRUD、计数、进度、详情、关联需求/用例/缺陷、按迭代查询 |
| **TCase Instance（用例实例）** | 5 | 指派、执行、移除、结果查询、解除关联 |
| **TestX Case（测试用例 2.0）** | 18 | 仓库/目录/用例 CRUD、批量创建/更新、搜索、历史/执行/评审/缺陷、模板 |
| **TestX Plan（测试计划 2.0）** | 21 | 目录/计划 CRUD、批量更新/归档、用例管理、缺陷关联、统计、历史、需求关联、模板 |
| **TestX Design（测试设计）** | 3 | 搜索设计、统计、标签 |
| **TestX Report（测试报告）** | 4 | 报告列表/详情/数据、报告模板 |
| **Comment（评论）** | 4 | 列表、添加、更新、计数 |
| **Wiki（文档）** | 12 | CRUD、计数、附件计数、Drawio、访问权限、关注者、标签 |
| **Timesheet（工时）** | 5 | 列表、添加、更新、计数、删除 |
| **Attachment（附件）** | 7 | 上传（普通/Base64）、列表、下载、图片链接、文档下载 |
| **Release（发布计划）** | 12 | 发布计划 CRUD、发布评审 CRUD、评审模板、评审日志、发布附件 |
| **Setting（配置）** | 24 | 模块/版本/基线/特性 CRUD、自定义字段配置、选项更新、级联字段、项目设置 |
| **Workflow（工作流）** | 8 | 状态流转、状态映射、结束状态、起始状态、工作流列表、步骤映射、添加缺陷步骤 |
| **Change（变更历史）** | 7 | 需求/缺陷/任务变更列表及计数、迭代变更 |
| **Workspace（项目）** | 17 | 项目列表/详情/更新、子项目、成员/角色、Mini 项目、工作日历、长短 ID 转换、文档列表、自定义字段 |
| **User（用户）** | 4 | 个人配置、第三方用户映射、视图列表、用户信息 |
| **Label（标签）** | 4 | 创建、更新、查询、计数 |
| **Pipeline（流水线）** | 3 | 创建/查询/删除关联 |
| **Board（看板）** | 4 | 卡片创建/列表/更新、板块列表 |
| **Storage（公共存储）** | 4 | 保存、查询、更新、删除 |
| **Program（项目集）** | 2 | 批量绑定/解绑业务对象 |
| **Source（源码）** | 3 | 保存提交、获取提交、提交对象 |
| **Report（报表）** | 1 | 项目报告 |
| **Measure（度量）** | 1 | 状态流转时间 |
| **Webhook** | 1 | 解析 webhook 事件载荷（JSON/form） |
| **Misc（杂项）** | 6 | 发布计划列表、提交关键字、待办需求/任务/缺陷、企业微信消息 |
| **CustomField** | 6 | 需求/缺陷字段标签及详情、自定义字段配置、工作项类型 |
| **Relation** | 2 | 需求关联缺陷、创建关联 |
| **Category** | 1 | 需求分类列表 |

## 自定义字段

TAPD API 中的自定义字段（`custom_field_*`、`custom_plan_field_*`、`cus_*`）通过 `CustomFields map[string]string` 透传，不会丢失：

```go
// 创建需求时设置自定义字段
story, err := client.CreateStory(ctx, &model.CreateStoryRequest{
    WorkspaceID: "12345678",
    Name:        "新功能",
    CustomFields: map[string]string{
        "custom_field_one": "值1",
        "cus_myfield":      "值2",
    },
})

// 读取响应中的自定义字段
fmt.Println(story.CustomFields["custom_field_one"])
```

支持自定义字段的实体：Story、Bug、Task、TCase、TestPlan、Iteration、LaunchForm、MiniItem。

## 分页

TAPD 单次最多返回 200 条数据，翻页示例：

```go
var all []model.Story
for page := 1; ; page++ {
    batch, err := client.ListStories(ctx, &model.ListStoriesRequest{
        WorkspaceID: "12345678",
        Limit:       200,
        Page:        page,
    })
    if err != nil {
        log.Fatal(err)
    }
    if len(batch) == 0 {
        break
    }
    all = append(all, batch...)
}
```

## 错误处理

SDK 返回结构化的 `*TAPDError`：

| HTTP 状态码 | ExitCode | 含义 |
|-------------|----------|------|
| 401 | 1 | 认证失败 |
| 404 | 2 | 资源未找到 |
| 422 | 3 | 参数错误 |
| 其他非 2xx | 4 | 服务端错误 |

```go
stories, err := client.ListStories(ctx, req)
if err != nil {
    if tapdErr, ok := err.(*tapd.TAPDError); ok {
        fmt.Printf("ExitCode: %d, HTTP: %d, Message: %s\n",
            tapdErr.ExitCode, tapdErr.HTTPStatus, tapdErr.Message)
    }
}
```

## 项目结构

```
tapd-sdk-go/
├── client.go              # Client 核心：认证、HTTP、响应解析
├── story.go               # 需求基础 CRUD
├── story_extras.go        # 需求扩展（分类、链接、批量、复制等 30 个方法）
├── bug.go / bug_extras.go # 缺陷 CRUD 及扩展
├── task.go                # 任务
├── iteration.go           # 迭代
├── mini_item.go           # 轻量工作项
├── tcase.go               # 测试用例
├── tcase_test_plan.go     # 测试计划
├── tcase_instance.go      # 用例实例
├── testx_case.go          # 测试用例 2.0
├── testx_plan.go          # 测试计划 2.0
├── testx_design.go        # 测试设计
├── testx_report.go        # 测试报告
├── comment.go             # 评论
├── wiki.go / wiki_extras.go # Wiki 文档
├── timesheet.go           # 工时
├── attachment.go          # 附件
├── release.go             # 发布计划
├── setting.go             # 配置（模块/版本/基线/特性/自定义字段）
├── workflow.go            # 工作流
├── change.go              # 变更历史
├── workspace.go           # 项目
├── board.go               # 看板
├── user.go                # 用户
├── label.go               # 标签
├── pipeline.go            # 流水线
├── storage.go             # 公共存储
├── source.go              # 源码提交
├── program.go             # 项目集
├── report.go              # 报表
├── measure.go             # 度量
├── webhook.go             # Webhook 解析
├── misc.go                # 杂项（待办、企微消息等）
├── custom_field.go        # 自定义字段查询
├── relation.go            # 实体关联
├── category.go            # 需求分类
├── parse.go               # 响应解析辅助函数
└── model/                 # 请求参数与响应数据结构
```

## 测试

所有 API 方法均有对应的单元测试，使用 `net/http/httptest` 模拟服务端，不依赖真实 TAPD API：

```bash
go test ./...
```

## 许可证

[Apache License 2.0](LICENSE)
