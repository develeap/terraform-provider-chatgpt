terraform {
  required_providers {
    chatgpt = {
      version = "0.0.1"
      source  = "hashicorp.com/develeap/chatgpt"
    }
  }
}

provider "chatgpt" {
  # api_key = "..."
}

resource "chatgpt_prompt" "query" {
  max_tokens = 256
  query      = "Who is the best cloud provider?"
}

output "query_result" {
  value = chatgpt_prompt.query.result
}
