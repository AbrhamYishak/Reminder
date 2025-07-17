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
    chrome.tabs.create({
    url: chrome.runtime.getURL("auth.html"),
   });
		console.log(data)
      }
    } catch (err) {
		navigate("/error")
    }
  };
	useEffect(() => {
	 handleCheck() 
	}, [])
	
}
