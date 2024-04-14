package bot

//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"tgBot/entities"
//)
//
//func (b *Bot) insertInfo() {
//	dataFile, err := ioutil.ReadFile("data.json")
//	if err != nil {
//		fmt.Println("Ошибка открытия файла:", err)
//		return
//	}
//
//	var data []entities.Course
//	err = json.Unmarshal(dataFile, &data)
//	if err != nil {
//		fmt.Println("Ошибка парсинга JSON:", err)
//		return
//	}
//	for _, obj := range data {
//		b.db.Create(&obj)
//	}
//}
