package boot_processor

import (
	"nmigo/internal/conversion"
)

// GetIntroductionMessage
// Returns the introduction message.
func GetIntroductionMessage() string {
	introductionMessage := "\n\n\tnmigo - the database migration tool."
	introductionMessage += "\n\tCopyright (C) 2015 - present, Anatoly Khaytovich <anatolyuss@gmail.com>\n"
	return introductionMessage
}

// Boot
// Boots the migration.
func Boot(conv conversion.Conversion) {
	connectionErrorMessage := checkConnection(conv)

	if connectionErrorMessage {
		//
	}
}

// checkConnection
// Checks correctness of connection details of both MySQL and PostgreSQL.
func checkConnection(conv conversion.Conversion) string {
	sql := "SELECT 1;"
	//
}
