import React, {useState,useEffect,useContext} from 'react'
import Newreminder from './Newreminder.jsx'
import Table from './Table.jsx'
import Loader from './Loader.jsx'
import {OpenContext, OpenProvider} from './MyContext.jsx'
function ActiveMessages()
{
  const { open, setOpen } = useContext(OpenContext);
  const [message_data, setmessage_data] = useState([])
  const [loading, setloading] = useState(true)
  const token = localStorage.getItem("ReminderToken")
  async function Getdata() {
    const endpoint = `${process.env.url}/getMessages`;
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
	setloading(false)
  }
  useEffect(() => {
     Getdata()
  }, [open])  
   return (
    <div className="flex flex-col gap-1 justify-center items-center">
      {loading ? (
		  <Loader/>
      ) : (
		  <Table datas = {message_data}/>
      )}
      <button type="button" class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800" onClick={()=> setOpen(prev => !prev)}>+ Add Reminder</button>
	  { open && 
		  <Newreminder/>
	  }
    </div>
  )
}

export default ActiveMessages
