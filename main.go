package main

import (
	"encoding/json"
	"net/http"

	"github.com/gbgomes/GoExpert/labCloudRun/internal/entity"
)

type Resposta struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/clima", Handle)
	http.ListenAndServe(":8080", mux)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	local := entity.NewLocalidadeViaCep()
	localidade, err := local.ColetaLocalidade(cep)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Err.Error()))
		return
	}

	weather := entity.NewWeather()
	tempo, err := weather.ColetaTempo(localidade.Localidade)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Resposta{tempo.Current.TempC, tempo.Current.TempF, tempo.Current.TempC + 273})
}
