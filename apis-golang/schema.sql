- Criar tabela pessoa_fisica
CREATE TABLE pessoa_fisica (
  pessoa_id SERIAL PRIMARY KEY, -- Substituído por SERIAL
  nome_completo VARCHAR(255) NOT NULL,
  cpf VARCHAR(11) NOT NULL UNIQUE,
  email VARCHAR(255),
  telefone VARCHAR(20),
  data_nascimento DATE NOT NULL,
  senha VARCHAR(255),
  sexo CHAR(1) NOT NULL
);


CREATE TABLE pessoa_estrangeira (
    pessoa_id SERIAL PRIMARY KEY,
    nome_completo VARCHAR(255) NOT NULL,
    num_documento VARCHAR(20) NOT NULL UNIQUE, -- Número de documento específico para estrangeiros (passaporte, RNE, etc.)
    email VARCHAR(255),
    pais VARCHAR(255) NOT NULL,
    senha VARCHAR(255),
    sexo CHAR(1)
);


-- Criar tabela pessoa_juridica
CREATE TABLE pessoa_juridica (
  pessoa_id SERIAL PRIMARY KEY,
  razao_social VARCHAR(255) NOT NULL,
  cnpj VARCHAR(14) NOT NULL UNIQUE,
  email VARCHAR(255),
  telefone VARCHAR(20)
);

-- Criar tabela consumidor
CREATE TABLE consumidor (
  consumidor_id SERIAL PRIMARY KEY,
  pessoa_id INT NOT NULL,
  senha VARCHAR(255),
  FOREIGN KEY (pessoa_id) REFERENCES pessoa_fisica(pessoa_id)
);

-- Criar tabela endereco
CREATE TABLE endereco (
  endereco_id SERIAL PRIMARY KEY,
  tipo CHAR(1) NOT NULL,
  cep VARCHAR(8),
  rua VARCHAR(255),
  numero INT,
  bairro VARCHAR(255),
  complemento VARCHAR(255),
  cidade VARCHAR(255),
  uf VARCHAR(2),
  pessoa_id INT NOT NULL,
  FOREIGN KEY (pessoa_id) REFERENCES pessoa_fisica(pessoa_id)
);

-- Criar tabela atendente
CREATE TABLE atendente (
  atendente_id SERIAL PRIMARY KEY,
  nome_completo VARCHAR(255) NOT NULL,
  cpf VARCHAR(11) NOT NULL UNIQUE,
  senha VARCHAR(255),
  email VARCHAR(255),
  data_nascimento DATE
);

-- Criar tabela atendimento
CREATE TABLE atendimento (
  atendimento_id SERIAL PRIMARY KEY,
  consumidor_id INT NOT NULL,
  atendente_id INT NOT NULL,
  tipo CHAR(1) NOT NULL,
  num_acompanhamento INT,
  origem CHAR(1) NOT NULL,
  fase CHAR(1) NOT NULL,
  data DATE NOT NULL,
  hora_abertura TIME NOT NULL,
  cred_atendimento INT,
  situacao CHAR(1) NOT NULL,
  FOREIGN KEY (consumidor_id) REFERENCES consumidor(consumidor_id),
  FOREIGN KEY (atendente_id) REFERENCES atendente(atendente_id)
);

-- Criar tabela tratativa
CREATE TABLE tratativa (
  tratativa_id SERIAL PRIMARY KEY,
  atendimento_id INT NOT NULL,
  data DATE NOT NULL,
  hora TIME NOT NULL,
  descricao TEXT,
  FOREIGN KEY (atendimento_id) REFERENCES atendimento(atendimento_id)
);

-- Criar tabela denuncia
CREATE TABLE denuncia (
  denuncia_id SERIAL PRIMARY KEY,
  atendimento_id INT NOT NULL,
  como_contratou CHAR(1) NOT NULL,
  area CHAR(1) NOT NULL,
  FOREIGN KEY (atendimento_id) REFERENCES atendimento(atendimento_id)
);

-- Criar tabela gestor
CREATE TABLE gestor (
  gestor_id SERIAL PRIMARY KEY,
  nome_completo VARCHAR(255) NOT NULL,
  cpf VARCHAR(11) NOT NULL UNIQUE,
  senha VARCHAR(255),
  email VARCHAR(255),
  data_nascimento DATE
);

-- Criar tabela relato
CREATE TABLE relato (
  relato_id SERIAL PRIMARY KEY,
  atendimento_id INT NOT NULL,
  descricao TEXT,
  FOREIGN KEY (atendimento_id) REFERENCES atendimento(atendimento_id)
);