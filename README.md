соцсеть

этапы разработки:
    ✅продумать и создать базу данных
    ✅написать orm модели
    написать redis_db для работы с кешем
        ✅кеширование постов, для вычисления самых популярных за последний час
    ✅написать Storage, для установки соединения с postgres и хранения redis структуры
    ✅написать Repository-файлы для каждой таблицы бд
        2 версии Update, для PUT и PATCH.
        тразакции на изменение полей
        использовать context
    ✅написать services
        сервисы должны принимать DTO модели
        написать логику обработки запросов в services
    ✅написать handlers
        добавить заголовок Location для Create
    ✅написать dto модели
    ✅переписать 3 слоя
        ❌ВОЗМОЖНО: Например, можно сделать UserDTO (без лишних данных) и UserDetailedDTO (с полной информацией).
    ✅написать мапперы orm -> dto
        ❌добавить самостоятельный маппинг других orm объектов
    ✅переписать приложение на Gin
    ✅добавить Hasher для хеширования паролей перед загрузкой в бд
    написать отдельный пакет auth
        ✅middleware
        ✅генерация refresh, access токенов
        ✅обновление токенов
        ✅обработчики
        ✅сервисы
        ✅sхранение refresh d Redis
        ✅Hasher
    написать Logger, заменить прошлую систему логирования
    ...

возможности:
    создать канал
    создавать посты в канале
    следить за каналами
    оставлять коментарии, реакции к постам
    единая лента популярных постов
    лента популярных постов отдельных категорий
    профиль пользователя
    меню приложения
    система текстового поиска: каналов, пользователей, постов внутри каналов
    добавить личные и груповые чаты между ползователями(используя WebSockets)
    добавлять в друзья других пользователей
    платежная система, для подписок, или платных подарков
    ...

прочий функционал:
    вывод красивых ошибок на экран пользователя
    ...

✅база данных:
    список таблиц:
        users
        channels
        posts
        subscriptions
        reactions
        commentaries