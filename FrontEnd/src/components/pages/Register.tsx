import { useState, useEffect } from "react";
import axios, {AxiosError} from "axios";
import AOS from "aos";
import "aos/dist/aos.css";
import { useNavigate } from "react-router-dom";
import { Button, Divider, Input } from "@nextui-org/react";
import { FaUser } from "react-icons/fa";
import { MdAttachEmail } from "react-icons/md";
import { TbEyeSearch, TbEyeX } from "react-icons/tb";
import { FcLinux } from "react-icons/fc";

export default function Register() {
  useEffect(()=>{
    AOS.init();
  },[]);
  let DomainName = import.meta.env.VITE_DOMAIN_NAME;
  let navigate = useNavigate();
  const [isVisible1, setIsVisible1] = useState(false);
  const [isVisible2, setIsVisible2] = useState(false);
  const [loading, setLoading] = useState(false);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmpassword, setConfirmPassword] = useState("");

  const [nameErr, setNamEerr] = useState("");
  const [emainErr, setEmailErr] = useState("");
  const [passwordErr, setPasswordErr] = useState("");
  const [confirmpasswordErr, setConfirmPasswordErr] = useState("");

  const handleSubmit = async () => {
    interface ErrorResponse {
      error_mail?: string;
    }
    setNamEerr("");
    setEmailErr("");
    setPasswordErr("");
    setConfirmPasswordErr("");

    if (name.length > 20 || !name) {
      setNamEerr("ชื่อไม่ควรเกิน 20 ตัวอักษรและไม่ควรเป็นค่าว่าง");
      return;
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      setEmailErr("รูปแบบอีเมลไม่ถูกต้อง");
      return;
    }

    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{6,}$/;
    if (!passwordRegex.test(password)) {
      setPasswordErr(
        "รหัสผ่านต้องประกอบด้วยตัวอักษรพิมพ์เล็ก, ตัวอักษรพิมพ์ใหญ่, ตัวเลข และความยาวอย่างน้อย 6 ตัวอักษร"
      );
      return;
    }

    if (password !== confirmpassword) {
      setConfirmPasswordErr("รหัสผ่านไม่ตรงกัน");
      return;
    }
    try {
      setLoading(true);
      const res = await axios.post(`${DomainName}/register`, {
        name,
        email,
        password,
      });
      if (res.status === 201) {
        //console.log("User registered successfully:", res.data);
        navigate("/signin");
      }
    } catch (error) {
      const err = error as AxiosError<ErrorResponse>; 
      setLoading(false);
      if (err.response?.data?.error_mail) {
        setEmailErr(err.response.data.error_mail);
      } else if (err.request) {
        console.error("No response received:", err.request);
      } else {
        console.error("Error setting up request:", err.message);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      className="md:w-[384px] w-full mx-10 flex flex-col gap-4"
      data-aos="zoom-in"
    >
      <div>
        <div className="flex items-end">
          <FcLinux size={64} />
          <p className="font-bold mb-1 text-3xl">Sign Up</p>
        </div>
        <Divider />
      </div>
      <div data-aos="fade-down" data-aos-delay="300">
        <Input
          label="Name"
          placeholder="1-20 character"
          size="sm"
          variant="faded"
          endContent={<FaUser size={33} className=" text-default-400" />}
          isInvalid={!!nameErr}
          errorMessage={nameErr}
          value={name}
          onValueChange={setName}
        />
      </div>
      <div
        className="flex flex-col gap-4"
        data-aos="zoom-in"
        //data-aos-delay="300"
      >
        <Input
          label="Email"
          placeholder="example@mail.com"
          size="sm"
          variant="faded"
          endContent={<MdAttachEmail size={33} className=" text-default-400" />}
          isInvalid={!!emainErr}
          errorMessage={emainErr}
          value={email}
          onValueChange={setEmail}
        />
        <Input
          label="Password"
          placeholder="a-z A-Z 0-9 >6 character"
          radius="sm"
          size="sm"
          variant="faded"
          type={isVisible1 ? "text" : "password"}
          endContent={
            <button
              className="focus:outline-none"
              type="button"
              onClick={() => setIsVisible1(!isVisible1)}
              aria-label="toggle password visibility"
            >
              {isVisible1 ? (
                <TbEyeX size={33} className=" text-default-400" />
              ) : (
                <TbEyeSearch size={33} className=" text-default-400" />
              )}
            </button>
          }
          isInvalid={!!passwordErr}
          errorMessage={passwordErr}
          value={password}
          onValueChange={setPassword}
        />
      </div>
      <div data-aos="fade-up" data-aos-delay="300">
        <Input
          label="Confirm Password"
          placeholder="Enter your password"
          radius="sm"
          size="sm"
          variant="faded"
          type={isVisible2 ? "text" : "password"}
          endContent={
            <button
              className="focus:outline-none"
              type="button"
              onClick={() => setIsVisible2(!isVisible2)}
              aria-label="toggle password visibility"
            >
              {isVisible2 ? (
                <TbEyeX size={33} className=" text-default-400" />
              ) : (
                <TbEyeSearch size={33} className=" text-default-400" />
              )}
            </button>
          }
          isInvalid={!!confirmpasswordErr}
          errorMessage={confirmpasswordErr}
          value={confirmpassword}
          onValueChange={setConfirmPassword}
        />
      </div>

      <div className="flex gap-2" data-aos="zoom-in-down">
        <Button
          radius="sm"
          color="primary"
          variant="faded"
          isLoading={loading}
          onPress={handleSubmit}
        >
          Sign Up
        </Button>
        <Button
          radius="sm"
          color="secondary"
          variant="light"
          onPress={() => navigate("/signin")}
        >
          Sign In
        </Button>
      </div>
    </div>
  );
}
