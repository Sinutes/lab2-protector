package main

import "flag"

//Сервер.
func main() {
  port := flag.String("port", "", "Порт в режиме 'Сервер'")
  key := flag.String("key", "", "Начальный ключ")
  flag.Parse()
  serverWork(*port, *key)
}
