package utils

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
	"os"
	"regexp"
)

func ExtractFieldFromDetail(detail string) string {
	// Регулярное выражение для извлечения поля из строки
	// Например, из строки "Key (email)=(example@gmail.com) already exists."
	re := regexp.MustCompile(`Key \((.*?)\)=`)
	matches := re.FindStringSubmatch(detail)

	// Если есть совпадения, вернуть поле
	if len(matches) > 1 {
		return matches[1]
	}

	// Если поле не найдено
	return "unknown"
}

func HandleDBError(err error) (int, map[string]string) {
	if err == nil {
		return http.StatusOK, nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // duplicate key value violates unique constraint
			return http.StatusConflict, map[string]string{
				"error":   "duplicate_key",
				"message": "Запись с такими данными уже существует",
				"field":   ExtractFieldFromDetail(pgErr.Detail),
			}
		case "23503": // foreign key violation
			return http.StatusBadRequest, map[string]string{
				"error":   "foreign_key_violation",
				"message": "Нарушена связь с другим объектом",
			}
		case "22P02": // invalid input syntax
			return http.StatusBadRequest, map[string]string{
				"error":   "invalid_input",
				"message": "Некорректные данные в запросе",
			}
		default:
			return http.StatusInternalServerError, map[string]string{
				"error":   "database_error",
				"message": "Произошла ошибка при работе с базой данных",
				"code":    pgErr.Code,
			}
		}
	}

	// Неизвестная ошибка
	return http.StatusInternalServerError, map[string]string{
		"error":   "unknown_error",
		"message": err.Error(),
	}
}

func GetDynamicHost() string {
	// Пример: хост может определяться через переменные окружения или другую логику
	if envHost := os.Getenv("SWAGGER_HOST"); envHost != "" {
		return envHost
	}
	return "localhost:8080" // значение по умолчанию
}
