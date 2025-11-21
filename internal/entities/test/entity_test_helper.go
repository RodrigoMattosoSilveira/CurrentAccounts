package test

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

const (
	NAME = 0
	EMAIL = 1
	CELL = 2
	PASSWORD = 3
	ROLE = 4
)

type TestCase struct {
	name string
	rest string
	path string
}
func setupTests (t *testing.T)  (*gin.Engine,  *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect test db: %v", err)
	}
	t.Helper()
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Add a cookie session store for tests (required!)
	store := cookie.NewStore([]byte("test-secret"))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("app_session", store))
	
	return router, db
}

func setupAuthenticationTests (t *testing.T, router *gin.Engine, db *gorm.DB) {
	if err := db.AutoMigrate(&people.Person{}); err != nil {
		t.Fatalf("failed to migrate test db: %v", err)
	}
	authentication.RegisterRoutes(router, db)
	PersonSeeder(db)
}
// func setupPeopleTests (t *testing.T, router *gin.Engine, db *gorm.DB) {
// 	if err := db.AutoMigrate(&people.Person{}); err != nil {
// 		t.Fatalf("failed to migrate test db: %v", err)
// 	}
// 	PersonSeeder(db)
// 	people.RegisterRoutes(router, db)
// }

// sanitizeFilename creates a safe filename from a test case name.
func sanitizeFilename(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	reg := regexp.MustCompile(`[^a-z0-9_-]`)
	return reg.ReplaceAllString(name, "")
}

// assertGoldenFile performs a golden file test for a given Gin handler.
// It makes a request to the specified path and compares the response body
// to the content of a golden file.
func assertGoldenFile(t *testing.T, router *gin.Engine, method, path string, testName string) {
	// Create the HTTP request
	req, err := http.NewRequest(method, path, nil )
	require.NoError(t, err)

	// Use the response recorder to capture the response
	w := httptest.NewRecorder( )
	router.ServeHTTP(w, req)

	// Assert that the request was successful
	require.Equal(t, http.StatusOK, w.Code, "Expected HTTP status 200" )

	// Get the actual HTML response body
	actualHTML := w.Body.String()

	// Generate the golden file path from the test name
	sanitizedName := sanitizeFilename(testName)
	goldenFileName := sanitizedName + ".golden"
	goldenFilePath := filepath.Join("testgolde", goldenFileName)

	// Update logic for golden files
	if os.Getenv("UPDATE_GOLDEN_FILES") != "" {
		t.Logf("Updating golden file: %s", goldenFilePath)
		err := os.MkdirAll(filepath.Dir(goldenFilePath), 0755)
		require.NoError(t, err)
		err = os.WriteFile(goldenFilePath, []byte(actualHTML), 0644)
		require.NoError(t, err)
	}

	// Read the golden file
	expectedHTML, err := os.ReadFile(goldenFilePath)
	require.NoError(t, err, "Failed to read golden file. Run with UPDATE_GOLDEN_FILES=true to create it.")

	// Compare the actual response to the golden file
	assert.Equal(t, string(expectedHTML), actualHTML)
}

func SetupTestRouter(route string, routeHandler  gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.GET(route, routeHandler)

	return r
}

func PersonSeeder (db *gorm.DB)  error {
	// Open the CSV file

	var count int64
	db.Model(&people.Person{}).Count(&count)
	if (count > 0) {
		log.Println("Database already seeded.")
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