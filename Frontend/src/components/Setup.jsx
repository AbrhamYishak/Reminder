import React, {useState} from 'react'
import {useNavigate} from 'react-router-dom'
function Setup() {
  let navigate = useNavigate()
  const [time , settime] = useState('')
  const token = localStorage.getItem("ReminderToken")
  const handleSetup = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch('http://localhost:8080/setup', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
        body: JSON.stringify({"TimeZone":time}),
      });

      const data = await res.json();
      if (res.ok) {
        alert('Setup Completed.');
        window.location.href = chrome.runtime.getURL('popup.html');
        chrome.tabs.getCurrent((tab) => {
        if (tab && tab.id) {
        chrome.tabs.remove(tab.id);
        }
        });
      } else {
        alert(data.message || 'Setup failed');
      }
    } catch (err) {
      console.error(err);
      alert('An error occurred check your network');
    }
  };
  return (
    <div className="flex flex-col justify-center items-center min-h-screen p-4">
      <div className="w-full max-w-md bg-white p-6 rounded-lg shadow-lg border border-gray-200">
        <div className="space-y-3 text-center flex flex-col">
          <h1 className="text-2xl font-bold">Setup your time zone</h1>
	      
          <div className="space-y-2 text-left flex flex-col">
            <label
              className="text-sm font-medium"
              htmlFor="email"
            >
              Time Zone
            </label>
  <select class="text-sm outline-none rounded-lg h-full" onChange = {(e)=>settime(e.target.value)}>
                        <option value="UTC-12:00"> (BIT) UTC-12:00</option>
   <option value="UTC-11:00"> (NST) UTC-11:00</option>
<option value="UTC-10:00"> (HST) UTC-10:00</option>
<option value="UTC-09:00"> (AKST) UTC-09:00</option>
<option value="UTC-08:00"> (PST) UTC-08:00</option>
<option value="UTC-07:00"> (MST) UTC-07:00</option>
<option value="UTC-06:00"> (CST) UTC-06:00</option>
<option value="UTC-05:00"> (EST) UTC-05:00</option>
<option value="UTC-04:00"> (AST) UTC-04:00</option>
<option value="UTC-03:00"> (ART) UTC-03:00</option>
<option value="UTC-02:00"> (GST) UTC-02:00</option>
<option value="UTC-01:00"> (CVT) UTC-01:00</option>
<option value="UTC±00:00"> (GMT) UTC±00:00</option>
<option value="UTC+01:00"> (CET) UTC+01:00</option>
<option value="UTC+02:00"> (EET) UTC+02:00</option>
<option value="UTC+03:00"> (EAT) UTC+03:00</option>
<option value="UTC+04:00"> (GST) UTC+04:00</option>
<option value="UTC+05:00"> (PKT) UTC+05:00</option>
<option value="UTC+06:00"> (BST) UTC+06:00</option>
<option value="UTC+07:00"> (ICT) UTC+07:00</option>
<option value="UTC+08:00"> (CST) UTC+08:00</option>
<option value="UTC+09:00"> (JST) UTC+09:00</option>
<option value="UTC+10:00"> (AEST) UTC+10:00</option>
<option value="UTC+11:00"> (SBT) UTC+11:00</option>
<option value="UTC+12:00"> (NZST) UTC+12:00</option>

                    </select>
           <button
            type="button"
            className="inline-flex self-center items-center justify-center mt-3 w-[40%] h-10 px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition text-lg"
	        onClick = {handleSetup}
          >
            Submit
          </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Setup
