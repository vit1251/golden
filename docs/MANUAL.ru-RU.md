# Инструкция по эксплуатации

Эта инструкция рассматривает наиболее популярные варианты использования Golden Point.

## Первичная настройка Goldnt Point

1. Убедитесь, что Golden Point запущен
2. Откройте в Web-браузере ссылку http://127.0.0.1:8080/setup

### Подробное описание параметров в разделе настроек

Описания параметров содержат краткое определение, а так же пример заключенный в кавычки (кавычки обозначают границы значения и не являються её частью).

#### RealName

Поле *RealName* содежит настоящее имя пользователя используемое в переписке..

Пример: "Ivan Petrov"

#### Origin

Эта строка появляется в нижней части сообщения и дает небольшой объем информации о системе, в которой возникла.

Примечание: В поле Origin можно указать путь до файла с несколькими строками записанный в кодировке UTF-8.
            В этом случае будет выбрана случайная строка из этого файла.
            Перед путем необходимо укзаать префикс "@"

Пример: "@C:\Users\vit12\Fido\Origin.txt"
Пример: "The Conference Mail BBS"

#### TearLine

Tearline provide person sign in all their messages

Пример: "Golden/{GOLDEN_PLATFORM}-{GOLDEN_ARCH} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})"

#### Inbound

Directory where store incoming packets

Пример: C:\Users\IvanP\Fido\Inbound

#### TempInbound

Directory where should be process incoming packets

Пример: C:\Users\IvanP\Fido\TempInbound

#### TempOutbound

Directory where process outbound packet

Пример: C:\Users\IvanP\Fido\TempOutbound

#### Temp

Temp directory where process packet

Пример: C:\Users\vit12\Fido\Temp

#### Outbound

Directory where store outbound packet

Пример: C:\Users\IvanP\Fido\Outbound

#### Address

Поинтовый адрес узла в FidoNet

Пример: "2:5030/1081.250"

#### NetAddr

Сетевой IP адрес узла предоставляющего услуги ФидоНет

Пример: f1081.n5030.ru:24554

#### Password

FidoNet point password

Пример: 1111abcd

#### Link

Сеетвой FTN адрес уза FidoNet uplink provide (i.e. BOSS address)

Пример: 2:5030/1081

#### Country

Страна месторасполжения Поинта

Пример: Russia

#### City

Город расположения Понита

Пример: Moscow

#### FileBox

Директория для расположения файловых эхоконференций полученных

Пример: "C:\Users\IvanP\Fido\Files

#### StationName

Ник нейм Поинта

Пример: ivanp1994



