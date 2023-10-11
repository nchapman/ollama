package format

import "fmt"

func HumanBytes(b int64) string {
	switch {
	case b > 1024*1024*1024:
		return fmt.Sprintf("%d GiB", b/1024/1024/1024)
	case b > 1024*1024:
		return fmt.Sprintf("%d MiB", b/1024/1024)
	case b > 1024:
		return fmt.Sprintf("%d KiB", b/1024)
	default:
		return fmt.Sprintf("%d B", b)
	}
}
