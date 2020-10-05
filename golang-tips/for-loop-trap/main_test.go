package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Entity struct {
	Name string
}

func ToPointerSlice1(ents []Entity) []*Entity {
	result := []*Entity{}
	var ent Entity
	for _, ent = range ents {
		fmt.Printf("%p\n", &ent)
		result = append(result, &ent)
	}
	return result
}

func TestToPointerSlice1(t *testing.T) {
	assert := assert.New(t)
	ents := []Entity{{Name: "a"}, {Name: "b"}, {Name: "c"}}

	rtnEnts := ToPointerSlice1(ents)
	rtnNames := []string{}
	for _, ent := range rtnEnts {
		rtnNames = append(rtnNames, ent.Name)
	}
	assert.Equal([]string{"c", "c", "c"}, rtnNames)
}

func ToPointerSlice2(ents []Entity) []*Entity {
	result := []*Entity{}
	for _, ent := range ents {
		tempEnt := ent
		result = append(result, &tempEnt)
	}
	return result
}

func TestToPointerSlice2(t *testing.T) {
	assert := assert.New(t)
	ents := []Entity{{Name: "a"}, {Name: "b"}, {Name: "c"}}

	rtnEnts := ToPointerSlice2(ents)
	rtnNames := []string{}
	for _, ent := range rtnEnts {
		rtnNames = append(rtnNames, ent.Name)
	}
	assert.Equal([]string{"a", "b", "c"}, rtnNames)
}

func ToPointerSlice3(ents []Entity) []*Entity {
	result := []*Entity{}
	for i := range ents {
		result = append(result, &ents[i])
	}
	return result
}

func TestToPointerSlice3(t *testing.T) {
	assert := assert.New(t)
	ents := []Entity{{Name: "a"}, {Name: "b"}, {Name: "c"}}

	rtnEnts := ToPointerSlice3(ents)
	rtnNames := []string{}
	for _, ent := range rtnEnts {
		rtnNames = append(rtnNames, ent.Name)
	}
	assert.Equal([]string{"a", "b", "c"}, rtnNames)
}

func TestA(t *testing.T) {
	ents := []Entity{{Name: "a"}, {Name: "b"}, {Name: "c"}}

	result := []*Entity{}

	{
		var value Entity
		for_temp := ents
		len_temp := len(for_temp)
		for index_temp := 0; index_temp < len_temp; index_temp++ {
			value_temp := for_temp[index_temp]
			// index = index_temp
			value = value_temp
			// origin body

			result = append(result, &value)

			t.Logf("%p", &value)
		}

	}

}
