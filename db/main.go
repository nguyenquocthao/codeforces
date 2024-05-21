package main

import (
	"fmt"
	"os"
	"strings"
)

type Entity struct {
	Entity    string   `json:"entity"`
	Name      string   `json:"name"`
	Questions []string `json:"questions"`
}

func main() {

	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	db := GetDb()
	records := []PubIntents{}
	db.Raw(`select * from intents where skillid='5855b77891fd47f09890b7480d14a72b'
		`).Find(&records)

	// PrintFileAndLineInfo(records)

	count := 0

	for _, v := range records {
		file.WriteString(fmt.Sprint(v.Skillid, " ", v.Intentid, " ", *NewJsonPsql(v.Data)))
		// fmt.Println(i, v.Name, v.Context)
		inputs := Convert[[]string](v.Data.Val["inputs"])
		changed := false
		for j, inp := range inputs {

			inputs[j] = strings.ReplaceAll(inp, "#you:c61687367fe4422c95bf7539ace11569", "@assistant:988c3818557d40a88e8925686e21f975")
			if inputs[j] != inp {
				changed = true
			}
			// fmt.Println("#", j, inp)
		}
		if changed {
			count += 1
			fmt.Println(v.Intentid)
		} else {
			continue
		}
		// continue
		v.Data.Val["inputs"] = inputs

		entities := Convert[[]Entity](v.Data.Val["entities"])
		found := false
		for _, e := range entities {
			if e.Entity == "988c3818557d40a88e8925686e21f975" {
				found = true
				break
			}
		}
		if !found {
			entities = append(entities, Entity{"988c3818557d40a88e8925686e21f975", "assistant", []string{}})
			v.Data.Val["entities"] = entities
		}
		fmt.Println(found, entities)

		db.Exec(`update intents set data=? where skillid=? and intentid=?`, v.Data, v.Skillid, v.Intentid)
		// fmt.Println(inputs)
	}
	fmt.Println(len(records), count)

}

// #you:c61687367fe4422c95bf7539ace11569
// @assistant:988c3818557d40a88e8925686e21f975
