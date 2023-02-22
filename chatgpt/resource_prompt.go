package chatgpt

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	chatgpt "github.com/sashabaranov/go-gpt3"
)

func resourcePrompt() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourcePromptRead,
		CreateContext: resourcePromptCreate,
		DeleteContext: resourcePromptDelete,
		Schema: map[string]*schema.Schema{
			"max_tokens": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The max number of characters to show in the result (max is 4000).",
				ValidateFunc: func(val any, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 0 || v > 4000 {
						errs = append(errs, fmt.Errorf("%q must be between 0 and 4000 inclusive, got: %d", key, v))
					}
					return
				},
			},
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The query you would like to send to ChatGPT.",
			},
			"result": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The result of your query as processed by ChatGPT.",
			},
		},
	}
}

func resourcePromptDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourcePromptCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resourcePromptRead(ctx, d, m)
	return nil
}

func resourcePromptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*chatgpt.Client)

	maxTokens := d.Get("max_tokens").(int)
	query := d.Get("query").(string)

	if !d.HasChange("query") && !d.HasChange("max_tokens") {
		return nil
	}

	req := chatgpt.CompletionRequest{
		Model:     AI_MODEL,
		MaxTokens: maxTokens,
		Prompt:    query,
	}
	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		return diag.FromErr(err)
	}

	result := strings.TrimSpace(resp.Choices[0].Text)

	if err := d.Set("result", result); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.ID)

	return diags
}
