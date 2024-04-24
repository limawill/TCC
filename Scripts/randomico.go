package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
    "math/rand"
    "time"
)

func main() {
    // Definir a pasta onde os arquivos estão localizados
    pasta := "/home/will/Documents/TCC/M_F_2/todos"

    // Abrir a pasta
    arquivos, err := ioutil.ReadDir(pasta)
    if err != nil {
        fmt.Println("Erro ao abrir a pasta:", err)
        return
    }

    // Contar o número de arquivos na pasta
    total := len(arquivos)

    // Inicializar o gerador de números aleatórios
    rand.Seed(time.Now().UnixNano())

    // Mapear nomes de arquivo gerados para evitar duplicatas
    gerados := make(map[string]bool)

    // Loop para renomear os arquivos
    for i, arquivo := range arquivos {
        extensao := filepath.Ext(arquivo.Name())
        
        // Gerar um nome aleatório único dentro do intervalo
        var novoNome string
        for {
            // Gerar um número aleatório dentro do intervalo do total de arquivos
            numeroAleatorio := rand.Intn(total) + 1
            novoNome = fmt.Sprintf("%04d%s", numeroAleatorio, extensao)
            
            // Verificar se o nome gerado já foi usado
            if _, ok := gerados[novoNome]; !ok {
                gerados[novoNome] = true
                break
            }
        }
        
        // Renomear o arquivo
        antigoCaminho := filepath.Join(pasta, arquivo.Name())
        novoCaminho := filepath.Join(pasta, novoNome)
        err = os.Rename(antigoCaminho, novoCaminho)
        if err != nil {
            fmt.Println("Erro ao renomear o arquivo:", err)
            return
        }

        fmt.Printf("Arquivo %d renomeado para: %s\n", i+1, novoNome)
    }

    fmt.Println("Renomeação concluída.")
}

