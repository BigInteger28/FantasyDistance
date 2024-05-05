package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

var x1000 = []string{"", "kU", "MU", "GU", "TU", "PU", "EU", "ZU", "YU", "RU", "QU", "X1U", "X2U", "X3U", "X4U", "X5U", "X6U"}

var x1000text = []string{"", " Duizend ", " Miljoen ", "  Miljard ", " Biljoen ", " Biljard ", " Triljoen ", " Triljard ", "  Quadriljoen ", " Quadriljard ", " Quintiljoen ", " Quintiljard ", " Septiljoen ", " Septiljard ", "Octiljoen ", "Octiljard ", " Noniljoen "}

type Number struct {
	exactValue     *big.Int
	thousands      *big.Int
	thousandsIndex int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func formatBigNumber(numberStr string) string {
	n := len(numberStr)
	if n <= 3 {
		return numberStr
	}

	var result strings.Builder
	start := n % 3
	if start > 0 {
		result.WriteString(numberStr[:start] + " ")
	}
	for i := start; i < n; i += 3 {
		if i > start {
			result.WriteString(" ")
		}
		result.WriteString(numberStr[i : i+3])
	}
	return result.String()
}

func setThousands(value *Number) {
	comp := big.NewInt(1000)
	x := 0
	for value.exactValue.Cmp(comp) > 0 {
		comp.Mul(comp, big.NewInt(1000))
		x++
	}
	value.thousandsIndex = x
	comp.Div(comp, big.NewInt(1000))
	value.thousands = new(big.Int).Div(value.exactValue, comp)
}

func setValue(value *Number, newValue string) {
	(*value).exactValue.SetString(newValue, 10)
	setThousands(value)
}

func convToNumber(value *big.Int) *Number {
	number := &Number{
		exactValue: new(big.Int).Set(value),
		thousands:  new(big.Int),
	}
	setThousands(number)
	return number
}

func getUserInput(prompt string) *big.Int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Fout bij het lezen van de input:", err)
		return big.NewInt(0)
	}
	input = strings.TrimSpace(input) // Verwijder whitespace
	value, ok := new(big.Int).SetString(input, 10)
	if !ok {
		fmt.Println("Ongeldige input. Gelieve een geldig getal in te voeren.")
		return big.NewInt(0)
	}
	return value
}

// calculateTotalSeconds berekent het totaal aantal seconden op basis van dagen, uren, minuten en seconden
func calculateTotalSeconds() *big.Int {
	days := getUserInput("Voer het aantal dagen in: ")
	hours := getUserInput("Voer het aantal uren in: ")
	minutes := getUserInput("Voer het aantal minuten in: ")
	seconds := getUserInput("Voer het aantal seconden in: ")
	secondsPerMinute := big.NewInt(60)
	secondsPerHour := new(big.Int).Mul(secondsPerMinute, big.NewInt(60))
	secondsPerDay := new(big.Int).Mul(secondsPerHour, big.NewInt(24))
	totalSeconds := new(big.Int)
	totalSeconds.Mul(days, secondsPerDay)
	totalSeconds.Add(totalSeconds, new(big.Int).Mul(hours, secondsPerHour))
	totalSeconds.Add(totalSeconds, new(big.Int).Mul(minutes, secondsPerMinute))
	totalSeconds.Add(totalSeconds, seconds)
	return totalSeconds
}

func distanceCovered(speed *big.Int, time *big.Int) {
	units := new(big.Int).Mul(speed, time)
	distance := convToNumber(units)
	fmt.Println("Afstand afgelegd: ", formatBigNumber(distance.exactValue.String()), " = ", distance.thousands.String(), x1000text[distance.thousandsIndex], " = ", x1000[distance.thousandsIndex])
	//CONVERSIE
}

func main() {
	for {
		fmt.Println("0. ")
		fmt.Println("1. Snelheid en tijd naar Afstand")
		fmt.Println("2. ")
		fmt.Println("3. ")
		var keuze int
		fmt.Print("Keuze: ")
		fmt.Scanln(&keuze)
		if keuze == 0 {

		} else if keuze == 1 {
			speed := getUserInput("Snelheid in Units per seconde: ")
			seconds := calculateTotalSeconds()
			distanceCovered(speed, seconds)
		} else if keuze == 2 {

		} else if keuze == 3 {

		} else if keuze == 4 {

		}
	}
}
