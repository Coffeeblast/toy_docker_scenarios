package database

import (
	"fmt"
	"log"
	"database/sql"
	"os"
	_ "modernc.org/sqlite"
	"strings"

	"example/story_maker/backend/model"
)

var BASE_PATH *LinuxPath = MakeLinuxPath("/home/somedude/user_data/story_maker/backend_go")
var DB_PATH *LinuxPath = BASE_PATH.Child("data/database.db")
var CREATE_ALL_SQL *LinuxPath = BASE_PATH.Child("database/queries/create_all.sql")

type Database struct {
	conn *sql.DB
}

func MakeDatabase() *Database {
	result := Database{}
	result.bootstrapDatabase()
	return &result
}

func (db *Database) bootstrapDatabase () {
	fmt.Println("Bootstraping database")
	isInitialized := DB_PATH.Exists()
	fmt.Printf("Db file exists: %v\n", isInitialized)
	fmt.Printf("Parent folder is: %v\n", DB_PATH.GetParent().path)
	if !isInitialized {
		DB_PATH.GetParent().MkDir(true)
	}
	var err error
	db.conn, err = sql.Open("sqlite", DB_PATH.GetPath())
	fmt.Printf("Opening database, value: %v error value: %v\n", db.conn, err)
	if !isInitialized {
		sql, _ := os.ReadFile(CREATE_ALL_SQL.GetPath())
		sql_text := string(sql)
		fmt.Printf("Running sql:\n%v\n", sql_text)
		createResult, createErr := db.conn.Exec(sql_text)
		fmt.Printf(
			"Creating tables: Result: %v, Error: %v\n", 
			createResult,
			createErr)
	}
}

func (db *Database) GetStories() []model.Story {
	log.Println("Getting stories")
	rows, _ := db.conn.Query("SELECT id, name, author FROM stories")
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    stories := make([]model.Story, 0)
	for rows.Next() {
        var story model.Story
        rows.Scan(&story.Id, &story.Name, &story.Author);
        stories = append(stories, story)
    }
	return stories
}

func (db *Database) AddStory(story model.Story) model.Story {
	log.Printf("Adding story %v\n", story)
	queryResult, _ := db.conn.Exec("INSERT INTO stories (name, author) VALUES (?, ?)", story.Name, story.Author)
	result := story
	newId, _ := queryResult.LastInsertId()
	result.Id = int(newId)
	return result
}

func (db *Database) DeleteStories(storyIds []int) {
	log.Printf("Deleting stories with IDs: %v\n", storyIds)
	if len(storyIds) > 0 {
		// Create placeholders for each ID
		placeholders := make([]string, len(storyIds))
		args := make([]any, len(storyIds)) // this is required so that args... construction may be used
		for i, id := range storyIds {
			placeholders[i] = "?"
			args[i] = id
		}

		// Construct the query
		query := fmt.Sprintf("DELETE FROM stories WHERE id IN (%s)", strings.Join(placeholders, ","))
		_, err := db.conn.Exec(query, args...)
		if err != nil {
			log.Printf("Deleting failed: %v\n", err)
			return
		}
	}
}

type LinuxPath struct {
	path string
}

func MakeLinuxPath(pathString string) *LinuxPath {
	var result LinuxPath
	if (
		len(pathString) > 1 && 
		pathString[len(pathString) - 1] == '/') {
		result = LinuxPath{pathString[:len(pathString) - 1]}
	} else {
		result = LinuxPath{pathString}
	}
	return &result
}

func (p *LinuxPath) Exists () bool {
	_, err := os.Stat(p.path)
	if err == nil {
		// File exists
		return true
	}
	if os.IsNotExist(err) {
		// File does not exist
		return false
	}
	// Some other error occurred (e.g., permission issues)
	return false
}

func (p *LinuxPath) GetPath() string {
	return p.path
}

func (p * LinuxPath) GetParent() *LinuxPath {
	idx := strings.LastIndex(p.path, "/")
	if idx <= 0 {
		// this should be handled with error in case idx == -1
		return MakeLinuxPath("/")
	} else {
		return MakeLinuxPath(p.path[:idx])
	}
}

func (p * LinuxPath) MkDir(parents bool) {
	// errors should be handled more gracefully :P
	perms := os.FileMode(0o700)
	if parents {
		os.MkdirAll(p.path, perms)
	} else {
		os.Mkdir(p.path, perms)
	}
}

func (p *LinuxPath) Child(childPath string) *LinuxPath {
	return MakeLinuxPath(p.path + "/" + childPath)
}

