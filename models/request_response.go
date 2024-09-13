package models

type InputRequest struct {
	InputText string `json:"input_text"`
}

type IBMRequest struct {
	Input      string           `json:"input"`
	Parameters IBMRequestParams `json:"parameters"`
	ModelID    string           `json:"model_id"`
	ProjectID  string           `json:"project_id"`
}

type IBMRequestParams struct {
	DecodingMethod    string   `json:"decoding_method"`
	MaxNewTokens      int      `json:"max_new_tokens"`
	MinNewTokens      int      `json:"min_new_tokens"`
	StopSequences     []string `json:"stop_sequences"`
	RepetitionPenalty float64  `json:"repetition_penalty"`
}

type IBMResponse struct {
	Text string `json:"generated_text"`
}
