package test

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

const (
	NAME     = 0
	ADDRESS  = 1
	EMAIL    = 2
	CELL     = 3
	PASSWORD = 4
	ROLE     = 5
)

type TestCase struct {
	Name string
	Rest string
	Path string
}		

func SetupFiberTests(t *testing.T) *fiber.App {
	db := SetupTestDB(t)

	// Setup the authentication controller
	app := SetupTestServerFiber(t, db)

	return app
}		

func SetupTestDB(t *testing.T) *gorm.DB	 {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect test db: %v", err)
	}
	t.Helper()

	// Migrate the schema
	err = db.AutoMigrate(&people.Person{})
	if err != nil {
		t.Fatalf("failed to migrate test db: %v", err)
	}

	err = PersonSeeder(db)
	if err != nil {
		t.Fatalf("failed to seed test db: %v", err)
	}
	return db
}

func PersonSeeder (db *gorm.DB)  error {
	// Open the CSV file

	var count int64
	db.Model(&people.Person{}).Count(&count)
	if (count > 0) {
		if err := db.Exec("DELETE FROM people").Error; err != nil {
			log.Fatalf("failed to clear table: %v", err)
		}
		if err := db.Exec("VACUUM").Error; err != nil {
			log.Fatalf("failed to vacuum: %v", err)
		}

		fmt.Println("Initialized the database.")
	}

	projectRoot, err := utilities.FindProjectRoot()
	if err != nil {
		fmt.Println("Error retrieving project root:", err)
		return err
	}
	peopleFile := filepath.Join(projectRoot, "internal/entities/test/testData/people.csv")
	file, err := os.Open(peopleFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the CSV file
	_people, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading CSV file:", err)
		return err
	}

	// Process the rows
	var person people.Person
	var persons []people.Person
	for _, row := range _people {
		// fmt.Printf("Row %d: %v\n", i, row)
		person.Name = row[NAME]
		person.Email = row[EMAIL]
		person.Cell = row[CELL]
		hashedPassword, err := HashPassword(row[PASSWORD])
		if err != nil {
			return errors.New("unable to hash password")
		}
		err = CheckPassword(row[PASSWORD], hashedPassword)
			if err != nil {
			fmt.Println("Invalid password")
		}
		person.Password = hashedPassword
		person.Role = row[ROLE]
		persons = append(persons, person)
	}
	db.Create(persons)
	log.Println("populated people test database")
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func SetupTestServerFiber(t *testing.T, db *gorm.DB) *fiber.App {
	t.Helper()

	app := fiber.New()
	store := session.New()
	app.Use(utilities.WithSession(store))

	// Register ONLY the routes needed for the test
	authentication.RegisterRoutes(app, db)

	return app
}

// sanitizeFilename creates a safe filename from a test case name.
func SanitizeFilename(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	reg := regexp.MustCompile(`[^a-z0-9_-]`)
	return reg.ReplaceAllString(name, "")
}
