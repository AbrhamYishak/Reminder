import useState from 'react'
import Newreminder from './Newreminder.jsx'

function ActiveMessages()
{
  const [add, setadd] = useState(false)
  return (
    <div className="flex flex-col justify-center items-center">
      <button
        type="button"
        className="cursor-pointer transition-all bg-green-500 text-white px-6 py-2 rounded-lg border-green-600 border-b-[4px]
                   hover:brightness-110 hover:-translate-y-[1px] hover:border-b-[6px]
                   active:border-b-[2px] active:brightness-90 active:translate-y-[2px]"
	    onClick = {()=> setadd(!add)}
      >
        + Add Reminder
      </button>
	   <Newreminder/>
    </div>
  )
}

export default ActiveMessages
