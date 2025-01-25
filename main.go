package main

import (
	"UserSystem/database"
	"UserSystem/models"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"sync"
)

func generateUsers(numUsers int) []models.User {
	gofakeit.Seed(0)
	users := make([]models.User, numUsers)

	for i := 0; i < numUsers; i++ {
		numAddresses := rand.Intn(5) + 1
		addresses := make([]models.Address, numAddresses)

		for j := 0; j < numAddresses; j++ {
			addresses[j] = models.Address{
				Street:  gofakeit.Street(),
				City:    gofakeit.City(),
				State:   gofakeit.State(),
				ZipCode: gofakeit.Zip(),
				Country: gofakeit.Country(),
			}
		}

		users[i] = models.User{
			ID:          uuid.New(),
			Name:        gofakeit.Name(),
			Email:       gofakeit.Email(),
			PhoneNumber: gofakeit.Phone(),
			Addresses:   addresses,
		}
	}

	return users
}

func saveToJSON(data []models.User, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("error closing file: %v", closeErr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

func readFromJson(fileName string) ([]models.User, error) {
	var users []models.User
	jsonData, _ := os.ReadFile(fileName)
	err := json.Unmarshal(jsonData, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func worker(db *gorm.DB, userChan <-chan models.User, wg *sync.WaitGroup, semaphore chan struct{}) {
	for user := range userChan {
		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("Error inserting user %s: %v\n", user.Name, err)
		}
		<-semaphore
		wg.Done()
	}
}

func main() {
	const numUsers = 1_000_000
	const fileName = "users_data.json"
	const maxWorkers = 10

	fmt.Printf("Generating %d users with multiple addresses...\n", numUsers)
	usersData := generateUsers(numUsers)
	fmt.Printf("Saving data to %s...\n", fileName)
	err := saveToJSON(usersData, fileName)
	if err != nil {
		fmt.Printf("Error saving data: %v\n", err)
		return
	}
	fmt.Println("Data generation and saving completed.")

	db := database.ConnectDB()
	database.Migrate(db)

	users, err := readFromJson(fileName)
	if err != nil {
		fmt.Printf("Error reading data from json: %v\n", err)
		return
	}
	fmt.Printf("Total users to process: %d\n", len(users))

	usersChannel := make(chan models.User, numUsers)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		go worker(db, usersChannel, &wg, semaphore)
	}

	for _, user := range users {
		semaphore <- struct{}{}
		wg.Add(1)
		usersChannel <- user
	}

	close(usersChannel)
	wg.Wait()
	fmt.Println("All users processed successfully.")

	Serve(db)
}
