package main

import (
    "fmt"
    "flag"
    "os/exec"
)

func main() {
  //Узнать: 1)какой режим будет использоваться
  //2) Какой порт (в обоих режимах)
  //3) Количество одновременных подключений (в режиме "Сервер")
  serverPort := flag.String("serverPort", "", "Порт в режиме 'Сервер'")
  clientPort := flag.String("clientPort", "", "Порт в режиме 'Клиент'")
  amountOfClients := flag.Int("n", 1, "Количество подключений")
  flag.Parse()
  //Начальный ключ
  initialKey := generateKey()

//Если режим "Сервер"
  if *serverPort != "" {
    fmt.Println("Режим 'Сервер'\n", "порт: ", *serverPort)

    cmd := exec.Command("start", "client.exe", "-port="+*serverPort, "-key="+initialKey)
    for i := 0; i < *amountOfClients; i++ {
      cmd = exec.Command("start", "client.exe", "-port="+*serverPort, "-key="+initialKey)
    cmd.Run()
    err := cmd.Run()
    if err != nil {
      fmt.Println("Command finished with error: ", err)
    }
  }

    serverWork(*serverPort, initialKey)
  }

//Если режим "Клиент"
  if *clientPort != "" {
    fmt.Println("Режим 'Клиент'\n", "порт: ", *clientPort)
    fmt.Println("Начальный ключ: ", initialKey)

    cmd := exec.Command("start", "server.exe", "-port="+*clientPort, "-key="+initialKey)
    cmd.Run()
    err := cmd.Run()
    if err != nil {
      fmt.Println("Command finished with error: ", err)
    }
    clientWork(*clientPort, initialKey)
  }
}
