package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	fmt.Println(data)
	jsonData, _ := json.Marshal(data)
	os.WriteFile("intents.json", jsonData, 0644)

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
