import { useState, useEffect } from "react";
import axios from "axios";
import TableEmp from "../tableEmp/table";
import { Users } from "../tableEmp/types";
function Employees() {
  const [users, setUsers] = useState<Users[]>([
    {
      id: 0,
      firstname: "",
      lastname: "",
      email: "",
      phone: "",
      department: "",
      positions: "",
      salary: 0,
    },
  ]);

  const getEmployees = async () =>{
    try {
      const DomainName = import.meta.env.VITE_DOMAIN_NAME;
      const res = await axios.get(`${DomainName}/employees`,{
        withCredentials: true
      });
      if (res.status === 200) setUsers(res.data);
    } catch (error) {
      console.log(error)
    }
  }
  useEffect(()=>{
    getEmployees();
  },[]);

  return (
    <div className="mt-10">
      <TableEmp users={users} />
    </div>
  );
}

export default Employees;
