package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Config     Sum
	CollUser   *mongo.Collection
	CollFile   *mongo.Collection
	CollCache  *mongo.Collection
	CollConfig *mongo.Collection
	Gdb        *gorm.DB
	Gmg        *mongo.Database
)
