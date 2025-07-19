package llm

import (
	"github.com/sashabaranov/go-openai"
	"github.com/wushiling50/aster/config"
)

type NationConfidence struct {
	Nation     string  `json:"nation"`
	Confidence float64 `json:"confidence"`
}

func NewDeepSeekClient(c config.DeepSeekModel) *openai.Client {
	config := openai.DefaultConfig(config.DeepseekAPIToken)
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

func BuildAnalysisSummaryReq(c config.DeepSeekModel, allText string) openai.ChatCompletionRequest {
	systemPrompt := `你是一名专业的 GitHub 数据分析师，对不同地区、语言和文化有深刻了解。
		你的任务是分析用户提供的 GitHub 用户的简介、贡献内容、使用语言及百分比, 为该用户做一个总结，
		建议包括：擅长的编程领域、文化/地区信息、能力水平、个人风格、个人性格等，不要求包含全部方面，可以自由发挥，总结更多方面。
		请严格按照以下标准回复：
			以中文回复,请直接回复总结内容，不要换行，不要分段,
			纯文本即可，不需要包含其他信息。字数控制在 600 字以内。`

	return openai.ChatCompletionRequest{
		Model: c.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "分析以下 GitHub 用户资料并进行总结:\n\n" + allText,
			},
		},
		MaxTokens:   c.MaxTokens,
		Temperature: c.Temperature,
		TopP:        c.TopP,
	}
}
