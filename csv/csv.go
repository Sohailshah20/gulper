package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5"
)

func ReadFile(pathStr string) ([]string, [][]string, error) {
	// filePath := "C:\\Users\\Sohail Shah\\Downloads\\customers-100.csv"
	path := filepath.Clean(pathStr)
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, errors.New("error opening file :" + err.Error())
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, nil, errors.New("error reading file :" + err.Error())
	}
	var columns []string
	for i := 0; i < len(data[0]); i++ {
		columns = append(columns, data[0][i])
	}
	return columns, data, nil
}

func GetRowString(arr []string) string {
	var res string
	for i := 1; i < len(arr); i++ {
		res += `"`
		res += arr[i]
		res += `"`
		if i != len(arr)-1 {
			res += `,`
		}
	}
	return res
}

func FormatValues(vals []string) (string, []string) {
	var values string
	var valuesArr []string
	for i := 1; i < len(vals); i++ {
		temp := "@"
		res := fmt.Sprintf(strings.ReplaceAll(vals[i], " ", ""))
		temp += res
		valuesArr = append(valuesArr, res)
		if i < len(vals)-1 {
			temp += ", "
		}
		values += temp

	}
	return values, valuesArr
}

func BatchArggs(strArr []string, valArr []string) pgx.NamedArgs {
	args := make(pgx.NamedArgs, len(strArr))
	for i := 0; i < len(strArr); i++ {
		args[strArr[i]] = valArr[i+1]
	}
	return args
}
