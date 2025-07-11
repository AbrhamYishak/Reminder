function Table({datas}) {
 const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/deleteMessage/${id}`, {
        method: 'DELETE'
      });
      if (!response.ok) {
        throw new Error('Delete failed');
      }

      setItems(items.filter(item => item.id !== id));
      console.log(`Deleted item with id ${id}`);
    } catch (err) {
      console.error(err);
    }
  }
  return (
    <div>
<div class="relative overflow-x-auto overflow-y-scroll h-[70vh] shadow-md sm:rounded-lg w-full m-3">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400 table-fixed">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
                <th scope="col" class="px-6 py-3">
                    Id
                </th>
                <th scope="col" class="px-6 py-3">
                    Link
                </th>
                <th scope="col" class="px-6 py-3">
                    Message
                </th>
                <th scope="col" class="px-6 py-3">
                    Time
                </th>
                <th scope="col" class="px-6 py-3">
                    Action
                </th>
            </tr>
        </thead>
        <tbody>
	     {datas.map(data => (
			  <tr class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700 border-gray-200">
                <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
					{data.ID}
                </th>
                <td class="px-6 py-4 max-w-1/6 break-words">
					{data.Email}
                </td>
                <td class="px-6 py-4 max-w-3/6 break-words">
					{data.Message}
                </td>
                <td class="px-6 py-4 max-w-1/6 break-words">
					{data.Time}
                </td>
                <td class="px-6 py-4">
				<div class="flex flex-col shadow-xs" role="group">
  <button type="button" class="px-2 py-1 text-sm font-medium text-gray-900 bg-white border border-gray-200 hover:bg-gray-100 hover:text-green-700 focus:z-10 focus:ring-2 focus:ring-green-700 focus:text-green-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-green-500 dark:focus:text-white">
    Completed
  </button>
  <button type="button" class="px-2 py-1 text-sm font-medium text-gray-900 bg-white border-t border-b border-gray-200 hover:bg-gray-100 hover:text-red-700 focus:z-10 focus:ring-2 focus:ring-red-700 focus:text-red-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-red-500 dark:focus:text-white"onClick = {() => handleDelete(data.ID)}>
    Delete
  </button>
  <button type="button" class="px-2 py-1 text-sm font-medium text-gray-900 bg-white border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-blue-500 dark:focus:text-white">
    Edit
  </button>
</div>
                </td>
            </tr>
		 ))}
        </tbody>
    </table>
</div>
	  </div>
  )
}

export default Table
