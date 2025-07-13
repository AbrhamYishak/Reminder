import React, {useState} from 'react'
import {useNavigate} from 'react-router-dom'
function Login() {
  let navigate = useNavigate()
  const [email, setemail] = useState('');
  const [password, setpassword] = useState('');

  const handleRegister = async (e) => {
	console.log(email)
	console.log(password)
    e.preventDefault();

    try {
      const res = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ "Email" : email, "Password": password }),
      });

      const data = await res.json();
      if (res.ok) {
        alert('Registered successfully. Check your email for verification.');
		navigate("/verifymail")
      } else {
        alert(data.message || 'Registration failed');
      }
    } catch (err) {
      console.error(err);
      alert('An error occurred check your network');
    }
  };
  return (
    <div className="flex flex-col justify-center items-center min-h-screen p-4">
      <div className="w-full max-w-md bg-white p-6 rounded-lg shadow-lg border border-gray-200">
        <div className="space-y-3 text-center flex flex-col">
          <h1 className="text-2xl font-bold">Welcome To Reminder</h1>
	      <h1 className="text-xl font-bold self-center">Login/Register</h1>
	      
          <div className="space-y-2 text-left flex flex-col">
            <label
              className="text-sm font-medium"
              htmlFor="email"
            >
              Email
            </label>
            <input
              className="flex h-10 w-full rounded-md border px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              type="email"
              placeholder="abrham@example.com"
	          value = {email}
	          onChange = {(e)=>setemail(e.target.value)}
              required
            />
            <label
              className="text-sm font-medium"
              htmlFor="password"
            >
              Password
            </label>
            <input
              className="flex h-10 w-full rounded-md border px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              type="password"
              placeholder="password"
	          value = {password}
	          onChange = {(e)=>setpassword(e.target.value)}
              required
            />
            <label
              className=" text-blue-500 text-sm font-medium"
              htmlFor="email"
            >
              forgot your password ?
            </label>
	  <div className="flex gap-2 justify-center items-center">
           <button
            type="button"
            className="inline-flex self-center items-center justify-center mt-3 w-[40%] h-10 px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition text-lg"
          >
            Login
          </button>
           <button
            type="button"
            className="inline-flex self-center items-center justify-center mt-3 w-[40%] h-10 px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition text-lg"
	        onClick = {handleRegister}
          >
            Register
          </button>
	  </div>
          </div>
          <div className="flex items-center space-x-2 my-4">
            <hr className="flex-grow border-gray-300" />
            <span className="text-gray-400 text-sm">OR</span>
            <hr className="flex-grow border-gray-300" />
          </div>
          <button
            type="button"
            className="inline-flex items-center justify-center w-full h-10 px-4 py-2 bg-[#4285F4] text-white rounded-md hover:bg-[#357AE8] transition"
          >
            Continue with Google
          </button>
        </div>
      </div>
    </div>
  )
}

export default Login
