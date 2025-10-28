package controllersrepositories

import (
	"anbima_backend/database"
	"anbima_backend/models"
	"encoding/json"
	"net/http"
)

func RegistroFundoAdmUnicoHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/registro_fundo_adm_unico.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.RegistroFundoAdmUnico
	for rows.Next() {
		var fundo models.RegistroFundoAdmUnico
		err = rows.Scan(
			&fundo.AdmUnico,
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

func RegistroFundoPatrimonioTotalHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/registro_fundo_patr_total.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.RegistroFundoPatrTotal
	for rows.Next() {
		var fundo models.RegistroFundoPatrTotal
		err = rows.Scan(
			&fundo.PatrTotal,
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

func RegistroFundoFundosTotaisHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	situacao := r.URL.Query().Get("situacao")

	if situacao != "" {
		// Executa a consulta
		rows, err := database.ExecutarConsultaWithOneParam(db, "sql/registro_fundo_fundos_totais_sit.sql", "sit", situacao)
		if err != nil {
			http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var fundos []models.RegistroFundoFundosTotais
		for rows.Next() {
			var fundo models.RegistroFundoFundosTotais
			err = rows.Scan(
				&fundo.TotalFundos,
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
	} else {
		// Executa a consulta
		rows, err := database.ExecutarConsulta(db, "sql/registro_fundo_fundos_totais.sql")
		if err != nil {
			http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		// Processa os resultados
		var fundos []models.RegistroFundoFundosTotais
		for rows.Next() {
			var fundo models.RegistroFundoFundosTotais
			err = rows.Scan(
				&fundo.TotalFundos,
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
}

func RegistroFundoTop10AdmHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/registro_fundo_top_10_adm.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.RegistroFundoTop10Adm
	for rows.Next() {
		var fundo models.RegistroFundoTop10Adm
		err = rows.Scan(
			&fundo.Adm,
			&fundo.QtdeFundos,
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

func RegistroFundoTop10FundosHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/registro_fundo_top_10_fundos.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.RegistroFundoTop10Fundos
	for rows.Next() {
		var fundo models.RegistroFundoTop10Fundos
		err = rows.Scan(
			&fundo.DenomSocial,
			&fundo.PatrimonioLiquido,
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

func RegistroFundoDistSitHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/registro_fundo_dist_situacao.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.RegistroFundoDistSit
	for rows.Next() {
		var fundo models.RegistroFundoDistSit
		err = rows.Scan(
			&fundo.Situacao,
			&fundo.Quantidade,
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

func RegistroFundoTimelineAberturaFundosHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/registro_fundo_abertura_fundos.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.RegistroFundoAberturaFundos
	for rows.Next() {
		var fundo models.RegistroFundoAberturaFundos
		err = rows.Scan(
			&fundo.Periodo,
			&fundo.Quantidade,
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
