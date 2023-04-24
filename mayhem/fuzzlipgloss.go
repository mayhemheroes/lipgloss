package fuzzlipgloss

import (
    "strconv"
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/charmbracelet/lipgloss"
)

func mayhemit(data []byte) int {

    var num int
    if len(data) > 2 {
        num, _ = strconv.Atoi(string(data[0]))
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            // case 0:
            //     testStringA, _ := fuzzConsumer.GetString()
            //     testStringB, _ := fuzzConsumer.GetString()
            //     testfloat, _ := fuzzConsumer.GetFloat64()
            //     testPos := lipgloss.Position(testfloat)
                
            //     lipgloss.JoinHorizontal(testPos, testStringA, testStringB)
            //     return 0

            // case 2:
            //     testStringA, _ := fuzzConsumer.GetString()
            //     testStringB, _ := fuzzConsumer.GetString()
            //     testfloat, _ := fuzzConsumer.GetFloat64()
            //     testPos := lipgloss.Position(testfloat)
                
            //     lipgloss.JoinVertical(testPos, testStringA, testStringB)
            //     return 0

            // case 3:
            //     testWidth, _ := fuzzConsumer.GetInt()
            //     thestHeight, _ := fuzzConsumer.GetInt()
            //     testfloat1, _ := fuzzConsumer.GetFloat64()
            //     testfloat2, _ := fuzzConsumer.GetFloat64()
            //     testString, _ := fuzzConsumer.GetString()
            //     testhPos := lipgloss.Position(testfloat1)
            //     testvPos := lipgloss.Position(testfloat2)

            //     lipgloss.Place(testWidth, thestHeight, testhPos, testvPos, testString)
            //     return 0

            case 4:
                testStyle := lipgloss.NewStyle()
                testString1, _ := fuzzConsumer.GetString()
                testString2, _ := fuzzConsumer.GetString()

                testStyle.SetString(testString1, testString2)
                return 0

            case 5:
                var testStyle lipgloss.Style
                fuzzConsumer.GenerateStruct(&testStyle)
                
                testStyle.Copy()
                return 0

            case 6:
                var testStyle lipgloss.Style
                fuzzConsumer.GenerateStruct(&testStyle)

                otherStyle := lipgloss.NewStyle()
                otherStyle.Inherit(testStyle)
                return 0

            case 7:
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