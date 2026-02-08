# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Terraform provider that wraps the OpenAI ChatGPT API, allowing queries to be executed as Terraform resources. Built with Terraform Plugin SDK v2 and the `go-gpt3` client library. Single resource: `chatgpt_prompt`.

## Build & Development Commands

```bash
# Build the provider binary
make build

# Install to local Terraform plugins directory (builds first)
# NOTE: OS_ARCH in Makefile defaults to darwin_arm64 — change for your platform
make install

# Run unit tests
make test

# Run acceptance tests (requires real API key)
CHATGPT_API_KEY=sk-xxx TF_ACC=1 make testacc

# Generate provider documentation
go generate ./...

# Run a single test
go test ./chatgpt/ -run TestFunctionName -v
```

## Architecture

Uses Terraform Plugin SDK v2 (not the newer Plugin Framework). The provider is minimal:

- `main.go` — Entry point, serves the provider via `plugin.Serve`
- `chatgpt/provider.go` — Provider schema (single `api_key` field), configures the OpenAI client. Auth via `CHATGPT_API_KEY` env var or HCL attribute. Uses `gpt.GPT3TextDavinci003` model constant.
- `chatgpt/resource_prompt.go` — The `chatgpt_prompt` resource. Create delegates to Read (which calls the completion API). Delete is a no-op. All attributes are `ForceNew`, so any change triggers recreation.

The provider has no Update operation, no data sources, and no import support. The OpenAI client (`*gpt.Client`) is passed through the provider's `meta` interface.

## Key Dependencies

- `github.com/hashicorp/terraform-plugin-sdk/v2` v2.25.0 — Terraform provider framework
- `github.com/sashabaranov/go-gpt3` v1.1.1 — OpenAI API client
- Go 1.20

## Release Process

Tags matching `v*` trigger the GitHub Actions workflow (`.github/workflows/release.yml`), which runs GoReleaser with GPG signing. Requires `GPG_PRIVATE_KEY` and `PASSPHRASE` secrets configured in the repository.

## Local Testing with Terraform

After `make install`, create a `~/.terraformrc` dev override pointing to the local binary, then run `terraform plan`/`apply` in `examples/`.
