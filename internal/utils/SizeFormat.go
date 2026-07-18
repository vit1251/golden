package utils

import "fmt"

func FormatBytes(size int64) string {
    switch {
    case size >= 1<<30:
        return fmt.Sprintf("%.1f GiB", float64(size)/(1<<30))
    case size >= 1<<20:
        return fmt.Sprintf("%.1f MiB", float64(size)/(1<<20))
    case size >= 1<<10:
        return fmt.Sprintf("%d KiB", size/(1<<10))
    default:
        return fmt.Sprintf("%d B", size)
    }
}
