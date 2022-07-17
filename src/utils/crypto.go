package utils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GenerateUuid(args ...string) uuid.UUID {
	argsLength := len(args)

	if argsLength == 0 {
		uuid := uuid.NewV4()
		return uuid
	}

	argsString := strings.Join(args[:], "")
	uuid := uuid.NewV5(uuid.NamespaceURL, argsString)
	return uuid
}
