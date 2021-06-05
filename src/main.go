package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/koenverburg/vaccine-notifier/src/notify"
	"github.com/koenverburg/vaccine-notifier/src/utils"
)

func sendMakeAppointment () {
  notify.NotifyViaSlack("vaccine", "Maak nu een afspraak!", "", "good")
} 
func sendCurrentYear(year string) {
  notify.NotifyViaSlack("vaccine", fmt.Sprintf("Nog even wachten het is nu %s", year), "", "bad")
} 

func isMatchingBirthYear(text string) (bool, string) {
  birthYear := utils.Env("BIRTH_YEAR")
  var re = regexp.MustCompile(`(?m)([\d]{4})`)
  result := re.FindString(text)
  fmt.Println(result)
  if (result == birthYear) {
    return true, result
  }
  return false, result
}

// Main func
func main() {
  c := colly.NewCollector()

  // Find and visit all links
  c.OnHTML(".info.text", func(e *colly.HTMLElement) {
    e.ForEach("h3", func(_ int, elem *colly.HTMLElement) {
      if strings.Contains(elem.Text, "Maak online uw vaccinatieafspraak!") {
        isMatch, text := isMatchingBirthYear(e.Text)
        if(isMatch) {
          sendMakeAppointment()
        } else {
          sendCurrentYear(text)
        }
      } 
    })
  })

  c.Visit(utils.Env("URL"))
}
