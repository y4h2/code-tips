package manage

import "fmt"

type Cloneable interface {
}

type OfficialDocument interface {
	Cloneable

	Clone() OfficialDocument
	Display()
}

type FAR struct {
}

func (far *FAR) Clone() OfficialDocument {
	return &FAR{}
}

func (far *FAR) Display() {
	fmt.Println("《可行性分析报告》")
}

func NewFAR() *FAR {
	return &FAR{}
}

type SRS struct {
}

func (srs *SRS) Clone() OfficialDocument {
	return &SRS{}
}

func (srs *SRS) Display() {
	fmt.Println("《软件需求规格说明书》")
}

func NewSRS() *SRS {
	return &SRS{}
}

type PrototypeManager struct {
	hashTable map[string]OfficialDocument
}

func (pm *PrototypeManager) GetOfficialDocument(key string) OfficialDocument {
	return pm.hashTable[key].Clone()
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		hashTable: map[string]OfficialDocument{
			"far": NewFAR(),
			"srs": NewSRS(),
		},
	}
}
