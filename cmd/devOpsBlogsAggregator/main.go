package main

import (
	"devOpsBlogsAggregator/internal/pkg/curl"
	"devOpsBlogsAggregator/internal/pkg/generator"
	"devOpsBlogsAggregator/internal/pkg/parsers"
	"devOpsBlogsAggregator/internal/pkg/parsers/kubernetes"
	"devOpsBlogsAggregator/internal/pkg/utils"
	"fmt"
	"io/ioutil"
	"sort"

	log "github.com/sirupsen/logrus"
)

func main() {
	// fmt.Println("Kek")
	// curlTest()
	kubernetesTest()
	// testUtil()
}

func curlTest() {
	doc, err := curl.GetHtmlNode("https://google.com")
	if err != nil {
		log.Error(err)
	}
	fmt.Println(doc)
}

func kubernetesTest() {
	// Get all acticles links
	// articleLinks := kubernetes.GetAllArticleURL(kubernetes.BlogsURL)

	articles := []parsers.Article{}

	// var i = 0
	// semaphore := make(chan int, 10)

	// for _, link := range articleLinks {
	// 	semaphore <- 1
	// 	go func(link string) {
	// 		defer func() {
	// 			<-semaphore
	// 		}()
	// 		articleURL := kubernetes.BaseURL + link
	// 		doc, _ := curl.GetHtmlNode(articleURL)

	// 		fmt.Println(i)
	// 		i++

	// 		var article = parsers.Article{}
	// 		kubernetes.ParserArticle(articleURL, doc, &article)
	// 		articles = append(articles, article)
	// 	}(link)
	// }
	// for len(semaphore) > 0 {
	// 	time.Sleep(time.Millisecond * 10)
	// }

	// TEST
	var article = parsers.Article{}
	articleURL := "https://kubernetes.io/blog/2016/07/Bringing-End-To-End-Kubernetes-Testing-To-Azure-2/"
	doc, _ := curl.GetHtmlNode(articleURL)
	kubernetes.ParserArticle(articleURL, doc, &article)
	//

	articles = append(articles, article)

	fmt.Println("Articles count: ", len(articles))

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].DateCreate.Before(articles[j].DateCreate)
	})

	for k, v := range articles {
		fmt.Println(k)
		tempaleMD, err := ioutil.ReadFile("./assets/template/article.gotmpl")
		if err != nil {
			panic(err)
		}

		fmt.Println(v.URL)
		fmt.Println(utils.TitleCreator(k, v.Title))

		outFile := "./blogs/kubernetes/" + utils.TitleCreator(k, v.Title)
		generator.Generate(string(tempaleMD), outFile, v)
	}
}

func testUtil() {
	t, err := utils.GetTimeFromString("Wednesday, December 01, 2021")
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Result: ", t)

	t2, err := utils.GetTimeFromString("2020.11.12")
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Result: ", t2)

	t3, err := utils.GetTimeFromString("2020.11.12 11:11")
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Result: ", t3)
}
