package config

import (
	"github.com/boltdb/bolt"
	"os"
)

const (
	conf         = "world.aur"
	pkgInstalled = "pkgInstalled"
	//pkgLatest    = "pkgLatest"
)

var (
	db *bolt.DB
)

func getConfigFilePath() string {
	c, err := os.UserConfigDir()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return c + "/" + conf
}

func init() {
	var err error
	db, err = bolt.Open(getConfigFilePath(), 0600, nil)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	//defer db.Close()

	if err := db.Batch(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists([]byte(pkgInstalled))
		return e
	}); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	//if err := db.Batch(func(tx *bolt.Tx) error {
	//	_, e := tx.CreateBucketIfNotExists([]byte(pkgLatest))
	//	return e
	//}); err != nil {
	//	println(err)
	//	os.Exit(1)
	//}
}

func AddPackage(pkg ...PkgInfo) {
	err := db.Batch(func(tx *bolt.Tx) error {
		for _, v := range pkg {
			e := tx.Bucket([]byte(pkgInstalled)).Put([]byte(v.Name), []byte(v.Installed))
			if e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func UpdatePackage(fn func(name, version string) string) {
	err := db.Batch(func(tx *bolt.Tx) error {
		change := make(map[string]string)
		e := tx.Bucket([]byte(pkgInstalled)).ForEach(func(k []byte, v []byte) error {
			if newVersion := fn(string(k), string(v)); newVersion != "" {
				change[string(k)] = newVersion
			}
			return nil
		})
		if e != nil {
			return e
		}

		for k, v := range change {
			e = tx.Bucket([]byte(pkgInstalled)).Put([]byte(k), []byte(v))
			if e != nil {
				return e
			}
		}

		return nil
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func RemovePackage(pkg ...string) {
	err := db.Batch(func(tx *bolt.Tx) error {
		for _, v := range pkg {
			e := tx.Bucket([]byte(pkgInstalled)).Delete([]byte(v))
			if e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func IsExist(pkg string) bool {
	var existed bool

	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		if tx.Bucket([]byte(pkgInstalled)).Get([]byte(pkg)) != nil {
			existed = true
		}

		return nil
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	return existed
}

func GetVersion(pkg string) string {
	var ver string

	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		if v := tx.Bucket([]byte(pkgInstalled)).Get([]byte(pkg)); v != nil {
			ver = string(v)
		}

		return nil
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	return ver
}

func WalkThroughKVs(fn func(k, v string) error) {
	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(pkgInstalled))

		err := b.ForEach(func(k, v []byte) error {
			return fn(string(k), string(v))
		})
		return err
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
