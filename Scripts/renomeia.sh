#!/bin/bash

# Pasta onde os arquivos estão localizados
pasta="/home/will/Documents/TCC/M_F_2/women"

# Verifica se a pasta existe
if [ ! -d "$pasta" ]; then
  echo "A pasta não existe."
  exit 1
fi

# Entra na pasta
cd "$pasta" || exit

# Conta o número de arquivos na pasta
total=$(ls -1 | grep -c .)

# Loop para renomear os arquivos
contador=1
for arquivo in *; do
  # Obtém a extensão do arquivo
  extensao="${arquivo##*.}"
  
  # Normaliza a extensão para .jpg
  novo_nome="${contador}.jpg"
  
  # Renomeia o arquivo
  mv "$arquivo" "$novo_nome"
  
  # Incrementa o contador
  ((contador++))
done

echo "Renomeação concluída."

