package cdn_log_parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// разбиение файла с логами cdn
func StartParse(file_init_path string, limitlog int) {
	// открываем файл с логами
	file1, err := os.OpenFile(file_init_path, os.O_RDONLY, 0111)
	if err != nil {
		logrus.Fatal("Error in StartParse: OpenFile ", err)
	}
	defer file1.Close()

	ParseLogs(file1, limitlog)
}

// ParseLogs чтение файла логов и запись в БД
func ParseLogs(file_init *os.File, limitlog int) error {

	cdn_path := "./cdn_file/cdn1.log"   

	os.Create(cdn_path)
	file_log, err := file_create_and_open(cdn_path)
	if err != nil {
		logrus.Fatal("Error in StartParse: OpenFile ", err)
	}

	// задаем значения, чтобы постоянно не выделять память под них
	count := 0
	count_file := 1
	// читаем файл
	scanner := bufio.NewScanner(file_init)
	for scanner.Scan() {
		count++
		file_log.Write(readerByteString(scanner.Text() + "\n"))
		if count == limitlog {
			// закрываю старый
			file_log.Close()
			rename(cdn_path, "./cdn_file/cdn"+fmt.Sprint(count_file)+".log")

			// создаю и открываю новый файл
			count_file++
			cdn_path = "./cdn_file/cdn" + fmt.Sprint(count_file) + ".log"
			file_log, err = file_create_and_open(cdn_path)
			if err != nil {
				logrus.Error(err)
			}
			count = 0
		}
	}

	log.Println("Количество обработанных логов:", limitlog*(count_file-1)+count)
	return nil
}

// создаем и открываем файл
func file_create_and_open(cdn_path string) (*os.File, error) {

	os.Create(cdn_path)

	file_log, err := os.OpenFile(cdn_path, os.O_WRONLY, 0111)
	if err != nil {
		return nil, err
	}

	return file_log, nil
}

// из строки в массив байт
func readerByteString(r string) []byte {
	b := bytes.NewReader([]byte(r))
	data, err := ioutil.ReadAll(b)
	if err != nil {
		log.Println(err)
	}
	return data
}

// перемещение файлов
func rename(srcPath, dstPath string) {
	err := os.Rename(srcPath, dstPath)
	if err != nil {
		log.Fatal(err)
	}
}
