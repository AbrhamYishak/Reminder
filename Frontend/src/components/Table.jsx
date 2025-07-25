import React, {useContext,useState} from 'react'
import { Pencil } from 'lucide-react';
import { Check } from 'lucide-react';
import { Trash } from 'lucide-react';
import Expands from './Expands.jsx'
import Edit from './Edit.jsx'
import {OpenContext} from './MyContext.jsx'
function Table({ datas }) {
  const token = localStorage.getItem("ReminderToken")
  const [link, setlink] = useState("")
  chrome.tabs.query({ active: true, currentWindow: true }, (tabs)=>{
    const currentTab = tabs[0];
    const url = currentTab.url;
    setlink(url)
   });
  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/deleteMessage/${id}`, {
        method: 'DELETE',
      headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
    });
      if (!response.ok) throw new Error('Delete failed');
      console.log(`Deleted item with id ${id}`);
	  location.reload()
    } catch (err) {
      console.error(err);
    }
  };
  const handleComplete = async (id,l) => {
	  console.log(link)
	  console.log(l)
	if (link !== l){
		  alert ("the link is not same go to the the link in the email to complete")
		  return
	  }
    try {
      const response = await fetch(`http://localhost:8080/completeMessage/${id}`, {
        method: 'POST',
      headers: { 'Content-Type': 'application/json',
		"Authorization": `Bearer ${token}`},
    });
      if (!response.ok) throw new Error('Complation failed');
		console.log(`Completed item with id ${id}`);
		location.reload()
    } catch (err) {
      console.error(err);
    }
  };
  return (
    <div className="w-full p-1">
      <div className="h-[70vh] shadow-md sm:rounded-lg w-full overflow-y-scroll">
        <table className="table-fixed w-full text-sm text-left text-gray-500 dark:text-gray-400 border-collapse">
          <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th className="w-[10%] px-2 py-3">Id</th>
              <th className="w-[25%] px-2 py-3">Link</th>
              <th className="w-[40%] px-2 py-3">Message</th>
              <th className="w-[15%] px-2 py-3">Time</th>
              <th className="w-[10%] px-2 py-3">Ac</th>
            </tr>
          </thead>
          <tbody>
            {datas.map((data) => (
              <tr key={data.ID} className="odd:bg-white even:bg-gray-50 dark:odd:bg-gray-900 dark:even:bg-gray-800 border-b dark:border-gray-700">
                <td className="px-2 py-2 font-medium text-gray-900 dark:text-white truncate whitespace-nowrap">
                  {data.ID}
                </td>
                <td className="px-2 py-2 truncate whitespace-nowrap overflow-hidden text-ellipsis" title={data.Link}>
                  {data.Link}
                </td>
                <td className="px-2 py-2 truncate whitespace-nowrap overflow-hidden text-ellipsis" title={data.Message}>
                  {data.Message}
                </td>
                <td className="px-2 py-2 truncate whitespace-nowrap" title={data.Time}>
                  {data.Time}
                </td>
                <td className="px-2 py-2">
                  <div className="flex flex-col gap-1 justify-center items-center">
                    <button type = "button" className="truncate text-[0.25rem] bg-white border border-gray-200 hover:bg-gray-100 hover:text-green-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:bg-gray-700"onClick={()=>handleComplete(data.ID,data.Link)}>
                      <Check/>
                    </button>
                    <button type = "button" onClick={() => handleDelete(data.ID)} className="truncate text-[0.25rem] bg-white border border-gray-200 hover:bg-gray-100 hover:text-red-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:bg-gray-700">
                      <Trash/>
                    </button>
				     <Edit id = {data.ID} link = {data.Link} message = {data.Message}/> 
				     <Expands link = {data.Link} message = {data.Message} time = {data.Time}/> 
				   </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Table;
