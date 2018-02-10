package dummy

import (
	"fmt"

	"github.com/eraclitux/trace"
)

func main() {
	fmt.Println("printed")
	trace.Println("not printed")
}
