package mongo

import "time"

type entry struct {
	CreatedAt time.Time `bson:"at"`
	Field     string
	Key       string  `bson:"_id"`
	Value     *string `bson:"data,omitempty"`
	IntVal    *int    `bson:"ival,omitempty"`
}

func (d *entry) IsExpired(lifetime time.Duration) bool {
	return time.Now().After(d.CreatedAt.Add(lifetime))
}
