# 📝 REST Microblog API

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Простой REST API для заметок с тегами** — мой первый самостоятельный проект на Go.

> ⚠️ **Статус проекта:** учебный пет-проект. Возможны баги, недоработки и внезапные падения.  
> 👨‍💻 **Автор:** [Nevrd](https://github.com/Nevrd) | 🎂 17 лет | 🕐 Go изучаю всего 4 месяца

---

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
- **Миграции:** нативные SQL-скрипты+[golang-migrate](https://github.com/golang-migrate/migrate)

---


# 📝 REST Microblog API

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**A simple REST API for notes with tags** — my first pet project on Go.

> ⚠️ **Project status:** educational / learning. Bugs, unexpected crashes, and missing features are expected.  
> 👨‍💻 **Author:** [Nevrd](https://github.com/Nevrd) | 🎂 17 y.o. | 🕐 Learning Go for 4 months

---

## 🌍 Choose your language | Выберите язык

- [English](#english)
- [Русский](#русский)

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
- **Database:** PostgreSQL 15
- **Router:** [gorilla/mux](https://github.com/gorilla/mux)
- **PostgreSQL driver:** [pgx](https://github.com/jackc/pgx)
- **Migrations:** native SQL scripts

---

## 🚀 Quick Start

### 1. Clone the repo
```bash
git clone https://github.com/Nevrd/REST-Microblog-Apis.git
cd REST-Microblog-Apis
