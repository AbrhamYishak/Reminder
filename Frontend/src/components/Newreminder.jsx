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

  async function HandleAi() {
    const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });

    const result = await chrome.scripting.executeScript({
      target: { tabId: tab.id },
      func: () => {
        const elements = document.querySelectorAll("h1, h2, h3, h4, h5, h6, p");
        let text = "";
        elements.forEach((el) => {
          if (el.innerText.trim()) {
            text += el.innerText.trim() + " ";
          }
        });
        return text;
      },
    });

    const extractedText = result[0].result;
	const endpoint = 'http://localhost:8080/getAiMessage';
	const res = await fetch(endpoint, {
	     method: 'POST',
	     headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
		body: JSON.stringify({ Message: extractedText}),
	   });

	   if (!res.ok) {
	     console.error('Could not Generate Message');
	     return;
	   }

	   const data = await res.json();
	setmessage(data.Message)
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
<div class="relative inline-flex items-center justify-center gap-4 group">
  <div
    class="absolute inset-0 duration-100 opacity-60 transitiona-all bg-gradient-to-r from-indigo-500 via-pink-500 to-yellow-400 rounded-xl blur-lg filter group-hover:opacity-100 group-hover:duration-200"
  ></div>
  <button
    type="button"
    class="group relative inline-flex items-center justify-center text-sm rounded-md bg-white px-2 py-1 font-semibold text-black transition-all duration-200 hover:bg-gray-100 hover:shadow-lg hover:-translate-y-0.5 hover:shadow-gray-600/30"
    title="payment"
	onClick = {HandleAi}
	  >
    Generate Message<svg
      aria-hidden="true"
      viewBox="0 0 10 10"
      height="10"
      width="10"
      fill="none"
      class="mt-0.5 ml-2 -mr-1 stroke-white stroke-2"
    >
      <path
        d="M0 5h7"
        class="transition opacity-0 group-hover:opacity-100"
      ></path>
      <path
        d="M1 1l4 4-4 4"
        class="transition group-hover:translate-x-[3px]"
      ></path>
    </svg>
  </button>
</div>
	  <button type="submit" className="bg-gradient-to-r from-green-300 to-green-500 text-white font-bold py-2 px-4 rounded-md mt-4 hover:to-green-700 transition ease-in-out duration-150" onClick = {Getdata} >+ Add</button>
    </form>
  </div>
  )
}

export default Newreminder
