package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "container/heap"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

type Point struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type PathRequest struct {
    Start Point `json:"start"`
    End   Point `json:"end"`
}

type PathResponse struct {
    Path []Point `json:"path"`
}

// Node for A* pathfinding
type Node struct {
    point    Point
    g        int     // Cost from start to current node
    h        int     // Estimated cost from current node to end
    f        int     // Total cost (g + h)
    parent   *Node
    index    int     // For heap implementation
}

// PriorityQueue implementation
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].f < pq[j].f }
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    node := x.(*Node)
    node.index = n
    *pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    node := old[n-1]
    old[n-1] = nil
    node.index = -1
    *pq = old[0 : n-1]
    return node
}

const GRID_SIZE = 20

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/find-path", findPathHandler).Methods("POST")
    
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"},
        AllowedMethods: []string{"POST", "GET", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })
    
    handler := c.Handler(r)
    log.Printf("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

func findPathHandler(w http.ResponseWriter, r *http.Request) {
    var req PathRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    path := findPath(req.Start, req.End)
    
    response := PathResponse{Path: path}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func manhattan(p1, p2 Point) int {
    return abs(p1.X - p2.X) + abs(p1.Y - p2.Y)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func findPath(start, end Point) []Point {
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)

    startNode := &Node{
        point: start,
        g: 0,
        h: manhattan(start, end),
    }
    startNode.f = startNode.g + startNode.h

    heap.Push(&pq, startNode)
    visited := make(map[string]bool)
    cameFrom := make(map[string]*Node)

    for pq.Len() > 0 {
        current := heap.Pop(&pq).(*Node)
        key := fmt.Sprintf("%d,%d", current.point.X, current.point.Y)

        if current.point == end {
            // Reconstruct path
            path := []Point{}
            for curr := current; curr != nil; curr = cameFrom[fmt.Sprintf("%d,%d", curr.point.X, curr.point.Y)] {
                path = append([]Point{curr.point}, path...)
            }
            return path
        }

        visited[key] = true

        // Check all adjacent cells
        directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
        for _, dir := range directions {
            next := Point{current.point.X + dir.X, current.point.Y + dir.Y}
            
            if next.X < 0 || next.X >= GRID_SIZE || next.Y < 0 || next.Y >= GRID_SIZE {
                continue
            }

            nextKey := fmt.Sprintf("%d,%d", next.X, next.Y)
            if visited[nextKey] {
                continue
            }

            g := current.g + 1
            h := manhattan(next, end)
            f := g + h

            neighbor := &Node{
                point: next,
                g: g,
                h: h,
                f: f,
            }

            cameFrom[nextKey] = current
            heap.Push(&pq, neighbor)
        }
    }

    return []Point{} // No path found
}