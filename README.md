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
