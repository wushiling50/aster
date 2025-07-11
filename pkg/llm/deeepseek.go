package llm

import (
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/wushiling50/aster/config"
	"github.com/wushiling50/aster/pkg/constants"
)

type NationConfidence struct {
	Nation     string  `json:"nation"`
	Confidence float64 `json:"confidence"`
}

func NewDeepSeekClient(c config.DeepSeekModel) *openai.Client {
	config := openai.DefaultConfig(os.Getenv(constants.DeepSeekAPIToken))
	config.BaseURL = c.Endpoint
	client := openai.NewClientWithConfig(config)
	return client
}

func BuildAnalysisNationReq(c config.DeepSeekModel, allText string) openai.ChatCompletionRequest {
	systemPrompt := `你是一名专业的 GitHub 数据分析师，对不同地区、语言和文化有深刻了解。
	你的任务是分析用户提供的 GitHub 资料内容，推测用户的最可能国籍或地区名。
	请严格按照以下格式回复：
		{"nation": "nationName", "confidence": confidenceValue}
		- nation: 国家/地区名（英文全称）
		- confidence: 置信度（0.0-1.0，保留两位小数）
	不要回复任何其他内容，只输出JSON格式结果！`

	return openai.ChatCompletionRequest{
		Model: c.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "分析以下 GitHub 用户资料并推测其地区:\n\n" + allText,
			},
		},
		MaxTokens:   c.MaxTokens,
		Temperature: c.Temperature,
		TopP:        c.TopP,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	}
}
