package fuzz_lipgloss_style

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/charmbracelet/lipgloss"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {

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