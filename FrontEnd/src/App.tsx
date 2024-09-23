import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Ripple from "./components/loading/Ripple";
import Root from "./Root";
import AuthLayout from "./AuthLayout";
import Home from "./components/pages/Home";
import SignIn from "./components/pages/SignIn";
import Register from "./components/pages/Register";
import Logout from "./components/pages/Logout";

export default function App() {

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Root />,
      children: [{
        path: "/",
        element: <Home />,
      }]
    },
    {
      element: <AuthLayout />,
      children: [
        {
          path: "signin",
          element: <SignIn />,
         // loader: redirectIfUser,
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
