# AGENTS.md

本文件为 AI 编码代理在本仓库中工作时提供指导。

## 项目概述

TAPD（腾讯敏捷产品研发平台）Open API 的 Go SDK。独立 Go 模块，**仅使用标准库，无外部依赖**。Go 1.24+，Apache 2.0 协议。

包结构：顶层 `tapd` 包（HTTP 客户端 + 按资源划分的方法）和 `model/` 包（请求/响应结构体）。

## API 文档（唯一事实来源）

本项目有两套 API 参考文档，它们是所有 Go SDK 代码的**唯一事实来源**：

- `docs/api_reference/` — 完整版 API 文档
- `docs/mini_api_reference/` — 精简版 API 文档

**核心原则：**

1. **不可杜撰**：所有 SDK 方法的参数、返回值、字段名称和类型必须严格来源于以上文档，不得凭空编造任何 API 接口或字段。
2. **不可缺少**：文档中定义的所有接口和字段都必须在 SDK 中完整实现，不得遗漏。
3. **数据结构以文档为准**：对应的对象（Request/Response/Entity）数据结构必须满足文档的定义。如果现有代码与文档存在不一致，**以文档为准**进行修正。
4. **文档优先级**：当两套文档之间存在冲突时，以 `docs/api_reference/`（完整版）为准。

在新增或修改任何 API 方法时，必须先查阅对应的 API 文档，确认接口路径、请求参数、响应结构等信息后再编写代码。

## 常用命令

```bash
# 运行所有测试（含竞态检测 + 覆盖率，与 CI 一致）
go test -race -coverprofile=coverage.out ./...

# 运行单个包的测试
go test ./model/...

# 按名称运行单个测试（正则）
go test -run TestNewClient_BearerAuth ./...
go test -run '^TestListStories$' .

# 静态分析（CI 门禁）
go vet ./...

# 格式检查（CI 会拒绝未格式化的文件）
gofmt -l .
gofmt -w .   # 应用格式化
```

CI（`.github/workflows/ci.yml`）依次运行 gofmt 检查、`go vet ./...`、`go test -race -coverprofile=coverage.out ./...`，使用 Go 1.24.12。推送前确保本地全部通过。

## 架构

### 请求流程

每个资源方法都遵循相同的管线流程——理解一次即可通用：

1. 调用者构建 `*model.XxxRequest` 结构体，调用 `client.XxxAction(ctx, req)`。
2. 方法调用 `req.ToParams()` 获取 `map[string]string`。
3. `Client.doGet` / `Client.doPost`（位于 `client.go`）构建 HTTP 请求，附加 `Authorization` 头（`Bearer <token>` 或 `Basic <base64>`），发送请求。
4. `doRequest` 读取响应体，将非 2xx HTTP 状态码映射为 `*TAPDError`（含 `ExitCode`，参见 `mapHTTPError`），然后反序列化为 `model.TAPDResponse{Status, Data, Info}`。若 `Status != 1`，返回 `ExitCode: 4` 的 `*TAPDError`。
5. 资源方法调用 `parseList[T]` / `parseOne[T]` / `parseCount`（`parse.go` 中的泛型辅助函数）来解包 TAPD 特有的信封格式 `[{"Story": {...}}, ...]` 或 `{"Story": {...}}`。
6. 对于单实体响应，方法通常会使用 `c.webURL` 设置 `entity.URL`，使模型携带可点击链接。

这意味着：新增资源时几乎不需要手写 JSON 解析——`parseList[model.Foo](data, "Foo")` 就是标准模式。

### 新增 API 方法

1. **查阅文档**：先在 `docs/api_reference/` 或 `docs/mini_api_reference/` 中找到对应的 API 定义，确认接口路径、请求参数、响应字段。
2. 在 `model/<resource>.go` 中定义请求结构体，实现 `ToParams() map[string]string` 方法。使用 `setOptional` / `setOptionalInt`（定义于 `model/wiki.go`）跳过空值。**所有字段必须与文档定义一致。**
3. 如果响应包含新实体，在同一文件中添加模型结构体。**Entity 的字段必须与文档中的响应结构完全对应。**
4. 在顶层包的 `<resource>.go` 中为 `*Client` 添加方法。读取用 `doGet`，写入用 `doPost`（form-urlencoded）。JSON POST（如企业微信 webhook）用 `doPostJSON`。
5. 使用 `parseList[T]` / `parseOne[T]` / `parseCount` 解码 `data`。
6. 在 `<resource>_test.go` 中添加测试，使用 `newMockServer`（参见 `client_test.go`）——处理器返回 `{"status":1,"data":...,"info":"success"}`。

### 文件拆分（extras 模式）

大型资源拆分为多个文件以保持可控：

- `<resource>.go` — 基础 CRUD（Create/Get/List/Update/Count）。
- `<resource>_extras.go` — 次要端点（分类、链接、关联、批量操作、复制、模板、视图配置查询等）。
- `tcase_test_plan.go` / `tcase_instance.go` — tcase 的领域特定子模块。

`model/` 中遵循同样的约定（如 `model/story_extras.go`、`model/bug_extras.go`、`model/wiki_extras.go`）。当为某个资源添加大量端点时，优先创建新的 `_extras.go` 文件，而非膨胀基础文件。**绝不**将新类型放入 `model/request.go` 或 `model/model.go`——这些是共享文件，已经很拥挤。

### 自定义字段（重要特性）

TAPD API 中存在三类动态字段前缀，SDK 必须全部支持：

| 前缀 | 说明 | 出现位置 |
|------|------|---------|
| `custom_field_*` | 标准自定义字段（如 `custom_field_one`、`custom_field_9`） | 请求参数（过滤）+ 响应数据 |
| `custom_plan_field_*` | 计划自定义字段（如 `custom_plan_field_1`） | 请求参数（过滤）+ 响应数据 |
| `cus_` | 自定义字段别名（如 `cus_myfield`） | 仅请求参数（创建/更新时设值） |

辅助函数位于 `model/custom_fields.go`：

- `IsCustomField(key)` — 判断 key 是否为以上三种前缀之一。
- `ExtractCustomFields(raw)` — 反序列化时从原始 JSON map 中提取所有自定义字段键值对。
- `MergeCustomFields(params, custom)` — 将自定义字段 map 注入请求参数 map。

#### 响应对象处理模式

携带自定义字段的实体结构体必须：

1. 添加 `CustomFields map[string]string` 字段，标记 `json:"-"`
2. 实现 `UnmarshalJSON`：先用 type alias 解析已知字段，再解析原始 JSON 为 `map[string]json.RawMessage`，通过 `ExtractCustomFields` 提取动态字段
3. 实现 `MarshalJSON`：先序列化已知字段，再将 `CustomFields` 合并进 JSON 对象

标准示例参见 `model/story.go` 中的 `Story` 结构体。

#### 请求对象处理模式

涉及自定义字段的请求结构体（Create/Update/List/Count）必须：

1. 添加 `CustomFields map[string]string` 字段
2. 在 `ToParams()` 方法末尾调用 `MergeCustomFields(params, r.CustomFields)`

调用者可通过此 map 传入 `custom_field_*`、`custom_plan_field_*`、`cus_*` 等任意动态字段。

#### 需要支持自定义字段的实体

以下实体的 API 文档中明确包含自定义字段，SDK 必须完整支持：

- **Story** — `custom_field_*`（最多 200 个）+ `custom_plan_field_*`（最多 10 个）
- **Bug** — `custom_field_*`（最多 100 个）+ `custom_plan_field_*`
- **Task** — `custom_field_*`（最多 50 个）
- **TCase** — `custom_field_*`（最多 50 个）
- **TestPlan** — `custom_field_*`（最多 50 个）
- **Iteration** — `custom_field_*`（最多 50 个）
- **LaunchForm** — `custom_field_*`（最多 40 个）
- **MiniItem** — `custom_field_*`（最多 100 个）

### 认证

`NewClient(accessToken, apiUser, apiPassword)` — accessToken 非空时使用 Bearer 认证；否则使用 user/pass 进行 Basic 认证。`NewClientWithBaseURL` 允许通过覆盖 `apiURL` 和 `webURL` 指向私有 TAPD 部署（传 `""` 则回退到 `https://api.tapd.cn` / `https://www.tapd.cn`）。

`Client` 是协程安全的。`nick` 字段（当前用户，通过 `FetchNick` 获取）由 `sync.RWMutex` 保护——使用 `GetNick`/`SetNick`，不要直接访问 `c.nick`。

### 错误处理

返回给调用者的错误要么是 `*TAPDError`（HTTP 非 2xx 或 TAPD 响应体中 `status != 1`），要么是包装后的标准库错误。`TAPDError.ExitCode` 映射：401→1、404→2、422→3、其他/TAPD 逻辑错误→4。调用者可通过 `errors.As` 或 `err.(*tapd.TAPDError)` 进行类型断言。

## 编码约定

- 所有公开 API 的第一个参数为 `context.Context`。
- 导出标识符的文档注释使用中文——与现有风格保持一致。每个资源方法在注释中链接到对应的 TAPD API 文档 URL。
- 请求结构体的 ID 字段使用 `string` 类型（TAPD 返回的 ID 是字符串），通过 `setOptional` 实现显式的空值省略。
- 测试与源文件并置（`foo.go` ↔ `foo_test.go`），使用 `httptest.NewServer` 模拟服务器——测试中绝不调用真实的 TAPD API。
- 零依赖规则是底线：不要在 `go.mod` 中添加 `require` 条目。任何标准库之外的依赖都违反项目设计契约。
