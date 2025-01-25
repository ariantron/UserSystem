package main

import (
	"UserSystem/database"
	"UserSystem/internal/models"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
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

func readFromJsonStream(fileName string) ([]models.User, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []models.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func worker(db *gorm.DB, userChan <-chan []models.User, wg *sync.WaitGroup, semaphore chan struct{}) {
	for userBatch := range userChan {
		err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&userBatch).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Printf("Error inserting batch: %v\n", err)
		}
		<-semaphore
		wg.Done()
	}
}

func setupDatabaseWorkers(db *gorm.DB, users []models.User, maxWorkers int, batchSize int) {
	userBatchChannel := make(chan []models.User, len(users)/batchSize+1)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		go worker(db, userBatchChannel, &wg, semaphore)
	}

	for i := 0; i < len(users); i += batchSize {
		end := i + batchSize
		if end > len(users) {
			end = len(users)
		}

		semaphore <- struct{}{}
		wg.Add(1)
		userBatchChannel <- users[i:end]
	}

	close(userBatchChannel)
	wg.Wait()
}

func generateAndSaveUserData(numUsers int, fileName string) error {
	fmt.Printf("Generating %d users with multiple addresses...\n", numUsers)
	usersData := generateUsers(numUsers)

	fmt.Printf("Saving data to %s...\n", fileName)
	if err := saveToJSON(usersData, fileName); err != nil {
		return fmt.Errorf("error saving data: %w", err)
	}

	fmt.Println("Data generation and saving completed....")
	return nil
}

func processUserData(fileName string, db *gorm.DB, maxWorkers int, batchSize int) error {
	users, err := readFromJsonStream(fileName)
	if err != nil {
		return fmt.Errorf("error reading data from json: %v", err)
	}

	fmt.Printf("Total users to process: %d\n", len(users))

	startTime := time.Now()
	setupDatabaseWorkers(db, users, maxWorkers, batchSize)
	duration := time.Since(startTime)

	fmt.Printf("Processed %d users in %v\n", len(users), duration)
	return nil
}

func main() {
	const numUsers = 1_000
	const fileName = "users_data.json"
	const maxWorkers = 20
	const batchSize = 100

	if err := generateAndSaveUserData(numUsers, fileName); err != nil {
		log.Fatal(err)
	}

	db := database.ConnectDB()
	database.Migrate(db)

	if err := processUserData(fileName, db, maxWorkers, batchSize); err != nil {
		log.Fatal(err)
	}

	Serve(db)
}
