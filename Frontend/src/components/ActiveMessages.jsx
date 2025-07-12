import React, {useState,useEffect} from 'react'
import Newreminder from './Newreminder.jsx'
import Table from './Table.jsx'
function ActiveMessages()
{
  const [add, setadd] = useState(false)
  const [message_data, setmessage_data] = useState([])
  async function Getdata() {
    const endpoint = 'http://localhost:8080/getMessages';
    const res = await fetch(endpoint, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });

    if (!res.ok) {
      console.error('Failed to fetch data');
      return;
    }

    const data = await res.json();
    console.log('Received data:', data);
    setmessage_data(data)
  }
  useEffect(() => {
     Getdata()
  }, [])
  
   return (
    <div className="flex flex-col gap-3 justify-center items-center">
      <Table datas = {message_data}/>
      <button
        type="button"
        className="cursor-pointer transition-all bg-green-500 text-white px-6 py-2 rounded-lg border-green-600 border-b-[4px]
                   hover:brightness-110 hover:-translate-y-[1px] hover:border-b-[6px]
                   active:border-b-[2px] active:brightness-90 active:translate-y-[2px]"
	    onClick = {()=> setadd(!add)}
      >
        + Add Reminder
      </button>

	  { add && <Newreminder/>}
    </div>
  )
}

export default ActiveMessages
