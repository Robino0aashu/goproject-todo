package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID 			int `json:"id"`
	Completed 	bool `json:"completed"`
	Body 		string `json:"body"`
}

func main() {
	fmt.Println("Server Started!!")
	app:= fiber.New()

	todos:=[]Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error{
		return c.Status(200).JSON(todos)
	})

	//Create a Todo
	app.Post("api/todos", func(c *fiber.Ctx) error{
		todo:=&Todo{} //{id:0, completed: false, body:""}

		if err:=c.BodyParser(todo); err!=nil{
			return err
		}
		if todo.Body==""{
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}
		todo.ID=len(todos)+1
		todos=append(todos, *todo)

		return c.Status(201).JSON(todos)
	})


	//update a Todo
	app.Patch("/api/updateTodo/:id", func(c *fiber.Ctx) error{
		id:=c.Params("id")
		
		for i, todo:=range todos{
			if fmt.Sprint(todo.ID)==id{
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"error": fmt.Sprint("No todo found with this id: ", id)})
	})

	//Delete a Todo
	app.Delete("api/deleteTodo/:id" ,func(c *fiber.Ctx) error{
		id:=c.Params("id")

		for i, todo := range todos{
			if fmt.Sprint(todo.ID)==id{
				todos=append(todos[:i],todos[i+1:]...)
				return c.Status(200).JSON(todos)
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"error": fmt.Sprint("No Todo found for the id: ",id)})
	})

	log.Fatal(app.Listen(":3000"))
}
