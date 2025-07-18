import React , {useState,useContext} from 'react'
import { X } from 'lucide-react';
import {OpenContext} from './MyContext.jsx'
function Newreminder() {
  const [link,setlink] = useState("")
  const [hour, sethour] = useState("")
  const [d, setd] = useState("")
  const [message, setmessage] = useState("")
  const [t, sett] = useState("AM")
  const token = localStorage.getItem("ReminderToken")
  const { open, setOpen } = useContext(OpenContext); 
  console.log(open)
	chrome.tabs.query({ active: true, currentWindow: true }, (tabs)=>{
    const currentTab = tabs[0];
    const url = currentTab.url;
    setlink(url)
   });
  async function Getdata() {
	setOpen(!open)
    const endpoint = 'http://localhost:8080/createMessage';
    const res = await fetch(endpoint, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
		body: JSON.stringify({ Message: message, Link: link, Hour: hour, Date:d, Meridiem : t }),
    });

    if (!res.ok) {
      console.error('Failed to fetch data');
      return;
    }

    const data = await res.json();
    console.log('Received data:', data);
    setmessage_data(data)
  }
  return (
  <div className="fixed w-full max-w-[300px] bg-white rounded-lg shadow-md p-6 z-10 top-1/6">
	<div className="flex flex-col text-red-500">
	  <button type="button" className="self-end" onClick={() => setOpen(!open)}><X/></button>
	  </div>
    <h2 className="text-2xl font-bold text-gray-800 mb-4">New Reminder</h2>
    <form className="flex flex-col">
      <input type="text" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Link" value = {link} onChange = {(e)=>setlink(e.target.value)}/>
	  <div className="flex justify-between items-center">
	  <input type="text" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="hour:min" value = {hour} onChange = {(e)=>sethour(e.target.value)}/>
	  <div class="inset-y-0 my-aut0 flex items-center bg-white pb-2">
	  <select class="text-lg outline-none rounded-lg h-full" value = {t} onChange={(e)=>sett(e.target.value)}>
                        <option>AM</option>
                        <option>PM</option>
                    </select>
        </div>
	  </div>
      <input type="date" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Date" value = {d} onChange = {(e)=>setd(e.target.value)}/>
      <textarea name="message" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Message" value = {message} onChange = {(e)=>setmessage(e.target.value)}>
	  </textarea>
	  <button type="submit" className="bg-gradient-to-r from-green-300 to-green-500 text-white font-bold py-2 px-4 rounded-md mt-4 hover:to-green-700 transition ease-in-out duration-150" onClick = {Getdata} >+ Add</button>
    </form>
  </div>
  )
}

export default Newreminder
