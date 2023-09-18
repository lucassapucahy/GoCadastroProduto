echo "Iniciando o banco de dados."
PowerShell.exe -ExecutionPolicy Bypass -File "run-compose.ps1"
echo "Iniciando o Aplicacao."
PowerShell.exe -ExecutionPolicy Bypass -File "run-app.ps1"

