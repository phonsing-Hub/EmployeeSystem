import { Outlet } from "react-router-dom";
import NavbarHeader from "./components/navbar/navbarHeader";
import { UserProvider, useUser } from "./UserProvider";
import Ripple from "./components/loading/Ripple";
function Root() {
  const user = useUser();

  return (
    <main>
      <UserProvider>
        {!user ? (
          <>
            <NavbarHeader />
            <main className="mx-2 px-6">
              <Outlet />
            </main>
          </>
        ) : (
          <div className="w-full h-screen flex justify-center items-center">
            <Ripple />
            <h1>Loading ...</h1>
          </div>
        )}
      </UserProvider>
    </main>
  );
}

export default Root;
