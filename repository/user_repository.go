package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/ToluwaniO/utube/global"
	"github.com/ToluwaniO/utube/helper"
	"github.com/ToluwaniO/utube/model"
	"log"
)

var collection *firestore.CollectionRef

func AddUser(id string, user *model.User) (*model.User, error)  {
	collection = global.Firestore.Collection("users")
	if helper.IsBlank(user.UID) || helper.IsBlank(user.Email) {
		return user, errors.New("invalid data")
	}
	result, err := collection.Doc(id).Set(context.Background(), user.ToMap())
	if err != nil {
		log.Fatal(err)
		return &model.User{}, err
	}
	log.Println(fmt.Sprintf("User with id = %s was added on %s", id, result.UpdateTime.String()))
	return user, nil
}
