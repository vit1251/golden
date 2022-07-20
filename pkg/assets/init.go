package assets

import (
        "embed"
)

//go:embed static/*
var Content embed.FS
