import React, {useEffect} from 'react'
import { useNavigate } from 'react-router-dom';

export default function Check() {
    const token = localStorage.getItem("ReminderToken");
	let navigate = useNavigate()
	if (token === null){
		const setuptoken = localStorage.getItem("ReminderSetupToken");
		console.log(setuptoken)
		if (setuptoken === null){
			chrome.tabs.create({
			url: chrome.runtime.getURL("auth.html"),
			});
		}else{
			const handleVerification = async () => {
				try {
				const res = await fetch("https://reminder-wgwj.onrender.com/getauthtoken", {
					method: 'POST',
					headers: { 'Content-Type': 'application/json',
					"Authorization": `Bearer ${setuptoken}`},
				});
				const data = await res.json();
				if (res.ok) {
				navigate("/dashboard")
                localStorage.setItem("ReminderToken", data.token);
				} else {
				if (data.message === "not verified") {
					navigate("/verify")
				}
			    else{
					chrome.tabs.create({
                    url: chrome.runtime.getURL("auth.html"),
                    });
				}
				return
				}
				} catch (err) {
					chrome.tabs.create({
                    url: chrome.runtime.getURL("auth.html"),
                    });
				}
				};
			handleVerification()
		}
    }
	const handleCheck = async () => {
	if (token !== null){
    try {
	  const endpoint = "https://reminder-wgwj.onrender.com/checktoken"
      const res = await fetch(endpoint, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
      });

      const data = await res.json();
	  console.log(res)
	  console.log(data)
      if (res.ok) {
		navigate("/dashboard")
		console.log(data)
      } else {
    chrome.tabs.create({
    url: chrome.runtime.getURL("auth.html"),
   });
	return
      }
    } catch (err) {
		console.log(err)
		alert(err)
		navigate("/error")
    }
}
  };
  useEffect(() => {
   handleCheck()
  }, [])
}
  
