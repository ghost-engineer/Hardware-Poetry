package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Haiku structure: 5-7-5 syllables
// We'll map soil moisture (0-100) to syllable counts for each line
// Example: dry = 5, moist = 7, wet = 5

// Example word banks for each syllable count
var words = map[int][]string{
	1: {"rain", "sun", "dirt", "leaf", "dew", "mud", "seed", "root", "bud", "worm"},
	2: {"petal", "garden", "water", "flower", "blossom", "morning", "evening", "shadow", "meadow", "beetle"},
	3: {"butterfly", "vegetable", "humidity", "tomorrow", "delicate", "beautiful", "sunflower", "cucumber", "tomato", "lavender"},
	4: {"evaporation", "photosynthesis", "transpiration", "fertilizer", "irrigation", "pollination", "germination", "cultivation", "sprinkling", "harvesting"},
	5: {"unfurling petals", "gentle summer rain", "roots beneath the earth", "shadows softly fall", "sunlight warms the leaves", "dew upon the grass", "earth embraces seed", "breeze whispers softly", "clouds drift overhead", "morning in the field"},
	7: {"flowers open in the dawn", "moisture clings to every leaf", "gentle rain falls on the earth", "roots drink deeply in the dark", "sunrise paints the garden gold", "petals shimmer in the breeze", "earth awakens with the light", "shadows dance across the field", "beetles wander through the grass", "dew collects on every blade"},
}

// Map soil moisture to syllable count for each line
func moistureToSyllables(moisture int) (int, int, int) {
	// 0-33: dry, 34-66: moist, 67-100: wet
	if moisture <= 33 {
		return 5, 7, 5
	} else if moisture <= 66 {
		return 7, 5, 7
	} else {
		return 5, 5, 7
	}
}

// Generate a line with a given syllable count
func generateLine(syllables int) string {
	if ws, ok := words[syllables]; ok {
		return ws[rand.Intn(len(ws))]
	}
	// Try to build a line from smaller words
	line := ""
	remaining := syllables
	for remaining > 0 {
		max := remaining
		if max > 3 {
			max = 3
		}
		w := words[max][rand.Intn(len(words[max]))]
		if line != "" {
			line += " "
		}
		line += w
		remaining -= max
	}
	return line
}

func generateHaiku(moisture int) string {
	l1, l2, l3 := moistureToSyllables(moisture)
	return fmt.Sprintf("%s\n%s\n%s", generateLine(l1), generateLine(l2), generateLine(l3))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var moisture int
	fmt.Print("Enter soil moisture level (0-100): ")
	fmt.Scan(&moisture)
	if moisture < 0 {
		moisture = 0
	}
	if moisture > 100 {
		moisture = 100
	}
	haiku := generateHaiku(moisture)
	fmt.Println("\nYour garden haiku:")
	fmt.Println(haiku)
}
