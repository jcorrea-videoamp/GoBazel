# GoBazel
First steps with Bazel for building projects in Go

## Steps

<ol>
  <li> Create a new project </li>
  <li> Add a `BUILD.bazel` file to the project </li>
  <li> Add a `WORKSPACE.bazel` file to the project </li>
  <li> Modify the `WORKSPACE.bazel` file to assign a name to the workspace </li>
  
  .. code:: bzl
    workspace(
        # How this workspace would be referenced with absolute labels from another workspace
        name = "gobazel",
    )
  
  <li> Modify the `WORKSPACE.bazel` file to include `io_bazel_rules_go` and `bazell_gazelle` building rules. Reference: https://github.com/bazelbuild/rules_go#generating-build-files</li>
  
  .. code:: bzl
    load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

    http_archive(
        name = "io_bazel_rules_go",
        sha256 = "f2dcd210c7095febe54b804bb1cd3a58fe8435a909db2ec04e31542631cf715c",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
        ],
    )

    http_archive(
        name = "bazel_gazelle",
        sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        ],
    )

    load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
    load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

    go_rules_dependencies()

    go_register_toolchains(version = "1.18")

    gazelle_dependencies()
  
  <li> Add a library to your project: Create a new folder that will contain the source and test files for your library. 
    For this example project I created the folder ``go_hello_world`` under the projects folder.</li>
  <li> Add a `BUILD.bazel` file to your library's folder. Modify the file to include the rules for building go library and test in bazel. </li>
  
  .. code:: bzl
    load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
  
    go_library(
        name = "foo_library",
        srcs = [
            "a.go",
            "b.go",
        ],
        importpath = "github.com/example/project/foo",
        deps = [
            "//tools",
            "@org_golang_x_utils//stuff",
        ],
        visibility = ["//visibility:public"],
    )
  
    go_test(
        name = "foo_test",
        srcs = [
            "a_test.go",
            "b_test.go",
        ],
        embed = [":foo_lib"],
        deps = [
            "//testtools",
            "@org_golang_x_utils//morestuff",
        ],
    )
  
  For this example project I used [goconvey](https://github.com/smartystreets/goconvey) for test assertions so I had to modify not only the `BUILD.bazel` file in the library but also the `WORKSPACE.bazel` file in the root folder to make the project work with repositoires.
  
  To create the go_repostory entry in the `WORKSPACE.bazel` file you need to create the name accordingly, `com_github_owner_repo`, the importpath
  is the repo url and the tag can be grabbed from the repo.
  
  To add the dependency to the deps in the go_test entry in the `BUILD.bazel` file of the library you need to fetch the name using this command
  `bazel query @com_github_owner_repo//...`. For goconvey: `bazel query @com_github_smartystreets_goconvey//...`
  
  <li> Add an entry point to your project: Create a folder that will contain the source file with the main functioin of your project (and tests?).
    For this example I created the folder `go_web` with a single main file. </li>
  
  <li> Add a `BUILD.bazel` file to your folder and modify it to include the rules to build a go binay. </li>
  
  .. code:: bzl
    load("@io_bazel_rules_go//go:def.bzl", "go_binary")
  
    go_binary(
        name = "foo",
        srcs = ["main.go"],
    )
  
  In this example project I used [gorilla/mux](https://github.com/gorilla/mux) as a dependency, so I had to add the repository in the `WORKSPACE.bazel` file and modify the `go_binary` entry in the `BUILD.bazel` file to include an additional external dependency.
  
  </ol>
  
  
  

