package cmd

import (
	"fmt"
	"os"
	"strconv"
	"taskmanager/internal/tasks"
	UIForm "taskmanager/ui/form"
	UITable "taskmanager/ui/table"
	"taskmanager/utils"

	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Comands(db *gorm.DB) {

	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Command to add a task",
		Long:  `This command add a task to the task list.`,
		Run: func(cmd *cobra.Command, args []string) {
			task := UIForm.Create()
			newTask := tasks.Add(db, task)

			fmt.Printf(`Tarea creada con éxito: %s(%d)`, task.Name, newTask.ID)
		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Command to list all tasks",
		Long:  `This command list all tasks.`,
		Run: func(cmd *cobra.Command, args []string) {
			tasksList := tasks.GetAll(db)
			if len(tasksList) == 0 {
				fmt.Println("No hay tareas por hacer")
				return
			}

			columns := []table.Column{
				{Title: "ID", Width: 4},
				{Title: "Nombre", Width: 10},
				{Title: "Descripción", Width: 20},
				{Title: "Estado", Width: 10},
			}

			values := []table.Row{}

			for i := 0; i < len(tasksList); i++ {
				values = append(values, table.Row{strconv.Itoa(tasksList[i].ID), tasksList[i].Name, tasksList[i].Description, strconv.FormatBool(tasksList[i].Completed)})
			}

			UITable.Create(columns, values)
		},
	}

	var cmdDetail = &cobra.Command{
		Use:   "detail [id]",
		Short: "Command to show a task detail",
		Long:  `This command show a task detail.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])
			task := tasks.GetByID(db, id)
			fmt.Printf("ID: %d \nNombre: %s \nDescripción: %s \nEstado: %t \n", task.ID, task.Name, task.Description, task.Completed)

		},
	}

	var cmdUpdate = &cobra.Command{
		Use:   "update [id] [string] [string]",
		Short: "Command to update a task",
		Long:  `This command update a task.`,
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])

			var tasksEdited = tasks.GetByID(db, id)

			if args[1] == "name" {
				tasksEdited.Name = args[2]
			} else if args[1] == "description" {
				tasksEdited.Description = args[2]
			} else {
				panic("Invalid argument")
			}

			task := tasks.UpdateByID(db, id, *tasksEdited)

			fmt.Printf("Actualización realizada con éxito \n\nID: %d \nNombre: %s \nDescripción: %s \nEstado: %t \n", task.ID, task.Name, task.Description, task.Completed)

		},
	}

	var cmdCompleted = &cobra.Command{
		Use:   "completed [id]",
		Short: "Command to mark a task as completed",
		Long:  `This command mark a task as completed.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])

			var tasksEdited = tasks.GetByID(db, id)

			tasksEdited.Completed = true

			task := tasks.UpdateByID(db, id, *tasksEdited)

			fmt.Printf("La tarea con ID: %d se ha marcado como completada", task.ID)
		},
	}

	var cmdDeleted = &cobra.Command{
		Use:   "delete [id]",
		Short: "Command to delete a task",
		Long:  `This command delete a task.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])

			tasks.DeleteByID(db, id)

			fmt.Printf("La tarea con ID: %d se ha eliminado", id)
		},
	}

	var rootCmd = &cobra.Command{Use: "taskmanager"}
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdDetail)
	rootCmd.AddCommand(cmdUpdate)
	rootCmd.AddCommand(cmdCompleted)
	rootCmd.AddCommand(cmdDeleted)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
