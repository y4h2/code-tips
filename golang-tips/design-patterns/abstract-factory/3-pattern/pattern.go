package pattern

import "fmt"

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

// === Factory

type SkinFactory interface {
	GetButton() Button
	GetTextField() TextFiled
	GetComboBox() ComboBox
	GetDialog() Dialog
}

type SpringSkinFactory struct{}

func (SpringSkinFactory) GetButton() Button {
	return &SpringButton{}
}
func (SpringSkinFactory) GetTextField() TextFiled {
	return SpringTextField{}
}
func (SpringSkinFactory) GetComboBox() ComboBox {
	return SpringComboBox{}
}
func (SpringSkinFactory) GetDialog() Dialog {
	return SpringDialog{}
}

type SummerSkinFactory struct{}

func (SummerSkinFactory) GetButton() Button {
	return &SummerButton{}
}
func (SummerSkinFactory) GetTextField() TextFiled {
	return SummerTextField{}
}
func (SummerSkinFactory) GetComboBox() ComboBox {
	return SummerComboBox{}
}
func (SummerSkinFactory) GetDialog() Dialog {
	return SummerDialog{}
}
