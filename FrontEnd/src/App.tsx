import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import Root from "./Root";
import AuthLayout from "./AuthLayout";
import Home from "./components/pages/Home";
import Employees from "./components/pages/Employees";
import SignIn from "./components/pages/SignIn";
import Register from "./components/pages/Register";
import Logout from "./components/pages/Logout";


export default function App() {

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Root />,
      children: [
        {
        path: "/",
        element: <Home />,
        },
        {
          path: "/employees",
          element: <Employees />,
          }
      ]
    },
    {
      element: <AuthLayout />,
      children: [
        {
          path: "signin",
          element: <SignIn />,
        },
        {
          path: "register",
          element: <Register />,
        },
        {
          path: "logout",
          element: <Logout />,
        },
      ],
    },
  ]);

  return (
    <RouterProvider router={router} />
  );
}
