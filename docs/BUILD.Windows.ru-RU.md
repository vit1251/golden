# Сборка из исходного кода

## Подготовка окружения

 1) Установите набор трансляторов **TDM-GCC**. Можно скачать по ссылке https://jmeubank.github.io/tdm-gcc/download/
 2) Установите систему контроля версий **GitSCM**. Можно скачать по ссылке https://git-scm.com/downloads

## Сборка с использованием сценария PowerShell

 1) Запустите команду сборки в директории с исходным кодом (например, в директории C:\Golden\src):

```
C:\Golden\src> powershell -executionpolicy RemoteSigned -file "build-windows.ps1"
```

## Сборка в ручном режиме

 1) (Опционально) Подготовтье Golang окружение для целевой платформы
 2) Выполните последовательно следующие команды

```
C:\Golden\src> go generate
C:\Golden\src> go build
```
