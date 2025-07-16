import React, {useEffect} from 'react'
import { useNavigate } from 'react-router-dom';

export default function Check() {
    const token = localStorage.getItem('ReminderToken');
	let navigate = useNavigate()
    const handleCheck = async () => {
    try {
      const res = await fetch('http://localhost:8080/checktoken', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
      });

      const data = await res.json();
      if (res.ok) {
		navigate("/dashboard")
		console.log(data)
      } else {
		navigate("/auth")
		console.log(data)
      }
    } catch (err) {
      console.error(err);
      alert('An error occurred check your network');
    }
  };
	useEffect(() => {
	 handleCheck() 
	}, [])
	
}
