package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/userAdityaa/BookEndBackEnd/config"
	"github.com/userAdityaa/BookEndBackEnd/models"
	"github.com/userAdityaa/BookEndBackEnd/utils"
)

func HandleGenerateText(w http.ResponseWriter, r *http.Request) {
	var inputReq models.InputRequest

	if err := json.NewDecoder(r.Body).Decode(&inputReq); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	ibmReq := models.IBMRequest{
		Input: fmt.Sprintf("<|user|>\n%s\n<|assistant|>", inputReq.InputText),
		Parameters: models.IBMRequestParams{
			DecodingMethod:    "greedy",
			MaxNewTokens:      200,
			MinNewTokens:      50,
			StopSequences:     []string{},
			RepetitionPenalty: 1.05,
		},
		ModelID:   "ibm/granite-13b-chat-v2",
		ProjectID: config.ProjectID,
	}

	client := utils.NewHttpClient()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+config.AuthToken).
		SetBody(ibmReq).
		Post(config.APIEndpoint)

	if err != nil {
		http.Error(w, "Failed to communicate with IBM API", http.StatusInternalServerError)
		log.Println("Error communicating with IBM API:", err)
		return
	}

	log.Println("IBM API Response Status:", resp.Status())
	log.Println("IBM API Response Body:", string(resp.Body()))

	var ibmResponse models.IBMResponse
	if err := json.Unmarshal(resp.Body(), &ibmResponse); err != nil {
		http.Error(w, "Failed to parse IBM response", http.StatusInternalServerError)
		log.Println("Error parsing IBM response:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": ibmResponse.Text})
}
