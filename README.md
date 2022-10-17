# test-scool

# 1

Написать сервис.
Стэк: Go, MySQL

Запуск сервиса с параметром порта. (port)

Endpoints:
GET /profile

Headers:
Content-type: application/json

Создать тестовую базу и импортировать структуру и данные.

scheme.sql - структура
data.sql – данные

Добавить constraints в таблицы

Написать middleware аутентификации, проверка header “Api-key” в таблице auth.
В случае неверного “Api-key”, ошибка 403.

GET /profile - отдать данные всех пользователей. Если присутствует параметр “username” отдать один объект.

На выходе обьект должен содержать: id, username, first_name, last_name, city, school  берем из таблиц user, user_profile, user_data.