package pageobjects

import (
	"github.com/tebeka/selenium"
	"practica3/support"
	"strings"
	"time"
)

var driver selenium.WebDriver

var(
	searchInput = "input#search_query_top"
	searchButton = "button[name='submit_search']"
	listViewButton = "i.icon-th-list"
	addToCartButton = "a.ajax_add_to_cart_button"
	cartMessage = "span[class*=ajax_cart_product_txt]:not([class*='unvisible'])"
)

func ByCss (locator string) (element selenium.WebElement, err error) {
	element, err = driver.FindElement(selenium.ByCSSSelector, locator)
	if err != nil {
		return
	}
	return element, nil
}

func InitWebDriver(){
	driver = support.WDInit()
}

func LoadPage (url string){
	driver.Get(url)
}

func WriteText (term string) (err error) {
	//driver.SetImplicitWaitTimeout(time.Second * 5)
	element, err := ByCss(searchInput)
	if err != nil{
		return
	}
	element.SendKeys(term)
	return nil
}

func ClickSearch () {
	element, err := ByCss(searchButton)
	if err != nil{
		return
	}
	element.Click()
}

func ClickListView () {
	element, err := ByCss(listViewButton)
	if err != nil{
		return
	}
	element.Click()
}

func AddToCart(){
	//driver.SetImplicitWaitTimeout(time.Second * 1000000)
	element, err := ByCss(addToCartButton)
	if err != nil{
		return
	}
	element.Click()
}

func GetAddToCartMessage() (message string){
	element, err := ByCss(cartMessage)
	if err != nil{
		return
	}
	elementInnerText, _ := element.GetAttribute("innerText")
	message = strings.TrimSpace(elementInnerText)

	//if elementInnerText != message {
	//	return fmt.Errorf("Esperaba: %v - Obtenido: %v", message, elementInnerText)
	//}
	return message
}

func TakeScreenShot(fileName string){
	//driver.SetImplicitWaitTimeout(time.Millisecond * 2000)
	time.Sleep(time.Millisecond * 2000)
	shot, _ := driver.Screenshot()
	support.SaveImage(shot, fileName)
}

func ClosePage (){
	driver.Quit()
}