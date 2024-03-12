package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/richardbahrii-emodo/go-rest/database"
	"github.com/richardbahrii-emodo/go-rest/helpers"
	"github.com/richardbahrii-emodo/go-rest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName string = "todos"

func GetAllTodo(c *fiber.Ctx) error {
	collection := database.GetCollection(collectionName)

	cursor, err := collection.Find(c.Context(), bson.M{}, options.Find())

	if err != nil {
		return helpers.SendMessage(c, fiber.StatusInternalServerError, err.Error(), false, nil)
	}

	todos := make([]models.Todo, 0)
	if err = cursor.All(c.Context(), &todos); err != nil {
		return helpers.SendMessage(c, fiber.StatusInternalServerError, err.Error(), false, nil)
	}

	return helpers.SendMessage(c, fiber.StatusOK, "", true, todos)
}

type createTodoDTO struct {
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(createTodoDTO)

	if err := c.BodyParser(todo); err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	collection := database.GetCollection(collectionName)
	res, err := collection.InsertOne(c.Context(), todo)

	if err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	return helpers.SendMessage(c, fiber.StatusOK, "Succesfully created todo.", true, fiber.Map{
		"id":        res.InsertedID,
		"completed": todo.Completed,
		"title":     todo.Title,
	})
}

type updateTodoDTO struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Completed bool               `json:"completed" bson:"completed"`
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	dbID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	updatedTodo := new(updateTodoDTO)

	if err = c.BodyParser(updatedTodo); err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	collection := database.GetCollection(collectionName)

	res := collection.FindOneAndUpdate(c.Context(), bson.M{"_id": dbID}, bson.M{"$set": updatedTodo})

	if res.Err() != nil {
		return helpers.SendMessage(c, fiber.StatusInternalServerError, res.Err().Error(), false, nil)
	}

	updatedTodo.ID = dbID

	return helpers.SendMessage(c, fiber.StatusOK, "Successfully updated todo.", true, updatedTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	dbID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.SendMessage(c, fiber.StatusBadGateway, err.Error(), false, nil)
	}

	collection := database.GetCollection(collectionName)

	_, err = collection.DeleteOne(c.Context(), bson.M{"_id": dbID})

	if err != nil {
		return helpers.SendMessage(c, fiber.StatusInternalServerError, err.Error(), false, nil)
	}

	return helpers.SendMessage(c, fiber.StatusOK, "Successfully deleted todo.", true, fiber.Map{
		"deletedId": dbID,
	})

}
