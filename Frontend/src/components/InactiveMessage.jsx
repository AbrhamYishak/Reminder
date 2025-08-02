import React , {useState, useEffect} from 'react'
import Loader from './Loader.jsx'
import InTable from './InTable.jsx'
function InactiveMessage() {
  const [message_data, setmessage_data] = useState([])
  const token = localStorage.getItem("ReminderToken")
  const [loading, setloading] = useState(true)
  async function Getdata() {
    const endpoint = 'http://localhost:8080/getInactiveMessages';
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
    console.log('Received data:', data);
    setmessage_data(data)
	setloading(false)
  }
  useEffect(() => {
     Getdata()
  }, []) 
  const handleDelete = async () => {
    try {
      const response = await fetch("http://localhost:8080/deleteInMessage", {
        method: 'DELETE',
      headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
    });
      if (!response.ok) throw new Error('Delete failed');
	  location.reload()
    } catch (err) {
      console.error(err);
    }
  };
   return (
    <div className="flex flex-col gap-1 justify-center items-center">
      {loading ? (
		  <Loader/>
      ) : (
		  <InTable datas = {message_data}/>
      )}
	<button type="button" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800" onClick={()=>handleDelete()}>Clear</button>
    </div>
  )
}


export default InactiveMessage
