package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	ccsv "github.com/tsak/concurrent-csv-writer"
)

type extractedJob struct {
	title string
	jobCondition []string
	jobSector []string
	companyName string
}


var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	c := make(chan []extractedJob)
	var totalJobs []extractedJob
	totalPages := getPages()
	for i := 1; i <= totalPages; i++{
		fmt.Println("Extracting page ", i)
		go getPage(i, c)
	}
	for i :=0; i < totalPages; i++ {
		extractedJobs := <- c
		totalJobs = append(totalJobs, extractedJobs...)
	}

	writeJobs(totalJobs)
	fmt.Println("Done, extracted", len(totalJobs), " jobs")
}



func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&recruitPage="+ strconv.Itoa(page)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	jobAreas := doc.Find(".item_recruit")

	jobAreas.Each(func(i int, s *goquery.Selection){
		go extractJob(s, c)
	})
	for i:=0; i<jobAreas.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(s *goquery.Selection, c chan<- extractedJob){
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
	c<-job
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

func writeJobs(jobs []extractedJob){
	ccsvWriter, err := ccsv.NewCsvWriter("jobs.csv")
	checkErr(err)
	defer ccsvWriter.Close()

	// file, err := os.Create("jobs.csv")
	// checkErr(err)
	// utf8bom := []byte{0xEF, 0xBB, 0xBF}
	// file.Write(utf8bom)

	// w, err := ccsv.NewCsvWriter(file)
	// defer w.Flush()

	headers := []string {"Title", "Job Condition", "Job Sector", "Company"}
	wErr := ccsvWriter.Write(headers)
	checkErr(wErr)

	done := make(chan bool)

	for _, job := range jobs {
		go func(job extractedJob){
			jobSlice := extractJobSlice(job)
			wErr := ccsvWriter.Write(jobSlice)
			checkErr(wErr)
			done <- true
		}(job)
	}
	
	for i:=0; i < len(jobs); i++ {
		<-done
	}
}

func extractJobSlice(job extractedJob) []string{
	jobSlice := []string {}
	jobSlice = append(jobSlice, job.title)
	jobSlice = append(jobSlice, job.companyName)
	jobConditionJoined := strings.Join(job.jobCondition, " ")
	jobSlice = append(jobSlice, jobConditionJoined)
	jobSectorJoined := strings.Join(job.jobSector, " ")
	jobSlice = append(jobSlice, jobSectorJoined)
	return jobSlice
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