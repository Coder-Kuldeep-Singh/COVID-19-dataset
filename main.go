package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://www.worldometers.info/coronavirus/"
)

type Countries struct {
	Name            string
	TotalCases      string
	NewCases        string
	TotalDeaths     string
	Newdeath        string
	TotalRecovered  string
	ActiveCases     string
	SeriousCritical string
	TotalCases1MPop string
	Deaths1MPop     string
	firstCase       string
}

func Visiturl(baseURL string) {
	response, err := http.Get(baseURL)
	ControlError("Error Url doesn't exists", err)
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	ControlError("Error while reading webpage", err)

	var countries []Countries
	document.Find("table").Each(func(i int, s *goquery.Selection) {
		tbody := s.Find("tbody")
		tbody.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			//Countries
			rowhtml.Find("tr td:nth-child(1)").Each(func(indexth int, tablecell *goquery.Selection) {
				name := tablecell.Text()
				country := Countries{
					Name: name,
				}
				countries = append(countries, country)
			})

			//Total Cases
			rowhtml.Find("tr td:nth-child(2)").Each(func(indexth int, tablecell *goquery.Selection) {
				totalcases := tablecell.Text()
				country := Countries{
					TotalCases: totalcases,
				}
				countries = append(countries, country)
			})

			//New Cases
			rowhtml.Find("tr td:nth-child(3)").Each(func(indexth int, tablecell *goquery.Selection) {
				newcases := tablecell.Text()
				country := Countries{
					NewCases: newcases,
				}
				countries = append(countries, country)
			})

			//Total Deaths
			rowhtml.Find("tr td:nth-child(4)").Each(func(indexth int, tablecell *goquery.Selection) {
				totaldeaths := tablecell.Text()
				country := Countries{
					TotalDeaths: totaldeaths,
				}
				countries = append(countries, country)
			})

			//Newdeath
			rowhtml.Find("tr td:nth-child(5)").Each(func(indexth int, tablecell *goquery.Selection) {
				newdeath := tablecell.Text()
				country := Countries{
					Newdeath: newdeath,
				}
				countries = append(countries, country)
			})

			//Total Recovered
			rowhtml.Find("tr td:nth-child(6)").Each(func(indexth int, tablecell *goquery.Selection) {
				totalrecovered := tablecell.Text()
				country := Countries{
					TotalRecovered: totalrecovered,
				}
				countries = append(countries, country)
			})

			//Active Cases
			rowhtml.Find("tr td:nth-child(7)").Each(func(indexth int, tablecell *goquery.Selection) {
				activecases := tablecell.Text()
				country := Countries{
					ActiveCases: activecases,
				}
				countries = append(countries, country)
			})

			//serious critical
			rowhtml.Find("tr td:nth-child(8)").Each(func(indexth int, tablecell *goquery.Selection) {
				seriouscritical := tablecell.Text()
				country := Countries{
					SeriousCritical: seriouscritical,
				}
				countries = append(countries, country)
			})

			//totalcase1mpop
			rowhtml.Find("tr td:nth-child(9)").Each(func(indexth int, tablecell *goquery.Selection) {
				totalcase1mpop := tablecell.Text()
				country := Countries{
					TotalCases1MPop: totalcase1mpop,
				}
				countries = append(countries, country)
			})

			//deaths1mpop
			rowhtml.Find("tr td:nth-child(10)").Each(func(indexth int, tablecell *goquery.Selection) {
				deaths1mpop := tablecell.Text()
				country := Countries{
					Deaths1MPop: deaths1mpop,
				}
				countries = append(countries, country)
			})

			//first case
			rowhtml.Find("tr td:nth-child(11)").Each(func(indexth int, tablecell *goquery.Selection) {
				firstcase := tablecell.Text()
				country := Countries{
					firstCase: firstcase,
				}
				countries = append(countries, country)
			})
		})
	})
	for _, country := range countries {
		fmt.Printf("%s\n", country)
	}

}
func ControlError(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
		return
	}
}
func main() {
	Visiturl(baseURL)
}
