import {BrowserRouter, Routes, Route} from "react-router-dom";
import MfosScreen from "../Screens/MfosModule/MfosScreen";
import MfoCardScreen from "../Screens/MfoCardModule/MfoCardScreen";
import AdminScreen from "../Screens/AdminModule/AdminScreen";

const Router = () => {
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<MfosScreen />} />
            <Route path="/admin" element={<AdminScreen />} />
            <Route path="/mfo/:id" element={<MfoCardScreen />} />
            <Route path="/mfo/" element={<MfoCardScreen />} />
            <Route path="*" element={<div>Not found</div>} />
            <Route path="/specification/:id" element={<MfosScreen />} />
        </Routes>
    </BrowserRouter>
}

export default Router;