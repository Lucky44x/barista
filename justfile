default:
    {{just_executable()}} --list

build:
    mkdir -p "./build/"
    go build -o "./build/barista"

install:
    go install
