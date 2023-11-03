package database

import (
	"fmt"
	"gorm.io/gorm"
	"io/fs"
	"metro-in/internal/common/entity"
	"metro-in/internal/common/custom_errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const migrationsPath = "internal/config/database/migrations"
const migrationContext = "migration"

// Migration entity
type Migration struct {
	ID    string
	Query string
}

var tables = []interface{}{
	&entity.SubwayLineStation{},
	&entity.Station{},
	&entity.SubwayLine{},
}

// RunMigrations runs all autoMigrates and new migrations listed in database/migrations
func RunMigrations(db *gorm.DB) error {

	if err := autoMigrateTables(db); err != nil {
		return custom_errors.Error{Context: migrationContext, Message: "Failed during autoMigrate", Err: err}
	}

	if err := db.AutoMigrate(&Migration{}); err != nil {
		return custom_errors.Error{Context: migrationContext, Message: "Could not create Migrations Table", Err: err}
	}

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return custom_errors.Error{Context: migrationContext, Message: "Could not read migration directory", Err: err}
	}

	migrations, err := buildMigrations(files)
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		if err = runMigration(&migration, db); err != nil {
			return err
		}
	}

	return nil
}

func runMigration(migration *Migration, db *gorm.DB) error {

	isNew, err := isNewMigration(migration, db)
	if err != nil {
		return err
	}

	if !isNew {
		return nil
	}

	tx := db.Begin()
	if tx.Error != nil {
		return custom_errors.Error{Context: migrationContext, Message: fmt.Sprintf("Error initiating transaction"), Err: tx.Error}
	}

	if err = executeMigration(migration, tx); err != nil {
		return rollbackOnError(tx, fmt.Sprintf("Error running migration %s", migration.ID))
	}

	if err = createMigrationReference(migration, tx); err != nil {
		return rollbackOnError(tx, fmt.Sprintf("Error creating migration reference: %s", migration.ID))
	}

	if err = tx.Commit().Error; err != nil {
		return rollbackOnError(tx, fmt.Sprintf("Error committing changes: %s", migration.ID))
	}

	return nil
}

func rollbackOnError(tx *gorm.DB, message string) error {
	tx.Rollback()
	return custom_errors.Error{
		Context: migrationContext,
		Message: message,
		Err:     tx.Error,
	}
}

func executeMigration(migration *Migration, tx *gorm.DB) error {
	return tx.Exec(migration.Query).Error
}

func createMigrationReference(migration *Migration, tx *gorm.DB) error {
	return tx.Create(migration).Error
}

func buildMigrations(files []fs.DirEntry) ([]Migration, error) {
	var migrations []Migration

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		query, err := readMigrationFile(file.Name())
		if err != nil {
			return migrations, err
		}

		migrationID, err := fileNameToFileID(file.Name())
		if err != nil {
			return migrations, err
		}
		migrations = append(migrations, Migration{migrationID, query})
	}

	return migrations, nil
}

func isNewMigration(migration *Migration, db *gorm.DB) (bool, error) {
	var count int64
	if err := db.Model(migration).Where("id = ?", migration.ID).Count(&count).Error; err != nil {
		return false, custom_errors.Error{
			Context: migrationContext,
			Message: fmt.Sprintf("Error looking for migration %s at the database", migration.ID),
			Err:     err,
		}
	}

	return count == 0, nil
}

func readMigrationFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(filepath.Join(migrationsPath, fileName))
	if err != nil {
		return "", custom_errors.Error{Context: migrationContext, Message: "Could not read migration file", Err: err}
	}

	return string(fileContent), nil
}

func fileNameToFileID(fileName string) (string, error) {
	if fileName == "" {
		return "", custom_errors.Error{Context: migrationContext, Message: "Migration file name can't be empty", Err: nil}
	}

	fileNameSplit := strings.Split(fileName, "_")
	if len(fileNameSplit) < 2 {
		return "", custom_errors.Error{Context: migrationContext, Message: "Migration file name must be formatted like '${date}_${migration_action}' ", Err: nil}
	}

	_, err := strconv.Atoi(fileNameSplit[0])
	if err != nil {
		return "", custom_errors.Error{Context: migrationContext, Message: "Migration file name must contain only numbers before the first underscore", Err: nil}
	}

	return fileNameSplit[0], nil
}
