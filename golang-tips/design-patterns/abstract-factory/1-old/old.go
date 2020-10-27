package old

import (
	"fmt"
)

// ==== button

type Button interface {
	Display()
}

type SpringButton struct{}

func (SpringButton) Display() {
	fmt.Println("显示浅绿色按钮。")
}

type SummerButton struct{}

func (SummerButton) Display() {
	fmt.Println("显示浅蓝色按钮。")
}

// func GetButton(buttonType string) (Button, error) {
// 	if buttonType == "spring" {
// 		return &SpringButton{}, nil
// 	} else if buttonType == "summer" {
// 		return &SummerButton{}, nil
// 	} else {
// 		return nil, errors.New("unknown type")
// 	}
// }

type ButtonFactory interface {
	GetButton() Button
}

type SpringButtonFactory struct{}

func (spring *SpringButtonFactory) GetButton() Button {
	return &SpringButton{}
}

type SummerButtonFactory struct {
}

func (summer *SummerButtonFactory) GetButton() Button {
	return &SpringButton{}
}

// ===== TextField
type TextFiled interface {
	Display()
}

type SpringTextField struct{}

func (SpringTextField) Display() {
	fmt.Println("显示绿色边框文本框。")
}

type SummerTextField struct{}

func (SummerTextField) Display() {
	fmt.Println("显示蓝色边框文本框。")
}

type TextFieldFactory interface {
	GetTextField() TextFiled
}

type SpringTextFieldFactory struct{}

func (SpringTextFieldFactory) GetTextField() TextFiled {
	return &SpringTextField{}
}

type SummerTextFieldFactory struct{}

func (SummerTextFieldFactory) GetTextField() TextFiled {
	return &SummerTextField{}
}

// === ComboBox

type ComboBox interface {
	Display()
}

type SpringComboBox struct{}

func (SpringComboBox) Display() {
	fmt.Printf("显示绿色边框组合框。")
}

type SummerComboBox struct{}

func (SummerComboBox) Display() {
	fmt.Println("显示蓝色边框组合框。")
}

type ComboBoxFactory interface {
	GetComboBox() ComboBox
}

type SpringComboBoxFactory struct{}

func (SpringComboBoxFactory) GetComboBox() TextFiled {
	return &SpringComboBox{}
}

type SummerComboBoxFactory struct{}

func (SummerComboBoxFactory) GetComboBox() TextFiled {
	return &SummerComboBox{}
}

// === Dialog

type Dialog interface {
	Display()
}

type SpringDialog struct{}

func (SpringDialog) Display() {
	fmt.Printf("显示绿色边框会话框。")
}

type SummerDialog struct{}

func (SummerDialog) Display() {
	fmt.Printf("显示蓝色边框会话框。")
}

type DialogFactory interface {
	GetDialog() Dialog
}

type SpringDialogFactory struct{}

func (SpringDialogFactory) GetDialog() Dialog {
	return &SpringDialog{}
}

type SummerDialogFactory struct{}

func (SummerDialogFactory) GetDialog() Dialog {
	return &SummerDialog{}
}
