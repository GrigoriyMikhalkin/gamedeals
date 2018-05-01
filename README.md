## Предварительные условия

Установить golang 1.10 и настройка среды. В ubuntu 16.04 исполнить команды в терминале:

    sudo apt-get install golang-1.10
    echo 'PATH="/usr/lib/go-1.10/bin:$PATH"' >> ~/.profile
    echo 'export GOPATH="$HOME/Go"' >> ~/.profile
    source ~/.profile

Установить и настроить Postgres: https://www.digitalocean.com/community/tutorials/postgresql-ubuntu-16-04-ru

Для dependency management используем [dep](https://github.com/golang/dep). Установка:

    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

## Установка проекта локально

Скачать проект из репозитория, в директорию `$GOPATH/src/`:

    git clone https://github.com/GrigoriyMikhalkin/gamedeals.git

Либо

    go get github.com/GrigoriyMikhalkin/gamedeals

В последнем случае, проект будет скачан в `$GOPATH/src/github/GrigoriyMikhalkin/gamedeals`. Нужно войти в эту директорию и из нее установить зависимости:

    dep ensure

Установить локальную базу данных:

    sudo su postgres
    createdb gamedeals

Также, в директории проекта, необходимо создать файл `dbconfig.json`, в котором будут определены параметры для подключения к базе данных. Пример файла:

    {
        "db_user": "postgres"
        "db_password": "postgres"
        "db_name": "gamedeals"
    }

## Разработка

Разработка ведется с использование [git flow](https://danielkummer.github.io/git-flow-cheatsheet/index.html). Для каждой фичи/бага создается отдельный бранч с названием типа `GD-100`, где 100 -- это id задачи(смотреть в trello). На каждую фичу/баг обязательно писать тесты. После завершения работы над задачей, merge request делать только после прохождения всех тестов на СI.

Код фронтенда находится в директории `frontend`.


## Запуск проекта локально

    go run main.go


