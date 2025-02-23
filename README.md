# FLEXIVE 

## Этапы разработки:
- ✅**Продумать и создать базу данных**
- ✅ **Написать ORM-модели**
- **Написать RedisDB для работы с кэшем**
    - ✅ Кэширование постов, для вычисления самых популярных за последний T промежуток времени
    - Сохранять refresh токены по userID
- ✅ **Написать Storage**, для установки соединения с Postgres и хранения Redis-структуры
- ✅ **Написать Repository-файлы для каждой таблицы БД**
    - Две версии Update: для **PUT** и **PATCH**
    - Транзакции на изменение полей
    - Использовать **context**
- ✅ **Написать Services**
    - Сервисы должны принимать **domain-модели**, их же и возвращать
    - мапить **domain** в **orm** и прокидывать на **repository**
    - мапить **orm** в **domain** и возвращать на **usecase**
    - Написать валидацию параметров
-  **Написать Controllers**
    - Добавить заголовок **Location** для **Create**
    - Принимать **JSON** и разбирать его в **dto**
    - кидать **dto** в **usecase**
    - принимать **dto** с **usecase** и возвращать его
- **Написать Usecases**
    - Принимать **dto** с **controllers**
    - Выполнять все логику приложения на этом слое
    - Мапить **dto** в **domain** и кидать в сервисы
    - Обрабатывать возможные ошибки с **services**
    - мапить **domain** в **dto** и возвращать в **controllers**
- **Написать DTO-модели, каждой(почти) функции контроллера**
- **Написать domain-слой**
    - Хранить **domain модели** 
    - Написать функции мапинга **orm <-> domain** 
- ✅ **Переписать приложение на Gin**
- ✅ **Добавить Hasher** для хеширования паролей перед загрузкой в БД
- **Написать слой `auth`**
    - ✅ Middleware
    - ✅ Генерация **refresh, access токенов**
    - ✅ Обновление токенов
    - ✅ Обработчики
    - ✅ Сервисы
    - ✅ Хранение **refresh в Redis**
- ✅ **Избавится от Storage**, разделить логику работы с redis и postgres на разные repository
- **Single логика приложения**
    - просмотр своего профиля
    - редактирование своего профиля
    - создание канала
    - публикация постов
    - хранение аватарки в S3
- **Написать Logger**, заменить прошлую систему логирования
- ...

## Возможности:
- Создать канал
- Создавать посты в канале
- Следить за каналами
- Оставлять комментарии, реакции к постам
- Единая лента популярных постов
- Лента популярных постов отдельных категорий
- Профиль пользователя
- Меню приложения
- Система текстового поиска: каналов, пользователей, постов внутри каналов
- Добавить **личные и групповые чаты** между пользователями (*используя WebSockets*)
- Добавлять в друзья других пользователей
- **Платежная система**, для подписок или платных подарков
- ...

## Прочий функционал:
- Вывод красивых ошибок на экран пользователя
- ...

## ✅ База данных:
### **Список таблиц:**
- `users`
- `channels`
- `posts`
- `subscriptions`
- `reactions`
- `commentaries`
