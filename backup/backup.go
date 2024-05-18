package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func PerformBackup(sourceDir, destinationDir string) error {
	// Caminhe pela árvore de diretórios da fonte
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			LogError(fmt.Sprintf("Erro ao caminhar pelo diretório: %v", err))
			return err
		}

		// Defina o caminho de destino correspondente
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			LogError(fmt.Sprintf("Erro ao obter o caminho relativo: %v", err))
			return err
		}
		destPath := filepath.Join(destinationDir, relPath)

		// Verifica se é um diretório ou arquivo
		if info.IsDir() {
			// Cria o diretório no destino
			if err := os.MkdirAll(destPath, info.Mode()); err != nil {
				LogError(fmt.Sprintf("Erro ao criar diretório: %v", err))
				return err
			}
		} else {
			// Copia o arquivo
			if err := copyFile(path, destPath); err != nil {
				LogError(fmt.Sprintf("Erro ao copiar arquivo: %v", err))
				return err
			}
		}
		return nil
	})
}

func copyFile(sourceFile, destFile string) error {
	// Abre o arquivo de origem
	src, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer src.Close()

	// Cria o arquivo de destino
	dst, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copia os dados de origem para o destino
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	// Copia as permissões do arquivo
	info, err := os.Stat(sourceFile)
	if err != nil {
		return err
	}
	return os.Chmod(destFile, info.Mode())
}
