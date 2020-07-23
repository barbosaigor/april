package april

import (
	"fmt"
	"testing"
)

func TestPickRandDepsYml(t *testing.T) {
	nodes, err := PickFromYaml("cmd/examples/conf.yml", 4)
	if err != nil {
		t.Errorf("PickRandDepsYml returned an error: %v", err)
	}
	fmt.Println(nodes)
}
