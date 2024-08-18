// File: cmd/migration_generator/main.go

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
)

func main() {
	// Generate migration
	migrationContent := generateMigration()

	// Write migration to file
	err := writeMigrationToFile(migrationContent)
	if err != nil {
		log.Fatalf("Failed to write migration: %v", err)
	}

	log.Println("Migration generated successfully")
}

func generateMigration() string {
    var migration strings.Builder

    // Add your models here
    models := []interface{}{
        &models.UserProfile{},
        &models.Creator{},
        &models.Article{},
        &models.Tag{},
        &models.Shop{},
        &models.Category{},
        &models.Product{},
        &models.Comment{},
        &models.ArticleLike{},
        &models.ArticleEmbedding{},
        &models.ProductEmbedding{},
        &models.UserArticleInteraction{},
        &models.UserProductInteraction{},
        &models.UserArticleRecommendation{},
        &models.UserProductRecommendation{},
        &models.KafkaEvent{},
    }

    for _, model := range models {
        t := reflect.TypeOf(model).Elem()
        migration.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", toSnakeCase(t.Name())))

        fields := []string{}
        for i := 0; i < t.NumField(); i++ {
            field := t.Field(i)
            tag := field.Tag.Get("gorm")
            if tag == "-" {
                continue
            }

            columnName := getColumnName(field)
            columnType := getColumnType(field)

            fieldDef := fmt.Sprintf("  %s %s", columnName, columnType)

            if isPrimaryKey(field) {
                fieldDef += " PRIMARY KEY"
            }
            if isNotNull(field) {
                fieldDef += " NOT NULL"
            }
            fields = append(fields, fieldDef)
        }

        migration.WriteString(strings.Join(fields, ",\n"))
        migration.WriteString("\n);\n\n")
    }

    return migration.String()
}


func writeMigrationToFile(content string) error {
	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("%s_create_schema.up.sql", timestamp)
	
	migrationDir := "./migrations"
	if err := os.MkdirAll(migrationDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create migrations directory: %v", err)
	}
	
	filepath := filepath.Join(migrationDir, filename)
	
	return os.WriteFile(filepath, []byte(content), 0644)
}

func toSnakeCase(str string) string {
	var result strings.Builder
	for i, r := range str {
		if i > 0 && (r >= 'A' && r <= 'Z') {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func getColumnName(field reflect.StructField) string {
	tag := field.Tag.Get("gorm")
	if tag != "" {
		parts := strings.Split(tag, ";")
		for _, part := range parts {
			if strings.HasPrefix(part, "column:") {
				return strings.TrimPrefix(part, "column:")
			}
		}
	}
	return toSnakeCase(field.Name)
}

func getColumnType(field reflect.StructField) string {
    tag := field.Tag.Get("gorm")
    if strings.Contains(tag, "type:vector") {
        return "vector"
    }

    switch field.Type.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return "INTEGER"
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return "INTEGER"
    case reflect.Float32, reflect.Float64:
        return "REAL"
    case reflect.String:
        return "TEXT"
    case reflect.Bool:
        return "BOOLEAN"
    case reflect.Struct:
        if field.Type.Name() == "Time" {
            return "TIMESTAMP"
        }
    case reflect.Slice:
        if field.Type.Elem().Kind() == reflect.Float32 || field.Type.Elem().Kind() == reflect.Float64 {
            return "vector"
        }
    }
    return "TEXT" // Default to TEXT for unknown types
}


func isPrimaryKey(field reflect.StructField) bool {
	tag := field.Tag.Get("gorm")
	return strings.Contains(tag, "primaryKey")
}

func isNotNull(field reflect.StructField) bool {
	tag := field.Tag.Get("gorm")
	return !strings.Contains(tag, "null")
}
