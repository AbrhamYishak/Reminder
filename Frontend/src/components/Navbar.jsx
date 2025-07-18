import {Outlet, Link} from "react-router-dom";
import { MessageSquareText } from 'lucide-react';
import { History } from 'lucide-react';
import { Settings } from 'lucide-react';
function Nav() {
  return (
	  <div>
<div class="flex flex-col m-4 items-center justify-center">
  <div class="border border-gray-300 py-3 flex gap-1 shadow-xl rounded-md">
    <div class="group relative px-4 cursor-pointer">
      <div class="flex h-10 items-center gap-1 justify-center rounded-full hover:text-blue-500">
          <MessageSquareText/> 
		  <Link to = "/dashboard"><p className="text-lg">Active</p></Link>
      </div>
    </div>
    <div class="group relative px-4 cursor-pointer">
      <div class="flex h-10 items-center gap-1 justify-center rounded-full hover:text-blue-500">
	     <History/>
         <Link to = "/dashboard/inactive"><p className="text-lg">History</p></Link>
      </div>
</div>
<div class="group relative px-4 cursor-pointer">
  <div class="flex h-10  items-center gap-1 justify-center rounded-full hover:text-blue-500"> 
	    <Settings/>
        <Link to = "/dashboard/setting"><p className="text-lg">Setting</p></Link>
  </div>
    </div>
</div>
</div>
<Outlet/>
	  </div>
  )}
export default Nav
