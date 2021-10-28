package dec

import (
	"fmt"
	"math/big"
        "os"
)

func dec() {
  var intString string
  intString = os.Args[1]
  var dValue *big.Int
  dValue = new(big.Int)
  value, ok := dValue.SetString(intString, 10)
  if !ok {
	fmt.Println("can't convert to decimal")
  }
  fmt.Println(value)
}
