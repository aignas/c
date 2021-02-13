"""
DO NOT EDIT!

This file is automatically @generated by blackjack.
"""
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def cargo_dependencies():


    http_archive(
        name = "crates_io_aho_corasick_0.7.15",
        url = "https://crates.io/api/v1/crates/aho-corasick/0.7.15/download",
        sha256 = "7404febffaa47dac81aa44dba71523c9d069b1bdc50a77db41195149e17f68e5",
        strip_prefix = "aho-corasick-0.7.15",
        type = "tar.gz",
        build_file_content = """
load("@io_bazel_rules_rust//rust:rust.bzl", "rust_library")

rust_library(
    name = "aho_corasick",
    aliases = {},
    srcs = glob(["**/*.rs"]),
    crate_type = "lib",
    deps = ["@crates_io_memchr_2.3.4//:memchr"],
    proc_macro_deps = [],
    edition = "2015",
    crate_features = ["default", "std"],
    rustc_flags = ["--cap-lints=allow"] + [],
    visibility = ["//visibility:public"],
)
    """,
    )
    

    http_archive(
        name = "crates_io_memchr_2.3.4",
        url = "https://crates.io/api/v1/crates/memchr/2.3.4/download",
        sha256 = "0ee1c47aaa256ecabcaea351eae4a9b01ef39ed810004e298d2511ed284b1525",
        strip_prefix = "memchr-2.3.4",
        type = "tar.gz",
        build_file_content = """
load("@io_bazel_rules_rust//rust:rust.bzl", "rust_library")

rust_library(
    name = "memchr",
    aliases = {},
    srcs = glob(["**/*.rs"]),
    crate_type = "lib",
    deps = [],
    proc_macro_deps = [],
    edition = "2015",
    crate_features = ["default", "std", "use_std"],
    rustc_flags = ["--cap-lints=allow"] + [],
    visibility = ["//visibility:public"],
)
    """,
    )
    

    http_archive(
        name = "crates_io_once_cell_1.5.2",
        url = "https://crates.io/api/v1/crates/once_cell/1.5.2/download",
        sha256 = "13bd41f508810a131401606d54ac32a467c97172d74ba7662562ebba5ad07fa0",
        strip_prefix = "once_cell-1.5.2",
        type = "tar.gz",
        build_file_content = """
load("@io_bazel_rules_rust//rust:rust.bzl", "rust_library")

rust_library(
    name = "once_cell",
    aliases = {},
    srcs = glob(["**/*.rs"]),
    crate_type = "lib",
    deps = [],
    proc_macro_deps = [],
    edition = "2018",
    crate_features = ["alloc", "default", "std"],
    rustc_flags = ["--cap-lints=allow"] + [],
    visibility = ["//visibility:public"],
)
    """,
    )
    

    http_archive(
        name = "crates_io_regex",
        url = "https://crates.io/api/v1/crates/regex/1.4.3/download",
        sha256 = "d9251239e129e16308e70d853559389de218ac275b515068abc96829d05b948a",
        strip_prefix = "regex-1.4.3",
        type = "tar.gz",
        build_file_content = """
load("@io_bazel_rules_rust//rust:rust.bzl", "rust_library")

rust_library(
    name = "regex",
    aliases = {},
    srcs = glob(["**/*.rs"]),
    crate_type = "lib",
    deps = ["@crates_io_aho_corasick_0.7.15//:aho_corasick", "@crates_io_memchr_2.3.4//:memchr", "@crates_io_regex_syntax_0.6.22//:regex_syntax", "@crates_io_thread_local_1.1.3//:thread_local"],
    proc_macro_deps = [],
    edition = "2015",
    crate_features = ["aho-corasick", "default", "memchr", "perf", "perf-cache", "perf-dfa", "perf-inline", "perf-literal", "std", "thread_local", "unicode", "unicode-age", "unicode-bool", "unicode-case", "unicode-gencat", "unicode-perl", "unicode-script", "unicode-segment"],
    rustc_flags = ["--cap-lints=allow"] + [],
    visibility = ["//visibility:public"],
)
    """,
    )
    

    http_archive(
        name = "crates_io_regex_syntax_0.6.22",
        url = "https://crates.io/api/v1/crates/regex-syntax/0.6.22/download",
        sha256 = "b5eb417147ba9860a96cfe72a0b93bf88fee1744b5636ec99ab20c1aa9376581",
        strip_prefix = "regex-syntax-0.6.22",
        type = "tar.gz",
        build_file_content = """
load("@io_bazel_rules_rust//rust:rust.bzl", "rust_library")

rust_library(
    name = "regex_syntax",
    aliases = {},
    srcs = glob(["**/*.rs"]),
    crate_type = "lib",
    deps = [],
    proc_macro_deps = [],
    edition = "2015",
    crate_features = ["default", "unicode", "unicode-age", "unicode-bool", "unicode-case", "unicode-gencat", "unicode-perl", "unicode-script", "unicode-segment"],
    rustc_flags = ["--cap-lints=allow"] + [],
    visibility = ["//visibility:public"],
)
    """,
    )
    

    http_archive(
        name = "crates_io_thread_local_1.1.3",
        url = "https://crates.io/api/v1/crates/thread_local/1.1.3/download",
        sha256 = "8018d24e04c95ac8790716a5987d0fec4f8b27249ffa0f7d33f1369bdfb88cbd",
        strip_prefix = "thread_local-1.1.3",
        type = "tar.gz",
        build_file_content = """
load("@io_bazel_rules_rust//rust:rust.bzl", "rust_library")

rust_library(
    name = "thread_local",
    aliases = {},
    srcs = glob(["**/*.rs"]),
    crate_type = "lib",
    deps = ["@crates_io_once_cell_1.5.2//:once_cell"],
    proc_macro_deps = [],
    edition = "2018",
    crate_features = [],
    rustc_flags = ["--cap-lints=allow"] + [],
    visibility = ["//visibility:public"],
)
    """,
    )
    
