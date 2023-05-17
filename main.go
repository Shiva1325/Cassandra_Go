package main

import (
	controller "Cassandra_Go/Controller"
	test "Cassandra_Go/Test"
	"Cassandra_Go/pb"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CSVData struct {
	X int
	Y float64
}

func main() {
	log.Println("Please specify server address to connect: ")
	for {
		commandlinearg, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Println("Error while reading command. Please try again.")
			continue
		}
		commandlinearg = strings.Trim(commandlinearg, "\r\n")
		if !strings.Contains(commandlinearg, ":") {
			log.Println("Please provide correct server address")
			continue
		}
		if strings.TrimSpace(commandlinearg) == "" {
			log.Println("Please specify address to connect")
		} else {
			server_address := strings.Split(commandlinearg, ":")
			address := server_address[0]
			port := server_address[1]
			ConnectToServer(address, port)
			return
		}
	}

}

func ConnectToServer(server_address string, port_number string) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	connection, err := grpc.Dial(server_address+":"+port_number, transportOption)
	if err != nil {
		log.Fatal("cannot connect to the server", err)
	}
	log.Printf("Connection established to server %s:%v", server_address, port_number)
	clientservice := controller.NewClientCassandraService(pb.NewCassandraServiceClient(connection))

	now := time.Now()
	log.Printf("Execution Started: %v", now.Local())

	//Uncomment this section for testing synchronous testing ----------------------------------------------------------------

	setData(clientservice)
	getData(clientservice)

	//Uncomment this section until here for testing synchronous testing ------------------------------------------------------

	// Uncomment this section for asynchronous testing -----------------------------------------------------------------------

	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// 	for j := 0; j < 1; j++ {
	// 		user_data := test.GenerateData(j + 1)
	// 		for i := 0; i < len(user_data.User)/3; i++ {
	// 			user_details := &pb.User{
	// 				Id:      user_data.User[i].Id,
	// 				Name:    user_data.User[i].Name,
	// 				Country: user_data.User[i].Country,
	// 				Hashkey: user_data.User[i].Hashkey,
	// 			}
	// 			response, _ := controller.SaveData(user_details, clientservice)
	// 			log.Printf("Response : %v", response)
	// 		}
	// 	}
	// 	wg.Done()
	// }()
	// go func() {
	// 	for j := 0; j < 1; j++ {
	// 		user_data := test.GenerateData(j + 1)
	// 		for i := 0; i < len(user_data.User)/3; i++ {
	// 			country := &pb.GetUserData{
	// 				Key:         int64(user_data.User[i].Id),
	// 				Countryname: user_data.User[i].Country,
	// 			}
	// 			response, err := controller.GetData(country, clientservice)
	// 			if err != nil {
	// 				log.Printf("Error: %v", err)
	// 			}
	// 			PrintRecords(response, user_data.User[i].Country)
	// 		}
	// 	}
	// 	wg.Done()
	// }()
	// wg.Wait()

	// Uncomment this section until here for asynchronous testing ------------------------------------------------------------

	//Uncomment this for normal testing from here ------------------------------------------------------------------------

	// for {
	// 	log.Println("Enter a comand: ")
	// 	commandlinearg, err := bufio.NewReader(os.Stdin).ReadString('\n')
	// 	if err != nil {
	// 		log.Println("Error while reading command. Please try again.")
	// 		continue
	// 	}
	// 	commandlinearg = strings.Trim(commandlinearg, "\r\n")
	// 	switch commandlinearg {
	// 	case "a":
	// 		log.Println("Please provide user name")
	// 		username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// 		for {
	// 			if username == "" || strings.TrimSpace(username) == "" {
	// 				log.Println("Please provide correct user name")
	// 				username, err = bufio.NewReader(os.Stdin).ReadString('\n')
	// 				continue
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 		log.Println("Please provide user country")
	// 		usercountry, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// 		for {
	// 			if usercountry == "" || strings.TrimSpace(usercountry) == "" {
	// 				log.Println("Please provide correct user name")
	// 				usercountry, err = bufio.NewReader(os.Stdin).ReadString('\n')
	// 				continue
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 		if err != nil {
	// 			log.Printf("Error: %v", err)
	// 			continue
	// 		}
	// 		user_details := &pb.User{
	// 			Id:      int32(0),
	// 			Name:    username,
	// 			Country: usercountry,
	// 			Hashkey: int64(0),
	// 		}
	// 		response, err := controller.SaveData(user_details, clientservice)
	// 		if err != nil {
	// 			log.Printf("Error: %v", err)
	// 		}
	// 		if response.Status {
	// 			log.Printf("User %v saved successfully.", username)
	// 		} else {
	// 			log.Printf("Failed to save user data.")
	// 		}
	// 	case "g":
	// 		log.Println("Please provide country name to get all records")
	// 		country_name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	// 		for {
	// 			if country_name == "" || strings.TrimSpace(country_name) == "" {
	// 				log.Println("Please provide correct country name")
	// 				country_name, err = bufio.NewReader(os.Stdin).ReadString('\n')
	// 				continue
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 		if err != nil {
	// 			log.Printf("Error: %v", err)
	// 		}
	// 		country := &pb.GetUserData{
	// 			Key:         0,
	// 			Countryname: country_name,
	// 		}
	// 		response, err := controller.GetData(country, clientservice)
	// 		if err != nil {
	// 			log.Printf("Error: %v", err)
	// 		}
	// 		PrintRecords(response, country_name)
	// 	default:
	// 		log.Println("Please enter correct command")
	// 	}
	// }

	// Uncomment until here for normal testing ---------------------------------------------------------------------------

	log.Printf("Time Elapsed for total execution: %v", time.Since(now))
}

// Function that loads data from test file and writes those records in to disk and in-memory
func setData(clientservice *controller.CassandraClientService) {
	log.Println("---------------------------------Set Data----------------------------------")
	csvFile, _ := os.Create("SetData.csv")

	csvWriter := csv.NewWriter(csvFile)
	for j := 0; j < 1; j++ {
		user_data := test.GenerateData(j + 1)
		for i := 0; i < len(user_data.User); i++ {
			time_data := CSVData{}
			time_data.X = i
			now1 := float64(time.Now().UnixNano()) / float64(time.Millisecond)
			user_details := &pb.User{
				Id:      user_data.User[i].Id,
				Name:    user_data.User[i].Name,
				Country: user_data.User[i].Country,
				Hashkey: user_data.User[i].Hashkey,
			}
			response, _ := controller.SaveData(user_details, clientservice)
			end := float64(time.Now().UnixNano()) / float64(time.Millisecond)
			time_data.Y = float64(end - now1)
			err := csvWriter.Write([]string{string(time_data.X), fmt.Sprintf("%f", time_data.Y)})
			if err != nil {
				panic(err)
			}
			//log.Printf("Time elapsed for record-%v is: %v", j+1, time.Since(now1))
			if response.Status {

			}
		}
	}
	csvWriter.Flush()
	csvFile.Close()
}

// Function that gets data from test file and queries those records in to in-memory and displays all results
func getData(clientservice *controller.CassandraClientService) {
	log.Println("---------------------------------Get Data----------------------------------")
	csvFile, _ := os.Create("GetData.csv")

	csvWriter := csv.NewWriter(csvFile)
	for j := 0; j < 1; j++ {
		//log.Printf("Record-%v started at: %v", j+1, now1.Local())
		user_data := test.GenerateData(j + 1)
		for i := 0; i < len(user_data.User); i++ {
			time_data := CSVData{}
			time_data.X = i
			now1 := float64(time.Now().UnixNano()) / float64(time.Millisecond)
			country := &pb.GetUserData{
				Key:         int64(user_data.User[i].Id),
				Countryname: user_data.User[i].Country,
			}
			response, err := controller.GetData(country, clientservice)
			if err != nil {
				log.Printf("Error: %v", err)
			}
			PrintRecords(response, user_data.User[i].Country)
			end := float64(time.Now().UnixNano()) / float64(time.Millisecond)
			time_data.Y = float64(end - now1)
			csvWriter.Write([]string{string(time_data.X), fmt.Sprintf("%f", time_data.Y)})
		}
		//log.Printf("Time elapsed for record-%v is: %v", j+1, time.Since(now1))
	}
	csvWriter.Flush()
	csvFile.Close()
}

// Printing the get records
func PrintRecords(response *pb.SavedRecords, country string) {
	if len(response.Records) == 0 {
		log.Printf("No records found for %v", country)
	} else {
		log.Printf("Records for %v:", country)
		for _, value := range response.Records {
			log.Printf("User Name: %v", value.Name)
		}
	}
}
