package main

import "github.com/pi6atv/winterhill-lib/pkg/fixtures"

func main() {
	fixtures.Send(fixtures.FullCycle[:], 9901)
}
