package parsernews

import (
	"io"
	"bytes"
	"net/http"
	// "strings"
	
	"github.com/PuerkitoBio/goquery"
	"tengri_news/models"
	"github.com/gin-gonic/gin"
)

func ParseOneNewsByUrl(c *gin.Context){
	url := "https://tengrinews.kz/"
	// if strings.HasPrefix(c.Param("url"),"world_news"){
	// 	url += "world_news"
	// 	url += c.Param("url")
	// }else if strings.HasPrefix(c.Param("url"), "kazakhstan_news"){
	// 	url += "kazakhstan_news"
	// 	url += c.Param("url")
	// }
	url += c.Param("url")
	

	resp, err := http.Get(url)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	
	TextNew := []models.OneNew{}

	htmlData, err := io.ReadAll(resp.Body)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlData))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	targetDiv := doc.Find(".content_main_inner")

	targetDiv.Find(".content_main_text").Each(func(i int, s *goquery.Selection){
		var OneNew models.OneNew
		OneNew.Text = s.Find("p").Text()
		

		TextNew = append(TextNew, OneNew)


	})
	
	c.JSON(http.StatusOK, TextNew)

	
}

func ParseLastNews(c *gin.Context)  {
	url := "https://tengrinews.kz"
	newsList := []models.News{}

	// Выполнение GET запроса к URL
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Чтение HTML кода из ответа
	htmlData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Извлекаем текст и атрибут href из всех ссылок
	targetDiv := doc.Find("#content-1")

	targetDiv.Find(".main-news_top_item").Each(func(i int, s *goquery.Selection) {
		var news models.News

		
		// Извлекаем ссылку
		link, _ := s.Find("a").Attr("href")

		news.Link = "news" + link
	

		// Извлекаем ссылку на фото
		news.Photo, _ = s.Find(".main-news_top_item_img").Attr("src")

		// Извлекаем название статьи
		news.Title = s.Find(".main-news_top_item_title a").Text()

		// Извлекаем время
		news.Time = s.Find("time").Text()

		newsList = append(newsList, news)
	})

	c.JSON(http.StatusOK, newsList)
}

// func main() {
// 	fmt.Println(ParseLastNews)
// }

