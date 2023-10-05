# shortner
A project to learn golang as I go

## Fisrt interaction (01 and 02 october 2023)
I used fiber and gorm looking for a fast development time, as I did not adapt to gorm, changes were made.

## Second interaction (03 to 05 october 2023)
Gorm was replaced by pgx and sqlx and models were ajusted accordingly.

Models no longer interact directly with database and a module for queries were add

Database connections and migrations were changed from package models to package db

Routes structure was planed

