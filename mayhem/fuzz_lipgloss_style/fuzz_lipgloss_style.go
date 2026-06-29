package fuzz_lipgloss_style

import (
	"fmt"
	"os"

	fuzz "github.com/AdaLogics/go-fuzz-headers"

	"charm.land/lipgloss/v2"
)

func mayhemit(data []byte) int {

	if len(data) > 2 {
		num := int(data[0])
		data = data[1:]
		fuzzConsumer := fuzz.NewConsumer(data)

		switch num {

		case 255:
			// KAT path for mayhem/test.sh — known input → known lipgloss.Size output.
			w, h := lipgloss.Size("hello")
			if err := os.WriteFile("/tmp/lipgloss-kat.out", []byte(fmt.Sprintf("KAT:%d:%d\n", w, h)), 0o644); err != nil {
				return 1
			}
			if w != 5 || h != 1 {
				return 1
			}
			return 0

		case 0:
			testStyle := lipgloss.NewStyle()
			testString1, _ := fuzzConsumer.GetString()
			testString2, _ := fuzzConsumer.GetString()

			testStyle.SetString(testString1, testString2)
			return 0

		case 1:
			var testStyle lipgloss.Style
			fuzzConsumer.GenerateStruct(&testStyle)

			testStyle.Copy()
			return 0

		case 2:
			var testStyle lipgloss.Style
			fuzzConsumer.GenerateStruct(&testStyle)

			otherStyle := lipgloss.NewStyle()
			otherStyle.Inherit(testStyle)
			return 0

		case 3:
			testString, _ := fuzzConsumer.GetString()

			lipgloss.Size(testString)
			return 0

		default:
			testString, _ := fuzzConsumer.GetString()
			arrSize, _ := fuzzConsumer.GetInt()

			var intArr []int
			for i := 0; i < arrSize; i++ {
				temp, _ := fuzzConsumer.GetInt()
				intArr = append(intArr, temp)
			}

			var matchStyle lipgloss.Style
			fuzzConsumer.GenerateStruct(&matchStyle)

			var unmatchStyle lipgloss.Style
			fuzzConsumer.GenerateStruct(&unmatchStyle)

			lipgloss.StyleRunes(testString, intArr, matchStyle, unmatchStyle)
			return 0
		}
	}
	return 0
}

func Fuzz(data []byte) int {
	_ = mayhemit(data)
	return 0
}
