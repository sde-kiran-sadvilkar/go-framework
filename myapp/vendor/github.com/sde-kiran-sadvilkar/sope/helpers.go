package sope

import "os"

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
