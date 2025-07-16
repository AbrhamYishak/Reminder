import React , {useState} from 'react'
import { X } from 'lucide-react';
function Newreminder() {
  const [link,setlink] = useState("")
  const [hour, sethour] = useState("")
  const [d, setd] = useState("")
  const [time, settime] = useState("")
  const [message, setmessage] = useState("")
  console.log(d)
  return (
  <div className="fixed w-full max-w-[300px] bg-white rounded-lg shadow-md p-6 z-10 top-1/6">
	<div className="flex flex-col text-red-500">
	  <button type="button" className="self-end"><X/></button>
	  </div>
    <h2 className="text-2xl font-bold text-gray-800 mb-4">New Reminder</h2>
    <form className="flex flex-col">
      <input type="text" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Link" value = {link} onChange = {(e)=>setlink(e.target.value)}/>
	  <input type="text" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="hour:min" value = {hour} onChange = {(e)=>sethour(e.target.value)}/>
      <input type="date" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Date" value = {d} onChange = {(e)=>setd(e.target.value)}/>
      <textarea name="message" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Message" value = {message} onChange = {(e)=>setmessage(e.target.value)}>
	  </textarea>
      <button type="submit" className="bg-gradient-to-r from-green-300 to-green-500 text-white font-bold py-2 px-4 rounded-md mt-4 hover:to-green-700 transition ease-in-out duration-150" >+ Add</button>
    </form>
  </div>
  )
}

export default Newreminder
