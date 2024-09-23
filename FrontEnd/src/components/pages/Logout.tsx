import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Ripple from "../loading/Ripple";
import axios from "axios";

export default function Logout() {
  const navigate = useNavigate();

  const checkout = async () => {
    try {
      const DomainName = import.meta.env.VITE_DOMAIN_NAME;
      const res = await axios.get(`${DomainName}/checkout`, {
        withCredentials: true,
      });
      if (res.status === 200) {
        navigate("/signin");
      }
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    setTimeout(() => {
      checkout();
    }, 2000);
  }, []); 

  return <Ripple />;
}
