# Yandex.Disk File Pusher

## English

This small CLI utility is useful when you want to upload single file to Yandex.Disk but don't want to fiddle with sync official client or WebDAV. It is especially useful for automating backups (e. g. in conjuction with Proxmox's ``vzdump``).

### Installation

The recommended install way is via Releases by invoking single command

```sh
# binary will be $(go env GOPATH)/bin/yapusher
curl -sfL https://install.goreleaser.com/github.com/fat0troll/yapusher.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z

# or install it into ./bin/
curl -sfL https://install.goreleaser.com/github.com/fat0troll/yapusher.sh | sh -s vX.Y.Z

# In alpine linux (as it does not come with curl by default)
wget -O - -q https://install.goreleaser.com/github.com/fat0troll/yapusher.sh | sh -s vX.Y.Z
```

If you're brave enough, or have Go installed, you can invoke

```sh
go get -u github.com/fat0troll/yapusher
```

The stability of master branch is questionable. Please consider using pre-built binaries except you facing some bugs that aren't fixed in newest release.

### Usage

Assuming that ``yapusher`` installed in your ``$PATH``, you can invoke

```sh
yapusher -h
```

to get small help on available arguments.

Before first use you need to authorize the application. Invoke ``yapusher`` without params to get URL for authorization. Open this URL in your favourite browser and provide access to your Yandex account. In return you will get the code for app. You have 10 minutes to invoke ``yapusher -authCode XXXXXXX`` to end autorization process.

For uploading single file you should run this command

```sh
yapusher -file /path/to/file -uploadPath "some/path"
```

There are some assumptions:

* ``uploadPath`` is the path from the root of your Yandex.Disk. You don't need to include first ``/``.
* This path should exist at the time of ``yapusher`` running. Currently this utility can't make directory for you before uploading.
* The file you uploading must be 10 gigabytes or less. This is Yandex restriction, not app's one.
* The file you uploading shouldn't exist in target directory. If you want to overwrite file, pass ``-force`` flag to params.

On success, the progress bar will be shown, and after the end of upload the file will appear in your Yandex.Disk.

There is no support for uploading entire directories (yet).

### Development and TODO

This utility is in early stages of development. Things may change or break. However, this utility is used by author for making ``vzdump`` backups uploads to Yandex.Disk in semi-production environment.

There are some things to do:

* Test coverage
* Creating upload path if it's not exist on Yandex.Disk
* Splitting large (more than 10 gigabytes) files to parts
* Maybe something else

### License

See [LICENSE](https://github.com/fat0troll/yapusher/blob/master/LICENSE).

## Russian

Эта маленькая консольная утилита полезна, если вы хотите загрузить единичный файл на Яндекс.Диск, но не хотите связываться с полноценным приложением для синхронизации или WebDAV. Особенно такая утилита полезна для автоматизации заливки резервных копий (например, в связке с ``vzdump`` из состава Proxmox).

### Установка

Рекомендуемый способ установки — с использованием релизов. Установить необходимый релиз можно следующей командой:

```sh
# исполняемый файл программы будет доступен в пути $(go env GOPATH)/bin/golangci-lint
curl -sfL https://install.goreleaser.com/github.com/fat0troll/yapusher.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z

# исполняемый файл программы будет доступен ./bin/
curl -sfL https://install.goreleaser.com/github.com/fat0troll/yapusher.sh | sh -s vX.Y.Z

# версия для wget (потому что Alpine Linux не имеет в поставке по умолчанию curl)
wget -O - -q https://install.goreleaser.com/github.com/fat0troll/yapusher.sh | sh -s vX.Y.Z
```

Если вы хотите собрать утилиту из исходников (и имеете установленный Go в системе), вы можете установить ``yapusher`` так:

```sh
go get -u github.com/fat0troll/yapusher
```

Стабильность ветки ``master`` находится под вопросом. Рекомендуется использовать собранную автоматически версию утилиты из релиза, если вы не являетесь разработчиком на Go или же не испытываете затруднений в работе с утилитой, исправления которых ещё не вошло в очередной релиз.

### Использование

Предполагая, что ``yapusher`` установлен в директорую из вашего ``$PATH``, его можно запустить так:

```sh
yapusher -h
```

В ответ вы получите краткую справку (на английском) о флагах, используемых в приложении.

Прежде чем использовать утилиту, её нужно авторизовать. Запустите ``yapusher`` без параметров для получения URL авторизации. Откройте полученную ссылку в вашем любимом браузере и дайте доступ к вашему аккаунту Яндекса приложению. Яндекс вернёт вам семизначный код, который в течение 10 минут необходимо предоставить приложению c помощью команды ``yapusher -authCode [полученный код]``.

Для загрузки единичного файла выполните следующую команду

```sh
yapusher -file /путь/к/файлу -uploadPath "путь/на/яндекс/диске"
```

Ожидается следующее

* ``uploadPath`` — путь на вашем Яндекс.Диске, начиная от его корня. Включать в него корневой слэш ``/`` не нужно.
* Путь на Яндекс.Диске должен существовать на момент запуска ``yapusher``. На данный момент утилита не умеет создавать себе директории для загрузки самостоятельно.
* Размер загружаемого файла не должен превышать 10 гигабайт. Это ограничение Яндекса.
* Файл не должен уже находиться в целевой директории Диска. Если он там уже есть, а вы хотите его перезаписать, добавьте к аргументам флаг ``-force``.

В случае успеха будет показан прогресс-бар, по заполнению которого загруженный файл появится в вашем Яндекс.Диске.

### Разработка и TODO

Эта утилита находится в самом начале разработки. Что-то может измениться или сломаться. Однако, автор этой программы уже использует её на пре-продакшен окружении для бекапа дампов ``vzdump``.

План разработки:

* Покрыть утилиту тестами
* Внедрить возможность создавать директорию для загрузки на Яндекс.Диске, если её ещё нет там
* Разделять большие (более 10 гигабайт) файлы на куски и загружать их по частям
* Что-нибудь ещё, список может быть расширен.

### Лицензия

См. [LICENSE](https://github.com/fat0troll/yapusher/blob/master/LICENSE).
