import { Routes, Route, Navigate } from 'react-router-dom';
import DashboardPage from "@/pages/DashboardPage";
import MainLayout from "@/layouts/MainLayout";
import "./App.css";

function App() {
    return (
        <Routes>
            <Route element={<MainLayout />}>
                <Route path="/" element={<DashboardPage />} />
            </Route>
        </Routes>
    );
}

export default App;
