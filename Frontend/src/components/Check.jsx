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
				const res = await fetch('http://localhost:8080/getauthtoken', {
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
					alert("token not working")
				}
				return
				}
				} catch (err) {
				navigate("/error")
				}
				};
			handleVerification()
		}
    }
	const handleCheck = async () => {
	if (token !== null){
    try {
      const res = await fetch('http://localhost:8080/checktoken', {
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
		navigate("/error")
    }
}
  };
  useEffect(() => {
   handleCheck()
  }, [])
}
  
