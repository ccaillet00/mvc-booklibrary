package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Globale Slices für Datenhaltung
var noten []float64
var faecher []string

func askUser(question string) string {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}
func istGueltigeNote(note float64) bool {
	return note >= 1.0 && note <= 6.0
}

func getNotengrade(note float64) string {
	if note >= 5.5 {
		return "Sehr gut"
	} else if note >= 5.0 {
		return "Gut"
	} else if note >= 4.0 {
		return "Genügend"
	} else {
		return "Ungenügend"
	}
}

func noteHinzufuegen() {
	fach := askUser("Fach:")
	if fach == "" {
		fmt.Println("❌ Fach darf nicht leer sein!")
		return
	}

	noteText := askUser("Note (1.0 - 6.0):")
	note, err := strconv.ParseFloat(noteText, 64)
	if err != nil {
		fmt.Println("❌ Ungültige Note! Bitte Zahl zwischen 1.0 und 6.0 eingeben.")
		return
	}

	if !istGueltigeNote(note) {
		fmt.Println("❌ Note muss zwischen 1.0 und 6.0 liegen!")
		return
	}

	// Note und Fach zu Slices hinzufügen
	noten = append(noten, note)
	faecher = append(faecher, fach)

	fmt.Printf("✅ Note %.1f für %s hinzugefügt! (%s)\n", note, fach, getNotengrade(note))
}

func alleNotenAnzeigen() {
	if len(noten) == 0 {
		fmt.Println("📝 Noch keine Noten vorhanden.")
		return
	}

	fmt.Println("\n=== ALLE NOTEN ===")
	for i := 0; i < len(noten); i++ {
		grade := getNotengrade(noten[i])
		fmt.Printf("%d. %s: %.1f (%s)\n", i+1, faecher[i], noten[i], grade)
	}
}

func durchschnittBerechnen() {
	if len(noten) == 0 {
		fmt.Println("📝 Noch keine Noten vorhanden.")
		return
	}

	summe := 0.0
	for i := 0; i < len(noten); i++ {
		summe += noten[i]
	}
	durchschnitt := summe / float64(len(noten))

	fmt.Printf("\n📊 DURCHSCHNITT\n")
	fmt.Printf("Anzahl Noten: %d\n", len(noten))
	fmt.Printf("Durchschnitt: %.2f (%s)\n", durchschnitt, getNotengrade(durchschnitt))
}

func main() {
	fmt.Println("=== 🎓 NOTENVERWALTUNG ===")

	for {
		fmt.Println("\n--- HAUPTMENÜ ---")
		fmt.Println("1. Note hinzufügen")
		fmt.Println("2. Alle Noten anzeigen")
		fmt.Println("3. Durchschnitt berechnen")
		fmt.Println("4. Beste/Schlechteste Note")
		fmt.Println("5. Notenstatistik")
		fmt.Println("6. Beenden")

		auswahl := askUser("Deine Wahl (1-6):")

		switch auswahl {
		case "1":
			noteHinzufuegen()
		case "2":
			alleNotenAnzeigen()
		case "3":
			durchschnittBerechnen()
		case "4":
			fmt.Println("Beste/Schlechteste Note") //not yet implemented
		case "5":
			fmt.Println("Notenstatistik") //not yet implemented
		case "6":
			fmt.Println("👋 Auf Wiedersehen!")
			return
		default:
			fmt.Println("❌ Ungültige Auswahl! Bitte wähle 1-6.")
		}
	}
}
