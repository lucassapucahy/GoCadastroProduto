
CREATE TABLE "produtos"(
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    quantidade integer
);

INSERT INTO "produtos" (nome,descricao,preco,quantidade) VALUES
('Camiseta', '100% algodao', 59.90, 5),
('Computador',  'Dell', 2500.00, 1),
('Tenis', 'Numero 41', 499.00, 10)