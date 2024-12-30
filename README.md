
# Shortest Pathfinding Visualization Using DFS(Depth first serach)

This project is a visualization tool for the A* pathfinding algorithm. It includes a backend implemented in Go and a frontend built using React.

## Features

- **Interactive Grid**: Users can select start and end points on a 20x20 grid.
- **Shortest Path Calculation**: Calculates the shortest path using the A* algorithm.
- **Backend API**: Powered by a Go backend that processes the pathfinding logic.
- **Frontend Interface**: Built with React to provide an intuitive and interactive UI.
- **Cross-Origin Resource Sharing**: Enabled to support local frontend-backend communication.

---

## Technology Stack

### Backend
- **Language**: Go
- **Frameworks/Libraries**: 
  - `gorilla/mux` for routing
  - `rs/cors` for CORS handling

### Frontend
- **Language**: JavaScript (React)
- **Styling**: CSS

---

## Installation and Setup

### Prerequisites
- Node.js installed for the frontend
- Go installed for the backend

### Backend Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/kanishka231/full_stack_devleopment_task.git
   cd full_stack_devleopment_task/backend
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the backend:
   ```bash
   go run main.go
   ```
4. The backend server will be available at `http://localhost:8080`.

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd full_stack_devleopment_task/pathfinder-frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```
4. The frontend will be available at `http://localhost:5173`.

---

## Usage

1. Open the frontend in your browser at `http://localhost:5173`.
2. Click a grid cell to set the **start point** (green).
3. Click another cell to set the **end point** (red).
4. Click the **"Find Shortest Path"** button to calculate and visualize the path.
5. Reset the grid using the **"Reset Grid"** button.

---

## API Reference

### Endpoint
- **POST** `/find-path`

### Request Body
```json
{
  "start": { "x": <int>, "y": <int> },
  "end": { "x": <int>, "y": <int> }
}
```

### Response Body
```json
{
  "path": [
    { "x": <int>, "y": <int> },
    ...
  ]
}
```

---

## Project Structure

### Backend
```
backend/
├── main.go
├── go.mod
└── go.sum
```

### Frontend
```
frontend/
├── src/
│   ├── components/
│   │   └── GridPathfinder.js
│   │   └── GridPathfinder.css
│   └── App.js
├── public/
└── package.json
```

---

## Future Enhancements

- Add obstacles to the grid for more complex pathfinding scenarios.
- Support for diagonal movement.
- Enhanced UI/UX with animations.
