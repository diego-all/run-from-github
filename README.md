# run-from-github

Local

    go run main.go init --frontend vue --db postgres my-ska

From Repo

    go run github.com/diego-all/run-from-github@latest init --frontend vue --db postgres my-ska


With Private Repo

    root@pho3nix:/home/diegoall/MAESTRIA_ING/PRUEBA/run-from-github# go run github.com/diego-all/run-from-github@latest init --frontend vue --db postgres my-ska
    go: github.com/diego-all/run-from-github@latest: module github.com/diego-all/run-from-github: git ls-remote -q origin in /home/diegoall/go/pkg/mod/cache/vcs/32adadc68215205c2e7993c9d546c078393bf6dddabf2ad3145e99429eddbf22: exit status 128:
            fatal: could not read Username for 'https://github.com': terminal prompts disabled
    Confirm the import path was entered correctly.
    If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.




