{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  # Имя окружения
  name = "go-gin-dev-env";

  # Зависимости, которые будут установлены в окружении
  buildInputs = [
    pkgs.go # Go компилятор и инструменты
    pkgs.git # Git для работы с репозиториями
  ];

  # Переменные окружения
  shellHook = ''
    echo "Welcome to Go + Gin development environment!"
    export GOPATH=$PWD/go
    export PATH=$GOPATH/bin:$PATH
    echo "GOPATH is set to $GOPATH"
    exec zsh
  '';
}

