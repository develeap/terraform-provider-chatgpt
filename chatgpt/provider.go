package chatgpt

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gpt "github.com/sashabaranov/go-gpt3"
)

const AI_MODEL = gpt.GPT3TextDavinci003

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHATGPT_API_KEY", nil),
				Description: "ChatGPT API Key from https://platform.openai.com/account/api-keys",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"chatgpt_prompt": resourcePrompt(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)
	var diags diag.Diagnostics

	if apiKey != "" {
		c := gpt.NewClient(apiKey)
		// if err != nil {
		// 	return nil, diag.FromErr(err)
		// }

		return c, diags
	}

	return nil, diags
}
