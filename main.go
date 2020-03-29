package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
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

func Visiturl(baseURL string) []Countries {
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
				name := strings.TrimSpace(tablecell.Text())
				country := Countries{
					Name: name,
				}
				countries = append(countries, country)
			})

			//Total Cases
			rowhtml.Find("tr td:nth-child(2)").Each(func(indexth int, tablecell *goquery.Selection) {
				totalcases := strings.TrimSpace(tablecell.Text())
				country := Countries{
					TotalCases: totalcases,
				}
				countries = append(countries, country)
			})

			//New Cases
			rowhtml.Find("tr td:nth-child(3)").Each(func(indexth int, tablecell *goquery.Selection) {
				newcases := strings.TrimSpace(tablecell.Text())
				country := Countries{
					NewCases: newcases,
				}
				countries = append(countries, country)
			})

			//Total Deaths
			rowhtml.Find("tr td:nth-child(4)").Each(func(indexth int, tablecell *goquery.Selection) {
				totaldeaths := strings.TrimSpace(tablecell.Text())
				country := Countries{
					TotalDeaths: totaldeaths,
				}
				countries = append(countries, country)
			})

			//Newdeath
			rowhtml.Find("tr td:nth-child(5)").Each(func(indexth int, tablecell *goquery.Selection) {
				newdeath := strings.TrimSpace(tablecell.Text())
				country := Countries{
					Newdeath: newdeath,
				}
				countries = append(countries, country)
			})

			//Total Recovered
			rowhtml.Find("tr td:nth-child(6)").Each(func(indexth int, tablecell *goquery.Selection) {
				totalrecovered := strings.TrimSpace(tablecell.Text())
				country := Countries{
					TotalRecovered: totalrecovered,
				}
				countries = append(countries, country)
			})

			//Active Cases
			rowhtml.Find("tr td:nth-child(7)").Each(func(indexth int, tablecell *goquery.Selection) {
				activecases := strings.TrimSpace(tablecell.Text())
				country := Countries{
					ActiveCases: activecases,
				}
				countries = append(countries, country)
			})

			//serious critical
			rowhtml.Find("tr td:nth-child(8)").Each(func(indexth int, tablecell *goquery.Selection) {
				seriouscritical := strings.TrimSpace(tablecell.Text())
				country := Countries{
					SeriousCritical: seriouscritical,
				}
				countries = append(countries, country)
			})

			//totalcase1mpop
			rowhtml.Find("tr td:nth-child(9)").Each(func(indexth int, tablecell *goquery.Selection) {
				totalcase1mpop := strings.TrimSpace(tablecell.Text())
				country := Countries{
					TotalCases1MPop: totalcase1mpop,
				}
				countries = append(countries, country)
			})

			//deaths1mpop
			rowhtml.Find("tr td:nth-child(10)").Each(func(indexth int, tablecell *goquery.Selection) {
				deaths1mpop := strings.TrimSpace(tablecell.Text())
				country := Countries{
					Deaths1MPop: deaths1mpop,
				}
				countries = append(countries, country)
			})

			//first case
			rowhtml.Find("tr td:nth-child(11)").Each(func(indexth int, tablecell *goquery.Selection) {
				firstcase := strings.TrimSpace(tablecell.Text())
				country := Countries{
					firstCase: firstcase,
				}
				countries = append(countries, country)
			})
		})
	})
	return countries

}

func ControlError(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
		return
	}
}

func RenderTable(w http.ResponseWriter, req *http.Request) {
	values := Visiturl(baseURL)
	html := `<table border="1">`
	html += `<thead>`
	html += `<th>Countries</th>`
	html += `<th>Total Cases</th>`
	html += `<th>New Cases</th>`
	html += `<th>Total Deaths</th>`
	html += `<th>New Deaths</th>`
	html += `<th>Total Recovered</th>`
	html += `<th>Active Cases</th>`
	html += `<th>Serious Critical</th>`
	html += `<th>Total Cases 1M Pop</th>`
	html += `<th>Death 1M Pop</th>`
	html += `<th>First Case</th>`
	html += `</thead>`
	html += `<tbody>`
	for _, value := range values {
		// fmt.Fprintf("")
		html += `<tr>`
		html += "<td>" + value.Name + "</td>"
		html += "<td>" + value.TotalCases + "</td>"
		html += "<td>" + value.NewCases + "</td>"
		html += "<td>" + value.TotalDeaths + "</td>"
		html += "<td>" + value.Newdeath + "</td>"
		html += "<td>" + value.TotalRecovered + "</td>"
		html += "<td>" + value.ActiveCases + "</td>"
		html += "<td>" + value.SeriousCritical + "</td>"
		html += "<td>" + value.TotalCases1MPop + "</td>"
		html += "<td>" + value.Deaths1MPop + "</td>"
		html += "<td>" + value.firstCase + "</td>"
		html += `</tr>`
	}
	html += `</tbody>`
	html += `</table>`
	fmt.Fprintf(w, html)
}

func DBConnect() (db *sql.DB) {
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DB")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname)
	if err != nil {
		log.Println("Connection String failed", err)
	}
	fmt.Println("connected")
	return db
}

func InsertStatistics(w http.ResponseWriter, req *http.Request) {
	db := DBConnect()
	countries := Visiturl(baseURL)
	for _, country := range countries {
		// inserted, err := db.Prepare("INSERT INTO statistics(country,total_cases,new_cases,total_deaths,new_death,total_recovered,active_cases,serious_critical,total_cases_1M_pop,death_1M_pop,first_death) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
		inserted, err := db.Prepare("INSERT INTO statistics(country) VALUES(?)")
		if err != nil {
			log.Println("Error while Inserting data", err.Error())
		}
		// executing, err := inserted.Exec(country.Name, country.TotalCases, country.NewCases, country.TotalDeaths, country.Newdeath, country.TotalRecovered, country.ActiveCases, country.SeriousCritical, country.TotalCases1MPop, country.Deaths1MPop, country.firstCase)
		executing, err := inserted.Exec(country.Name)
		if err != nil {
			log.Println("Error to Executing the Insert Statement", err)
			return
		}
		fmt.Println(executing)

	}
	fmt.Fprintf(w, "Query Executed Successfully")
	defer db.Close()

}

func RunCrawler(w http.ResponseWriter, req *http.Request) {
	_ = Visiturl(baseURL)
	fmt.Fprintf(w, "Crawling finished")
	// fmt.Println(value)
}

func LandingPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Landing Page")
}

func main() {
	http.HandleFunc("/", LandingPage)
	http.HandleFunc("/run", RunCrawler)
	http.HandleFunc("/table", RenderTable)
	http.HandleFunc("/query", InsertStatistics)
	fmt.Println("Development Server started :8080")
	http.ListenAndServe(":8080", nil)

}

// CREATE TABLE statistics ( id INT() UNSIGNED AUTO_INCREMENT PRIMARY KEY, counntry VARCHAR(30) NOT NULL, total_cases VARCHAR(30) NOT NULL,new_cases VARCHAR(50) NOT NULL,total_deaths  VARCHAR(255) NOT NULL, new_death VARCHAR(255) NOT NULL,total_recovered VARCHAR(255) NOT NULL,active_cases VARCHAR(255) NOT NULL,serious_critical VARCHAR(255) NOT NULL, total_cases_1M_pop VARCHAR(255) NOT NULL,death_1M_pop VARCHAR(255) NOT NULL,fisrt_death VARCHAR(255) NOT NULL)
