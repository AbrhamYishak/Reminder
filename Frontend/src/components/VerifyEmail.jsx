import React , {useState} from 'react'
import {useNavigate,useLocation} from 'react-router-dom'
function Login() {
  let navigate = useNavigate()
  let location = useLocation()
  const email = location.state.email
  const [token, settoken] = useState('');
  const handleVerification = async (e) => {
	console.log(email)
    e.preventDefault();

    try {
      const res = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
		  body: JSON.stringify({ "Email": email , "VerificationToken" : token}),
      });

      const data = await res.json();
      if (res.ok) {
        alert('Verfication Completed');
		navigate("/dashboard")
      } else {
        alert(data.message || 'Verification failed');
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
          <h1 className="text-2xl font-bold">Verify Email</h1>
          <div className="space-y-2 text-left flex flex-col">
            <label
              className="text-sm font-medium"
              htmlFor="email"
            >
              Token
            </label>
            <input
              className="flex h-10 w-full rounded-md border px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
              type="text"
	          value = {token}
	          onChange = {((e)=>settoken(e.target.value))}
              placeholder="token"
              required
            />
           <button
            type="button"
	        onClick = {handleVerification}
            className="inline-flex self-center items-center justify-center mt-3 w-[40%] h-10 px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition text-lg"
          >
            Verify
          </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Login
