package gocpu

import "github.com/shirou/gopsutil/cpu"
import "fmt"
import "time"
import "errors"

func GetCpuUsage() ([]string, error) {
    percentages, err := cpu.Percent(time.Second, false)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return nil, errors.New("failed to load percentages")
    }

    var output_percentages []string
    for i, percent := range percentages {
        output_percentages = append(output_percentages, fmt.Sprintf("CPU: %d - %f", i, percent))
    }

    return output_percentages, nil
}
