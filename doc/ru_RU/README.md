# VK Api wrapper

![help wanted](https://img.shields.io/badge/-help%20wanted-success)
[![godoc reference](https://godoc.org/github.com/xelaj/vk?status.svg)](https://godoc.org/github.com/xelaj/vk)
[![license MIT](https://img.shields.io/badge/license-MIT-green)](https://github.com/xelaj/vk/blob/master/README.md)
[![chat telegram](https://img.shields.io/badge/chat-telegram-0088cc)](https://bit.ly/2xlsVsQ)
![version v1.0.0](https://img.shields.io/badge/version-v0.1.0-red)
![unstable](https://img.shields.io/badge/stability-unstable-yellow)
<!--
code quality
golangci
contributors
go version
gitlab pipelines
-->
 

[english](https://github.com/xelaj/vk/blob/master/doc/en_US/README.md) **русский**

Простой способ интегрировать ваше приложение с Вконтакте!

<p align="center">
<img src="https://i.ibb.co/wdDc6Fm/logo.jpg"/>
</p>

## Как установить

Установка такая же простая как и у других проектов:

`go get -v github.com/xelaj/vk`

Ага, так просто!

## Как использовать

![preview](/doc/assets/short.gif)

**Примеры кода [здесь](https://github.com/xelaj/vk/blob/master/examples)**

### Simple How-To

В общем и целом обертка устроена следующим образом:

Для начала перейдите на [страницу редактирования приложений](https://vk.com/apps?act=manage) и скопируйте ID приложения, secure key (необходим только для secure.* методов, необязателен) и service key. Далее, создайте новый инстанс клиента:

``` go
client, err := vk.NewClient(vk.ClientConfig{
    ID:         clientID,
    SecureKey:  yourAppSecureKey,
    ServiceKey: yourAppAccessKey,
    KeyStorage: keystorage.NewPrimitive("vk").Set(userNameInStorage, userKey),
})
```

Обратите внимание: клиент использует keystorage, благодаря которому можно использовать один и тот же клиент для запросов от имени разных пользователей или групп. Поэтому, если вам не нужно использовать клиент во множестве сервисов, используйте примитивное хранилище, которое предоставляет `github.com/xelaj/keystorage`

Теперь у вас есть доступ ко всем методам, которые на данный момент имплементированы в библиотеке. Использовать методы очень просто: например, таким образом можно получить список сообщений от имени определенного пользователя:

``` go
client.By("vasya pupkin").MessagesGetConversations(methods.MessagesGetConversationsRequest{
    Count: 200,
    Filter: "unread",
})
```

Все! таким образом можно использовать абсолютно все методы, которые доступны на данный момент.

### Longpoll API

Longpoll API разделен на две части: для групп и для пользователей (пользовательский longpoll это legacy версия, которая до сих пор нужна что бы получать обновления)

Пакет longpoll разделен на эти части, дополнительно реализуя интерфейс общего бота. Однако у пользовательского варианта есть несколько особенностей:

1) документация не описывает четко формат сообщений, список из enum очень сложно грамотно обработать, поэтому до сих пор реализована поддержка не всех эвентов.
2) у апи есть очень много legacy кода, который кардинально отличается от группового, поэтому нет никакой возможности написать общую реализацию, хотя интерфейсы у них почти идентичные.

* [Пример пользовательского бота](https://github.com/xelaj/vk/blob/master/examples/detect_spam_words)
* [Пример бота для групп](https://github.com/xelaj/vk/blob/master/examples/simple_bot)

### Клавиатура

Клавиатура вынесена в отдельный пакет vk/keyboard. По сути, в нем описаны только типы и утилитарные функции, однако такой подход более удобен, чем ручное заполнение json.

Пример использования клавиатуры:

``` go
keyboard.New(false,
    keyboard.Col(
        keyboard.Text("green button!", "", "success"),
        keyboard.Text("red button...", "", "failure"),
    ),
    keyboard.Col(
        keyboard.Geo(`{"geo":"button"}`), // there is a bug in vk api, you must use payload only json.
    ),
)
```

### Streaming API

Streaming API был реализован частично в формате pre-alpha версии. Технически, поддержка реализована полностью, однако написать тесты не удалось, поскольку документация четко не описывает, как сделать тестовый стрим.

Если вам необходим функционал стриминга эвентов, и у вас есть желание помочь, пожалуйста, создайте ишью. Мы попробуем дописать тесты и довести код до идеального состояния!

## Нужного мне метода нет! Это норм?

Да, это норм. У API vk в вресии 5.103 существует не менее 512 методов, при чем для каждого необходимо описать параметры, ограничения и прочее. Поэтому, если вам необходимо имплементировать какой-либо метод, сделайте следующее:

1) раскомментируйте нужный метод в файле `methods.go`, а в директории `methods` проимплементируйте его
2) создайте пулл реквест с изменениями, которые вы сделали. Скорее всего, реквест мы примем в день же его создания.
3) все! Пользуйтесь! И вам хорошо, и нам приятно.

## Выглядит, что интерфейс библиотеки не очень стабилен, это так?

Да, это так. Текущий интерфейс к методам будет переделан, когда появится возможность переделать. Однако библиотека production-ready, так что беспокоиться об этом не стоит.

## Вклад в проект

пожалуйста, прочитайте [информацию о помощи](https://github.com/xelaj/vk/blob/master/doc/ru_RU/CONTRIBUTING.md), если хотите помочь. А помощь очень нужна!

## TODO

* проимплементировать все методы
* получить тестовый стенд от поддержки вк, и настроить gitlab-ci
* дописать тесты для Streaming API
* Реализовать VkPay методы
* Подготовить сервер для Callback API
* унифицировать Longpoll API для групп и пользователей

## Авторы

* **Richard Cooper** — [ololosha228](https://github.com/ololosha228)

## Лицензия

This project is licensed under the MIT License - see the [LICENSE](https://github.com/xelaj/vk/blob/master/doc/ru_RU/LICENSE.md) file for details
