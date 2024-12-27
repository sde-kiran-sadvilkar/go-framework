package sope

import (
	"crypto/rand"
	"os"
)


const (
	randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321_+"
)

// RandomString generates a random string length n from values in the const randomString
func (s *Sope) RandomString(n int) string {
	si, r := make([]rune, n), []rune(randomString)

	for i := range si {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		si[i] = r[x%y]
	}
	return string(si)
}


func (s *Sope) CreateDirIfNotExist(path string) error {

	const mode = 0755

	_,err := os.Stat(path)
	if os.IsNotExist(err) {
		err:= os.Mkdir(path,mode)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sope) CreateFileIfNotExist(path string) error{
	_,err := os.Stat(path)
	if os.IsNotExist(err){
		file,err:= os.Create(path)

		if err != nil {
			return err
		}

		defer func (file *os.File)  {
			_ = file.Close()
		}(file)

	}
	
	return nil
}


func (s *Sope) connectToDB(){

	if os.Getenv("DATABASE_TYPE") != "" {
		
		db, err := s.OpenDB(os.Getenv("DATABASE_TYPE"),s.BuildDSN())

		if err != nil {
			s.ErrorLog.Println(err)
			os.Exit(1)
		}

		s.DB = Database{
			DataType: os.Getenv("DATABASE_TYPE"),
			Pool:  db,
		}

	}

}
