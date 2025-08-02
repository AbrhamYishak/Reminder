import Expands from './Expands.jsx'
function InTable({ datas }) {
  const handleDelete = async (id) => {
    try {
      const response = await fetch(`https://reminder-wgwj.onrender.com/deleteMessage/${id}`, {
        method: 'DELETE'
      });
      if (!response.ok) throw new Error('Delete failed');
      console.log(`Deleted item with id ${id}`);
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
              <th className="w-[20%] px-2 py-3">Link</th>
              <th className="w-[40%] px-2 py-3">Message</th>
              <th className="w-[20%] px-2 py-3">Time</th>
			  <th className="w-[10%] px-2 py-3"> Ac</th>
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
				<td>
                   <Expands time = {data.Time} message = {data.Message} link = {data.Link} />
				</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default InTable;
