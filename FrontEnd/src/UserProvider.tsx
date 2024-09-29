import { useState, useEffect, createContext, useContext, ReactNode } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

interface User {
  id: number;
  firstname: string;
  lastname: string;
  email: string;
  phonenumber: string;
  hiredate: string;
  departmentname: string;
  positions: string;
  salary: number;
}
const UserContext = createContext<User | null>(null);

interface UserProviderProps {
  children: ReactNode; 
}

export function UserProvider({ children }: UserProviderProps) {
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null); 

  const getUser = async () => {
    const DomainName = import.meta.env.VITE_DOMAIN_NAME;
    try {
      const res = await axios.get(`${DomainName}/auth`, {
        withCredentials: true,
      });
      if (res.status === 200) setUser(res.data);
    } catch (error) {
      console.error(error);
      navigate("/signin");
    }
  };

  useEffect(() => {
    getUser();
  }, []);

  // useEffect(() => {
  //  console.log(user)
  // }, [user]);

  return <UserContext.Provider value={user}>{children}</UserContext.Provider>;
}

export const useUser = () => {
  return useContext(UserContext);
};
