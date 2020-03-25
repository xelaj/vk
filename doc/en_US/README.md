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

**english** [русский](https://github.com/xelaj/vk/blob/master/doc/ru_RU/README.md)

The Simplest way to integrate your app with Vk!

<p align="center">
<img src="https://i.ibb.co/wdDc6Fm/logo.jpg"/>
</p>

## Getting started

Installing is simple as for other packages:

`go get -v github.com/xelaj/vk`

Yeup! So easy and simple!

## How to use

![preview](/doc/assets/short.gif)

**Code examples are [here](https://github.com/xelaj/vk/blob/master/examples)**

### Simple How-To

In general, wrapper working something like this:

First of all, go to [managing apps page](https://vk.com/apps?act=manage), and copy appID, Sevure key (it is required only for secure.* methods, not required) and service key. Next, create new client instance:

``` go
client, err := vk.NewClient(vk.ClientConfig{
    ID:         clientID,
    SecureKey:  yourAppSecureKey,
    ServiceKey: yourAppAccessKey,
    KeyStorage: keystorage.NewPrimitive("vk").Set(userNameInStorage, userKey),
})
```

Warn: client uses `keystorage` package, which helps using single instance of client for multiple users requesting. If you don't need this feature, just use `PrimitiveStorage`, which provide `github.com/xelaj/keystorage` package, for single user (or group) requesting.

Now, you have access to all methods, which are implemented for current time. Use methods is so simple: for example, code below shows, how to get all converstaions by wanted user:

``` go
client.By("vasya pupkin").MessagesGetConversations(methods.MessagesGetConversationsRequest{
    Count: 200,
    Filter: "unread",
})
```

That's all! You can request absolutely all methods same way, which are implemented at this moment.

### Longpoll API

Longpoll API is splitted for two parts: for the groups and for the users (user longpoll is legacy version, which still required for getting updates)

Package longpoll is splitted for these part, additionally implements common `Bot` interface. But user version has two differences:

1) the documentation does not clearly describe messages format, so list of enums is difficult to correctly handle, so cause of this, not of all event types is supported correctly (technicaly it is exist, but in experemental way)
2) api has **A LOT** of legacy code is fundamentally different from group longpoll. So there is no way to create realization, but theoretically, interfaces are near the same.

* [example of user bot](https://github.com/xelaj/vk/blob/master/examples/detect_spam_words)
* [example of group bot](https://github.com/xelaj/vk/blob/master/examples/simple_bot)

### Keyboard

Message keyboard is separated package. In general, it describes only types and utility functions, but this approach is more usable, than raw json items.

Example of package usage:

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

Streaming API is partially implemented as pre-alplha version. Technically, support is fully implemented, howver, it wasn't possible to  write unit tests, cause vk api docs doesn't clearly describe how to make a test stream.

If you need the functionality of streaming events, and you have a desire to help, please create issue. We will try to add tests and bring the code to perfect condition!

## I can't find some method! Is this okay?

Yeup, it's okay. Vk API has at least 512 methods in v5.103, and for every method we need to implement parameters, argument restrictions, etc. for each. Therefore, if you need to implement any method, required for you, do the following:

1) uncomment the desired method in the `methods.go` file, and implement it in the` methods` directory
2) create a pull request with the changes you made. Most likely, we will accept the request on the day of its creation.
3) That's it! Use it! Both you and we feel good and pleased 😘

## Looks like the library interface is not very stable, is it?

Yes it is. The current interface to methods will be refactor when the authors will have some time to doi it. However, this lib is production-ready, so don't worry about that.

## Contributing

Please read [contributing guide](https://github.com/xelaj/vk/blob/master/doc/en_US/CONTRIBUTING.md) if you want to help. And the help is very necessary!

## TODO

* implements all methods
* get the test stand by vk support, and setup gitlab-ci
* done unit tests for Streaming APi
* Implement VkPay API Methods
* Prepare custom http server for Callback API
* unify Longpoll API for groups and users

## Authors

* **Richard Cooper** — [ololosha228](https://github.com/ololosha228)

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/xelaj/vk/blob/master/doc/ru_RU/LICENSE.md) file for details
