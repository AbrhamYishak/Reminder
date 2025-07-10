function Newreminder() {
  return (
  <div class="w-full max-w-[300px] bg-white rounded-lg shadow-md p-6 z-10">
    <h2 class="text-2xl font-bold text-gray-800 mb-4">New Reminder</h2>

    <form class="flex flex-col">
      <input type="text" class="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Link"/>
      <input type="text" class="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Title"/>
      <input type="text" class="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="min/hour"/>
      <input type="date" class="bg-gray-100 text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Date"/>
      <textarea name="message" class="bg-gray-100 h-[30vh] text-gray-800 border-0 rounded-md p-2 mb-4 focus:bg-gray-200 focus:outline-none focus:ring-1 focus:ring-blue-500 transition ease-in-out duration-150" placeholder="Message">
	  </textarea>
      <button type="submit" class="bg-gradient-to-r from-green-300 to-green-500 text-white font-bold py-2 px-4 rounded-md mt-4 hover:to-green-700 transition ease-in-out duration-150">+ Add</button>
    </form>
  </div>
  )
}

export default Newreminder
