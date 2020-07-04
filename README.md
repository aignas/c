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
2. Setup [gazelle](https://github.com/bazelbuild/bazel-gazelle) and run:
```
$ bazel build ...
$ bazel run //:gazelle
```
3. Setup `direnv` and `gazelle` shell script
```
$ sudo apt-get install direnv
$ cat <<EOF >> ~/.zshrc
export DIRENV_LOG_FORMAT=
eval "\$(direnv hook zsh)"
EOF
$ direnv allow
$ gazelle
```
4. Setup a CI check to test that `gazelle` does not need to be run.
5. Add `bazel buildifier` and add it to the CI script.
6. Add shellcheck. Note, my setup only works on linux, but adding Mac would not be too difficult.
7. Pin buildifier to a particular tag.
8. Add bazelisk in order to not need a system `bazel` installation.
9. Add github actions to bootstrap bazelisk and run the ci script