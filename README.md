Это простое веб-приложение для поиска информации о спортивных клубах. Backend реализован на Go с использованием фреймворка Gin для упрощения создания правления http-сервером и базы данных PostgreSQL. Frontend — это HTML-страница с формой поиска, использующая Bootstrap для стилей и JavaScript для взаимодействия с API. Приложение позволяет искать клубы по названию или городу, отображая результаты в таблице: название клуба, город, количество титулов и средний возраст игроков.


# Создание таблиц и демо-данных

После создания пользователя и базы данных создаются следующие таблицы:
```
CREATE TABLE clubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    city VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE club_info (
    id SERIAL PRIMARY KEY,
    club_id INTEGER REFERENCES clubs(id),
    city_id INTEGER REFERENCES cities(id),
    titles INTEGER NOT NULL,
    average_age NUMERIC(4,1) NOT NULL
);
```

А также используются следующие демо-данные:

```
INSERT INTO clubs (name) VALUES 
('Real Madrid'), ('Barcelona'), ('Manchester United'), ('Bayern Munich'), ('Juventus'), 
('Liverpool'), ('Paris Saint-Germain'), ('Chelsea'), ('AC Milan'), ('Inter Milan'), 
('Ajax'), ('Benfica'), ('Porto'), ('Atlético Madrid'), ('Arsenal'), 
('Tottenham Hotspur'), ('Monaco'), ('Atalanta'), ('Manchester City'), ('Bayer Leverkusen');

INSERT INTO cities (city) VALUES 
('Madrid'), ('Barcelona'), ('Manchester'), ('Munich'), ('Turin'), 
('Liverpool'), ('Paris'), ('London'), ('Milan'), ('Amsterdam'), 
('Lisbon'), ('Porto'), ('Moscow');

INSERT INTO club_info (club_id, city_id, titles, average_age) VALUES
(1, 1, 35, 27.5),  
(14, 1, 11, 28.3), 
(2, 2, 27, 26.8),  
(3, 3, 20, 28.2),  
(19, 3, 8, 25.8),  
(4, 4, 32, 27.0),  
(20, 4, 0, 26.2),  
(5, 5, 36, 29.1),  
(6, 6, 19, 25.9),  
(7, 7, 15, 26.4),  
(17, 7, 8, 25.4),  
(8, 8, 19, 27.3),  
(15, 8, 13, 26.7), 
(16, 8, 2, 27.1),  
(9, 9, 19, 28.5),  
(10, 9, 19, 27.7),
(18, 9, 0, 26.9),  
(11, 10, 4, 25.2), 
(12, 11, 37, 26.9),
(13, 12, 32, 27.4);
```

# Установка зависимостей

Клонируйте репозиторий 

Установите зависимости: go mod tidy
...
# Запуск

Запустите backend: go run main.go

Откройте в браузере: http://localhost:8080 

# Демонстрация работы программы

<img src="/Users/andrejkozura/Documents/Андрюшка-программист 4.1/WB/sports_clubs/demonstration.gif">
