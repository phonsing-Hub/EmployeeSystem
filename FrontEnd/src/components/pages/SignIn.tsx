import { useState, useEffect } from "react";
import axios,{AxiosError} from "axios";
import AOS from "aos";
import "aos/dist/aos.css";
import { useNavigate } from "react-router-dom";
import { Button, Checkbox, Divider, Input } from "@nextui-org/react";
import { MdAttachEmail } from "react-icons/md";
import { TbEyeSearch, TbEyeX } from "react-icons/tb";
import { FcLinux } from "react-icons/fc";

export default function SignIn() {
  useEffect(()=>{
    AOS.init();
  },[]);

  let DomainName = import.meta.env.VITE_DOMAIN_NAME;
  let navigate = useNavigate();
  const [isVisible, setIsVisible] = useState(false);
  const [loading, setLoading] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [emainErr, setEmailErr] = useState("");
  const [passwordErr, setPasswordErr] = useState("");

  const handleSubmit = async () => {
    interface ErrorResponse {
      email?: string;
      password?: string;
    }
    setEmailErr("");
    setPasswordErr("");
    try {
      setLoading(true);
      const res = await axios.post(
        `${DomainName}/auth/login`,
        {
          email,
          pass: password,
        },
        { withCredentials: true }
      );
      if (res.status === 200) {
        console.log("User registered successfully:", res.data);
        navigate("/");
      }
    } catch (error) {
      const err = error as AxiosError<ErrorResponse>;  
      if (err.response?.data?.email) {
        setEmailErr(err.response.data.email);
      }
      if (err.response?.data?.password) {
        setPasswordErr(err.response.data.password);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    
      <div
        className="md:w-[384px] w-full mx-10 flex flex-col gap-4"
        data-aos="zoom-in"
        data-aos-delay="300"
      >
        <div>
          <div className="flex items-end">
            <FcLinux size={64} />
            <p className="font-bold mb-1 text-3xl">Sign In</p>
          </div>
          <Divider />
        </div>
        <Input
          label="Email"
          placeholder="example@mail.com"
          size="sm"
          variant="faded"
          endContent={<MdAttachEmail size={33} className=" text-default-400" />}
          value={email}
          onValueChange={setEmail}
          isInvalid={!!emainErr}
          errorMessage={emainErr}
          type="email"
        />
        <Input
          label="Password"
          placeholder="a-z A-Z 0-9 >6 character"
          radius="sm"
          size="sm"
          variant="faded"
          type={isVisible ? "text" : "password"}
          endContent={
            <button
              className="focus:outline-none"
              type="button"
              onClick={() => setIsVisible(!isVisible)}
              aria-label="toggle password visibility"
            >
              {isVisible ? (
                <TbEyeX size={33} className=" text-default-400" />
              ) : (
                <TbEyeSearch size={33} className=" text-default-400" />
              )}
            </button>
          }
          value={password}
          onValueChange={setPassword}
          isInvalid={!!passwordErr}
          errorMessage={passwordErr}
        />
        <Checkbox color="success">Remember me</Checkbox>
        <div className="flex gap-2">
          <Button
            radius="sm"
            color="primary"
            variant="faded"
            isLoading={loading}
            onPress={handleSubmit}
          >
            Sign In
          </Button>
          <Button
            radius="sm"
            color="secondary"
            variant="light"
            onPress={() => navigate("/register")}
          >
            Sign Up
          </Button>
        </div>
      </div>
   
  );
}
