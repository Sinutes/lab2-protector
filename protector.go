package main

import (
  "math/rand"
  "time"
  "crypto/md5"
  "encoding/hex"
  "strconv"
)

//Тип "Защитник соединения".
type protector struct {}

//Функция, порождающая новый экземпляр типа.
func NewProtector(key string) *protector  {
  p := new(protector)
  return p
}

//Функция, генерирующая новый ключ на основе предыдущего.
func (protector *protector) nextSessionKey(hashKey string) string  {
  hasher := md5.New()
  hasher.Write([]byte(hashKey))
  return hex.EncodeToString(hasher.Sum(nil))
}

//Функция генерирует случайную строку.
func generateKey() string  {
  rand.Seed(time.Now().UTC().UnixNano())
  result := ""
  for i := 0; i < 10; i++ {
    result += strconv.Itoa(rand.Intn(10))
  }
  return result
}
