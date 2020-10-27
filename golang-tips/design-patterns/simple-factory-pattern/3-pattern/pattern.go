package pattern

import "fmt"

type Chart interface {
	Display()
}

type HistogramChart struct {
}

func (chart *HistogramChart) Display() {
	fmt.Println("display histogram chart")
}

func NewHistogramChart() *HistogramChart {
	fmt.Println("create histogram chart")

	return &HistogramChart{}
}

type PieChart struct {
}

func (chart *PieChart) Display() {
	fmt.Println("display pie chart")
}

func NewPieChart() *PieChart {
	fmt.Println("create pie chart")

	return &PieChart{}
}

type LineChart struct {
}

func (chart *LineChart) Display() {
	fmt.Println("display line chart")
}

func NewLineChart() *LineChart {
	fmt.Println("create line chart")

	return &LineChart{}
}

type BarChart struct {
}

func (chart *BarChart) Display() {
	fmt.Println("display line chart")
}

func NewBarChart() *BarChart {
	fmt.Println("create line chart")

	return &BarChart{}
}

var _ Chart = &HistogramChart{}
var _ Chart = &PieChart{}
var _ Chart = &LineChart{}

// 由于是static方法, 无须使用struct
func GetChart(t string) Chart {
	if t == "histogram" {
		return NewHistogramChart()
	} else if t == "pie" {
		return NewPieChart()
	} else if t == "line" {
		return NewLineChart()
	} else if t == "bar" {
		return NewBarChart()
	}

	return nil
}
