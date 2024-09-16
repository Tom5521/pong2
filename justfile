
skip-compress := env_var_or_default("SKIP_COMPRESS","0")
devel := env_var_or_default("devel","")

run:
    CC=gcc go run -v -tags rgfw .
run-devel:
    CC=clang go run -v -tags "rgfw devel" .

release:
    just compile
    gh release create $(git describe --tags --abbrev=0) --generate-notes ./builds/*

clean:
    rm -rf builds ./*.exe ./pong

compile:
    just build windows amd64
    just build linux amd64

build os arch:
    SKIP_COMPRESS={{skip-compress}} ./scripts/build.sh {{os}} {{arch}}

compress bin:
    ./scripts/compress.sh {{bin}}
