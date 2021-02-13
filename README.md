# My code monorepo

This is my code monorepo.  I don't have a lot of code, but I want to have it
all in a single repository, so that I can reuse tooling.  I also want to use
`bazel`.

## Vision

The idea is to have a single repository for all personal code I may write. Since
I am interested in organization of code and how we can do more with less effort, I
am tempted to think that a monorepo built with `bazel` may be an answer to that.

I should be able to `git clone` and build the source code after bootstrapping
the environment with the scripts within the repo. Nothing should be needed on
the system in addition to `git` and potentially a compiler.

### Structure

The structure of the monorepo is:
```
.
├── artifacts     # artifacts that are built by the CI scripts, in .gitignore.
├── src           # All source code.
├── target        # Rust build dir, in .gitignore.
├── third_party   # third_party dependencies
├── tools         # scripts and tools to make life easier, should be in PATH.
...
```

## Setup

1. Clone the repo
2. Install the build/env dependencies:
  1. `direnv`
  2. `g++`
3. Run bootstrap:
```sh
$ direnv allow
$ bootstrap.sh
```

## Known issues

There is a list of known issues, that I would like to address before I deem this to be working perfectly:
* `rules_go` does not provide an IDE integration for generated files. `gopls`
  works pretty well for source code and jumping to definition for standard
  library or dependencies used via go modules though.
  Tracked in https://github.com/bazelbuild/rules_go/issues/512
* `cargo-raze` does not work well with the IDE. There have been a few
  improvements that may be used to make this better, but I have not made it to
  work yet. Related links:
  * https://github.com/google/cargo-raze/issues/42
  * https://fuchsia.dev/fuchsia-src/development/languages/rust/editors
  * https://github.com/bazelbuild/rules_rust/issues/71
  * https://github.com/bazelbuild/rules_rust/pull/384
  
  It is possible to get autocompletion working if one symlinks //cargo/{target,Cargo.{toml,lock}} files to the library directory and moves the `.rs` files into a `src` subdir.

## Links

- [Awesome monorepo](https://github.com/korfuri/awesome-monorepo)
- [Awesome bazel](https://github.com/jin/awesome-bazel)
- [Trunk based development: Monorepos](https://trunkbaseddevelopment.com/monorepos/)
