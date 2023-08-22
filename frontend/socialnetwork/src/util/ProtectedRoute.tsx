import React, { ReactNode, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../context/AuthContext';
import { getCookie } from '../controllers/SetUserContextAndCookie';

interface ProtectedRouteProps {
    element: ReactNode
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
    element
}) => {
    const userContext = useContext(UserContext);
    const navigate = useNavigate();
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const checkUserToken = () => {
        const userToken = getCookie("sessionToken");
        if (!userToken || userToken === 'undefined') {
            setIsLoggedIn(false);
            return navigate('/auth/login');
        }
        setIsLoggedIn(true);
    }
    useEffect(() => {
            checkUserToken();
        }, [isLoggedIn]);
    return (
        <React.Fragment>
            {
                isLoggedIn ? element : null
            }
        </React.Fragment>
    );
}

export default ProtectedRoute