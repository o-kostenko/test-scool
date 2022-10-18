# test-school

# Що було зроблено

## 1 Запуск сервиса с параметром порта. (port)
зчитується змінна Environment: PORT
якщо змінна не встановлена, то за замовчуванням: 8090

## 2 створені ендпойнти 
GET /profile

GET /profile/{id}

## 3 до схеми бази данних додані обмеження:
поле таблиці `user` унікальне

поля `user_id` таблиць `user_profile` та `user_data`
є зовнішній ключ на `id` таблиці 'user'

## 4 база даних запускається в докері з параметрами:
docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:8.0


## 5 створена мідлвар функція перевірки наявності значення
header “Api-key” в таблиці auth

## 6 перевірка після запуску додатку на порту 8090 
curl 'http://localhost:8090/profile'   -H 'Accept: application/json, text/plain, */*'   -H 'Api-key: ffff-2918-xcas'

відповідь:
[{"id":1,"username":"test","first_name":"Александр","last_name":"Школьный","city":"Киев","school":"гімназія №179 міста Києва"},{"id":2,"username":"admin","first_name":"Дмитрий","last_name":"Арбузов","city":"Харьков","school":"ліцей №227"},{"id":3,"username":"guest","first_name":"Василий","last_name":"Шпак","city":"Житомир","school":"Медична гімназія №33 міста Києва"}]

curl 'http://localhost:8090/profile'   -H 'Accept: application/json, text/plain, */*'   -H 'Api-key: ffff-'

відповідь:
"wrong api-key"


# task
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


