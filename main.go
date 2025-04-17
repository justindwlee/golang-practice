package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	title string
	jobCondition []string
	jobSector []string
	companyName string
}

var extractedJobs []extractedJob

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	totalPages := getPages()
	for i := 1; i <= totalPages; i++{
		fmt.Println("Extracting page ", i)
		getPage(i)
	}
	for _, job := range extractedJobs {
		fmt.Println(job)
	}
	writeJobs(extractedJobs)
	fmt.Println("Done, extracted", len(extractedJobs))
}

func writeJobs(jobs []extractedJob){
	file, err := os.Create("jobs.csv")
	checkErr(err)
	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string {"Title", "Job Condition", "Job Sector", "Company"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string {}
		jobSlice = append(jobSlice, job.title)
		jobSlice = append(jobSlice, job.companyName)
		jobConditionJoined := strings.Join(job.jobCondition, " ")
		jobSlice = append(jobSlice, jobConditionJoined)
		jobSectorJoined := strings.Join(job.jobSector, " ")
		jobSlice = append(jobSlice, jobSectorJoined)
		wErr := w.Write(jobSlice)
		checkErr(wErr)
	}
}

func getPage(page int){
	pageURL := baseURL + "&recruitPage="+ strconv.Itoa(page)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobAreas := doc.Find(".item_recruit")

	jobAreas.Each(func(i int, s *goquery.Selection){
		job := extractJob(s)
		extractedJobs = append(extractedJobs, job)
	})
}

func extractJob(s *goquery.Selection) extractedJob {
	jobTitle := s.Find(".area_job > .job_tit > a").AttrOr("title", "")

	var jobCondition []string
	s.Find(".area_job > .job_condition > span").Each(func(i int, s *goquery.Selection){
		s.Find("a").Each(func(i int, s *goquery.Selection){
			jobCondition = append(jobCondition, s.Text())
		})
		spanText := strings.TrimSpace(s.Contents().Not("a").Text())
		if spanText != "" {
			jobCondition = append(jobCondition, spanText)
		}
	})

	var jobSector []string
	s.Find(".area_job > .job_sector").Each(func(i int, s *goquery.Selection){
		s.Find("b").Each(func(i int, b *goquery.Selection){
			jobSector = append(jobSector, b.Text())
		})
	})

	companyName := s.Find(".area_corp > .corp_name > a").Text()
	companyName = strings.TrimSpace(companyName)

	job := extractedJob{
		title: jobTitle,
		jobCondition: jobCondition,
		jobSector: jobSector,
		companyName: companyName,
	}
	return job
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response){
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}