package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type Point struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
type RequestBody struct {
	Start Point `json:"start"`
	End Point `json:"end"`
}

type ResponseBody struct {
	Path []Point `json:"path"`
}

func dfs(grid[][]int,current Point,end Point ,path []Point, visited map[string]bool) []Point{
   if current.Row == end.Row && current.Col == end.Col {
	 return append(path, current)
   }
   key := fmt.Sprintf("%d,%d", current.Row, current.Col)
   if visited[key] {
	 return nil
   }
   visited[key] = true
   directions := []Point{
	{Row: current.Row + 1, Col: current.Col},
	{Row: current.Row - 1, Col: current.Col},
	{Row: current.Row, Col: current.Col + 1},
	{Row: current.Row, Col: current.Col - 1},
   }
   for _,dir := range directions {
	if dir.Row >=0 && dir.Row < len(grid) && dir.Col >=0 && dir.Col <len(grid[0]){
		newPath := dfs(grid,dir,end,append(path,current),visited)
		if newPath != nil {
			return newPath
		}
	}
	}
	return nil
}
func handleShortestPath (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	if r.Method == http.MethodOptions {
		return
	}
	var body RequestBody
	if err:= json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gridSize := 20
	grid:= make([][]int,gridSize)
	for i:=range grid{
		grid[i] = make([]int,gridSize)
	}
	visited := make(map[string]bool)
	path := dfs(grid,body.Start,body.End,[]Point{},visited)
	if path != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponseBody{Path: path})
	}else{
		http.Error(w, "No path found", http.StatusNotFound)
		return
	}
}

func main() {
	http.HandleFunc("/shortest-path",handleShortestPath)
	log.Fatal(http.ListenAndServe(":5000", nil))
	fmt.Println("Hello, World!")
}