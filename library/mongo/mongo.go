package mongo

import (
	"encoding/json"
	"fmt"
	"time"

	c "github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	indexName            = "expire_index"
	timeFieldName        = "at"
	MongoDupKeyErrorCode = 11000
)

type Store struct {
	session        *mgo.Session
	col            *mgo.Collection
	lifetime       time.Duration
	isTransient    bool
	ensureAccuracy bool
}

func createMongoConn() *mgo.Session {
	session, err := mgo.Dial(c.GetString("mongodb"))
	if err != nil {
		panic(err)
	}
	return session
}

func New(name string) *Store {
	ses := createMongoConn()
	col := ses.DB(c.GetString("mongo.database")).C(name)

	return &Store{
		ses,
		col,
		0,
		false,
		false,
	}
}

//create mongo index
func (s *Store) CreateIndex(d time.Duration) *Store {
	index := mgo.Index{
		Key:         []string{timeFieldName},
		Unique:      false,
		Background:  true,
		ExpireAfter: d,
		Name:        indexName,
	}
	err := s.col.EnsureIndex(index)
	if err != nil {
		return nil
	}
	return s
}

func (s *Store) Close() {
	s.session.Close()
}

//
func (s *Store) Add(field, key string, value interface{}) error {
	doc := entry{
		time.Now(),
		field,
		key,
		nil,
		nil,
	}

	switch t := value.(type) {
	case int:
		doc.IntVal = &t
	case *int:
		doc.IntVal = t
	case string:
		doc.Value = &t
	case *string:
		doc.Value = t
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		strValue := string(b)
		doc.Value = &strValue
	}

	if err := s.col.Insert(&doc); err != nil {
		mgoerr := err.(*mgo.LastError)
		if mgoerr.Code == MongoDupKeyErrorCode {
			return fmt.Errorf("Could not get the '%s' key because it does not exist or it is expired", key)
		}

		return err
	}

	return nil
}

func (s *Store) Count() (int, error) {
	return s.col.Count()
}

func (s *Store) Delete(key string) error {
	if s.ensureAccuracy {
		if err := s.testExpiration(key); err != nil {
			return err
		}
	}

	err := s.col.RemoveId(key)
	if err == mgo.ErrNotFound {
		return fmt.Errorf("Could not get the '%s' key because it does not exist or it is expired", key)
	}

	return err
}

func (s *Store) EnsureAccuracy(value bool) {
	s.ensureAccuracy = value
}

func (s *Store) Flush() error {
	_, err := s.col.RemoveAll(bson.M{})
	return err
}

func (s *Store) Get(key string, ref interface{}) error {
	if s.ensureAccuracy {
		if err := s.testExpiration(key); err != nil {
			return err
		}
	}

	if !s.isTransient {
		query := bson.M{"$currentDate": bson.M{"at": true}}
		if err := s.col.UpdateId(key, query); err != nil {
			if err == mgo.ErrNotFound {
				return fmt.Errorf("Could not get the '%s' key because it does not exist or it is expired", key)
			}
			return err
		}
	}

	doc := entry{}
	err := s.col.FindId(key).One(&doc)
	if err != nil {
		if err == mgo.ErrNotFound {
			return fmt.Errorf("Could not get the '%s' key because it does not exist or it is expired", key)
		}
		return err
	}

	switch t := ref.(type) {
	case *int:
		if doc.IntVal == nil {
			return fmt.Errorf("Unexpected type: %T", ref)
		}
		*t = *doc.IntVal
	case *string:
		if doc.Value == nil {
			return fmt.Errorf("Unexpected type: %T", ref)
		}
		*t = *doc.Value
	default:
		if doc.Value == nil {
			return fmt.Errorf("Unexpected type: %T", ref)
		}
		err = json.Unmarshal([]byte(*doc.Value), ref)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) Set(key string, value interface{}) error {
	qSet := bson.M{}
	unset := bson.M{}
	switch t := value.(type) {
	case int:
		qSet["ival"] = t
		unset["val"] = ""
	case *int:
		qSet["ival"] = *t
		unset["val"] = ""
	case string:
		qSet["val"] = t
		unset["ival"] = ""
	case *string:
		qSet["val"] = *t
		unset["ival"] = ""
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		qSet["val"] = string(b)
		unset["ival"] = ""
	}

	query := bson.M{"$set": qSet, "$unset": unset}
	if !s.isTransient {
		query["$currentDate"] = bson.M{"at": true}
	}

	if s.ensureAccuracy {
		if err := s.testExpiration(key); err != nil {
			return err
		}
	}

	if err := s.col.UpdateId(key, query); err != nil {
		if err == mgo.ErrNotFound {
			return fmt.Errorf("Could not get the '%s' key because it does not exist or it is expired", key)
		}
		return err
	}

	return nil
}

func (s *Store) SetLifetime(d time.Duration) error {
	fmt.Println(s.col)
	fmt.Println(d)
	s.col.DropIndexName(indexName)

	index := mgo.Index{
		Key:         []string{timeFieldName},
		Unique:      false,
		Background:  true,
		ExpireAfter: d,
		Name:        indexName,
	}
	s.col.EnsureIndex(index)

	s.lifetime = d
	return nil
}

func (s *Store) SetTransient(value bool) {
	s.isTransient = value
}

func (s *Store) testExpiration(key string) error {
	doc := entry{}

	err := s.col.FindId(key).One(&doc)
	if err != nil {
		if err == mgo.ErrNotFound {
			return fmt.Errorf("Could not get the '%s' key because it does not exist or it is expired", key)
		}
		return err
	}
	if doc.IsExpired(s.lifetime) {
		return fmt.Errorf("Could not get the '%s' key because it does not exist or it is expired", key)
	}

	return nil
}
