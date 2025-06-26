package parsers

import (
	"io"
	"regexp"
	"strings"
	"utils"

	"golang.org/x/net/html"
)

const (
	OutOfStockStatus = iota
	CriticalStatus
	LowStatus
	HighStatus
	Undefined
)

func GetHtmlResponseParser() ResponseParser {
	return &htmlResponseParser{regexpForBrandLink: regexp.MustCompile(`(https:\/\/[a-z]+\.[a-z]+\/((\w|\d)*\/){1})`)}
}

type ResponseParser interface {
	ParseBrandLinks(r io.Reader) ([]string, error)
	ParseProductLinks(r io.Reader) ([]string, error)
	ParseProductStatus(r io.Reader) (int, error)
}

type htmlResponseParser struct {
	regexpForBrandLink *regexp.Regexp
}

func (parser htmlResponseParser) ParseBrandLinks(r io.Reader) ([]string, error) {
	uniqueLinks := make(map[string]int)
	readStreamLine := utils.GetReadStreamLine(r)
	var err error
	for err != io.EOF {
		line := ""
		line, err = readStreamLine()
		if err != nil && err != io.EOF {
			break
		}

		link := parser.regexpForBrandLink.FindString(line)
		if link == "" || isGarbageBrandLink(link) {
			continue
		}

		if uniqueLinks[link] == 0 {
			uniqueLinks[link]++
		}
	}

	links := make([]string, 0)
	for key := range uniqueLinks {
		links = append(links, key)
	}

	if err == io.EOF || err == nil {
		return links, nil
	}

	return links, utils.GetStreamReadingError(err.Error())
}

func isGarbageBrandLink(link string) bool {
	garbageMark := []string{
		"/my/",
		"/order/",
		"/compare/",
		"/search/",
		"/drugie/",
		"/raznoe/",
	}

	for _, garbageMark := range garbageMark {
		if strings.Contains(link, garbageMark) {
			return true
		}
	}

	return false
}

func (p htmlResponseParser) ParseProductLinks(r io.Reader) ([]string, error) {
	rootNode, err := html.Parse(r)
	if err != nil {
		return nil, HtmlParsingError{message: err.Error()}
	}

	productsListNode := findProductsListNode(rootNode)
	if productsListNode == nil {
		return nil, NoProductListNodeParsingError{}
	}

	productLinks := make([]string, 0)
	listItem := productsListNode.FirstChild.NextSibling
	productBlock := listItem.FirstChild.NextSibling
	imageWrapper := productBlock.FirstChild.NextSibling
	link := imageWrapper.FirstChild.NextSibling.Attr[1].Val // Link from image navigation is used for better perfomance
	for listItem != nil {
		if !isGarbageBrandLink(p.regexpForBrandLink.FindString(link)) {
			productLinks = append(productLinks, link)
		}
		if listItem.NextSibling != nil && listItem.NextSibling.NextSibling != nil {
			listItem = listItem.NextSibling.NextSibling
		} else {
			break
		}
		productBlock = listItem.FirstChild.NextSibling
		imageWrapper = productBlock.FirstChild.NextSibling
		link = imageWrapper.FirstChild.NextSibling.Attr[1].Val
	}

	return productLinks, nil
}

func findProductsListNode(node *html.Node) *html.Node {
	var result *html.Node

	if node.FirstChild != nil {
		result = findProductsListNode(node.FirstChild)
	}

	if result != nil {
		return result
	}

	if node != nil && len(node.Attr) > 0 && node.Attr[0].Key == "class" && node.Attr[0].Val == "s-products-list thumbs-view" {
		return node
	}

	if node.NextSibling != nil {
		result = findProductsListNode(node.NextSibling)
	}

	return result
}

func (htmlResponseParser) ParseProductStatus(r io.Reader) (int, error) {
	readStreamLine := utils.GetReadStreamLine(r)
	var err error
	for err != io.EOF {
		line := ""
		line, err = readStreamLine()
		if err != nil && err != io.EOF {
			return Undefined, utils.GetStreamReadingError(err.Error())
		}

		status, ok := tryGetStatus(line)
		if !ok {
			continue
		}
		return status, nil
	}

	return Undefined, NoStatusParsingError{}
}

func tryGetStatus(s string) (int, bool) {
	switch {
	case strings.Contains(s, "stock-transparent"):
		return OutOfStockStatus, true
	case strings.Contains(s, "stock-critical"):
		return CriticalStatus, true
	case strings.Contains(s, "stock-low"):
		return LowStatus, true
	case strings.Contains(s, "stock-high"):
		return HighStatus, true
	default:
		return Undefined, false
	}
}
