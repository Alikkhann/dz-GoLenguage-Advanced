package repository

import (
    "encoding/json"
    "go/project-3/internal/payload"
    "os"
)

type Repo interface {
	Add(jsonStr payload.DataRequest) error
	LoadFile() ([]payload.DataRequest, error)
	WriteFile(data *[]payload.DataRequest) error
}

type VerifyRepository struct {
  FilePath string
}


func (v *VerifyRepository) Add(jsonStr payload.DataRequest) error {
  //чтение
  fileJson, err := v.LoadFile()
  if err != nil {
    fileJson = []payload.DataRequest{} // если данных нет то просто создаем новый пустой список
  }
  //добавление к прочитанному новые данные
  fileJson = append(fileJson, jsonStr)
  err = v.WriteFile(&fileJson) 
  if err != nil {
    return err
  }
  return nil
  }

func (l *VerifyRepository) LoadFile() ([]payload.DataRequest, error) {
  fileData, err := os.ReadFile(l.FilePath)
  if err != nil {
    return nil, err
  }
  return jsonDecode(fileData)  //читаю файл и сразу перевожу в стуктуру и возращаю ответ от Json c ошибкой
}

func jsonDecode(file []byte) ([]payload.DataRequest, error) {
  dataStruct := []payload.DataRequest{}
  err := json.Unmarshal(file, &dataStruct)
  if err != nil {
  	return dataStruct, err
  }
  return dataStruct, nil
}

func (w *VerifyRepository) WriteFile(data *[]payload.DataRequest) error {
  dataFile, err := json.Marshal(data)
  if err != nil {
    return err
  }
  err = os.WriteFile(w.FilePath, dataFile, 0644)
  return err
}













// func (v *VerifyRepository) ReadandWriteFile(jsonStr payload.DataRequest, hash string) error {
// 		//чтение
// 	fileJson, err := os.ReadFile(v.FilePath)
// 	if err != nil {
// 		data, _ := json.Marshal(jsonStr)
// 		os.WriteFile(v.FilePath, data , 0644)
// 	}else{
// 		//перевод прочитанных данных в слайс структур
// 		jsonStruct, err := Json(fileJson)
// 		if err != nil {
// 			fmt.Println("Не удалось перевести  json в структуру")
// 		}
// 		//добавление к прочитанному новые данные
// 		jsonStruct = append(jsonStruct, jsonStr)
// 		byteJson, err := json.Marshal(jsonStruct)
// 		err = os.WriteFile(v.FilePath, byteJson, 0644)
// 		if err != nil {
// 			fmt.Println("Не удалось записать данные в Json файл")
// 			return err
// 		}
// 	}
// 	return nil

