сервис реализован на go с использованием gin+gorm 
приложение контейнеризированно docker

для запуска:

1) настройте .env в корне проекта

.env example:
---------------------------------------------------------------------------------------------------
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=pr_reviewer

SERVER_PORT=8080
SSL_MODE=disable

DATABASE_URL=postgres://postgres:password@postgres:5432/pr_reviewer?sslmode=disable
---------------------------------------------------------------------------------------------------

2) запустите docker compose -d --build

Во время разработки была изменена структура бд: добавлена доп таблица для хранения reviewers
Не был реализованн рандомный выбор ревьюверов
   
