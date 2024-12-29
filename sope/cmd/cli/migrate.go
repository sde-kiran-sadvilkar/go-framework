package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	switch arg2 {
	case "up":
		err := sop.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":

		if arg3 == "all" {
			err := sop.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := sop.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := sop.MigrateDownAll(dsn)
		if err != nil {
			return err
		}

		err = sop.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil

}
