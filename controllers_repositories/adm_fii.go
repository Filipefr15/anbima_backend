package controllersrepositories

import (
	"anbima_backend/database"
	"anbima_backend/models"
	"encoding/json"
	"net/http"
)

func AdmFiiHandler(w http.ResponseWriter, r *http.Request) {
	// Conecta ao banco
	db, err := database.ConectaDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Executa a consulta
	rows, err := database.ExecutarConsulta(db, "sql/adm_fii.sql")
	if err != nil {
		http.Error(w, "Erro ao executar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Processa os resultados
	var admFiis []models.AdmFii
	for rows.Next() {
		var admFii models.AdmFii
		err = rows.Scan(
			&admFii.TpFundoClasse,
			&admFii.CnpjFundoClasse,
			&admFii.DenomSocial,
			&admFii.DtReg,
			&admFii.DtConst,
			&admFii.CdCvm,
			&admFii.DtCancel,
			&admFii.Sit,
			&admFii.DtIniSit,
			&admFii.DtIniAtiv,
			&admFii.DtIniExerc,
			&admFii.DtFimExerc,
			&admFii.Classe,
			&admFii.DtIniClasse,
			&admFii.RentabFundo,
			&admFii.Condom,
			&admFii.FundoCotas,
			&admFii.FundoExclusivo,
			&admFii.TribLprazo,
			&admFii.PublicoAlvo,
			&admFii.EntidInvest,
			&admFii.TaxaPerfm,
			&admFii.InfTaxaPerfm,
			&admFii.TaxaAdm,
			&admFii.InfTaxaAdm,
			&admFii.VlPatrimLiq,
			&admFii.DtPatrimLiq,
			&admFii.Diretor,
			&admFii.CnpjAdmin,
			&admFii.Admin,
			&admFii.PfPjGestor,
			&admFii.CpfCnpjGestor,
			&admFii.Gestor,
			&admFii.CnpjAuditor,
			&admFii.Auditor,
			&admFii.CnpjCustodiante,
			&admFii.Custodiante,
			&admFii.CnpjControlador,
			&admFii.Controlador,
			&admFii.InvestCemprExter,
			&admFii.ClasseAnbima,
		)
		if err != nil {
			http.Error(w, "Erro ao fazer scan dos dados: "+err.Error(), http.StatusInternalServerError)
			return
		}

		admFiis = append(admFiis, admFii)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro durante iteração das linhas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Configura o header para JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica e retorna o JSON
	if err := json.NewEncoder(w).Encode(admFiis); err != nil {
		http.Error(w, "Erro ao codificar JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
