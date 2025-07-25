import React, { useState, useRef, useEffect } from 'react';
import {Expand} from "lucide-react"
export default function Expands({link, message, time}) {
  const [open, setOpen] = useState(false);

  const togglePopover = () => setOpen(prev => !prev);
  const popoverRef = useRef(null);
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (
        popoverRef.current &&
        !popoverRef.current.contains(event.target)
      ) {
        setOpen(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  return (
    <>
      <button
        onClick={togglePopover}
        type="button"
		className="truncate text-[0.25rem] bg-white border border-gray-200 hover:bg-gray-100 hover:text-blue-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:bg-gray-700"
      >
	    <Expand/>
      </button>

      {open && (
        <div
          ref={popoverRef}
          className="fixed top-1/4 left-1/2 z-50 transform -translate-x-1/2 -translate-y-1/2 w-80 max-w-full bg-white border border-gray-300 rounded-lg shadow-lg dark:bg-gray-800 dark:border-gray-700"
        >
<div className="fixed my-auto mx-auto w-full max-w-[400px] max-h-[80vh] overflow-y-auto overflow-x-hidden scrollbar-hide bg-white rounded-lg shadow-md p-6 z-50">
  <form className="flex flex-col">
    <input
      type="text"
      className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150 w-full break-words"
      value={link}
      disabled
    />

    <div className="flex justify-between items-center gap-4">
      <input
        type="text"
        className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150 w-full"
        placeholder="hour:min"
        value={time}
        disabled
      />
    </div>

    <textarea
      name="message"
      className="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150 w-full resize-none overflow-y-scroll"
      placeholder="Message"
      value={message}
      disabled
      rows={4}
    />
  </form>
</div>
        </div>
      )}
    </>
  );
}
