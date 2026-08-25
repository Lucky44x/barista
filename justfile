default:
    {{just_executable()}} --list

build:
    mkdir -p "./build/bin/"
    go build -o "./build/barista"

install:
    go install
