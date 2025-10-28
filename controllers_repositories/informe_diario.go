package controllersrepositories

import (
	"anbima_backend/database"
	"anbima_backend/models"
	"encoding/json"
	"net/http"
)

func InfDiarioUltimosDiasFundosHandler(w http.ResponseWriter, r *http.Request) {

	cnpj_fundo := r.URL.Query().Get("cnpj_fundo")
	mes := r.URL.Query().Get("mes")
	ano := r.URL.Query().Get("ano")

	var noParams bool = false
	if cnpj_fundo != "" || mes != "" || ano != "" {
		noParams = true
	}

	mapInfo := map[string]string{
		"cnpj_fundo_classe": cnpj_fundo,
		"mes":               mes,
		"ano":               ano,
	}

	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsultaWithParams(db, "sql/inf_diario_ultimos_dias.sql", mapInfo, noParams)
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.InformeDiarioLastDays
	for rows.Next() {
		var fundo models.InformeDiarioLastDays
		err = rows.Scan(
			&fundo.TpFundoClasse,
			&fundo.CNPJFundoClasse,
			&fundo.IdSubclasse,
			&fundo.DtComptc,
			&fundo.VlTotal,
			&fundo.VlQuota,
			&fundo.VlPatrimLiq,
			&fundo.CaptcDia,
			&fundo.ResgDia,
			&fundo.NrCotst,
			&fundo.Dia,
			&fundo.Mes,
			&fundo.Ano,
		)
		if err != nil {
			http.Error(w, "Erro ao fazer scan dos dados: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fundos = append(fundos, fundo)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro durante iteração das linhas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Configura o header para JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica e retorna o JSON
	if err := json.NewEncoder(w).Encode(fundos); err != nil {
		http.Error(w, "Erro ao codificar JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
