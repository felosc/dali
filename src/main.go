package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

// Estructura incial de los datos email

type email struct {
	Message_ID                string `json:"email_id"`
	Date                      string `json:"email_date"`
	From                      string `json:"email_from"`
	To                        string `json:"email_to"`
	Subject                   string `json:"email_subject"`
	Mime_Version              string `json:"email_version"`
	Content_Type              string `json:"email_content_type"`
	Content_Transfer_Encoding string `json:"email_content_transfer_encodin"`
	X_From                    string `json:"email_x_from"`
	X_To                      string `json:"email_x_to"`
	X_cc                      string `json:"email_x_cc"`
	X_bcc                     string `json:"email_x_bcc"`
	X_Folder                  string `json:"email_x_folder"`
	X_Origin                  string `json:"email_x_origin"`
	X_FileName                string `json:"email_x_filename"`
	Message_Content           string `json:"email_content"`
}

func main() {
	//Strcut que almacena los datos
	countDir := 0
	var dataEmail email
	var dataIndex = `{"index":{"_index":"emails"}}`
	//ruta del documetno
	//`Downloads/enro_email_20110402/maildir/allen-p/inbox/1.txt`
	var pathEmailToWalk = "/home/vboxuser/Downloads/enron_mail_20110402/maildir/"
	//var pathToWalk = "/home/vboxuser/Downloads/enron_mail_20110402/maildir/"
	//var fileEmail = "/home/vboxuser/Downloads/enron_mail_20110402/maildir/allen-p/all_documents/10."
	//prueba de que encontro el archivo
	dirFiles := os.DirFS(pathEmailToWalk)

	fs.WalkDir(dirFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		println("-----------", pathEmailToWalk+"/"+path, "----")
		if err != nil {
			panic(err)
		}

		//comprueba si existe el archivo
		if d.Type().IsDir() {
			countDir++
			println(path, " es carpeta N°: ", countDir, "\n")

		} else {

			fileOpen := pathEmailToWalk + "/" + path
			file, err := os.Open(fileOpen)
			if err != nil {
				fmt.Println("&&: ", err)
			}
			//cierra el archivo
			defer file.Close()
			//variable array para guardar datos del archivo
			var data = []string{}
			//leemos el archivo
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				//convertimos bytes a texto
				scanText := scanner.Text()
				//guardamos datos del archivo en el array data
				data = append(data, scanText)
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}

			//caracter punto de inico o indice
			pointCaracter := ": "
			pointToSaveContent_1 := ".nsf"
			pointToSaveContent_2 := ".pst"
			var dataForStruct = []string{}
			var indexForStruct = []string{}
			indexForStruct = append(indexForStruct, "emails")

			//recorrer datos de archivo escaneados
			for i := 0; i < len(data); i++ {
				//busca si hay un ":" en el valor de ese campo
				findStrin := strings.Index(data[i], pointCaracter)
				//comprueba que no haya mas de 15 posiciones en el dataForStruct
				if len(dataForStruct)-1 < 15 {

					//si no esta ":" en ese campo seguir agregando datos a ese campo
					if findStrin == -1 {
						//impresion para saber la posicion en el archivo en caso de no encontrar ":"
						//fmt.Println("No se encontro ", pointCaracter, "//", i, "//", findStrin)
						//referncia de posicion para seguir añadiendo datos
						countAdd := len(dataForStruct) - 1
						//refenrecia de string ".pst y nsf para saber cuando termina la cabecera del correo"
						findStrinToSaveContent2 := strings.Index(strings.ToLower(dataForStruct[countAdd]), pointToSaveContent_2)
						findStrinToSaveContent1 := strings.Index(strings.ToLower(dataForStruct[countAdd]), pointToSaveContent_1)
						//comprueba si en dato guardado esta .pst o .nsf
						if findStrinToSaveContent1 != -1 || findStrinToSaveContent2 != -1 {
							//referencia de cuando llega al contenido
							//fmt.Println("se agregara el contenido del email")
							//numero en el que va el ciclo y en que posicion se encuentra dentro del dataFprStruct
							//fmt.Println(dataForStruct[countAdd], "ciclo #: ", i, "Posicion donde esta el dato: ", countAdd)
							//print V:
							//fmt.Println("Creacion de nueva pisicion para el contenido del correo")
							//despues de haber comprobado que el dato es .pst o .nsf
							//procede a guardar el inicio del contenido del correo en el dataForStruct
							//para crear la nueva posicion en dataForStruc
							dataForStruct = append(dataForStruct, data[i])
						}
						//se concatenan los datos en una ṕosicion ya esxistente
						//ya que no se encontraron ":" o "es el contenido del email"
						dataForStruct[countAdd] = dataForStruct[countAdd] + data[i]
						continue
					}

					//en caso de que si encontrara ":" separa los datos encontrados antes ":" y despues
					//si encuentra mas de 2 ":" separa desde el primero hacia adelante
					cutString := strings.SplitN(data[i], pointCaracter, 2)
					//impresion de que se encontro el puntero
					//fmt.Println("i= ", i, " se encontro el puntero en posicion ", findStrin)
					//muestra el dato capturado despues de ":"
					//fmt.Println("--", cutString, "-- ", len(cutString))
					//se guardan los datos encontrados despues de ":" en un array nuevo
					dataForStruct = append(dataForStruct, cutString[1])

				} else {
					//si dataForStruct llega a tener 15 posiciones
					//todo lo que encuentre el archivo despues de esa
					//posicion el texto se agrega al la posicion 15 en dataForStrict
					//es el Message_content basicamente
					dataForStruct[15] = dataForStruct[15] + data[i]
				}
			}
			if len(dataForStruct)-1 < 14 {
				fmt.Println(len(dataForStruct))
				for i, strcu := range dataForStruct {
					fmt.Println("--", i, "--", strcu)
				}
			}

			if len(dataForStruct)-1 < 15 {
				dataEmail.Message_ID = dataForStruct[0]
				dataEmail.Date = dataForStruct[1]
				dataEmail.From = dataForStruct[2]
				dataEmail.To = ""
				dataEmail.Subject = dataForStruct[3]
				dataEmail.Mime_Version = dataForStruct[4]
				dataEmail.Content_Type = dataForStruct[5]
				dataEmail.Content_Transfer_Encoding = dataForStruct[6]
				dataEmail.X_From = dataForStruct[7]
				dataEmail.X_To = dataForStruct[8]
				dataEmail.X_cc = dataForStruct[9]
				dataEmail.X_bcc = dataForStruct[10]
				dataEmail.X_Folder = dataForStruct[11]
				dataEmail.X_Origin = dataForStruct[12]
				dataEmail.X_FileName = dataForStruct[13]
				dataEmail.Message_Content = dataForStruct[14]
			} else {

				//se guardan los datos en el struct basdo en el formato de los documentos

				dataEmail.Message_ID = dataForStruct[0]
				dataEmail.Date = dataForStruct[1]
				dataEmail.From = dataForStruct[2]
				dataEmail.To = dataForStruct[3]
				dataEmail.Subject = dataForStruct[4]
				dataEmail.Mime_Version = dataForStruct[5]
				dataEmail.Content_Type = dataForStruct[6]
				dataEmail.Content_Transfer_Encoding = dataForStruct[7]
				dataEmail.X_From = dataForStruct[8]
				dataEmail.X_To = dataForStruct[9]
				dataEmail.X_cc = dataForStruct[10]
				dataEmail.X_bcc = dataForStruct[11]
				dataEmail.X_Folder = dataForStruct[12]
				dataEmail.X_Origin = dataForStruct[13]
				dataEmail.X_FileName = dataForStruct[14]
				dataEmail.Message_Content = dataForStruct[15]
			}
			//fmt.Println(dataEmail)
			//crear json
			convertJson(dataEmail, dataIndex)

		}
		return nil
	})

}

func convertJson(data email, dataindex string) {
	var pathJson string = "dataEmail.ndjson"
	showEmail, err := json.Marshal(data)

	if err != nil {
		println("error convirtiendo a json showEmail")
		panic(err)
	}
	dataIndex := dataindex

	if _, err := os.Stat(pathJson); err == nil {
		openTxt, err := ioutil.ReadFile(pathJson)
		if err != nil {
			fmt.Println("error leyendo el archivo json")
			panic(err)
		}

		openTxt = append(openTxt, "\n"+dataIndex+"\n"...)
		openTxt = append(openTxt, showEmail...)
		err = os.WriteFile(pathJson, openTxt, 0644)
		if err != nil {
			fmt.Println("error guardando nuevo contenido json")
			panic(err)
		}
		println("datos agregados a json con exito")

	} else {
		_, err := os.Create("dataEmail.ndjson")
		if err != nil {
			panic(err)
		}

		fmt.Println("nuevo Jso Creado")

	}

}
