import { Navigate } from 'react-router-dom';

export default function Check() {
    const token = localStorage.getItem('token');
	console.log("in check")
    return token ? <Navigate to = "/dashboard" /> : <Navigate to="/auth" />;
}
