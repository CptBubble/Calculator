package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"strconv"
)

func main() {
	calculator := app.New()
	window := calculator.NewWindow("Калькулятор кредита")
	window.Resize(fyne.NewSize(1000, 1000)) //установление размера экрана

	//	label := widget.NewLabel("Кредит")
	labelStoimNedvizhi := widget.NewLabel("Стоимость недвижимости")
	labelFirstVznos := widget.NewLabel("Первоначальный взнос")
	labelSrokKredita := widget.NewLabel("Срок кредита")
	labelProcentStavka := widget.NewLabel("Процентная ставка")

	entryStoimNedvizhi := widget.NewEntry()
	entryFirstVznos := widget.NewEntry()
	entrySrokKredita := widget.NewEntry()
	entryProcentStavka := widget.NewEntry()

	radio := widget.NewRadioGroup([]string{"Аннуитетный", "Дифференцированный"}, func(s string) {

	})

	answer := widget.NewLabel("")
	answer1 := widget.NewLabel("")
	answer2 := widget.NewLabel("")

	buttonIpoteka := widget.NewButton("Рассчитать", func() {
		perSN, err := strconv.ParseFloat(entryStoimNedvizhi.Text, 64)
		perFV, err1 := strconv.ParseFloat(entryFirstVznos.Text, 64)
		perSK, err2 := strconv.ParseFloat(entrySrokKredita.Text, 64)
		perPS, err3 := strconv.ParseFloat(entryProcentStavka.Text, 64)
		typeKredit := radio.Selected
		if typeKredit == "Аннуитетный" {

			if err != nil || err1 != nil || err2 != nil || err3 != nil {
				answer.SetText("Ошибка")
			} else {
				summaOstatka := perSN - perFV
				monthProcent := perPS / 12 / 100
				KoefAnnuenta := monthProcent * math.Pow(1+monthProcent, perSK) / ((math.Pow(1+monthProcent, perSK)) - 1)
				ItogPlatesh := summaOstatka * KoefAnnuenta
				answer.SetText(fmt.Sprintf("%.2f", ItogPlatesh))
				answer1.SetText(fmt.Sprintf("%.2f", ItogPlatesh*24))
				answer2.SetText(fmt.Sprintf("%.2f", ItogPlatesh*24-summaOstatka))
			}
		} else if typeKredit == "Дифференцированный" {
			if err != nil || err1 != nil || err2 != nil || err3 != nil {
				answer.SetText("Ошибка")
			} else {
				summaOstatka := perSN - perFV

				srok := int(perSK)
				var chetchuk float64 = 0
				var nachilprocent float64 = 0
				stringg := ""
				stringh := ""
				stringprocent := ""
				for i := 0; i < srok; i++ {

					mainPlatezh := summaOstatka / perSK
					ostatokDolga := summaOstatka - (mainPlatezh * chetchuk)
					procenti := ostatokDolga * (perPS / 100) / 12
					vznos := procenti + mainPlatezh
					nachilprocent = nachilprocent + procenti

					stringh = fmt.Sprintf(" - Последний  взнос: %.2f ", vznos)
					stringprocent = fmt.Sprintf("Проценты:  %.2f ", nachilprocent)

					if i == 0 {

						vznos1 := vznos
						stringg = fmt.Sprintf("Первый взнос: %.2f ", vznos1)
						chetchuk = chetchuk + 1

					} else {

						chetchuk = chetchuk + 1

					}

				}

				answer.SetText(stringg + stringh)
				answer1.SetText(fmt.Sprintf(stringprocent))
				answer2.SetText(fmt.Sprintf("Долг с процентами: %.2f", summaOstatka+nachilprocent))
			}
		} else {
			answer.SetText("Ошибка")
		}

	})

	window.SetContent(container.NewVBox(
		labelStoimNedvizhi,
		entryStoimNedvizhi,
		labelFirstVznos,
		entryFirstVznos,
		labelSrokKredita,
		entrySrokKredita,
		labelProcentStavka,
		entryProcentStavka,
		radio,
		buttonIpoteka,
		answer,
		answer1,
		answer2,
	))
	window.ShowAndRun()
}
