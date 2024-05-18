package main

import (
	"backup-app/backup"
	"fmt"
	"os"
)

func main() {
	// Inicializa o logger
	err := backup.InitLogger()
	if err != nil {
		fmt.Println("Erro ao inicializar o logger:", err)
		os.Exit(1)
	}

	// Definindo o diretório de origem e destino
	sourceDir := "/home/ana.carolyne/Documentos/testes-minhaApp-go/testes-app-go"
	destinationDir := "/home/ana.carolyne/Documentos/testes-minhaApp-go/testes/testes-app2-go"

	// Executa o backup
	err = backup.PerformBackup(sourceDir, destinationDir)
	if err != nil {
		backup.LogError(fmt.Sprintf("Erro ao realizar o backup: %v", err))
		os.Exit(1)
	}
	backup.LogInfo("Backup concluído com sucesso!")
	fmt.Println("Backup concluído com sucesso!")
}
