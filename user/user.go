package user

import (
	"golang-fiber/config"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// GetAll users
func GetUserName(c *fiber.Ctx) error {
	mg := config.MGConn
	userName := c.Params("username")
	query := bson.D{{Key: "username", Value: userName}}

	cursor, err := mg.Db.Collection("users").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []User = make([]User, 0)

	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var username string
	for user, _ := range users {
		username = users[user].UserName
	}
	return c.JSON(username)
}

func GetPassword(c *fiber.Ctx) error {
	mg := config.MGConn
	passWord := c.Params("password")
	userName := c.Params("username")

	query := bson.D{{"username", userName}, {"password", passWord}}
	cursor, err := mg.Db.Collection("users").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []User = make([]User, 0)

	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var idUser string
	for u, _ := range users {
		idUser = users[u].ID
	}
	return c.JSON(idUser)
}

// Post -> New User
func Post(c *fiber.Ctx) error {
	mg := config.MGConn

	collection := mg.Db.Collection("users")

	newUser := new(User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newUser.ID = ""

	insertionResult, err := collection.InsertOne(c.Context(), newUser)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createduser := &User{}
	createdRecord.Decode(createduser)

	return c.Status(201).JSON(createduser)
}

///// End Post user

// Put -> update user
func Put(c *fiber.Ctx) error {
	mg := config.MGConn

	idParam := c.Params("id")

	userID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.SendStatus(400)
	}

	newUser := new(User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "password", Value: newUser.Password},
			},
		},
	}

	err = mg.Db.Collection("users").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}

	newUser.ID = idParam

	return c.Status(200).JSON(newUser)
}

///// end put user

// Delete -> delete user
func Delete(c *fiber.Ctx) error {
	mg := config.MGConn

	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := mg.Db.Collection("users").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(employeeID)

}
