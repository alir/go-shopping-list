// Package main runs the shopping list
package main

import (
	"go-shopping-list/pkg/gui"
	"go-shopping-list/pkg/recipe"
	"log"
	"runtime"
)

func main() {
	//Fetch Recipes
	allRecipes, recipeMap, err := recipe.ProcessRecipes(&recipe.FileInteractionImpl{})
	if err != nil {
		log.Fatalf("error getting all recipes err=%e", err)
	}

	//Calculate Workflow Struct
	wf, err := gui.NewWorkflow(&gui.WorkflowChecker{}, runtime.GOOS)
	if err != nil {
		log.Fatalf("error calculating workflow to use err=%e", err)
	}

	// wf = &gui.TerminalFakeWorkflow{} // Uncomment me if you want to just print to terminal!

	// Show Window
	myWindow := gui.NewApp(allRecipes, recipeMap, wf)
	myWindow.ShowAndRun()
}
