package delhash

import (
    "fmt"
    "go/project-3/internal/payload"
    "go/project-3/internal/repository"
)

func DeleteHash(repo repository.Repo, fileData *[]payload.DataRequest, hash string) (bool, error) {
  found := false
  for i, h := range *fileData {
  if hash == h.Hash {
    found = true
    *fileData = append((*fileData)[:i], (*fileData)[i+1:]...)
    break
  }
  }
   if found == true {
   fmt.Println("Подтверждение прошло Успешно")
   err := repo.WriteFile(fileData)
   if err != nil {
      return found, err
      }
   }
   return found, nil
}