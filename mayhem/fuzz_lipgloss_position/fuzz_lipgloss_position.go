package fuzz_lipgloss_position

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
                testStringA, _ := fuzzConsumer.GetString()
                testStringB, _ := fuzzConsumer.GetString()
                testfloat, _ := fuzzConsumer.GetFloat64()
                testPos := lipgloss.Position(testfloat)
                
                lipgloss.JoinHorizontal(testPos, testStringA, testStringB)
                return 0

            case 2:
                testStringA, _ := fuzzConsumer.GetString()
                testStringB, _ := fuzzConsumer.GetString()
                testfloat, _ := fuzzConsumer.GetFloat64()
                testPos := lipgloss.Position(testfloat)
                
                lipgloss.JoinVertical(testPos, testStringA, testStringB)
                return 0

            default:
                testWidth, _ := fuzzConsumer.GetInt()
                thestHeight, _ := fuzzConsumer.GetInt()
                testfloat1, _ := fuzzConsumer.GetFloat64()
                testfloat2, _ := fuzzConsumer.GetFloat64()
                testString, _ := fuzzConsumer.GetString()
                testhPos := lipgloss.Position(testfloat1)
                testvPos := lipgloss.Position(testfloat2)

                lipgloss.Place(testWidth, thestHeight, testhPos, testvPos, testString)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}