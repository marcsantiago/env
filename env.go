// Package env contains a set of utility functions to read environment
// variables.
package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Var reads an environment variable as a string.
func Var(envVarName, defaultValue string) string {
	return getVar(envVarName, defaultValue, false)
}

// ImportantVar reads an environment variable as a string, emitting a warning
// when one is not available and it is forced to return the fallback value.
func ImportantVar(envVarName, defaultValue string) string {
	return getVar(envVarName, defaultValue, true)
}

// MandatoryVar reads an environment variable as a string, emitting a fatal
// log statement when one is not available.
func MandatoryVar(envVarName string) string {
	v := os.Getenv(envVarName)
	if len(v) == 0 {
		panic(fmt.Sprintf("Missing mandatory env var name %s", envVarName))
	}

	return v
}

// VarAsBool reads an environment variable as a boolean.
func VarAsBool(envVarName string, defaultValue bool) bool {
	v := os.Getenv(envVarName)
	if len(v) == 0 {
		return defaultValue
	}

	bv, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}

	return bv
}

// MandatoryVarAsBool is the Mandatory version of VarAsBool.
func MandatoryVarAsBool(envVarName string) bool {
	str := MandatoryVar(envVarName)

	v, err := strconv.ParseBool(str)
	if err != nil {
		panic(fmt.Sprintf("Invalid bool value for env var name %s value %s", envVarName, str))
	}

	return v
}

// VarAsInt reads an environment variable as an integer.
func VarAsInt(envVarName string, defaultValue int) int {
	return int(VarAsInt64(envVarName, int64(defaultValue)))
}

// MandatoryVarAsInt is the version of VarAsInt.
func MandatoryVarAsInt(envVarName string) int {
	return int(MandatoryVarAsInt64(envVarName))
}

// VarAsInt64 reads an environment variable as a 64 bit integer.
func VarAsInt64(envVarName string, defaultValue int64) int64 {
	v := os.Getenv(envVarName)
	if len(v) == 0 {
		return defaultValue
	}

	iv, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}

	return iv
}

// MandatoryVarAsInt64 is the Mandatory version of VarAsInt64.
func MandatoryVarAsInt64(envVarName string) int64 {
	str := MandatoryVar(envVarName)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Invalid int value for env var name %s value %s", envVarName, str))
	}

	return v
}

// VarAsStringSlice reads an environment variable as a []string
func VarAsStringSlice(envVarName string, defaultValue []string, deliminator rune) []string {
	v := os.Getenv(envVarName)
	if len(v) == 0 {
		return defaultValue
	}

	if vv := strings.Split(v, string(deliminator)); len(vv) > 0 {
		return vv
	}
	return defaultValue
}

// MandatoryVarAsStringSlice is the Mandatory version of VarAsStringSlice.
func MandatoryVarAsStringSlice(envVarName string, deliminator rune) []string {
	str := MandatoryVar(envVarName)
	vv := strings.Split(str, string(deliminator))
	if len(vv) == 0 {
		panic(fmt.Sprintf("Invalid []string value for env var name %s value %s", envVarName, str))
	}
	return vv
}

func getVar(name, defaultValue string, warn bool) string {
	v := os.Getenv(name)
	if len(v) == 0 {
		if warn {
			log.Println("Using fallback default value for env var.", "var", name, "default", defaultValue)
		}
		return defaultValue
	}
	return v
}
