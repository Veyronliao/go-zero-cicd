package main

import models "Bolog/User/Model"

func main() {
	models.NewDB("root:root@tcp(192.168.211.151:3306)/Bolog_User?charset=utf8mb4&parseTime=True&loc=Local")
}
