package translate

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const OllamaAPI = "http://localhost:11434/api/predict"

type OllamaRequest struct {
	Model  string `json:"model"`
	Inputs string `json:"inputs"`
}

type OllamaResponse struct {
	Outputs string `json:"outputs"`
}

// TranslateWithOllama 使用Ollama API进行实时翻译
func TranslateWithOllama(text string) (string, error) {
	reqBody := OllamaRequest{
		Model:  "qwen2.5:3b-instruct-q3_K_M", // 根据实际模型名称调整
		Inputs: "翻译以下内容到中文:\n" + text,
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := http.Post(OllamaAPI, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Outputs, nil
}

// 在现有Translate函数中调用新方法（假设存在基础翻译逻辑）
func Translate(text string) (string, error) {
	// 可在此处添加其他翻译逻辑，或直接使用Ollama
	return TranslateWithOllama(text)
}
