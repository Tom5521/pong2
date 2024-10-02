
devel := env_var_or_default("devel","")

run:
    CC=gcc go run -v -tags sdl .
run-devel:
    CC=gcc go run -v -tags "sdl devel" .

release:
    just compile
    gh release create $(git describe --tags --abbrev=0) --generate-notes ./builds/*

clean:
    rm -rf builds ./*.exe ./pong

compile:
    just build windows amd64
    just build linux amd64

build os arch:
    ./scripts/build.sh {{os}} {{arch}}


