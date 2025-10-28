package controllersrepositories

import (
	"anbima_backend/database"
	"anbima_backend/models"
	"encoding/json"
	"net/http"
)

func RegistroFundoHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/registro_fundo.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var fundos []models.RegistroFundo
	for rows.Next() {
		var fundo models.RegistroFundo
		err = rows.Scan(
			&fundo.IdRegistroFundo,
			&fundo.CNPJFundo,
			&fundo.CodigoCVM,
			&fundo.DataRegistro,
			&fundo.DataConstituicao,
			&fundo.TipoFundo,
			&fundo.DenominacaoSocial,
			&fundo.DataCancelamento,
			&fundo.Situacao,
			&fundo.DataInicioSituacao,
			&fundo.DataAdaptacaoRCVM175,
			&fundo.DataInicioExercicio,
			&fundo.DataFimExercicio,
			&fundo.PatrimonioLiquido,
			&fundo.DataPatrimonioLiquido,
			&fundo.Diretor,
			&fundo.CNPJAdministrador,
			&fundo.Administrador,
			&fundo.TipoPessoaGestor,
			&fundo.CPFCNPJGestor,
			&fundo.Gestor,
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

func RegistroFundoNomeFundoHandler(w http.ResponseWriter, r *http.Request) {

	cnpj_fundo := r.URL.Query().Get("cnpj_fundo")
	if cnpj_fundo == "" {
		http.Error(w, "Parâmetro 'cnpj_fundo' é obrigatório", http.StatusBadRequest)
		return
	}

	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	cnpj_fundo = "'" + cnpj_fundo + "'"

	rows, err := database.ExecutarConsultaWithOneParam(db, "sql/registro_fundo_nome_fundo.sql", "cnpj_fundo", cnpj_fundo)
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Processa os resultados
	var fundos []models.RegistroFundoNomeFundo
	for rows.Next() {
		var fundo models.RegistroFundoNomeFundo
		err = rows.Scan(
			&fundo.DenominacaoSocial,
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
