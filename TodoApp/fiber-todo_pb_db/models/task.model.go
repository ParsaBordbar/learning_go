package models

type Todo struct {
	ID	 		int    `json:"id"`
	Title 		string `json:"title"`
	AssignedTo 	string `json:"assigned_to"`
	State 		string `json:"state"`
}

var Todos = []Todo{
	{ID: 1, Title: "Learn Go", AssignedTo: "Parsa", State: "in-progress 8%"},
	{ID: 2, Title: "Learn Rust", AssignedTo: "Parsa", State: "in-progress chapter 9"},
}