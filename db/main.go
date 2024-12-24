package main

import (
	"encoding/json"
	"fmt"
)

// normal table: development
// pub: pub_*
// production: prod_*

func main() {

	// publicskillid := "5855b77891fd47f09890b7480d14a72b"
	// fptskillid := "8ade0166715f495fad52aec4032b3a7f"
	// ralliskillid := "53407b3bfdd4440291f7f97e4676d78b"

	// file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	db := GetDb()
	records := []Intents{}
	db.Raw(`select * from intents where skillid='53407b3bfdd4440291f7f97e4676d78b'
		`).Find(&records)
	// PrintFileAndLineInfo(records)
	count := 0
	// return

	data := []map[string]any{}
	fmt.Println(29, len(records))

	for _, v := range records {
		// v.Skillid = "8ade0166715f495fad52aec4032b3a7f"
		// db.Model(&v).Create(&v)
		dt := Data{}
		json.Unmarshal([]byte(*v.Data), &dt)

		data = append(data, map[string]any{
			"id":      v.Intentid,
			"name":    v.Name,
			"actions": *NewJsonPsql(dt.Actions),
		})
	}

	WriteJson(data, "x.json")
	return

	// for _, v := range records {
	// 	// file.WriteString(fmt.Sprint(v.Skillid, " ", v.Intentid, " ", *NewJsonPsql(v.Data)))
	// 	// fmt.Println(i, v.Name, v.Context)
	// 	// fmt.Println(v.Data.Val)
	// 	actions := v.Data.Val.Actions
	// 	texts := []string{}
	// 	for _, ac := range actions {
	// 		for _, t := range ac.Texts {
	// 			texts = append(texts, *NewJsonPsql(t))
	// 		}
	// 		// texts = append(texts, ac.Texts...)
	// 	}

	// 	hasmaika, haspersona := []string{}, false

	// 	for _, t := range texts {
	// 		txt := strings.ToLower(t)
	// 		if strings.Contains(txt, "persona") && !haspersona {
	// 			haspersona = true
	// 			hasmaika = []string{}
	// 		}
	// 		if haspersona && !strings.Contains(txt, "persona") {
	// 			continue
	// 		}
	// 		if strings.Contains(txt, "maika") || strings.Contains(txt, "olli") {
	// 			hasmaika = append(hasmaika, txt)
	// 		}
	// 	}

	// 	inputhasmaika := []string{}
	// 	for _, t := range v.Data.Val.Inputs {
	// 		txt := strings.ToLower(t)
	// 		if strings.Contains(txt, "maika") || strings.Contains(txt, "olli") {
	// 			inputhasmaika = append(inputhasmaika, txt)
	// 		}
	// 	}

	// 	if len(hasmaika) > 0 || len(inputhasmaika) > 0 {
	// 		// fmt.Println(v.Intentid, "\t", hasmaika, "\t", inputhasmaika)
	// 		fmt.Println(v.Intentid)
	// 	}

	// }
	fmt.Println(len(records), count)

}

// #you:c61687367fe4422c95bf7539ace11569
// @assistant:988c3818557d40a88e8925686e21f975

// https://skills.iviet.com/skills/5855b77891fd47f09890b7480d14a72b/intents/dcc5a4b347f44b6e9f7640c7c4180c74

// func main0() {

// 	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

// 	db := GetDb()
// 	records := []Intents{}
// 	db.Raw(`select * from intents where skillid='5855b77891fd47f09890b7480d14a72b'
// 		`).Find(&records)

// 	// PrintFileAndLineInfo(records)

// 	count := 0

// 	for _, v := range records {
// 		file.WriteString(fmt.Sprint(v.Skillid, " ", v.Intentid, " ", *NewJsonPsql(v.Data)))
// 		// fmt.Println(i, v.Name, v.Context)
// 		// inputs := Convert[[]string](v.Data.Val["inputs"])
// 		inputs := v.Data.Val.Inputs
// 		changed := false
// 		for j, inp := range inputs {

// 			inputs[j] = strings.ReplaceAll(inp, "maika", "@assistant:988c3818557d40a88e8925686e21f975")
// 			if inputs[j] != inp {
// 				changed = true
// 			}
// 			// fmt.Println("#", j, inp)
// 		}
// 		if changed {
// 			count += 1
// 			fmt.Println(v)
// 		} else {
// 			continue
// 		}
// 		// continue
// 		v.Data.Val.Inputs = inputs

// 		// entities := Convert[[]Entity](v.Data.Val["entities"])
// 		entities := v.Data.Val.Entities
// 		found := false
// 		for _, e := range entities {
// 			if e.Entity == "988c3818557d40a88e8925686e21f975" {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			entities = append(entities, Entity{"988c3818557d40a88e8925686e21f975", "assistant", []string{}})
// 			v.Data.Val.Entities = entities
// 		}
// 		fmt.Println(found, entities)

// 		// db.Exec(`update intents set data=? where skillid=? and intentid=?`, v.Data, v.Skillid, v.Intentid)
// 		// fmt.Println(inputs)
// 	}
// 	fmt.Println(len(records), count)

// }
