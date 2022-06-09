package package2

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*3.Для закрепления практических навыков программирования, напишите программу, которая создаёт
один миллион пустых файлов в известной, пустой директории файловой системы используя вызов os.Create.
Ввиду наличия определенных ограничений операционной системы на число открытых файлов, такая программа
должна выполнять аварийную остановку. Запустите программу и дождитесь полученной ошибки. Используя
отложенный вызов функции закрытия файла, стабилизируйте работу приложения. Критерием успешного выполнения
программы является успешное создание миллиона пустых файлов в директории*/

var dPath = flag.String("dirPath", "./test", "folder for files")

func main() {

	flag.Parse()

	dirPath := strings.TrimSpace(*dPath)
	dirPathWithout := strings.Replace(dirPath, "./", "", 1)

	err := os.MkdirAll(dirPathWithout, 0777)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	for i := 0; i < 1000000; i++ {
		fileName := *dPath + "/" + "file" + fmt.Sprint(i)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(errors.New("error: cannot create file"), fileName)
		}

		fileClose(file)
	}

}
func fileClose(file *os.File) {
	defer file.Close()
}
