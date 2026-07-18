package utils

import (
    "io"
    "os"
)

func Move(src, dst string) error {
    err := os.Rename(src, dst)
    if err == nil {
        return nil
    }

    // Rename failed (cross-device, permissions, etc.), try copy+delete
    input, err := os.Open(src)
    if err != nil {
        return err
    }
    defer input.Close()

    output, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer output.Close()

    _, err = io.Copy(output, input)
    if err != nil {
        os.Remove(dst)
        return err
    }

    os.Remove(src)
    return nil
}
