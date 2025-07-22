import { HashRouter, Routes, Route } from "react-router-dom";
import ActiveMessages from "./components/ActiveMessages.jsx";
import Navbar from "./components/Navbar.jsx";
import InactiveMessage from './components/InactiveMessage.jsx';
import Setting from './components/Setting.jsx';
import Check from './components/Check.jsx';
import ErrorTv from './components/ErrorTv.jsx';
import { OpenProvider } from './components/MyContext.jsx';

function App() {
  return (
    <div className="flex flex-col m-1 min-w-[460px]">
      <HashRouter>
        <OpenProvider>
          <Routes>
            <Route path="/" element={<Check />} />
            <Route path="/error" element={<ErrorTv />} />
            <Route path="/dashboard" element={<Navbar />}>
              <Route index element={<ActiveMessages />} />
              <Route path="inactive" element={<InactiveMessage />} />
              <Route path="setting" element={<Setting />} />
            </Route>
          </Routes>
        </OpenProvider>
      </HashRouter>
    </div>
  );
}

export default App;
