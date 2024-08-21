package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type Student struct {
	Name     string
	LastName string
	Email    string
	Id       int32
}

// func PorcentagemPresenca() float64 {
// 	totalAulas++
// 	if presente {
// 		Presenca++
// 	}
// }

func enviarEmail(destinatario, assunto, corpo string) {
	from := "rebeca.gazola@sap.com"
	pass := "gg"

	msg := "From: " + from + "\n" +
		"To :" + destinatario + "\n" +
		"Subject: " + assunto + "\n\n" +
		corpo

	smtpHost := "smtp.office365.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{destinatario}, []byte(msg))
	if err != nil {
		log.Fatalf("Erro ao enviar e-mail: %v", err)

	}
	fmt.Println("E-mail enviado com sucesso!")

}

// func verificarPrsenca(alunas []Aluna) {
// 	for _, aluna := range alunas {
// 		if aluna.PorcentagemPresenca() < 75 {
// 			enviarEmail(aluna.Email,
// 				"Alerta de Presença",
// 				fmt.Sprintf("Olá %s, sua presença no curso é de %.2f%%, que está abaixo do mínimo exigido de 75%%.", aluna.Nome, aluna.PorcentagemPresenca()))
// 		}
// 	}
// }

// func attendance() {
// 	alunas := []Aluna{
// 		{"Ana", "ana@example.com", 0, 0},
// 		{"Beatriz", "beatriz@example.com", 0, 0},
// 	}

// 	// Simulação de presença em 10 aulas
// 	for i := 0; i < 10; i++ {
// 		alunas[0].MarcarPresenca(true)  // Ana presente
// 		alunas[1].MarcarPresenca(false) // Beatriz ausente
// 	}

// 	verificarPresenca(alunas)

// 	for _, aluna := range alunas {
// 		fmt.Printf("Nome: %s, E-mail: %s, Presenças: %d/%d\n",
// 			aluna.Nome, aluna.Email, aluna.Presenca, aluna.TotalAulas)
// 	}
// }

func CreateStudents() error {
	// name := stud.Name
	// lastName := stud.LastName
	// email := stud.Email
	// id := stud.Id

	importStudentsFromCSV()
	return nil

}

func importStudentsFromCSV() {
	file, err := os.Open("participants.csv")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1000)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
