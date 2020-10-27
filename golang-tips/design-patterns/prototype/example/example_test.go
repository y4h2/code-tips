package example

import "fmt"

func ExampleWeeklyLog() {
	previousLog := &WeeklyLog{
		name:    "张无忌",
		date:    "第12周",
		content: "这周工作很忙，每天加班！",
	}

	fmt.Println("****周报****")
	fmt.Println("周次：" + previousLog.GetDate())
	fmt.Println("姓名：" + previousLog.GetName())
	fmt.Println("内容：" + previousLog.GetContent())
	fmt.Println("--------------------------------")

	newLog := previousLog.Clone()
	newLog.SetDate("第13周")

	fmt.Println("****周报****")
	fmt.Println("周次：" + newLog.GetDate())
	fmt.Println("姓名：" + newLog.GetName())
	fmt.Println("内容：" + newLog.GetContent())
	fmt.Println("--------------------------------")

	// Output:
	// ****周报****
	// 周次：第12周
	// 姓名：张无忌
	// 内容：这周工作很忙，每天加班！
	// --------------------------------
	// ****周报****
	// 周次：第13周
	// 姓名：张无忌
	// 内容：这周工作很忙，每天加班！
	// --------------------------------
}

func ExampleWeeklyLog1() {
	previousLog := &WeeklyLog{
		name:    "张无忌",
		date:    "第12周",
		content: "这周工作很忙，每天加班！",
	}

	fmt.Println("****周报****")
	fmt.Println("周次：" + previousLog.GetDate())
	fmt.Println("姓名：" + previousLog.GetName())
	fmt.Println("内容：" + previousLog.GetContent())
	fmt.Println("--------------------------------")

	// newLog := previousLog.Clone()
	newLog := *previousLog
	newLog.SetDate("第13周")

	fmt.Println("****周报****")
	fmt.Println("周次：" + newLog.GetDate())
	fmt.Println("姓名：" + newLog.GetName())
	fmt.Println("内容：" + newLog.GetContent())
	fmt.Println("--------------------------------")

	fmt.Println("****周报****")
	fmt.Println("周次：" + previousLog.GetDate())
	fmt.Println("姓名：" + previousLog.GetName())
	fmt.Println("内容：" + previousLog.GetContent())
	fmt.Println("--------------------------------")

	// Output:
	// ****周报****
	// 周次：第12周
	// 姓名：张无忌
	// 内容：这周工作很忙，每天加班！
	// --------------------------------
	// ****周报****
	// 周次：第13周
	// 姓名：张无忌
	// 内容：这周工作很忙，每天加班！
	// --------------------------------
	// ****周报****
	// 周次：第12周
	// 姓名：张无忌
	// 内容：这周工作很忙，每天加班！
	// --------------------------------
}
