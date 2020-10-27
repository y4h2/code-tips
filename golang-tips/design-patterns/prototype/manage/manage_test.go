package manage

func ExamplePrototypeManager() {

	pm := NewPrototypeManager()

	srsDoc := pm.GetOfficialDocument("srs")
	srsDoc.Display()

	farDoc := pm.GetOfficialDocument("far")
	farDoc.Display()

	// Output:
	// 《软件需求规格说明书》
	// 《可行性分析报告》
}
