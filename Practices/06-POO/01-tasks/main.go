package main

import "fmt"

/* --------------------------------- Tareas --------------------------------- */

// Objeto
type Task struct {
	name        string
	description string
	completed   bool
}

// Métodos
func (t *Task) showAttributes() {
	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletada: %t\n\n", t.name, t.description, t.completed)
}

func (t *Task) complete() {
	t.completed = true
}

/* ----------------------------- Lista De Tareas ---------------------------- */

// Objeto
type TaskList struct {
	tasks []*Task
}

// Métodos
func (tl *TaskList) addTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(i int) {
	tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
}

func (tl *TaskList) showTasks() {
	for i := 0; i < len(tl.tasks); i++ {
		fmt.Printf("%v\n\n", *tl.tasks[i])
	}
}

/* --------------------------- Programa Principal --------------------------- */

func main() {
	taskList := TaskList{}

	firstTask := Task{
		name:        "Primer Tarea (N°1)",
		description: "Esta es la descripción de mi primer tarea.",
		completed:   false,
	}

	secondTask := Task{
		name:        "Segunda Tarea (N°2)",
		description: "Esta es la descripción de mi segunda tarea.",
		completed:   false,
	}

	firstTask.showAttributes()
	secondTask.showAttributes()

	taskList.addTask(&firstTask)
	taskList.addTask(&secondTask)
	taskList.showTasks()

	taskList.deleteTask(1)
	taskList.showTasks()
}
