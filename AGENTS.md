# AGENTS.md

This file provides guidance to AI coding agents when working with code in this repository.

## Project

Go SDK for TAPD (腾讯敏捷产品研发平台) Open API. Standalone Go module, **standard library only — no external dependencies**. Go 1.24+. Apache 2.0.

Package layout: top-level `tapd` package (HTTP client + per-resource methods) and `model/` package (request/response structs).

## API 文档（唯一事实来源）

本项目有两套 API 参考文档，它们是所有 Go SDK 代码的**唯一事实来源**：

- `docs/api_reference/` — 完整版 API 文档
