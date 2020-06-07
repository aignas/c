# My code monorepo

This is my code monorepo.  I don't have a lot of code, but I want to have it
all in a single repository, so that I can reuse tooling.  I also want to use
`bazel`.

## Setup notes

1. Initial bazel setup
```
$ touch WORKSPACE
$ touch BUILD.bazel
$ bazel build ...
# configure .gitignore
```

1. Setup [gazelle](https://github.com/bazelbuild/bazel-gazelle) and run:
```
$ bazel build ...
$ bazel run //:gazelle
```
2. Setup `direnv` and `gazelle` shell script
```
$ sudo apt-get install direnv
$ cat <<EOF >> ~/.zshrc
export DIRENV_LOG_FORMAT=
eval "\$(direnv hook zsh)"
EOF
$ direnv allow
$ gazelle
```
3. Setup a CI check to test that `gazelle` does not need to be run.
4. Add `bazel buildifier` and add it to the CI script.
5. Add shellcheck. Note, my setup only works on linux, but adding Mac would not be too difficult.
6. Profit?
