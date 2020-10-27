package old

import "fmt"

type Chart struct {
	Type string
}

func (c *Chart) Display() {
	if c.Type == "histogram" {
		fmt.Println("display histogram chart")
	} else if c.Type == "pie" {
		fmt.Println("display pie chart")
	} else if c.Type == "line" {
		fmt.Println("display line chart")
	} else if c.Type == "bar" {
		fmt.Println("display bar chart")
	}
}

func NewChart(t string) *Chart {
	c := &Chart{
		Type: t,
	}

	if t == "histogram" {
		fmt.Println("create histogram chart")
	} else if t == "pie" {
		fmt.Println("create pie chart")
	} else if t == "line" {
		fmt.Println("create line chart")
	} else if t == "bar" {
		fmt.Println("create line chart")
	}

	return c
}
