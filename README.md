```markdown
# 📝 REST Microblog API

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Простой REST API для заметок с тегами** — мой первый самостоятельный проект на Go.  
**A simple REST API for notes with tags** — my first self-made Go project.

> ⚠️ **Статус проекта / Project status:** учебный пет-проект. Возможны баги, недоработки и внезапные падения.  
> 👨‍💻 **Автор / Author:** [Nevrd](https://github.com/Nevrd) | 🎂 17 лет / y.o. | 🕐 Go изучаю всего 4 месяца / Learning Go for 4 months

---

## 🌍 Выберите язык | Choose your language

- [Русский](#русский)
- [English](#english)

---

# Русский

## 🎯 О проекте

Это бэкенд-сервис для ведения микроблога: можно создавать заметки, добавлять к ним теги и искать по тексту.  
Проект написан **полностью с нуля, без гайдов и туториалов**, чтобы закрепить знания по REST API, Go и PostgreSQL.

**Что уже умеет:**
- CRUD для постов (создание, чтение, обновление, удаление)
- Хранение данных в PostgreSQL
- SQL-миграции (голые .sql файлы)
- Контексты для graceful shutdown
- Валидация и обработка ошибок (в процессе)

**Чего пока нет (в планах):**
- Полноценная система тегов (many-to-many)
- Полнотекстовый поиск
- Пагинация и фильтрация
- Авторизация
- Тесты

---

## 🛠 Стек

- **Язык:** Go 1.23
- **База данных:** PostgreSQL 18
- **Роутер:** [gorilla/mux](https://github.com/gorilla/mux)
- **Драйвер PostgreSQL:** [pgx](https://github.com/jackc/pgx)
- **Миграции:** нативные SQL-скрипты + [golang-migrate](https://github.com/golang-migrate/migrate)

---

## 🚀 Быстрый старт

### 1. Клонируй репозиторий
```bash
git clone https://github.com/Nevrd/REST-Microblog-Apis.git
cd REST-Microblog-Apis
```


### 2. Примени миграции
```bash
migrate -path internal/database/migrate -database "postgres://postgres:0404@localhost:5432/postgres?sslmode=disable" up
```

### 3. Запусти сервер
```bash
go run cmd/main.go
```
Сервер слушает порт `:8080`.

---

## 📡 API (базовые эндпоинты)

| Метод | Путь              | Что делает                |
|-------|-------------------|----------------------------|
| POST  | `/posts`          | Создать новый пост         |
| GET   | `/posts`          | Получить список всех постов|
| GET   | `/posts/{title}`  | Получить пост по заголовку |
| PUT   | `/posts/{title}`  | Обновить пост (требует доработки) |
| DELETE| `/post/{title}`   | Удалить пост               |

---

## 🐞 Известные баги

- ❌ `INTSERT` вместо `INSERT` в SQL (исправляется в ближайшем коммите)
- ❌ `nil map panic` при получении списка постов
- ❌ Update не находит пост, если меняется заголовок
- ❌ Пути API неоднородны (`/posts` и `/post`)
- ❌ Пароль БД зашит в коде (скоро вынесу в `.env`)

Я знаю о них и постепенно исправляю.  
Если найдёшь ещё — смело открывай Issue, буду рад любой обратной связи!

---

## 📚 Чему я научился

- Работать с `net/http` и сторонним роутером
- Использовать `database/sql` и pgx
- Писать сырые SQL-запросы и миграции
- Проектировать структуру проекта (handler → service → repository)
- Передавать `context` для управления жизненным циклом запросов
- Принимать критику кода и не сдаваться 😊

---

## 📌 Почему этот проект важен для меня

Мне 17 лет, я самостоятельно изучаю Go четвёртый месяц. Это мой первый опыт, когда я не повторял за кем-то, а придумал и реализовал всё сам. Пусть код не идеален, но он **мой**.  
Я выложил его сюда, чтобы видеть свой прогресс и не бояться показывать незаконченное. Если ты наткнулся на этот репозиторий — спасибо, что заглянул! Советы, критика и pull requests приветствуются.

---

## 📧 Контакты

- GitHub: [@Nevrd](https://github.com/Nevrd)
- (можно добавить Telegram/Email)

---

⭐ Если не жалко, поставь звёздочку — для молодого разработчика это очень мотивирует!  
Спасибо, что зашёл.

---

# English

## 🎯 About

This is a backend service for a microblog: create posts, attach tags (coming soon), and search through content.  
The project was written **entirely from scratch, without guides or tutorials**, as a hands-on way to solidify my knowledge of REST API, Go, and PostgreSQL.

**What's already working:**
- CRUD for posts (create, read, update, delete)
- PostgreSQL storage
- SQL migrations (plain .sql files)
- Context usage for graceful shutdown
- Basic error handling (work in progress)

**What's not there yet (on the roadmap):**
- Proper many-to-many tag system
- Full-text search
- Pagination and filtering
- Authentication
- Tests

---

## 🛠 Tech Stack

- **Language:** Go 1.23
- **Database:** PostgreSQL 18
- **Router:** [gorilla/mux](https://github.com/gorilla/mux)
- **PostgreSQL driver:** [pgx](https://github.com/jackc/pgx)
- **Migrations:** native SQL scripts + [golang-migrate](https://github.com/golang-migrate/migrate)

---

## 🚀 Quick Start

### 1. Clone the repo
```bash
git clone https://github.com/Nevrd/REST-Microblog-Apis.git
cd REST-Microblog-Apis
```

### 2. Apply migrations
```bash
migrate -path migrations -database "postgres://postgres:0404@localhost:5432/postgres?sslmode=disable" up
```

### 3. Start the server
```bash
go run cmd/main.go
```
The server listens on port `:8080`.

---

## 📡 API (basic endpoints)

| Method | Path              | Action                     |
|--------|-------------------|----------------------------|
| POST   | `/posts`          | Create a new post          |
| GET    | `/posts`          | Get all posts              |
| GET    | `/posts/{title}`  | Get a post by title        |
| PUT    | `/posts/{title}`  | Update a post (needs fix)  |
| DELETE | `/post/{title}`   | Delete a post              |

---

## 🐞 Known Bugs

- ❌ `INTSERT` instead of `INSERT` in SQL (fixing in the next commit)
- ❌ `nil map panic` when retrieving the list of posts
- ❌ Update fails if the title is changed
- ❌ Inconsistent API paths (`/posts` vs `/post`)
- ❌ Database password hardcoded (will move to `.env`)

I'm aware of them and gradually fixing.  
If you spot something else — feel free to open an Issue, I'd appreciate any feedback!

---

## 📚 What I Learned

- Working with `net/http` and a third-party router
- Using `database/sql` and pgx
- Writing raw SQL queries and migrations
- Structuring a Go project (handler → service → repository)
- Using `context` to manage request lifecycles
- Accepting code reviews and not giving up 😊

---

## 📌 Why This Project Matters to Me

I am 17 years old and have been teaching myself Go for 4 months. This is my first experience of not just following a tutorial, but designing and building everything on my own. The code isn't perfect, but it's **mine**.  
I made it public to track my progress and to stop being afraid of showing unfinished work. If you've found this repo — thank you for stopping by! Advice, critique, and pull requests are more than welcome.

---

## 📧 Contacts

- GitHub: [@Nevrd](https://github.com/Nevrd)
- (add your Telegram / Email if you want)

---

⭐ A star means a lot to a young developer — thank you for the support!
```