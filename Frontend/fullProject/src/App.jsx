import {HashRouter, Routes, Route} from "react-router-dom"
import ActiveMessages from "./components/ActiveMessages.jsx"
import Navbar from "./components/Navbar.jsx"
import InactiveMessage from './components/InactiveMessage.jsx'
import Setting from './components/Setting.jsx'
function App() {
  return (
    <div className="w-full h-[90vh] flex flex-col m-3 min-w-[500px]">
	<HashRouter >
	  <Routes>
	  <Route path = "/" element = {<Navbar/>}>
		<Route index element = {<ActiveMessages/>}/>
	    <Route path = "/inactive" element = {<InactiveMessage/>}/>
	    <Route path = "/setting" element = {<Setting/>}/>
	  </Route>
	  </Routes>
	  </HashRouter>
    </div>
  )
}

export default App
