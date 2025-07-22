import React, {useState,useEffect,useContext} from 'react'
import Newreminder from './Newreminder.jsx'
import Table from './Table.jsx'
import {OpenContext, OpenProvider} from './MyContext.jsx'
function ActiveMessages()
{
  const { open, setOpen } = useContext(OpenContext);
  const [message_data, setmessage_data] = useState([])
  const token = localStorage.getItem("ReminderToken")
  async function Getdata() {
    const endpoint = 'http://localhost:8080/getMessages';
    const res = await fetch(endpoint, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
    });

    if (!res.ok) {
      console.error('Failed to fetch data');
      return;
    }

    const data = await res.json();
    setmessage_data(data)
  }
  useEffect(() => {
     Getdata()
  }, [open])  
   return (
    <div className="flex flex-col gap-1 justify-center items-center">
      <Table datas = {message_data}/>
      <button
        type="button"
        className="cursor-pointer transition-all bg-green-500 text-white px-6 py-2 rounded-lg border-green-600 border-b-[4px]
                   hover:brightness-110 hover:-translate-y-[1px] hover:border-b-[6px]
                   active:border-b-[2px] active:brightness-90 active:translate-y-[2px]"
	    onClick = {()=> setOpen(prev => !prev)}
      >
        + Add Reminder
      </button>

	  { open && 
		  <Newreminder/>
	  }
    </div>
  )
}

export default ActiveMessages
