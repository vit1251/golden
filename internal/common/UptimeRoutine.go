package commonfunc

import "time"

var processStartTime = time.Now()

func GetUptime() time.Duration {
    return time.Since(processStartTime)
}
