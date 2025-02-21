# Library API

📚 Простой API для управления библиотекой

## Основные возможности:

*   **Авторы:** Управление данными об авторах.
*   **Книги:** Управление данными о книгах, включая связь с авторами.
*   **Пользователи:** Управление данными о пользователях, включая возможность выдачи и возврата книг.
*   **Docker-развертывание:** Готов к развертыванию с помощью `docker-compose up --build`.
*   **Документация API:** Документация Swagger доступна по адресу `http://localhost:8080/swagger/index.html` после успешного развертывания.
*   **Настройка:** Детали подключения к базе данных и другие настройки могут быть настроены через файл `.env`.
*   **Порты по умолчанию:**
    *   Web-сервер: порт `8080`
    *   База данных: порт `5432` 
