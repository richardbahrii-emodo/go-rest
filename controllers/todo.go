package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/richardbahrii-emodo/go-rest/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName string = "todos"

func (h *HandlerWithDb) GetAllTodo(c *fiber.Ctx) error {
	todos, err := h.DB.FindAll(collectionName, nil)

	if err != nil {
		return helpers.SendMessage(c, fiber.StatusInternalServerError, err.Error(), false, nil)
	}

	return helpers.SendMessage(c, fiber.StatusOK, "", true, todos)
}

type createTodoDTO struct {
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}

func (h *HandlerWithDb) CreateTodo(c *fiber.Ctx) error {
	todo := new(createTodoDTO)

	if err := c.BodyParser(todo); err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	res, err := h.DB.InsertOne(collectionName, todo)

	if err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	return helpers.SendMessage(c, fiber.StatusOK, "Succesfully created todo.", true, fiber.Map{
		"id":        res,
		"completed": todo.Completed,
		"title":     todo.Title,
	})
}

type updateTodoDTO struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Completed bool               `json:"completed" bson:"completed"`
}

func (h *HandlerWithDb) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	dbID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	updatedTodo := new(updateTodoDTO)

	if err = c.BodyParser(updatedTodo); err != nil {
		return helpers.SendMessage(c, fiber.StatusBadRequest, err.Error(), false, nil)
	}

	_, err = h.DB.UpdateOne(collectionName, bson.M{"_id": dbID}, bson.M{"$set": updatedTodo})

	if err != nil {
		return helpers.SendMessage(c, fiber.StatusInternalServerError, "Failed to Update todo", false, nil)
	}

	updatedTodo.ID = dbID

	return helpers.SendMessage(c, fiber.StatusOK, "Successfully updated todo.", true, updatedTodo)
}

func (h *HandlerWithDb) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	dbID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.SendMessage(c, fiber.StatusBadGateway, err.Error(), false, nil)
	}

	err = h.DB.DeleteOne(collectionName, bson.M{"_id": dbID})

	if err != nil {
		return helpers.SendMessage(c, fiber.StatusInternalServerError, err.Error(), false, nil)
	}

	return helpers.SendMessage(c, fiber.StatusOK, "Successfully deleted todo.", true, fiber.Map{
		"deletedId": dbID,
	})
}
