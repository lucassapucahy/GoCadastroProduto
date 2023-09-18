cd migration
psql -U postgres -c "CREATE DATABASE gocadastroproduto;"
psql -U postgres -d gocadastroproduto -f ./InsertInitialProducts.sql