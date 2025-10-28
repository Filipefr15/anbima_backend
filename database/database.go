package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectaDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar arquivo .env: %v", err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DATABASE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com banco de dados: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("erro ao conectar com banco de dados: %v", err)
	}

	return db, nil
}

// ExecutarConsulta executa uma consulta SQL genérica a partir de um arquivo
func ExecutarConsulta(db *sql.DB, arquivoSQL string) (*sql.Rows, error) {
	// Lê a consulta SQL do arquivo
	sqlBytes, err := ioutil.ReadFile(arquivoSQL)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo SQL: %v", err)
	}

	query := strings.TrimSpace(string(sqlBytes))

	// Executa a consulta
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar consulta: %v", err)
	}

	return rows, nil
}

func ExecutarConsultaWithOneParam(db *sql.DB, arquivoSQL string, varName, param string) (*sql.Rows, error) {
	// Lê a consulta SQL do arquivo
	sqlBytes, err := ioutil.ReadFile(arquivoSQL)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo SQL: %v", err)
	}

	varName = ":" + varName

	query := strings.TrimSpace(string(sqlBytes))
	query = strings.Replace(query, varName, param, 1)

	// Executa a consulta
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar consulta: %v", err)
	}

	return rows, nil
}

func ExecutarConsultaWithParams(db *sql.DB, arquivoSQL string, params map[string]string, noParams bool) (*sql.Rows, error) {
	// Lê a consulta SQL do arquivo
	sqlBytes, err := ioutil.ReadFile(arquivoSQL)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo SQL: %v", err)
	}

	query := strings.TrimSpace(string(sqlBytes))
	if noParams {
		query += " WHERE "
		for key, value := range params {
			if value != "" {
				query += fmt.Sprintf(" %s = '%s' AND", key, value)
			}
		}
		query = strings.TrimSuffix(query, " AND")
	}

	// Executa a consulta
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar consulta: %v", err)
	}

	return rows, nil
}
