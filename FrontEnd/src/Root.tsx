import { useState, useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import Ripple from "./components/loading/Ripple";
import NavbarHeader from "./components/navbar/navbarHeader";
function Root() {
  const [auth, setAuth] = useState(true);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  useEffect(() => {
    if(!auth){
      navigate("/signin");
    }
  }, []);

  if (loading)
    return (
      <div className="w-screen h-screen flex justify-center items-center">
         <Ripple/>
        Loading ...
      </div>
    );
  return (
    <main>
      <NavbarHeader />
      <main className="container mx-auto max-w-screen-2xl px-6 br">
        <Outlet />
      </main>
    </main>
  );
}

export default Root;
