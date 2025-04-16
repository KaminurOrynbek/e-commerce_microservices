# E-commerce Microservices Platform

Реализация микросервисной платформы для электронной коммерции с использованием чистый архитектуры (Clean Architecture) и Golang.

## 📌 Обзор проекта

Система состоит из трех микросервисов:
1. **API Gateway** - обработка маршрутизации, логирования и аутентификации
2. **Inventory Service** - управление продуктами и категориями
3. **Order Service** - обработка заказов и платежей

## 🛠 Технологический стек
- **Язык программирования**: Golang
- **Фреймворки**: Gin
- **Базы данных**: 
  - Inventory Service: PostgreSQL
  - Order Service: MongoDB
- **Инструменты**: Docker, Swagger (документация API)

## 🚀 Запуск проекта
1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/KaminurOrynbek/e-commerce_microservices.git
