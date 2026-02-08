# Terraform Provider ChatGPT

<img src="https://raw.githubusercontent.com/hashicorp/terraform-website/d841a1e5fca574416b5ca24306f85a0f4f41b36d/content/source/assets/images/logo-terraform-main.svg" width="300px">

A Terraform provider for interacting with OpenAI's ChatGPT API, enabling Infrastructure-as-Code management of AI-powered queries.

## Features

- Execute ChatGPT queries as Terraform resources
- Automatic response tracking and state management
- Support for GPT-3.5-turbo and GPT-4 models
- Configurable token limits and query parameters

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.2.x or higher
- [Go](https://golang.org/doc/install) 1.22.x or higher (to build the provider plugin)
- OpenAI API key (get yours at [platform.openai.com](https://platform.openai.com/account/api-keys))

## Installation

### From Terraform Registry

```hcl
terraform {
  required_providers {
    chatgpt = {
      source  = "develeap/chatgpt"
      version = "~> 0.0.1"
    }
  }
}
```

### Manual Installation

```bash
git clone https://github.com/develeap/terraform-provider-chatgpt
cd terraform-provider-chatgpt
make install
```

## Usage

### Provider Configuration

```hcl
provider "chatgpt" {
  api_key = var.openai_api_key  # or set CHATGPT_API_KEY environment variable
}
```

### Example: Query ChatGPT

```hcl
resource "chatgpt_prompt" "example" {
  query      = "Explain Infrastructure as Code in 3 sentences"
  max_tokens = 256
}

output "chatgpt_response" {
  value = chatgpt_prompt.example.result
}
```

### Example: Technical Documentation Generation

```hcl
resource "chatgpt_prompt" "api_docs" {
  query      = "Generate REST API documentation for a user management endpoint"
  max_tokens = 500
}

resource "local_file" "docs" {
  content  = chatgpt_prompt.api_docs.result
  filename = "${path.module}/docs/api.md"
}
```

## Resources

### `chatgpt_prompt`

Executes a query against the ChatGPT API and stores the response.

#### Arguments

- `query` (Required, ForceNew) - The prompt/question to send to ChatGPT
- `max_tokens` (Required, ForceNew) - Maximum tokens in the response (0-4000)

#### Attributes

- `id` - The completion ID returned by the API
- `result` - The ChatGPT response text

## Development

### Building

```bash
go build -o terraform-provider-chatgpt
```

### Testing

```bash
make test
```

### Acceptance Tests

```bash
CHATGPT_API_KEY=your-api-key TF_ACC=1 make testacc
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Issues

If you encounter any issues, please report them on the [issue tracker](https://github.com/develeap/terraform-provider-chatgpt/issues).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Made with ❤️ in develeap

[<img src="https://media.licdn.com/dms/image/C4D0BAQFXwTP7SFX0QQ/company-logo_200_200/0/1583831070407?e=2147483647&v=beta&t=bWP52NuMxHiQyhMIEe9D7xTNcQMuQDbrTy-ZiVVLCv0" width="50px">](https://www.develeap.com/)
[<img src="https://upload.wikimedia.org/wikipedia/commons/8/81/LinkedIn_icon.svg" width="50px">](https://www.linkedin.com/company/develeap/mycompany/)
