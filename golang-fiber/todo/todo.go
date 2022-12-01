package todo

import (
	"golang-fiber/config"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ----------------------------------------------------------------

type Todo struct {
	TodoID  string `json:"todoID,omitempty" bson:"_id,omitempty"`
	Done    bool   `json:"done"`
	Content string `json:"content"`
	UserID  string `json:"userid"`
}

// GetAll users
func GetAll(c *fiber.Ctx) error {
	mg := config.MGConn

	userId := c.Params("userid")
	query := bson.D{{Key: "userid", Value: userId}}

	cursor, err := mg.Db.Collection("todos").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var todos []Todo = make([]Todo, 0)

	if err := cursor.All(c.Context(), &todos); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(todos)
}

// Post -> New User
func Post(c *fiber.Ctx) error {
	mg := config.MGConn

	collection := mg.Db.Collection("todos")

	newTodo := new(Todo)

	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newTodo.TodoID = ""

	insertionResult, err := collection.InsertOne(c.Context(), newTodo)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createduser := &Todo{}
	createdRecord.Decode(createduser)

	return c.Status(201).JSON(createduser)
}

// ---------------------- End Post user

// Ppdate content
func UpdateContent(c *fiber.Ctx) error {
	mg := config.MGConn

	idParam := c.Params("id")
	todoID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.SendStatus(400)
	}

	newTodo := new(Todo)

	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: todoID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "content", Value: newTodo.Content},
			},
		},
	}

	err = mg.Db.Collection("todos").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}

	newTodo.TodoID = idParam

	return c.Status(200).JSON(newTodo)
}

// --------------------- end update content

// Update status done
func UpdateDone(c *fiber.Ctx) error {
	mg := config.MGConn

	idParam := c.Params("id")

	todoID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.SendStatus(400)
	}

	newTodo := new(Todo)

	if err := c.BodyParser(newTodo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_todoID", Value: todoID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "done", Value: true},
			},
		},
	}

	err = mg.Db.Collection("todos").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}

	newTodo.TodoID = idParam

	return c.Status(200).JSON(newTodo)
}

// ------------------ end update status done

// Delete -> delete user
func Delete(c *fiber.Ctx) error {
	mg := config.MGConn

	todoID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: todoID}}
	result, err := mg.Db.Collection("todos").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(todoID)
}

// ---------------------- end delete todo
