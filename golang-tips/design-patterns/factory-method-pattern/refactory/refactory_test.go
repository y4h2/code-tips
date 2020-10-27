package refactory

func ExampleGetChart() {
	pieChart := (&PieFactory{}).GetChart()
	pieChart.Display()

	// Output:
	// create pie chart
	// display pie chart
}
