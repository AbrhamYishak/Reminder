import {MemoryRouter, Routes, Route} from "react-router-dom"
import ActiveMessages from "./components/ActiveMessages.jsx"
import Navbar from "./components/Navbar.jsx"
import InactiveMessage from './components/InactiveMessage.jsx'
import Setting from './components/Setting.jsx'
import Login from './components/Login.jsx'
import VerfiyMail from './components/VerifyEmail.jsx'
import {createMemoryHistory} from "history"
function App() {
  return (
    <div className="flex flex-col m-1 min-w-[600px]">
	<MemoryRouter >
	<Routes>
	  <Route path = "/" element = {<Login/>}/>
	  <Route path = "/verfiymail" element = {<VerfiyMail/>}/>
	  <Route path = "/dashboard" element = {<Navbar/>}>
		<Route index element = {<ActiveMessages/>}/>
	    <Route path = "inactive" element = {<InactiveMessage/>}/>
	    <Route path = "setting" element = {<Setting/>}/>
	  </Route>
	  </Routes>
	  </MemoryRouter>
    </div>
  )
}

export default App
