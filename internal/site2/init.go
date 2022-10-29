package site2

import (
        "embed"
)

//go:embed static/*
var staticContent embed.FS

//go:embed public/*
var publicContent embed.FS
