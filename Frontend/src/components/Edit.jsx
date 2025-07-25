import React, { useState, useRef, useEffect } from 'react';
import {Pencil} from "lucide-react"
export default function Edit({id, link, message}) {
  const [open, setOpen] = useState(false);
  const [hour, sethour] = useState("")
  const [d, setd] = useState("")
  const [t, sett] = useState("AM")
  const [m, setm] = useState(message)
  const [l, setl] = useState(link)
  const token = localStorage.getItem("ReminderToken")
  const togglePopover = () => setOpen(prev => !prev);
  const popoverRef = useRef(null);
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (
        popoverRef.current &&
        !popoverRef.current.contains(event.target)
      ) {
        setOpen(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);
  const handleEdit = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/editMessage/${id}`, {
        method: 'PATCH',
      headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
		body: JSON.stringify({ Message: m, Link: l, Hour: hour, Date:d, Meridiem : t }),
    });
      if (!response.ok) throw new Error('Delete failed');
      console.log(`Deleted item with id ${id}`);
	  location.reload()
    } catch (err) {
      console.error(err);
    }
  };
  return (
    <>
      <button
        onClick={togglePopover}
        type="button"
		className="truncate text-[0.25rem] bg-white border border-gray-200 hover:bg-gray-100 hover:text-blue-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:bg-gray-700"
      >
        <Pencil/>
      </button>

      {open && (
        <div
          ref={popoverRef}
          className="fixed top-1/5 left-1/2 z-50 transform -translate-x-1/2 -translate-y-1/2 w-80 max-w-full bg-white border border-gray-300 rounded-lg shadow-lg dark:bg-gray-800 dark:border-gray-700"
        >
<div className="fixed my-auto mx-auto w-full max-w-[400px] max-h-[80vh] overflow-y-auto overflow-x-hidden scrollbar-hide bg-white rounded-lg shadow-md p-6 z-50">
  <form className="flex flex-col">
    <input
      type="text"
      className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150 w-full break-words"
      value={l}
	  onChange = {(e)=>setl(e.target.value)}
    />
	<div className="flex justify-between items-center">
	<input type="text" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="hour:min" value = {hour} onChange = {(e)=>sethour(e.target.value)}/>
	<div class="inset-y-0 my-aut0 flex items-center bg-white pb-2">
	<select class="text-lg outline-none rounded-lg h-full" value = {t} onChange={(e)=>sett(e.target.value)}>
        <option>AM</option>
        <option>PM</option>
    </select>
    </div>
	</div>
    <input type="date" className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-50 transition ease-in-out duration-150" placeholder="Date" value = {d} onChange = {(e)=>setd(e.target.value)}/>
    <textarea
      name="message"
      className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150 w-full resize-none overflow-y-scroll"
      placeholder="Message"
      value={m}
      rows={4}
	  onChange = {(e)=>setm(e.target.value)}
    />
<button type="button" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800" onClick={()=>handleEdit(id)}>Edit</button>
  </form>
</div>
</div>
      )}
    </>
  );
}
