import React from 'react'

function Login() {
  return (
    <div className="flex justify-center items-center min-h-screen p-4">
      <div className="w-full max-w-md bg-white p-6 rounded-lg shadow-lg border border-gray-200">
        <div className="space-y-4 text-center">
          <h1 className="text-3xl font-bold">Welcome</h1>
          <div className="space-y-2 text-left">
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
              required
            />
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
