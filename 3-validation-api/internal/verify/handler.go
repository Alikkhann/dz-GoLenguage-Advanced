package verify

import (
    "strings"
    "net/http"
    "go/project-3/internal/emailer"
    "go/project-3/internal/payload"
    "go/project-3/internal/pkg/req"
    "go/project-3/internal/pkg/resp"
    "go/project-3/internal/repository"
    "go/project-3/internal/createHash"
    "go/project-3/internal/delHash"
)



type VerifyHandler struct {
    Email string
    Repo repository.Repo
}

func (p *VerifyHandler) PostHandl(w http.ResponseWriter, r *http.Request) {
    //декодинг и валидация
    req, err := req.HandleBody[payload.LoginRequest](w, r)
    if err != nil {
    resp.Json(w, err.Error(), 400)
        return 
    }
    // создание хеша
    hash, err := createHash.CreateHash()
    if err != nil {
        resp.Json(w, err.Error(), 400)
        return
    }
    //создание файла и запись почты и хеша
    JsonStr := payload.DataRequest {
        Email: req.Email,
        Hash: hash,
    }
    //чтение
    p.Repo.Add(JsonStr)
    //создание конфигурации email для отправки письма
    err = emailer.MailerData(JsonStr)
    if err != nil {
      resp.Json(w, err.Error(), 400)
      return 
    }
}


func (g *VerifyHandler) GetHandl(w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/verify/") // "{hash}"
	fileData, err := g.Repo.LoadFile()
	if err != nil {
		resp.Json(w, err , 400)
		return 
	}
  found, err := delhash.DeleteHash(g.Repo, &fileData, hash)
	if err != nil {
		resp.Json(w, err , 400)
		return
	}
	resp.Json(w, found, 200)
	}





















// package verify

// import (
// 	"encoding/json"
// 	"fmt"
// 	"go/project-3/configs"
// 	"go/project-3/internal/payload"
// 	"go/project-3/internal/pkg"
// 	"go/project-3/internal/pkg/resp"
// 	"net/http"
// 	"net/smtp"
// 	"os"
// 	"strings"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/jordan-wright/email"
// )

// type DataRequest struct {
// 	Email string `json:"email"`
// 	Hash string  `json:"hash"`
// }

// type VerifyHandler struct {
// 	Email string
// }

// func (p *VerifyHandler) PostHandl(w http.ResponseWriter, r *http.Request) {
// 	//декодинг и валидация
// 	req := payload.LoginRequest{}
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		resp.Json(w, err.Error(), 400)
// 		return
// 	}
// 	validate := validator.New()
// 	err = validate.Struct(req)
// 	if err != nil {
// 		resp.Json(w, err.Error(), 400)
// 		return
// 	}
// 	resp.Json(w, req, 200)
	
// 	// создание хеша
// 	hash, err := pkg.CreateHash()
// 	if err != nil {
// 		resp.Json(w, err.Error(), 400)
// 		return
// 	}
	
// 	//создание файла и запись почты и хеша
// 	JsonStr := DataRequest {
// 		Email: req.Email,
// 		Hash: hash,
// 	}

// 	//чтение
// 	fileJson, err := os.ReadFile("dataRequest.json")
// 	if err != nil {
// 		data, _ := json.Marshal(JsonStr)
// 		os.WriteFile("dataRequest.json", data , 0644)
// 	}else{
// 	//перевод прочитанных данных в слайс структур
// 	dataStruct := []DataRequest{}
// 	err = json.Unmarshal(fileJson, &dataStruct)
// 	if err != nil {
// 		fmt.Println("Не удалось перевести данные в Json")
// 	}
//   //добавление к прочитанному новые данные
// 	dataStruct = append(dataStruct, JsonStr)
// 	byteJson, err := json.Marshal(dataStruct)
// 	err = os.WriteFile("dataRequest.json", byteJson, 0644)
// 	if err != nil {
// 		fmt.Println("Не удалось записать данные в Json файл")
// 	}
// }

// 	//создание конфигурации email для отправки письма
// 	smtpConfig := configs.NewConfig()
// 	e := email.NewEmail()
// 	e.From = smtpConfig.Email
// 	e.To = []string{JsonStr.Email}
// 	e.Subject = "Подтверждение почты"
// 	e.Text = []byte("Для подтверждения почты перейдите по ссылке http://localhost:8081/verify/" + JsonStr.Hash )
	
// 	addr := smtpConfig.Address
// 	host := strings.Split(smtpConfig.Address, ":")[0]
// 	err = e.Send(addr, smtp.PlainAuth("", smtpConfig.Email, smtpConfig.Password, host))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
	
// }




// func (g *VerifyHandler) GetHandl(w http.ResponseWriter, r *http.Request) {
// 	hash := strings.TrimPrefix(r.URL.Path, "/verify/") // "{hash}"

// 	file, err := os.ReadFile("dataRequest.json")
// 	if err != nil {
// 		resp.Json(w, err, 400)
// 	}

// 	dataJson, err := repository.Json(file)
// 	if err != nil {
// 		resp.Json(w, err, 400)
// 	}

// 	found := false
// 	for i, h := range dataJson {
// 		if hash == h.Hash {
// 			found = true
// 			dataJson = append(dataJson[:i], dataJson[i+1:]...)
// 			break
// 		}
// 	}
// 	if found == true {
// 		fmt.Println("Подтверждение прошло Успешно")
// 		dataFile, err := json.Marshal(dataJson)
// 		os.WriteFile("dataRequest.json", dataFile, 0644)
// 		if err != nil {
// 			resp.Json(w, err , 500)
// 		}
// 	}

// 	resp.Json(w, found, 200)
// }







// func (g *VerifyHandler) GetHandl(w http.ResponseWriter, r *http.Request) {
// 	hash := strings.TrimPrefix(r.URL.Path, "/verify/") // "{hash}"
// 	file, err := os.ReadFile("dataRequest.json")
// 	if err != nil {
// 		resp.Json(w, err, 400)
// 	}
// 	data := []DataRequest{}
	
// 	err = json.Unmarshal(file, &data)
// 	if err != nil {
// 		resp.Json(w, err, 400)
// 	}
// 	found := false
// 	for i, h := range data {
// 		if hash == h.Hash {
// 			found = true
// 			data = append(data[:i], data[i+1:]...)
// 			break
// 		}
// 	}
// 	if found == true {
// 		fmt.Println("Подтверждение прошло Успешно")
// 		dataFile, err := json.Marshal(data)
// 		os.WriteFile("dataRequest", dataFile, 0644)
// 		if err != nil {
// 			resp.Json(w, err , 500)
// 		}
// 	}

// 	resp.Json(w, found, 200)
// }