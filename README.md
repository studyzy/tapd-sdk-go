# tapd-sdk-go

[![Go](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)

TAPD（腾讯敏捷产品研发平台）Open API 的 Go SDK，零外部依赖，仅使用标准库实现。

## API 文档来源

本 SDK 的设计基于 TAPD 官方开放平台 API 文档：

> <https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/>

TAPD Open API 共提供约 192 个接口，覆盖 20 个模块。本 SDK 封装了其中常用的核心接口。

## 功能概览

SDK 当前覆盖 20 种 TAPD 资源类型，提供 **110 个 API 方法**：

| 资源 | 方法数 | 支持操作 |
|------|--------|----------|
| **Story（需求）** | 5 | 列表、详情、创建、更新、计数 |
| **Task（任务）** | 5 | 列表、详情、创建、更新、计数 |
| **Bug（缺陷）** | 5 | 列表、详情、创建、更新、计数 |
| **Iteration（迭代）** | 6 | 列表、创建、更新、计数、锁定、解锁 |
| **Comment（评论）** | 4 | 列表、添加、更新、计数 |
| **Wiki（文档）** | 5 | 列表、详情、创建、更新、计数 |
| **TCase（测试用例）** | 4 | 列表、创建、批量创建、计数 |
| **Timesheet（工时）** | 5 | 列表、添加、更新、计数、删除 |
| **Attachment（附件）** | 4 | 图片下载链接、附件列表、单个附件下载、文档下载 |
| **Workflow（工作流）** | 6 | 状态流转、状态映射、结束状态、所有结束状态、起始状态、工作流列表 |
| **Change（变更历史）** | 7 | 需求变更/计数、缺陷变更/计数、任务变更/计数、迭代变更 |
| **Board（看板）** | 4 | 工作项创建、列表、更新、板块列表 |
| **Release（发布计划）** | 3 | 创建、更新、计数 |
| **Setting（配置）** | 16 | 模块/版本/基线/特性的创建、更新、列表、计数 |
| **Report（报表）** | 1 | 项目报告 |
| **Measure（度量）** | 1 | 状态流转时间 |
| **Source（源码）** | 2 | 保存提交数据、获取提交数据 |
| **Category（需求分类）** | 1 | 需求分类列表 |
| **CustomField（自定义字段）** | 6 | 字段配置、需求/缺陷字段标签、需求/缺陷字段详情、需求类别 |
| **Relation（实体关联）** | 2 | 需求关联缺陷、创建关联 |
| **Workspace（项目）** | 7 | 列表、详情、子项目、公司项目列表、项目成员、添加成员、角色列表 |
| **Misc（杂项）** | 6 | 发布计划列表、源码提交关键字、待办需求/任务/缺陷、企业微信消息 |

## 安装

```bash
go get github.com/studyzy/tapd-sdk-go
```

## 快速开始

```go
package main

import (
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

    // 方式三：连接自定义 TAPD 站点
    // client := tapd.NewClientWithBaseURL(
    //     "https://api.my-tapd.com",   // API 地址
    //     "https://www.my-tapd.com",   // 前端页面地址
    //     "your-access-token", "", "",
    // )

    // 获取用户参与的项目列表
    workspaces, err := client.ListWorkspaces()
    if err != nil {
        log.Fatal(err)
    }
    for _, ws := range workspaces {
        fmt.Printf("项目: %s (ID: %s)\n", ws.Name, ws.ID)
    }

    // 查询需求列表
    stories, err := client.ListStories(&model.ListStoriesRequest{
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
}
```

## 架构设计

### 模块结构

```
tapd-sdk-go/
├── go.mod                 # 独立模块，零外部依赖
├── client.go              # Client 核心：认证、HTTP 封装、响应解析
├── story.go               # 需求 API
├── task.go                # 任务 API
├── bug.go                 # 缺陷 API
├── iteration.go           # 迭代 API（含锁定/解锁）
├── comment.go             # 评论 API
├── wiki.go                # Wiki 文档 API
├── tcase.go               # 测试用例 API
├── timesheet.go           # 工时 API（含计数/删除）
├── attachment.go          # 附件/图片 API（含单个下载/文档下载）
├── workflow.go            # 工作流 API（含起始/结束状态、列表）
├── change.go              # 变更历史 API（需求/缺陷/任务/迭代）
├── board.go               # 看板 API
├── release.go             # 发布计划 API
├── setting.go             # 配置 API（模块/版本/基线/特性）
├── report.go              # 报表 API
├── measure.go             # 度量 API
├── source.go              # 源码提交 API
├── category.go            # 需求分类 API
├── custom_field.go        # 自定义字段 API
├── relation.go            # 实体关联 API
├── workspace.go           # 项目 API（含成员/角色/子项目）
├── misc.go                # 杂项：发布计划列表、待办事项、企业微信等
└── model/                 # 数据模型与请求参数
    ├── model.go           # 通用模型（TAPDResponse、Workspace、UserWorkspace 等）
    ├── request.go         # 通用请求结构体（WorkflowRequest、TodoRequest 等）
    ├── story.go           # Story 及其请求结构体
    ├── task.go            # Task 及其请求结构体
    ├── bug.go             # Bug 及其请求结构体
    ├── iteration.go       # Iteration 及其请求结构体
    ├── comment.go         # Comment 及其请求结构体
    ├── wiki.go            # Wiki 及其请求结构体
    ├── tcase.go           # TCase 及其请求结构体
    ├── timesheet.go       # Timesheet 及其请求结构体
    ├── attachment.go      # Attachment/ImageInfo 模型及请求结构体
    ├── board.go           # BoardCard/BoardColumn 及其请求结构体
    ├── change.go          # WorkitemChange/BugChange/IterationChange 及请求结构体
    ├── release.go         # Release 及其请求结构体
    ├── setting.go         # Module/Version/Baseline/Feature 及 IterationLockRequest
    ├── workflow.go        # WorkflowTransition 及相关类型
    ├── category.go        # Category 模型
    └── custom_fields.go   # 自定义字段辅助函数（ExtractCustomFields 等）
```

### 设计原则

1. **零外部依赖**：仅使用 Go 标准库，可被任何项目无冲突引入。
2. **独立模块**：SDK 作为独立 Go module（`github.com/studyzy/tapd-sdk-go`），可被第三方项目直接引用。
3. **统一模式**：所有资源 API 遵循一致的调用模式：
   - 请求参数定义为 `model/` 下的结构体，通过 `ToParams()` 方法转为 `map[string]string`
   - Client 方法调用内部 `doGet`/`doPost` 发起 HTTP 请求
   - 自动解析 TAPD 统一响应格式 `{status, data, info}`
   - 自动处理 TAPD 特有的数据包裹格式（如 `[{"Story": {...}}]`）
4. **自定义字段透传**：`Story`、`Task`、`Bug` 通过自定义 `UnmarshalJSON`/`MarshalJSON` 将 `custom_field_*` 和 `custom_plan_field_*` 字段自动捕获到 `CustomFields map[string]string`，避免数据丢失。
5. **双模式认证**：支持 Bearer Token（推荐）和 Basic Auth 两种认证方式。
6. **可配置站点地址**：API 和前端 URL 均可自定义，支持连接非 `tapd.cn` 的 TAPD 部署。
7. **结构化错误**：返回 `TAPDError` 包含 HTTP 状态码、退出码和错误消息，便于调用方精确处理。

### 认证方式

| 方式 | 环境变量 | 说明 |
|------|----------|------|
| Bearer Token（推荐） | `TAPD_ACCESS_TOKEN` | 通过 TAPD 个人设置获取 |
| Basic Auth | `TAPD_API_USER` + `TAPD_API_PASSWORD` | API 账号密码 |

当同时提供两种凭据时，Bearer Token 优先。

### 自定义站点地址

SDK 默认连接 `https://api.tapd.cn`，如需连接其他 TAPD 部署，使用 `NewClientWithBaseURL`：

```go
client := tapd.NewClientWithBaseURL(
    "https://api.my-tapd.com",   // apiURL：API 请求地址
    "https://www.my-tapd.com",   // webURL：前端页面地址（用于生成条目链接）
    "your-access-token", "", "",
)

// 获取前端页面基础地址
fmt.Println(client.WebURL()) // https://www.my-tapd.com
```

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `apiURL` | `https://api.tapd.cn` | API 请求基础地址，传空字符串使用默认值 |
| `webURL` | `https://www.tapd.cn` | 前端页面基础地址，传空字符串使用默认值 |

### 错误处理

SDK 返回结构化的 `TAPDError`，包含以下退出码映射：

| HTTP 状态码 | 退出码 | 含义 |
|-------------|--------|------|
| 401 | 1 | 认证失败 |
| 404 | 2 | 资源未找到 |
| 422 | 3 | 参数错误 |
| 429/500/502 | 4 | 服务端错误/限流 |
| 其他 | 4 | API 错误 |

```go
stories, err := client.ListStories(req)
if err != nil {
    if tapdErr, ok := err.(*tapd.TAPDError); ok {
        fmt.Printf("退出码: %d, HTTP: %d, 消息: %s\n",
            tapdErr.ExitCode, tapdErr.HTTPStatus, tapdErr.Message)
    }
}
```

## API 覆盖对照

下表对比 TAPD Open API 全量接口与本 SDK 的覆盖情况：

| TAPD 模块 | 官方接口数 | SDK 已覆盖 | 覆盖范围 |
|-----------|-----------|-----------|----------|
| 需求 (Story) | 33 | 5 | 基础 CRUD + 计数 |
| 缺陷 (Bug) | 21 | 5 | 基础 CRUD + 计数 |
| 任务 (Task) | 10 | 5 | 基础 CRUD + 计数 |
| 迭代 (Iteration) | 8 | 6 | 列表、创建、更新、计数、锁定、解锁 |
| 测试 (TCase) | 27 | 4 | 列表、创建、批量创建、计数 |
| Wiki | 11 | 5 | 基础 CRUD + 计数 |
| 评论 (Comment) | 4 | 4 | 全部覆盖 |
| 工时 (Timesheet) | 5 | 5 | 全部覆盖（列表、添加、更新、计数、删除） |
| 附件 (Attachment) | 3 | 4 | 图片链接、附件列表、单个附件下载、文档下载 |
| 项目 (Workspace) | 17 | 7 | 列表、详情、子项目、公司项目、成员、添加成员、角色 |
| 工作流 (Workflow) | 6 | 6 | 全部覆盖 |
| 发布 (Release) | 11 | 4 | 列表、创建、更新、计数 |
| 配置 (Setting) | 21 | 16 | 模块/版本/基线/特性的 CRUD |
| 变更历史 (Change) | - | 7 | 需求/缺陷/任务变更及计数、迭代变更 |
| 看板 (Board) | - | 4 | 工作项创建、列表、更新、板块列表 |
| 报表 (Report) | - | 1 | 项目报告 |
| 度量 (Measure) | - | 1 | 状态流转时间 |
| 源码 (Source) | - | 2 | 保存/获取提交数据 |
| 自定义字段 | - | 6 | 字段配置、字段标签/详情（需求+缺陷）、需求类别 |
| 关联 (Relation) | - | 2 | 需求关联缺陷、创建关联 |
| 分类 (Category) | - | 1 | 需求分类列表 |
| 其他 | - | 5 | 源码提交关键字、待办需求/任务/缺陷、企业微信消息 |

## 测试

SDK 使用 `net/http/httptest` 进行 mock server 测试，每个 API 文件都有对应的测试文件：

```bash
go test ./...
```

## 许可证

[Apache License 2.0](LICENSE)
