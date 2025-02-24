package site2

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed static
var staticFS embed.FS

//go:embed public
var publicFS embed.FS

func staticFileSystem() http.FileSystem {
    fsys, err := fs.Sub(staticFS, "static")
    if err != nil {
        panic(err)
    }
    return http.FS(fsys)
}

func publicFileSystem() http.FileSystem {
    fsys, err := fs.Sub(publicFS, "public")
    if err != nil {
        panic(err)
    }
    return http.FS(fsys)
}
