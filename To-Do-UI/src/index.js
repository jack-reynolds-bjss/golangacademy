import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import './index.css';
import App from './App';
import CreateUpdate from './CreateUpdate';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/create",
    element: <CreateUpdate />,
  },
  {
    path: "/edit/:id",
    element: <CreateUpdate edit />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <>
    <h1>Todo App</h1>
    <RouterProvider router={router} />
  </>
);