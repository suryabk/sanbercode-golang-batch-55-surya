-- Active: 1713531721450@@127.0.0.1@5432@db-go-sql
CREATE db-go-sql;

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    division VARCHAR(20) NOT NULL
)