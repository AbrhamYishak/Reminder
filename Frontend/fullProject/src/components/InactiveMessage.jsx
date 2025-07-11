import React , {useState, useEffect} from 'react'
import Table from './Table.jsx'
function InactiveMessage() {
  const [message_data, setmessage_data] = useState([])
  async function Getdata() {
    const endpoint = 'http://localhost:8080/getInactiveMessages';
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
    </div>
  )
}


export default InactiveMessage
