import React, { useState, useCallback } from 'react';
import './GridPathfinder.css';

const GridPathfinder = () => {
  const GRID_SIZE = 20;
  const [startPoint, setStartPoint] = useState(null);
  const [endPoint, setEndPoint] = useState(null);
  const [path, setPath] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  const handleCellClick = (x, y) => {
    if (!startPoint) {
      setStartPoint({ x, y });
      setPath([]);
    } else if (!endPoint && !(x === startPoint.x && y === startPoint.y)) {
      setEndPoint({ x, y });
      setPath([]);
    }
  };

  const getCellClass = useCallback((x, y) => {
    if (startPoint && x === startPoint.x && y === startPoint.y) {
      return 'cell start-point';
    }
    if (endPoint && x === endPoint.x && y === endPoint.y) {
      return 'cell end-point';
    }
    if (path && path.some(point => point.x === x && point.y === y)) {
      return 'cell path-point';
    }
    return 'cell';
  }, [startPoint, endPoint, path]);

  const findPath = async () => {
    if (!startPoint || !endPoint) {
      alert('Please select both start and end points first');
      return;
    }

    setIsLoading(true);
    try {
      const response = await fetch('http://localhost:8080/find-path', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ start: startPoint, end: endPoint }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log('Path found:', data.path);
      
      if (data.path && data.path.length > 0) {
        setPath(data.path);
      } else {
        alert('No path found between the selected points');
      }
    } catch (error) {
      console.error('Error finding path:', error);
      alert('Error finding path. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const resetGrid = () => {
    setStartPoint(null);
    setEndPoint(null);
    setPath([]);
  };

  return (
    <div className="container">
      <div className="grid-container">
        <div className="grid">
          {Array.from({ length: GRID_SIZE }).map((_, y) => (
            <div key={y} className="grid-row">
              {Array.from({ length: GRID_SIZE }).map((_, x) => (
                <div
                  key={`${x}-${y}`}
                  className={getCellClass(x, y)}
                  onClick={() => handleCellClick(x, y)}
                  data-coordinates={`${x},${y}`}
                />
              ))}
            </div>
          ))}
        </div>
        <div className="controls">
          <button
            onClick={findPath}
            className="find-path-button"
            disabled={isLoading || !startPoint || !endPoint}
          >
            {isLoading ? 'Finding Path...' : 'Find Shortest Path'}
          </button>
          <button
            onClick={resetGrid}
            className="reset-button"
            disabled={isLoading}
          >
            Reset Grid
          </button>
        </div>
        <div className="instructions">
          <p>1. Click to set start point (green)</p>
          <p>2. Click again to set end point (red)</p>
          <p>3. Click "Find Shortest Path" to calculate the route</p>
        </div>
      </div>
    </div>
  );
};

export default GridPathfinder;