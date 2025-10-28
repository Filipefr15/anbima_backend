package main

import (
	controllersrepositories "anbima_backend/controllers_repositories"
	"fmt"
	"net/http"
)

// corsMiddleware adiciona headers CORS para aceitar requisições de qualquer origem
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Define headers CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Chama o próximo handler
		next(w, r)
	}
}

func main() {

	http.HandleFunc("/registrofundo", corsMiddleware(controllersrepositories.RegistroFundoHandler))
	http.HandleFunc("/registrofundo/nomefundo", corsMiddleware(controllersrepositories.RegistroFundoNomeFundoHandler))
	http.HandleFunc("/dashboard/registrofundo/admunico", corsMiddleware(controllersrepositories.RegistroFundoAdmUnicoHandler))
	http.HandleFunc("/dashboard/registrofundo/patrimonioTotal", corsMiddleware(controllersrepositories.RegistroFundoPatrimonioTotalHandler))
	http.HandleFunc("/dashboard/registrofundo/fundosTotais", corsMiddleware(controllersrepositories.RegistroFundoFundosTotaisHandler))
	http.HandleFunc("/dashboard/registrofundo/top10Adm", corsMiddleware(controllersrepositories.RegistroFundoTop10AdmHandler))
	http.HandleFunc("/dashboard/registrofundo/top10Fundos", corsMiddleware(controllersrepositories.RegistroFundoTop10FundosHandler))
	http.HandleFunc("/dashboard/registrofundo/distsit", corsMiddleware(controllersrepositories.RegistroFundoDistSitHandler))
	http.HandleFunc("/dashboard/registrofundo/timeline", corsMiddleware(controllersrepositories.RegistroFundoTimelineAberturaFundosHandler))
	http.HandleFunc("/infdiarioultimosdias", corsMiddleware(controllersrepositories.InfDiarioUltimosDiasFundosHandler))
	http.HandleFunc("/admfii", corsMiddleware(controllersrepositories.AdmFiiHandler)) //unused
	fmt.Println("Servidor RegistroFundo iniciado em :8080")
	fmt.Println("Acesse: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
